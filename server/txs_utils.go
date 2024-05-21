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

		_, isPending, err := tx.client.TransactionByHash(context.Background(), tx.txHash)
		if err != nil {
			log.Printf("Failed to get transaction by hash: %v", err)
			continue
		}

		if isPending {
			log.Printf("transaction %s is pending", tx.txHash)
			continue
		}

		// Check if it's time for the next retry
		if time.Since(tx.lastRetry) >= tx.backoffDuration {
			receipt, err := tx.client.TransactionReceipt(context.Background(), tx.txHash)
			if receipt == nil {
				continue
			}

			customHash := generateCustomHash(tx.methodName, tx.timestamp, tx.value)
			if err == nil && receipt.Status == types.ReceiptStatusSuccessful {
				log.Printf("Transaction %s with method %s was successful", tx.txHash.Hex(), tx.methodName)
				delete(transactionPool, common.HexToHash(customHash))
				continue
			}

			revertReason, err := getRevertReason(tx.client, tx.tx, receipt.BlockNumber)
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
		jsonErr, ok := err.(JsonError)
		if ok {
			errorData = jsonErr.ErrorData().(string)
		} else {
			return "", fmt.Errorf("non-json error: %s", err)
		}
	}

	revertReason, err := decodeCustomError(errorData)
	if err != nil {
		return "", fmt.Errorf("failed to unpack revert reason: %v", err)
	}

	return revertReason, nil
}

// curl -X POST -H "Content-Type: application/json" --data '{"jsonrpc":"2.0","method":"eth_call","params": [{"from": "0xf39fd6e51aad88f6f4ce6ab8827279cfffb92266","to": "0x1f2c6e90f3df741e0191eabb1170f0b9673f12b3","gas": "0x2dc6c0","gasPrice": "0x3b9aca07","value": "0x0","data": "0x90095e9d00000000000000000000000000000000000000000000000000000000664c83982c8b5307f3dfb5f32fe626f23be4dabbd607f217a19511a67ad101dc4a06e04c"}, "0x2093"],"id":1}' http://localhost:8545
// curl -X POST -H "Content-Type: application/json" --data "{\"jsonrpc\":\"2.0\",\"id\":2,\"method\":\"eth_call\",\"params\":[{\"from\":\"0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266\",\"input\":\"0x90095e9d00000000000000000000000000000000000000000000000000000000664c83982c8b5307f3dfb5f32fe626f23be4dabbd607f217a19511a67ad101dc4a06e04c\",\"to\":\"0x1f2c6e90f3df741e0191eabb1170f0b9673f12b3\"},\"0x2093\"]}" http://localhost:8545

// curl -X POST -H "Content-Type: application/json" --data "{\"jsonrpc\":\"2.0\",\"id\":2,\"method\":\"eth_call\",\"params\":[{\"from\":\"0xf39fd6e51aad88f6f4ce6ab8827279cfffb92266\",\"input\":\"0x90095e9d00000000000000000000000000000000000000000000000000000000664c83982c8b5307f3dfb5f32fe626f23be4dabbd607f217a19511a67ad101dc4a06e04c\",\"to\":\"0x1f2c6e90f3df741e0191eabb1170f0b9673f12b3\"},\"0x2093\"]}" http://localhost:8545 | jq .
func abiUnpackRevertReason(data []byte) (string, error) {
	if len(data) < 4 || data[0] != 0x08 || data[1] != 0xc3 || data[2] != 0x79 || data[3] != 0xa0 {
		return "", fmt.Errorf("not a revert reason")
	}

	revertReasonBytes := data[4:]
	revertReason := new(big.Int).SetBytes(revertReasonBytes).String()

	return revertReason, nil
}
