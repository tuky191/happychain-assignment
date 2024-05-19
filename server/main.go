package main

import (
	"crypto/ecdsa"
	"encoding/hex"
	"log"
	"net/http"
	"sync"
	"time"

	httpclient "github.com/drand/drand/client/http"
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
			sequencerRandomValue := generateSequencerRandom()
			log.Printf("Generated sequencer random value: %s", sequencerRandomValue)

			drandValue, err := getDrandValue(drandClient, currentTime-delay)
			if err != nil {
				log.Printf("Error fetching Drand value: %v", err)
				continue
			}
			drandValueHex := hex.EncodeToString(drandValue)

			randomness, err := calculateSequencerRandomness(drandValueHex, sequencerRandomValue)
			if err != nil {
				log.Fatalf("unable to calculate randomness, %v", err)
			}
			log.Printf("Generated sequencer randomness: %s", randomness)

			commitments[currentTime] = randomness
			commitSequencerRandom(client, privateKey, currentTime, randomness)
		}

		if currentTime%3 == 0 {
			drandValue, err := getDrandValue(drandClient, currentTime)
			if err != nil {
				log.Printf("Error fetching Drand value: %v", err)
				continue
			}
			drandValueHex := hex.EncodeToString(drandValue)
			log.Printf("Fetched Drand value: %s", drandValueHex)
			updateDrandRandomness(client, privateKey, currentTime, drandValueHex)
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
		}
	}

}

func revealSequencerRandom(client *ethclient.Client, privateKey *ecdsa.PrivateKey, timestamp int64, randomValue string) {
	log.Printf("Revealing sequencer random value: %s at timestamp: %d", randomValue, timestamp)
	sendTransaction(client, privateKey, contractAddresses.SequencerRandomOracleAddress, "revealSequencerRandomness", timestamp, randomValue)
}

func updateDrandRandomness(client *ethclient.Client, privateKey *ecdsa.PrivateKey, timestamp int64, randomValue string) {
	log.Printf("Updating Drand value: %s at timestamp: %d", randomValue, timestamp)
	sendTransaction(client, privateKey, contractAddresses.DrandOracleAddress, "postDrandRandomness", timestamp, randomValue)
}

func commitSequencerRandom(client *ethclient.Client, privateKey *ecdsa.PrivateKey, timestamp int64, randomValueHash string) {
	log.Printf("Posting commitment to contract: %s at timestamp: %d", contractAddresses.SequencerRandomOracleAddress, timestamp)
	sendTransaction(client, privateKey, contractAddresses.SequencerRandomOracleAddress, "postRandomnessCommitment", timestamp, randomValueHash)
}
