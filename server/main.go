package main

import (
	"context"
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil/hdkeychain"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	bip39 "github.com/tyler-smith/go-bip39"
)

type DrandResponse struct {
	Round      int64  `json:"round"`
	Randomness string `json:"randomness"`
}

type ContractAddresses struct {
	DrandOracleAddress           string `json:"drandOracleAddress"`
	SequencerRandomOracleAddress string `json:"sequencerRandomOracleAddress"`
	RandomnessOracleAddress      string `json:"randomnessOracleAddress"`
}

var httpGet = http.Get

var (
	anvilURL          = getEnv("ANVIL_URL", "http://anvil:8545")
	mnemonic          = getEnv("MNEMONIC", "")
	delay             = getEnvAsInt("DELAY", 9)
	precommitDelay    = getEnvAsInt("PRECOMMIT_DELAY", 10)
	drandInterval     = getEnvAsInt("DRAND_INTERVAL", 3)
	drandGenesis      = getEnvAsInt("DRAND_GENESIS", 1609459200)
	contractAddresses ContractAddresses
)

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func getEnvAsInt(name string, defaultValue int64) int64 {
	valueStr := getEnv(name, "")
	if value, err := strconv.ParseInt(valueStr, 10, 64); err == nil {
		return value
	}
	return defaultValue
}

func derivePrivateKeyFromMnemonic(mnemonic string) (*ecdsa.PrivateKey, error) {
	// Generate seed from mnemonic
	seed := bip39.NewSeed(mnemonic, "")

	// Derive the master key from the seed
	masterKey, err := hdkeychain.NewMaster(seed, &chaincfg.MainNetParams)
	if err != nil {
		return nil, err
	}

	// Derive the account key from the master key
	purpose, _ := masterKey.Child(44 + hdkeychain.HardenedKeyStart) // Purpose: BIP44
	coinType, _ := purpose.Child(60 + hdkeychain.HardenedKeyStart)  // CoinType: 60 for Ethereum
	account, _ := coinType.Child(0 + hdkeychain.HardenedKeyStart)   // Account: 0

	// Derive the first address key
	change, _ := account.Child(0) // External chain
	addressIndex, _ := change.Child(0)

	// Get the private key
	privateKey, err := addressIndex.ECPrivKey()
	if err != nil {
		return nil, err
	}

	return crypto.ToECDSA(privateKey.Serialize())
}

func loadContractAddresses(filename string) error {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	return json.Unmarshal(file, &contractAddresses)
}

func main() {
	if mnemonic == "" {
		log.Fatal("MNEMONIC environment variable is not set")
	}

	privateKey, err := derivePrivateKeyFromMnemonic(mnemonic)
	if err != nil {
		log.Fatalf("Failed to derive private key: %v", err)
	}

	client, err := ethclient.Dial(anvilURL)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	err = loadContractAddresses("/app/shared/addresses.json")
	if err != nil {
		log.Fatalf("Failed to load contract addresses: %v", err)
	}

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		currentTime := time.Now().Unix()

		if currentTime%2 == 0 {
			sequencerRandomValue := generateSequencerRandom()
			commitSequencerRandom(client, privateKey, currentTime, sequencerRandomValue)
		}

		if currentTime%3 == 0 {
			drandValue, err := getDrandValue(currentTime)
			if err != nil {
				log.Printf("Error fetching Drand value: %v", err)
				continue
			}
			addDrandValue(client, privateKey, currentTime, drandValue)
		}

		if (currentTime-precommitDelay)%2 == 0 {
			revealTime := currentTime - precommitDelay
			revealValue := generateSequencerRandom() // Fetch the actual committed value in practice
			revealSequencerRandom(client, privateKey, revealTime, revealValue)
		}
	}
}

func getDrandValue(timestamp int64) (string, error) {
	round := (timestamp - drandGenesis) / drandInterval
	url := fmt.Sprintf("https://api.drand.sh/public/%d", round)

	resp, err := httpGet(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var drandResp DrandResponse
	if err := json.Unmarshal(body, &drandResp); err != nil {
		return "", err
	}

	return drandResp.Randomness, nil
}

func generateSequencerRandom() string {
	h := sha256.New()
	h.Write([]byte(fmt.Sprintf("%d", time.Now().UnixNano())))
	return hex.EncodeToString(h.Sum(nil))
}

func commitSequencerRandom(client *ethclient.Client, privateKey *ecdsa.PrivateKey, timestamp int64, randomValue string) {
	commitment := sha256.Sum256([]byte(randomValue))
	commit(client, privateKey, contractAddresses.SequencerRandomOracleAddress, timestamp, commitment[:])
}

func revealSequencerRandom(client *ethclient.Client, privateKey *ecdsa.PrivateKey, timestamp int64, randomValue string) {
	reveal(client, privateKey, contractAddresses.SequencerRandomOracleAddress, timestamp, randomValue)
}

func addDrandValue(client *ethclient.Client, privateKey *ecdsa.PrivateKey, timestamp int64, value string) {
	sendTransaction(client, privateKey, contractAddresses.DrandOracleAddress, "addDrandValue", timestamp, value)
}

func commit(client *ethclient.Client, privateKey *ecdsa.PrivateKey, contractAddress string, timestamp int64, commitment []byte) {
	sendTransaction(client, privateKey, contractAddress, "commit", timestamp, hex.EncodeToString(commitment))
}

func reveal(client *ethclient.Client, privateKey *ecdsa.PrivateKey, contractAddress string, timestamp int64, value string) {
	sendTransaction(client, privateKey, contractAddress, "reveal", timestamp, value)
}

func sendTransaction(client *ethclient.Client, privateKey *ecdsa.PrivateKey, contractAddress string, methodName string, timestamp int64, value string) {
	contractABI, err := abi.JSON(strings.NewReader(`[{"constant":false,"inputs":[{"name":"timestamp","type":"uint256"},{"name":"value","type":"bytes32"}],"name":"` + methodName + `","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"}]`))
	if err != nil {
		log.Fatalf("Failed to parse contract ABI: %v", err)
	}

	fromAddress := crypto.PubkeyToAddress(privateKey.PublicKey)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatalf("Failed to get nonce: %v", err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatalf("Failed to suggest gas price: %v", err)
	}

	var data []byte
	if methodName == "commit" || methodName == "addDrandValue" {
		data, err = contractABI.Pack(methodName, big.NewInt(timestamp), common.HexToHash(value))
	} else if methodName == "reveal" {
		data, err = contractABI.Pack(methodName, big.NewInt(timestamp), value)
	}

	if err != nil {
		log.Fatalf("Failed to pack method call data: %v", err)
	}

	tx := types.NewTransaction(nonce, common.HexToAddress(contractAddress), big.NewInt(0), uint64(3000000), gasPrice, data)
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatalf("Failed to get network ID: %v", err)
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatalf("Failed to sign transaction: %v", err)
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatalf("Failed to send transaction: %v", err)
	}

	log.Printf("Sent transaction: %s with method: %s", signedTx.Hash().Hex(), methodName)
}
