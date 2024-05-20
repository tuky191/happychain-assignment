package main

import (
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/drand/drand/client"
	httpclient "github.com/drand/drand/client/http"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
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

type Transaction struct {
	client          *ethclient.Client
	privateKey      *ecdsa.PrivateKey
	contractAddress string
	methodName      string
	timestamp       int64
	value           string
	txHash          common.Hash
	attempts        int
	lastRetry       time.Time
	initialSend     time.Time
	backoffDuration time.Duration
	tx              *types.Transaction
	contractABI     abi.ABI
}

var transactionPool = make(map[common.Hash]*Transaction)
var poolMutex = &sync.Mutex{}

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
	revealTimestamps  = make(map[int64]int64)
)

func startConfirmationWorker() {
	go func() {
		for {
			time.Sleep(10 * time.Second) // Check every 10 seconds
			retryTransactions()
		}
	}()
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

	startConfirmationWorker()

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		currentTime := time.Now().Unix()
		log.Printf("Current time: %d", currentTime)

		if currentTime%2 == 0 {
			err := commitSequencerRandom(client, privateKey, currentTime, drandClient)
			if err != nil {
				log.Printf("commitSequencerRandom error: %v", err)
				continue
			}
		}

		if currentTime%3 == 0 {
			err := updateDrandRandomness(client, privateKey, currentTime, drandClient)
			if err != nil {
				log.Printf("updateDrandRandomness error: %v", err)
				continue
			}
		}

		for revealTime, commitTime := range revealTimestamps {
			if currentTime >= revealTime {
				err := revealSequencerRandom(client, privateKey, commitTime)
				if err != nil {
					log.Printf("revealSequencerRandom error: %v", err)
				}
				delete(revealTimestamps, revealTime)
			}
		}
	}
}

func revealSequencerRandom(client *ethclient.Client, privateKey *ecdsa.PrivateKey, commitTime int64) error {
	contractABI, err := abi.JSON(strings.NewReader(`[{"constant":false,"inputs":[{"name":"T","type":"uint256"},{"name":"randomness","type":"bytes32"}],"name":"revealSequencerRandomness","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"}]`))

	revealValue, exists := commitments[commitTime]
	if !exists {
		return fmt.Errorf("No committed value found for time: %d", commitTime)
	}
	log.Printf("Revealing sequencer random value: %s at timestamp: %d", revealValue, commitTime)

	if err != nil {
		return fmt.Errorf("failed to parse contract ABI: %v", err)
	}
	return sendTransaction(client, privateKey, contractAddresses.SequencerRandomOracleAddress, "revealSequencerRandomness", contractABI, commitTime, revealValue)
}

func commitSequencerRandom(client *ethclient.Client, privateKey *ecdsa.PrivateKey, currentTime int64, drandClient client.Client) error {
	sequencerRandomValue := generateSequencerRandom()
	log.Printf("Generated sequencer random value: %s", sequencerRandomValue)
	drandValue, err := getDrandValue(drandClient, currentTime-precommitDelay)
	if err != nil {
		return fmt.Errorf("Error fetching Drand value: %v", err)
	}
	drandValueHex := bufferToHex(drandValue)

	// This generates the randomness by adding the keccak256(drandValueHex and sequencerRandomValue)
	randomness, err := calculateSequencerRandomness(drandValueHex, sequencerRandomValue)
	if err != nil {
		return fmt.Errorf("unable to calculate randomness, %v", err)
	}
	log.Printf("Generated sequencer randomness: %s", randomness)
	commitments[currentTime] = randomness

	randomnessHash, err := calculateSequencerRandomnessHash(randomness)
	if err != nil {
		return fmt.Errorf("unable to calculate randomnessHash, %v", err)
	}
	log.Printf("Generated sequencer randomnessHash: %s for randomness: %s, randomness at timestamp: %d", randomnessHash, randomness, currentTime)
	log.Printf("Posting commitment to contract: %s at timestamp: %d", contractAddresses.SequencerRandomOracleAddress, currentTime)

	contractABI, err := abi.JSON(strings.NewReader(`[{"constant":false,"inputs":[{"name":"T","type":"uint256"},{"name":"randomnessHash","type":"bytes32"}],"name":"postRandomnessCommitment","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"}]`))
	if err != nil {
		return fmt.Errorf("failed to parse contract ABI: %v", err)
	}

	// Set the reveal timestamp
	revealTimestamps[currentTime+precommitDelay] = currentTime

	return sendTransaction(client, privateKey, contractAddresses.SequencerRandomOracleAddress, "postRandomnessCommitment", contractABI, currentTime, randomnessHash)
}

func updateDrandRandomness(client *ethclient.Client, privateKey *ecdsa.PrivateKey, currentTime int64, drandClient client.Client) error {
	randomValue, err := getDrandValue(drandClient, currentTime)
	if err != nil {
		return fmt.Errorf("Error fetching Drand value: %v", err)
	}
	randomValueHex := bufferToHex(randomValue)
	log.Printf("Fetched Drand value: %s", randomValueHex)
	log.Printf("Updating Drand value: %s at timestamp: %d", randomValueHex, currentTime)

	contractABI, err := abi.JSON(strings.NewReader(`[{"constant":false,"inputs":[{"name":"T","type":"uint256"},{"name":"randomness","type":"bytes32"}],"name":"postDrandRandomness","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"}]`))
	if err != nil {
		return fmt.Errorf("failed to parse contract ABI: %v", err)
	}
	return sendTransaction(client, privateKey, contractAddresses.DrandOracleAddress, "postDrandRandomness", contractABI, currentTime, randomValueHex)
}
