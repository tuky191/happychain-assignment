package server

import (
	"crypto/ecdsa"
	"sync"
	"time"

	"github.com/drand/drand/client"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type ContractAddresses struct {
	DrandOracleAddress           string `json:"drandOracleAddress"`
	SequencerRandomOracleAddress string `json:"sequencerRandomOracleAddress"`
	RandomnessOracleAddress      string `json:"randomnessOracleAddress"`
}

type Transaction struct {
	Client          *ethclient.Client
	PrivateKey      *ecdsa.PrivateKey
	ContractAddress string
	MethodName      string
	Timestamp       int64
	Value           string
	TxHash          common.Hash
	Attempts        int
	LastRetry       time.Time
	InitialSend     time.Time
	BackoffDuration time.Duration
	Tx              *types.Transaction
	ContractABI     abi.ABI
}

type Config struct {
	PrivateKey          *ecdsa.PrivateKey
	DrandClient         client.Client
	Commitments         map[int64]string
	TransactionPool     map[common.Hash]*Transaction
	PoolMutex           *sync.Mutex
	AnvilURL            string
	Mnemonic            string
	Delay               int64
	PrecommitDelay      int64
	DrandInterval       int64
	DrandGenesis        int64
	DrandChainHashBytes []byte
	DrandURL            string
	ContractAddresses   ContractAddresses
}

type JsonError interface {
	Error() string
	ErrorCode() int
	ErrorData() interface{}
}
