package server

import (
	"crypto/ecdsa"
	"server/v0/pkg/contracts"
	"sync"
	"time"

	"github.com/drand/drand/client"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type ContractAddresses struct {
	DrandOracleAddress           string `json:"drandOracleAddress"`
	SequencerRandomOracleAddress string `json:"sequencerRandomOracleAddress"`
	RandomnessOracleAddress      string `json:"randomnessOracleAddress"`
}

type Contract[T any] struct {
	Instance T
}

type Contracts struct {
	DrandOracle           *Contract[*contracts.DrandOracle]
	RandomnessOracle      *Contract[*contracts.RandomnessOracle]
	SequencerRandomOracle *Contract[*contracts.SequencerRandomOracle]
}
type Transaction struct {
	Client *ethclient.Client

	Contracts       *Contracts
	TransactOpts    *bind.TransactOpts
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
	ChainID             int64
	Delay               int64
	PrecommitDelay      int64
	DrandInterval       int64
	DrandGenesis        int64
	DrandChainHashBytes []byte
	DrandURL            string
	ContractAddresses   ContractAddresses
	TransactOpts        *bind.TransactOpts
}

type JsonError interface {
	Error() string
	ErrorCode() int
	ErrorData() interface{}
}
