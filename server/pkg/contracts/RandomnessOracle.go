// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// RandomnessOracleMetaData contains all meta data concerning the RandomnessOracle contract.
var RandomnessOracleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_drandOracle\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_sequencerRandomOracle\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"drandOracle\",\"outputs\":[{\"internalType\":\"contractDrandOracle\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sequencerRandomOracle\",\"outputs\":[{\"internalType\":\"contractSequencerRandomOracle\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"T\",\"type\":\"uint256\"}],\"name\":\"computeRandomness\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"T\",\"type\":\"uint256\"}],\"name\":\"isRandomnessEverAvailable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"T\",\"type\":\"uint256\"}],\"name\":\"simpleGetRandomness\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"T\",\"type\":\"uint256\"}],\"name\":\"unsafeGetRandomness\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// RandomnessOracleABI is the input ABI used to generate the binding from.
// Deprecated: Use RandomnessOracleMetaData.ABI instead.
var RandomnessOracleABI = RandomnessOracleMetaData.ABI

// RandomnessOracle is an auto generated Go binding around an Ethereum contract.
type RandomnessOracle struct {
	RandomnessOracleCaller     // Read-only binding to the contract
	RandomnessOracleTransactor // Write-only binding to the contract
	RandomnessOracleFilterer   // Log filterer for contract events
}

// RandomnessOracleCaller is an auto generated read-only Go binding around an Ethereum contract.
type RandomnessOracleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RandomnessOracleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RandomnessOracleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RandomnessOracleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RandomnessOracleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RandomnessOracleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RandomnessOracleSession struct {
	Contract     *RandomnessOracle // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RandomnessOracleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RandomnessOracleCallerSession struct {
	Contract *RandomnessOracleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// RandomnessOracleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RandomnessOracleTransactorSession struct {
	Contract     *RandomnessOracleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// RandomnessOracleRaw is an auto generated low-level Go binding around an Ethereum contract.
type RandomnessOracleRaw struct {
	Contract *RandomnessOracle // Generic contract binding to access the raw methods on
}

// RandomnessOracleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RandomnessOracleCallerRaw struct {
	Contract *RandomnessOracleCaller // Generic read-only contract binding to access the raw methods on
}

// RandomnessOracleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RandomnessOracleTransactorRaw struct {
	Contract *RandomnessOracleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRandomnessOracle creates a new instance of RandomnessOracle, bound to a specific deployed contract.
func NewRandomnessOracle(address common.Address, backend bind.ContractBackend) (*RandomnessOracle, error) {
	contract, err := bindRandomnessOracle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RandomnessOracle{RandomnessOracleCaller: RandomnessOracleCaller{contract: contract}, RandomnessOracleTransactor: RandomnessOracleTransactor{contract: contract}, RandomnessOracleFilterer: RandomnessOracleFilterer{contract: contract}}, nil
}

// NewRandomnessOracleCaller creates a new read-only instance of RandomnessOracle, bound to a specific deployed contract.
func NewRandomnessOracleCaller(address common.Address, caller bind.ContractCaller) (*RandomnessOracleCaller, error) {
	contract, err := bindRandomnessOracle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RandomnessOracleCaller{contract: contract}, nil
}

// NewRandomnessOracleTransactor creates a new write-only instance of RandomnessOracle, bound to a specific deployed contract.
func NewRandomnessOracleTransactor(address common.Address, transactor bind.ContractTransactor) (*RandomnessOracleTransactor, error) {
	contract, err := bindRandomnessOracle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RandomnessOracleTransactor{contract: contract}, nil
}

// NewRandomnessOracleFilterer creates a new log filterer instance of RandomnessOracle, bound to a specific deployed contract.
func NewRandomnessOracleFilterer(address common.Address, filterer bind.ContractFilterer) (*RandomnessOracleFilterer, error) {
	contract, err := bindRandomnessOracle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RandomnessOracleFilterer{contract: contract}, nil
}

// bindRandomnessOracle binds a generic wrapper to an already deployed contract.
func bindRandomnessOracle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RandomnessOracleMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RandomnessOracle *RandomnessOracleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RandomnessOracle.Contract.RandomnessOracleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RandomnessOracle *RandomnessOracleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RandomnessOracle.Contract.RandomnessOracleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RandomnessOracle *RandomnessOracleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RandomnessOracle.Contract.RandomnessOracleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RandomnessOracle *RandomnessOracleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RandomnessOracle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RandomnessOracle *RandomnessOracleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RandomnessOracle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RandomnessOracle *RandomnessOracleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RandomnessOracle.Contract.contract.Transact(opts, method, params...)
}

// ComputeRandomness is a free data retrieval call binding the contract method 0xb2b372b1.
//
// Solidity: function computeRandomness(uint256 T) view returns(bytes32)
func (_RandomnessOracle *RandomnessOracleCaller) ComputeRandomness(opts *bind.CallOpts, T *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _RandomnessOracle.contract.Call(opts, &out, "computeRandomness", T)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ComputeRandomness is a free data retrieval call binding the contract method 0xb2b372b1.
//
// Solidity: function computeRandomness(uint256 T) view returns(bytes32)
func (_RandomnessOracle *RandomnessOracleSession) ComputeRandomness(T *big.Int) ([32]byte, error) {
	return _RandomnessOracle.Contract.ComputeRandomness(&_RandomnessOracle.CallOpts, T)
}

// ComputeRandomness is a free data retrieval call binding the contract method 0xb2b372b1.
//
// Solidity: function computeRandomness(uint256 T) view returns(bytes32)
func (_RandomnessOracle *RandomnessOracleCallerSession) ComputeRandomness(T *big.Int) ([32]byte, error) {
	return _RandomnessOracle.Contract.ComputeRandomness(&_RandomnessOracle.CallOpts, T)
}

// DrandOracle is a free data retrieval call binding the contract method 0x09c21de0.
//
// Solidity: function drandOracle() view returns(address)
func (_RandomnessOracle *RandomnessOracleCaller) DrandOracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RandomnessOracle.contract.Call(opts, &out, "drandOracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DrandOracle is a free data retrieval call binding the contract method 0x09c21de0.
//
// Solidity: function drandOracle() view returns(address)
func (_RandomnessOracle *RandomnessOracleSession) DrandOracle() (common.Address, error) {
	return _RandomnessOracle.Contract.DrandOracle(&_RandomnessOracle.CallOpts)
}

// DrandOracle is a free data retrieval call binding the contract method 0x09c21de0.
//
// Solidity: function drandOracle() view returns(address)
func (_RandomnessOracle *RandomnessOracleCallerSession) DrandOracle() (common.Address, error) {
	return _RandomnessOracle.Contract.DrandOracle(&_RandomnessOracle.CallOpts)
}

// IsRandomnessEverAvailable is a free data retrieval call binding the contract method 0x06055269.
//
// Solidity: function isRandomnessEverAvailable(uint256 T) view returns(bool)
func (_RandomnessOracle *RandomnessOracleCaller) IsRandomnessEverAvailable(opts *bind.CallOpts, T *big.Int) (bool, error) {
	var out []interface{}
	err := _RandomnessOracle.contract.Call(opts, &out, "isRandomnessEverAvailable", T)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRandomnessEverAvailable is a free data retrieval call binding the contract method 0x06055269.
//
// Solidity: function isRandomnessEverAvailable(uint256 T) view returns(bool)
func (_RandomnessOracle *RandomnessOracleSession) IsRandomnessEverAvailable(T *big.Int) (bool, error) {
	return _RandomnessOracle.Contract.IsRandomnessEverAvailable(&_RandomnessOracle.CallOpts, T)
}

// IsRandomnessEverAvailable is a free data retrieval call binding the contract method 0x06055269.
//
// Solidity: function isRandomnessEverAvailable(uint256 T) view returns(bool)
func (_RandomnessOracle *RandomnessOracleCallerSession) IsRandomnessEverAvailable(T *big.Int) (bool, error) {
	return _RandomnessOracle.Contract.IsRandomnessEverAvailable(&_RandomnessOracle.CallOpts, T)
}

// SequencerRandomOracle is a free data retrieval call binding the contract method 0xc55c4262.
//
// Solidity: function sequencerRandomOracle() view returns(address)
func (_RandomnessOracle *RandomnessOracleCaller) SequencerRandomOracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RandomnessOracle.contract.Call(opts, &out, "sequencerRandomOracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SequencerRandomOracle is a free data retrieval call binding the contract method 0xc55c4262.
//
// Solidity: function sequencerRandomOracle() view returns(address)
func (_RandomnessOracle *RandomnessOracleSession) SequencerRandomOracle() (common.Address, error) {
	return _RandomnessOracle.Contract.SequencerRandomOracle(&_RandomnessOracle.CallOpts)
}

// SequencerRandomOracle is a free data retrieval call binding the contract method 0xc55c4262.
//
// Solidity: function sequencerRandomOracle() view returns(address)
func (_RandomnessOracle *RandomnessOracleCallerSession) SequencerRandomOracle() (common.Address, error) {
	return _RandomnessOracle.Contract.SequencerRandomOracle(&_RandomnessOracle.CallOpts)
}

// SimpleGetRandomness is a free data retrieval call binding the contract method 0xaf2556b5.
//
// Solidity: function simpleGetRandomness(uint256 T) view returns(bytes32)
func (_RandomnessOracle *RandomnessOracleCaller) SimpleGetRandomness(opts *bind.CallOpts, T *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _RandomnessOracle.contract.Call(opts, &out, "simpleGetRandomness", T)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SimpleGetRandomness is a free data retrieval call binding the contract method 0xaf2556b5.
//
// Solidity: function simpleGetRandomness(uint256 T) view returns(bytes32)
func (_RandomnessOracle *RandomnessOracleSession) SimpleGetRandomness(T *big.Int) ([32]byte, error) {
	return _RandomnessOracle.Contract.SimpleGetRandomness(&_RandomnessOracle.CallOpts, T)
}

// SimpleGetRandomness is a free data retrieval call binding the contract method 0xaf2556b5.
//
// Solidity: function simpleGetRandomness(uint256 T) view returns(bytes32)
func (_RandomnessOracle *RandomnessOracleCallerSession) SimpleGetRandomness(T *big.Int) ([32]byte, error) {
	return _RandomnessOracle.Contract.SimpleGetRandomness(&_RandomnessOracle.CallOpts, T)
}

// UnsafeGetRandomness is a free data retrieval call binding the contract method 0xea6ae42a.
//
// Solidity: function unsafeGetRandomness(uint256 T) view returns(bytes32)
func (_RandomnessOracle *RandomnessOracleCaller) UnsafeGetRandomness(opts *bind.CallOpts, T *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _RandomnessOracle.contract.Call(opts, &out, "unsafeGetRandomness", T)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// UnsafeGetRandomness is a free data retrieval call binding the contract method 0xea6ae42a.
//
// Solidity: function unsafeGetRandomness(uint256 T) view returns(bytes32)
func (_RandomnessOracle *RandomnessOracleSession) UnsafeGetRandomness(T *big.Int) ([32]byte, error) {
	return _RandomnessOracle.Contract.UnsafeGetRandomness(&_RandomnessOracle.CallOpts, T)
}

// UnsafeGetRandomness is a free data retrieval call binding the contract method 0xea6ae42a.
//
// Solidity: function unsafeGetRandomness(uint256 T) view returns(bytes32)
func (_RandomnessOracle *RandomnessOracleCallerSession) UnsafeGetRandomness(T *big.Int) ([32]byte, error) {
	return _RandomnessOracle.Contract.UnsafeGetRandomness(&_RandomnessOracle.CallOpts, T)
}
