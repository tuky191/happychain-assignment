package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func sendTransaction(client *ethclient.Client, privateKey *ecdsa.PrivateKey, contractAddress string, methodName string, contractABI abi.ABI, timestamp int64, value string) error {
	tx, err := createTransaction(client, privateKey, contractAddress, methodName, contractABI, timestamp, value)
	if err != nil {
		return fmt.Errorf("Failed to create transaction: %v", err)
	}

	signedTx, err := signAndSendTransaction(client, tx, privateKey)
	if err != nil {
		return fmt.Errorf("Failed to send transaction: %v", err)
	}

	poolMutex.Lock()
	customHash := generateCustomHash(methodName, timestamp, value)
	transactionPool[common.HexToHash(customHash)] = &Transaction{
		client:          client,
		privateKey:      privateKey,
		contractAddress: contractAddress,
		methodName:      methodName,
		timestamp:       timestamp,
		value:           value,
		txHash:          signedTx.Hash(),
		attempts:        0,
		initialSend:     time.Now(),
		tx:              signedTx,
		contractABI:     contractABI,
	}
	poolMutex.Unlock()
	return nil
}

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
		log.Printf("Create transaction: method: %s, timestamp: %d, hashedValueBytes32Hex, %s, value: %s", methodName, timestamp, bufferToHex(valueBytes32[:]), value)
	case "revealSequencerRandomness":
		valueBytes32 := stringToBytes32(value)
		data, err = contractABI.Pack(methodName, big.NewInt(timestamp), valueBytes32)
		log.Printf("Create transaction: method: %s, timestamp: %d, hashedValueBytes32Hex, %s, value: %s", methodName, timestamp, bufferToHex(valueBytes32[:]), value)
	case "postDrandRandomness":
		valueBytes32 := stringToBytes32(value)
		data, err = contractABI.Pack(methodName, big.NewInt(timestamp), valueBytes32)
		log.Printf("Create transaction: method: %s, timestamp: %d, hashedValueBytes32Hex, %s, value: %s", methodName, timestamp, bufferToHex(valueBytes32[:]), value)
	default:
		return nil, fmt.Errorf("Create transaction: Unknown method: %s", methodName)
	}

	if err != nil {
		return nil, fmt.Errorf("failed to pack method call data for method %s: with value: %s, %v", methodName, value, err)
	}

	tx := types.NewTransaction(nonce, common.HexToAddress(contractAddress), big.NewInt(0), uint64(3000000), gasPrice, data)
	return tx, nil
}

func retryTransactions() {
	poolMutex.Lock()
	defer poolMutex.Unlock()

	for _, tx := range transactionPool {
		//	Wait at least 10 seconds before the first retry attempt
		if tx.attempts == 0 && time.Since(tx.initialSend) < 10*time.Second {
			continue
		}
		receipt, err := tx.client.TransactionReceipt(context.Background(), tx.txHash)
		if receipt == nil {
			log.Printf("transaction %s is pending", tx.txHash)
			continue
		}
		// Check if it's time for the next retry
		if time.Since(tx.lastRetry) >= tx.backoffDuration {
			customHash := generateCustomHash(tx.methodName, tx.timestamp, tx.value)
			if err == nil && receipt.Status == types.ReceiptStatusSuccessful {
				log.Printf("Transaction %s with method %s was successful", tx.txHash.Hex(), tx.methodName)
				delete(transactionPool, common.HexToHash(customHash))
				continue
			}

			revertReason, err := getRevertReason(tx.client, tx, receipt.BlockNumber)
			if err != nil {
				log.Fatalf("Failed to get revert reason for %s: %v", tx.txHash.Hex(), err)
			}
			log.Printf("Tx: %s, failed with revert reason: %s", tx.txHash.Hex(), revertReason)
			if tx.attempts >= 3 {
				log.Printf("Transaction %s with method %s failed after %d attempts", tx.txHash.Hex(), tx.methodName, tx.attempts)
				delete(transactionPool, common.HexToHash(customHash))
				continue
			}

			// Retry the transaction
			tx.attempts++
			tx.lastRetry = time.Now()
			tx.backoffDuration = time.Duration(math.Pow(2, float64(tx.attempts))) * time.Second

			newTx, err := createTransaction(tx.client, tx.privateKey, tx.contractAddress, tx.methodName, tx.contractABI, tx.timestamp, tx.value)
			if err != nil {
				log.Fatalf("Failed to create new transaction: %v", err)
			}

			signedTx, err := signAndSendTransaction(tx.client, newTx, tx.privateKey)
			if err != nil {
				fmt.Printf("Failed to resend transaction: %v", err)
				continue
			}
			log.Printf("Resent transaction: %s with method: %s (attempt %d)", signedTx.Hash().Hex(), tx.methodName, tx.attempts)
			tx.txHash = signedTx.Hash()
			tx.tx = signedTx
			transactionPool[common.HexToHash(customHash)] = tx
		}
	}
}

func getRevertReason(client *ethclient.Client, tx *Transaction, blockNumber *big.Int) (string, error) {
	msg := ethereum.CallMsg{
		To:   tx.tx.To(),
		Data: tx.tx.Data(),
	}

	ctx := context.Background()
	raw, err := client.CallContract(ctx, msg, blockNumber)
	if err != nil {
		return fmt.Sprintf("err: %v - raw: %s", err, string(raw)), nil
	}

	revertReason, err := abiUnpackRevertReason(raw)
	if err != nil {
		return "", fmt.Errorf("failed to unpack revert reason: %v", err)
	}

	return revertReason, nil
}

func abiUnpackRevertReason(data []byte) (string, error) {
	if len(data) < 4 || data[0] != 0x08 || data[1] != 0xc3 || data[2] != 0x79 || data[3] != 0xa0 {
		return "", fmt.Errorf("not a revert reason")
	}

	revertReasonBytes := data[4:]
	revertReason := new(big.Int).SetBytes(revertReasonBytes).String()

	return revertReason, nil
}
