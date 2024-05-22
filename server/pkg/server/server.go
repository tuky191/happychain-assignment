package server

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"net/http"
	"server/v0/types/server"
	"strings"
	"sync"
	"time"

	"server/v0/pkg/utils"

	"github.com/drand/drand/client"
	httpclient "github.com/drand/drand/client/http"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Server struct {
	DrandClient      client.Client
	EthClient        *ethclient.Client
	Contracts        *server.Contracts
	PrivateKey       *ecdsa.PrivateKey
	Commitments      map[int64]string
	TransactionPool  map[common.Hash]*server.Transaction
	RevealTimestamps map[int64]int64
	PoolMutex        *sync.Mutex
	Config           server.Config
}

var transactionPool = make(map[common.Hash]*server.Transaction)

func NewServer(serverCfg server.Config) (*Server, error) {

	log.Println("Starting server with the following configuration:")
	log.Printf("Anvil URL: %s", serverCfg.AnvilURL)
	log.Printf("Drand URL: %s", serverCfg.DrandURL)
	log.Printf("Drand Chain Hash: %s", string(serverCfg.DrandChainHashBytes))

	ethClient, err := ethclient.Dial(serverCfg.AnvilURL)
	if err != nil {
		return nil, fmt.Errorf("Failed to connect to the Ethereum client: %v", err)
	}
	log.Println("Connected to the Ethereum client successfully.")

	drandClient, err := httpclient.New(serverCfg.DrandURL, serverCfg.DrandChainHashBytes, http.DefaultTransport)
	if err != nil {
		return nil, fmt.Errorf("Failed to create Drand HTTP client: %v", err)
	}
	log.Println("Created Drand HTTP client successfully.")

	contracts, err := utils.LoadContracts(&serverCfg.ContractAddresses, ethClient)
	if err != nil {
		return nil, fmt.Errorf("Failed to load Contracts: %v", err)
	}

	return &Server{
		DrandClient:      drandClient,
		EthClient:        ethClient,
		Commitments:      make(map[int64]string),
		TransactionPool:  make(map[common.Hash]*server.Transaction),
		PoolMutex:        &sync.Mutex{},
		RevealTimestamps: make(map[int64]int64),
		Config:           serverCfg,
		Contracts:        contracts,
	}, nil
}

func (s *Server) Start() {
	s.startRetryWorker()

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	for range ticker.C {
		currentTime := time.Now().Unix()
		log.Printf("Current time: %d", currentTime)

		if currentTime%2 == 0 {
			err := s.commitSequencerRandom(currentTime)
			if err != nil {
				log.Printf("commitSequencerRandom error: %v", err)
				continue
			}
		}

		if currentTime%3 == 0 {
			err := s.updateDrandRandomness(currentTime)
			if err != nil {
				log.Printf("updateDrandRandomness error: %v", err)
				continue
			}
		}

		for revealTime, commitTime := range s.RevealTimestamps {
			if currentTime >= revealTime {
				err := s.revealSequencerRandom(commitTime)
				if err != nil {
					log.Printf("revealSequencerRandom error: %v", err)
				}
				delete(s.RevealTimestamps, revealTime)
			}
		}
	}
}

func (s *Server) startRetryWorker() {
	go func() {
		for {
			time.Sleep(10 * time.Second) // Check every 10 seconds
			retryTransactions(s.PoolMutex)
		}
	}()
}

func (s *Server) revealSequencerRandom(commitTime int64) error {
	contractABI, err := abi.JSON(strings.NewReader(`[{"constant":false,"inputs":[{"name":"T","type":"uint256"},{"name":"randomness","type":"bytes32"}],"name":"revealSequencerRandomness","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"}]`))

	revealValue, exists := s.Commitments[commitTime]
	if !exists {
		return fmt.Errorf("No committed value found for time: %d", commitTime)
	}
	log.Printf("Revealing sequencer random value: %s at timestamp: %d", revealValue, commitTime)

	if err != nil {
		return fmt.Errorf("failed to parse contract ABI: %v", err)
	}

	return s.sendTransaction("revealSequencerRandomness", contractABI, commitTime, revealValue)
}

func (s *Server) commitSequencerRandom(currentTime int64) error {
	sequencerRandomValue := utils.GenerateSequencerRandom()
	log.Printf("Generated sequencer random value: %s", sequencerRandomValue)
	drandValue, err := s.getDrandValue(currentTime - s.Config.PrecommitDelay)
	if err != nil {
		return fmt.Errorf("Error fetching Drand value: %v", err)
	}
	drandValueHex := utils.BufferToHex(drandValue)

	// This generates the randomness by adding the keccak256(drandValueHex and sequencerRandomValue)
	randomness, err := utils.CalculateSequencerRandomness(drandValueHex, sequencerRandomValue)
	if err != nil {
		return fmt.Errorf("unable to calculate randomness, %v", err)
	}
	log.Printf("Generated sequencer randomness: %s", randomness)
	s.Commitments[currentTime] = randomness

	randomnessHash, err := utils.CalculateSequencerRandomnessHash(randomness)
	if err != nil {
		return fmt.Errorf("unable to calculate randomnessHash, %v", err)
	}

	log.Printf("Generated sequencer randomnessHash: %s for randomness: %s, randomness at timestamp: %d", randomnessHash, randomness, currentTime)
	log.Printf("Posting commitment to contract: %s at timestamp: %d", s.Config.ContractAddresses.SequencerRandomOracleAddress, currentTime)

	contractABI, err := abi.JSON(strings.NewReader(`[{"constant":false,"inputs":[{"name":"T","type":"uint256"},{"name":"randomnessHash","type":"bytes32"}],"name":"postRandomnessCommitment","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"}]`))
	if err != nil {
		return fmt.Errorf("failed to parse contract ABI: %v", err)
	}

	// Set the reveal timestamp
	s.RevealTimestamps[currentTime+s.Config.PrecommitDelay] = currentTime

	return s.sendTransaction("postRandomnessCommitment", contractABI, currentTime, randomnessHash)
}

func (s *Server) updateDrandRandomness(currentTime int64) error {
	randomValue, err := s.getDrandValue(currentTime)
	if err != nil {
		return fmt.Errorf("Error fetching Drand value: %v", err)
	}
	randomValueHex := utils.BufferToHex(randomValue)
	log.Printf("Fetched Drand value: %s", randomValueHex)
	log.Printf("Updating Drand value: %s at timestamp: %d", randomValueHex, currentTime)

	contractABI, err := abi.JSON(strings.NewReader(`[{"constant":false,"inputs":[{"name":"T","type":"uint256"},{"name":"randomness","type":"bytes32"}],"name":"postDrandRandomness","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"}]`))
	if err != nil {
		return fmt.Errorf("failed to parse contract ABI: %v", err)
	}
	return s.sendTransaction("postDrandRandomness", contractABI, currentTime, randomValueHex)
}

func (s *Server) getDrandValue(timestamp int64) ([]byte, error) {
	round := (timestamp - s.Config.DrandGenesis) / s.Config.DrandInterval
	log.Printf("Fetching Drand value for round: %d", round)

	result, err := s.DrandClient.Get(context.Background(), uint64(round))
	if err != nil {
		return nil, err
	}

	return result.Randomness(), nil
}

func (s *Server) sendTransaction(methodName string, contractABI abi.ABI, timestamp int64, value string) error {
	tx, err := callContract(methodName, s.Contracts, s.Config.TransactOpts, timestamp, value)
	if err != nil {
		return err
	}
	s.PoolMutex.Lock()
	customHash := utils.GenerateCustomHash(methodName, timestamp, value)
	transactionPool[common.HexToHash(customHash)] = &server.Transaction{
		Client:       s.EthClient,
		Contracts:    s.Contracts,
		TransactOpts: s.Config.TransactOpts,
		MethodName:   methodName,
		Timestamp:    timestamp,
		Value:        value,
		TxHash:       tx.Hash(),
		Attempts:     0,
		InitialSend:  time.Now(),
		Tx:           tx,
		ContractABI:  contractABI,
	}
	s.PoolMutex.Unlock()
	return nil
}
