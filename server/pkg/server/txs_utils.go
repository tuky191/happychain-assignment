package server

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math"
	"math/big"
	"server/v0/pkg/utils"
	"server/v0/types/server"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func signAndSendTransaction(client *ethclient.Client, tx *types.Transaction, privateKey *ecdsa.PrivateKey) (*types.Transaction, error) {
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to get network ID: %v", err)
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		return nil, fmt.Errorf("failed to sign transaction: %v", err)
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return nil, fmt.Errorf("failed to send transaction: %v", err)
	}

	log.Printf("Sent transaction: %s", signedTx.Hash().Hex())
	return signedTx, nil
}

func createTransaction(client *ethclient.Client, privateKey *ecdsa.PrivateKey, contractAddress string, methodName string, contractABI abi.ABI, timestamp int64, value string) (*types.Transaction, error) {

	fromAddress := crypto.PubkeyToAddress(privateKey.PublicKey)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to get nonce: %v", err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to suggest gas price: %v", err)
	}

	var data []byte
	switch methodName {
	case "postRandomnessCommitment":
		valueBytes32 := common.HexToHash(value)
		data, err = contractABI.Pack(methodName, big.NewInt(timestamp), valueBytes32)
		log.Printf("Create transaction: method: %s, timestamp: %d, hashedValueBytes32Hex, %s, value: %s", methodName, timestamp, utils.BufferToHex(valueBytes32[:]), value)
	case "revealSequencerRandomness":
		valueBytes32 := utils.StringToBytes32(value)
		data, err = contractABI.Pack(methodName, big.NewInt(timestamp), valueBytes32)
		log.Printf("Create transaction: method: %s, timestamp: %d, hashedValueBytes32Hex, %s, value: %s", methodName, timestamp, utils.BufferToHex(valueBytes32[:]), value)
	case "postDrandRandomness":
		valueBytes32 := utils.StringToBytes32(value)
		data, err = contractABI.Pack(methodName, big.NewInt(timestamp), valueBytes32)
		log.Printf("Create transaction: method: %s, timestamp: %d, hashedValueBytes32Hex, %s, value: %s", methodName, timestamp, utils.BufferToHex(valueBytes32[:]), value)
	default:
		return nil, fmt.Errorf("Create transaction: Unknown method: %s", methodName)
	}

	if err != nil {
		return nil, fmt.Errorf("failed to pack method call data for method %s: with value: %s, %v", methodName, value, err)
	}

	tx := types.NewTransaction(nonce, common.HexToAddress(contractAddress), big.NewInt(0), uint64(3000000), gasPrice, data)
	return tx, nil
}

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

			newTx, err := createTransaction(tx.Client, tx.PrivateKey, tx.ContractAddress, tx.MethodName, tx.ContractABI, tx.Timestamp, tx.Value)
			if err != nil {
				log.Fatalf("Failed to create new transaction: %v", err)
			}

			signedTx, err := signAndSendTransaction(tx.Client, newTx, tx.PrivateKey)
			if err != nil {
				fmt.Printf("Failed to resend transaction: %v", err)
				continue
			}
			log.Printf("Resent transaction: %s with method: %s (attempt %d)", signedTx.Hash().Hex(), tx.MethodName, tx.Attempts)
			tx.TxHash = signedTx.Hash()
			tx.Tx = signedTx
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
	var errorData string
	if err != nil {
		jsonErr, ok := err.(server.JsonError)
		if ok {
			errorData = jsonErr.ErrorData().(string)
		} else {
			return "", fmt.Errorf("non-json error: %s", err)
		}
	}

	revertReason, err := utils.DecodeCustomError(errorData)
	if err != nil {
		return "", fmt.Errorf("failed to unpack revert reason: %v", err)
	}

	return revertReason, nil
}
