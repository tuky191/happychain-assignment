package server

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/big"
	"server/v0/pkg/utils"
	"server/v0/types/server"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func retryTransactions(poolMutex *sync.Mutex) {
	poolMutex.Lock()
	defer poolMutex.Unlock()

	for _, tx := range transactionPool {
		//	Wait at least 10 seconds before the first retry attempt
		if tx.Attempts == 0 && time.Since(tx.InitialSend) < 10*time.Second {
			continue
		}

		_, isPending, err := tx.Client.TransactionByHash(context.Background(), tx.TxHash)
		if err != nil {
			log.Printf("Failed to get transaction by hash: %v", err)
			continue
		}

		if isPending {
			log.Printf("transaction %s is pending", tx.TxHash)
			continue
		}

		// Check if it's time for the next retry
		if time.Since(tx.LastRetry) >= tx.BackoffDuration {
			receipt, err := tx.Client.TransactionReceipt(context.Background(), tx.TxHash)
			if receipt == nil {
				continue
			}

			customHash := utils.GenerateCustomHash(tx.MethodName, tx.Timestamp, tx.Value)
			if err == nil && receipt.Status == types.ReceiptStatusSuccessful {
				log.Printf("Transaction %s with method %s was successful", tx.TxHash.Hex(), tx.MethodName)
				delete(transactionPool, common.HexToHash(customHash))
				continue
			}

			revertReason, err := getRevertReason(tx.Client, tx.Tx, receipt.BlockNumber)
			if err != nil {
				log.Fatalf("Failed to get revert reason for %s: %v", tx.TxHash.Hex(), err)
			}

			log.Printf("Tx: %s, failed with revert reason: %s", tx.TxHash.Hex(), revertReason)

			if tx.Attempts >= 3 {
				log.Printf("Transaction %s with method %s failed after %d attempts", tx.TxHash.Hex(), tx.MethodName, tx.Attempts)
				delete(transactionPool, common.HexToHash(customHash))
				continue
			}

			// Retry the transaction
			tx.Attempts++
			tx.LastRetry = time.Now()
			tx.BackoffDuration = time.Duration(math.Pow(2, float64(tx.Attempts))) * time.Second

			newTx, err := callContract(tx.MethodName, tx.Contracts, tx.TransactOpts, tx.Timestamp, tx.Value)
			log.Printf("Resent transaction: %s with method: %s (attempt %d)", newTx.Hash().Hex(), tx.MethodName, tx.Attempts)
			tx.TxHash = newTx.Hash()
			tx.Tx = newTx
			transactionPool[common.HexToHash(customHash)] = tx
		}
	}
}

func getRevertReason(client *ethclient.Client, tx *types.Transaction, blockNumber *big.Int) (string, error) {

	from, err := types.Sender(types.LatestSignerForChainID(tx.ChainId()), tx)
	if err != nil {
		return "", fmt.Errorf("unable to get sender, %v", err)
	}
	msg := ethereum.CallMsg{
		To:   tx.To(),
		Data: tx.Data(),
		From: from,
	}
	ctx := context.Background()
	_, err = client.CallContract(ctx, msg, blockNumber)

	revertReason, err := utils.DecodeCustomError(err)
	if err != nil {
		return "", fmt.Errorf("failed to unpack revert reason: %v", err)
	}

	return revertReason, nil
}

func callContract(methodName string, contracts *server.Contracts, transactOpts *bind.TransactOpts, timestamp int64, value string) (*types.Transaction, error) {
	var tx *types.Transaction
	var err error
	switch methodName {
	case "postRandomnessCommitment":
		var hashBytes32 [32]byte
		copy(hashBytes32[:], common.HexToHash(value).Bytes())
		tx, err = contracts.SequencerRandomOracle.Instance.SequencerRandomOracleTransactor.PostRandomnessCommitment(transactOpts, big.NewInt(timestamp), hashBytes32)
	case "revealSequencerRandomness":
		tx, err = contracts.SequencerRandomOracle.Instance.SequencerRandomOracleTransactor.RevealSequencerRandomness(transactOpts, big.NewInt(timestamp), utils.StringToBytes32(value))
	case "postDrandRandomness":
		tx, err = contracts.DrandOracle.Instance.PostDrandRandomness(transactOpts, big.NewInt(timestamp), utils.StringToBytes32(value))
	default:
		return nil, fmt.Errorf("Create transaction: Unknown method: %s", methodName)
	}
	log.Printf("sendTransaction: method: %s, timestamp: %d, value: %s", methodName, timestamp, value)

	if err != nil {
		decodedErr, decodingErr := utils.DecodeCustomError(err)
		if decodingErr != nil {
			return nil, fmt.Errorf("Failed execute transaction: %v", err)
		}
		return nil, fmt.Errorf("Failed execute transaction: %v", decodedErr)
	}
	return tx, nil
}
