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
	"github.com/drand/drand/client"
	httpclient "github.com/drand/drand/client/http"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	bip39 "github.com/tyler-smith/go-bip39"
	"golang.org/x/crypto/sha3"
)

var commitments = make(map[int64]string)

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
	drandChainHash    = getEnv("DRAND_CHAIN_HASH", "52db9ba70e0cc0f6eaf7803dd07447a1f5477735fd3f661792ba94600c84e971")
	drandURL          = getEnv("DRAND_URL", "https://api.drand.sh")
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
	log.Println("Starting server with the following configuration:")
	log.Printf("Anvil URL: %s", anvilURL)
	log.Printf("Drand URL: %s", drandURL)
	log.Printf("Drand Chain Hash: %s", drandChainHash)

	privateKey, err := derivePrivateKeyFromMnemonic(mnemonic)
	if err != nil {
		log.Fatalf("Failed to derive private key: %v", err)
	}
	log.Println("Derived private key from mnemonic successfully.")

	client, err := ethclient.Dial(anvilURL)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	log.Println("Connected to the Ethereum client successfully.")

	err = loadContractAddresses("/app/shared/addresses.json")
	if err != nil {
		log.Fatalf("Failed to load contract addresses: %v", err)
	}
	log.Println("Loaded contract addresses successfully.")

	chainHash, err := hex.DecodeString(drandChainHash)
	if err != nil {
		log.Fatalf("Failed to decode Drand chain hash: %v", err)
	}

	drandClient, err := httpclient.New(drandURL, chainHash, http.DefaultTransport)
	if err != nil {
		log.Fatalf("Failed to create Drand HTTP client: %v", err)
	}
	log.Println("Created Drand HTTP client successfully.")

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		currentTime := time.Now().Unix()
		log.Printf("Current time: %d", currentTime)

		if currentTime%2 == 0 {
			sequencerRandomValue := generateSequencerRandom()
			log.Printf("Generated sequencer random value: %s", sequencerRandomValue)
			commitments[currentTime] = sequencerRandomValue
			commitSequencerRandom(client, privateKey, currentTime, sequencerRandomValue)
		}

		if currentTime%3 == 0 {
			drandValue, err := getDrandValue(drandClient, currentTime)
			if err != nil {
				log.Printf("Error fetching Drand value: %v", err)
				continue
			}
			drandValueHex := hex.EncodeToString(drandValue)
			log.Printf("Fetched Drand value: %s", drandValueHex)
			addDrandValue(client, privateKey, currentTime, drandValueHex)
		}

		if (currentTime-precommitDelay)%2 == 0 {
			revealTime := currentTime - precommitDelay
			revealValue, exists := commitments[revealTime]
			if !exists {
				log.Printf("No committed value found for time: %d", revealTime)
				continue
			}
			log.Printf("Revealing sequencer random value: %s for time: %d", revealValue, revealTime)
			revealSequencerRandom(client, privateKey, revealTime, revealValue)

			// Calculate randomness(T)
			drandValue, err := getDrandValue(drandClient, revealTime)
			if err != nil {
				log.Printf("Error fetching Drand value for randomness calculation: %v", err)
				continue
			}
			drandValueHex := hex.EncodeToString(drandValue)
			randomness := calculateRandomness(drandValueHex, revealValue)
			log.Printf("Calculated randomness for time %d: %s", revealTime, randomness)
		}
	}

}

func getDrandValue(drandClient client.Client, timestamp int64) ([]byte, error) {
	round := (timestamp - drandGenesis) / drandInterval
	log.Printf("Fetching Drand value for round: %d", round)

	result, err := drandClient.Get(context.Background(), uint64(round))
	if err != nil {
		return nil, err
	}

	return result.Randomness(), nil
}

func keccak256(data []byte) []byte {
	hash := sha3.NewLegacyKeccak256()
	hash.Write(data)
	return hash.Sum(nil)
}

func calculateRandomness(drandValue, commitmentValue string) string {
	combined := []byte(drandValue + commitmentValue)
	hash := keccak256(combined)
	return hex.EncodeToString(hash)
}

func generateSequencerRandom() string {
	h := sha256.New()
	h.Write([]byte(fmt.Sprintf("%d", time.Now().UnixNano())))
	return hex.EncodeToString(h.Sum(nil))
}

func commitSequencerRandom(client *ethclient.Client, privateKey *ecdsa.PrivateKey, timestamp int64, randomValue string) {
	commitment := sha256.Sum256([]byte(randomValue))
	log.Printf("Committing sequencer random value: %s at timestamp: %d", randomValue, timestamp)
	commit(client, privateKey, contractAddresses.SequencerRandomOracleAddress, timestamp, commitment[:])
}

func revealSequencerRandom(client *ethclient.Client, privateKey *ecdsa.PrivateKey, timestamp int64, randomValue string) {
	log.Printf("Revealing sequencer random value: %s at timestamp: %d", randomValue, timestamp)
	reveal(client, privateKey, contractAddresses.SequencerRandomOracleAddress, timestamp, randomValue)
}

func addDrandValue(client *ethclient.Client, privateKey *ecdsa.PrivateKey, timestamp int64, value string) {
	log.Printf("Adding Drand value: %s at timestamp: %d", value, timestamp)
	sendTransaction(client, privateKey, contractAddresses.DrandOracleAddress, "addDrandValue", timestamp, value)
}

func commit(client *ethclient.Client, privateKey *ecdsa.PrivateKey, contractAddress string, timestamp int64, commitment []byte) {
	log.Printf("Committing value to contract: %s at timestamp: %d", contractAddress, timestamp)
	sendTransaction(client, privateKey, contractAddress, "commit", timestamp, hex.EncodeToString(commitment))
}

func reveal(client *ethclient.Client, privateKey *ecdsa.PrivateKey, contractAddress string, timestamp int64, value string) {
	log.Printf("Revealing value to contract: %s at timestamp: %d", contractAddress, timestamp)
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
		data, err = contractABI.Pack(methodName, big.NewInt(timestamp), common.HexToHash(value))
	}

	if err != nil {
		log.Fatalf("Failed to pack method call data for method %s: with value: %s, %v", methodName, value, err)
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
