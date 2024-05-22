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

// SequencerRandomOracleMetaData contains all meta data concerning the SequencerRandomOracle contract.
var SequencerRandomOracleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"expectedHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"computedHash\",\"type\":\"bytes32\"}],\"name\":\"InvalidRandomnessReveal\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"T\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"committedBlock\",\"type\":\"uint256\"}],\"name\":\"PrecommitDelayNotPassed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SequencerEntryAlreadyCommitted\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"T\",\"type\":\"uint256\"}],\"name\":\"SequencerRandomnessAlreadyRevealed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"T\",\"type\":\"uint256\"}],\"name\":\"SequencerRandomnessNotCommitted\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"T\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"randomnessHash\",\"type\":\"bytes32\"}],\"name\":\"SequencerRandomnessPosted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"T\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"randomness\",\"type\":\"bytes32\"}],\"name\":\"SequencerRandomnessRevealed\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"sequencerEntries\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"randomnessHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"randomness\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"committed\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"revealed\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"T\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"randomnessHash\",\"type\":\"bytes32\"}],\"name\":\"postRandomnessCommitment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"T\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"randomness\",\"type\":\"bytes32\"}],\"name\":\"revealSequencerRandomness\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"T\",\"type\":\"uint256\"}],\"name\":\"getSequencerRandomness\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// SequencerRandomOracleABI is the input ABI used to generate the binding from.
// Deprecated: Use SequencerRandomOracleMetaData.ABI instead.
var SequencerRandomOracleABI = SequencerRandomOracleMetaData.ABI

// SequencerRandomOracle is an auto generated Go binding around an Ethereum contract.
type SequencerRandomOracle struct {
	SequencerRandomOracleCaller     // Read-only binding to the contract
	SequencerRandomOracleTransactor // Write-only binding to the contract
	SequencerRandomOracleFilterer   // Log filterer for contract events
}

// SequencerRandomOracleCaller is an auto generated read-only Go binding around an Ethereum contract.
type SequencerRandomOracleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SequencerRandomOracleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SequencerRandomOracleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SequencerRandomOracleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SequencerRandomOracleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SequencerRandomOracleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SequencerRandomOracleSession struct {
	Contract     *SequencerRandomOracle // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// SequencerRandomOracleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SequencerRandomOracleCallerSession struct {
	Contract *SequencerRandomOracleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// SequencerRandomOracleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SequencerRandomOracleTransactorSession struct {
	Contract     *SequencerRandomOracleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// SequencerRandomOracleRaw is an auto generated low-level Go binding around an Ethereum contract.
type SequencerRandomOracleRaw struct {
	Contract *SequencerRandomOracle // Generic contract binding to access the raw methods on
}

// SequencerRandomOracleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SequencerRandomOracleCallerRaw struct {
	Contract *SequencerRandomOracleCaller // Generic read-only contract binding to access the raw methods on
}

// SequencerRandomOracleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SequencerRandomOracleTransactorRaw struct {
	Contract *SequencerRandomOracleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSequencerRandomOracle creates a new instance of SequencerRandomOracle, bound to a specific deployed contract.
func NewSequencerRandomOracle(address common.Address, backend bind.ContractBackend) (*SequencerRandomOracle, error) {
	contract, err := bindSequencerRandomOracle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SequencerRandomOracle{SequencerRandomOracleCaller: SequencerRandomOracleCaller{contract: contract}, SequencerRandomOracleTransactor: SequencerRandomOracleTransactor{contract: contract}, SequencerRandomOracleFilterer: SequencerRandomOracleFilterer{contract: contract}}, nil
}

// NewSequencerRandomOracleCaller creates a new read-only instance of SequencerRandomOracle, bound to a specific deployed contract.
func NewSequencerRandomOracleCaller(address common.Address, caller bind.ContractCaller) (*SequencerRandomOracleCaller, error) {
	contract, err := bindSequencerRandomOracle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SequencerRandomOracleCaller{contract: contract}, nil
}

// NewSequencerRandomOracleTransactor creates a new write-only instance of SequencerRandomOracle, bound to a specific deployed contract.
func NewSequencerRandomOracleTransactor(address common.Address, transactor bind.ContractTransactor) (*SequencerRandomOracleTransactor, error) {
	contract, err := bindSequencerRandomOracle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SequencerRandomOracleTransactor{contract: contract}, nil
}

// NewSequencerRandomOracleFilterer creates a new log filterer instance of SequencerRandomOracle, bound to a specific deployed contract.
func NewSequencerRandomOracleFilterer(address common.Address, filterer bind.ContractFilterer) (*SequencerRandomOracleFilterer, error) {
	contract, err := bindSequencerRandomOracle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SequencerRandomOracleFilterer{contract: contract}, nil
}

// bindSequencerRandomOracle binds a generic wrapper to an already deployed contract.
func bindSequencerRandomOracle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SequencerRandomOracleMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SequencerRandomOracle *SequencerRandomOracleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SequencerRandomOracle.Contract.SequencerRandomOracleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SequencerRandomOracle *SequencerRandomOracleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SequencerRandomOracle.Contract.SequencerRandomOracleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SequencerRandomOracle *SequencerRandomOracleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SequencerRandomOracle.Contract.SequencerRandomOracleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SequencerRandomOracle *SequencerRandomOracleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SequencerRandomOracle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SequencerRandomOracle *SequencerRandomOracleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SequencerRandomOracle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SequencerRandomOracle *SequencerRandomOracleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SequencerRandomOracle.Contract.contract.Transact(opts, method, params...)
}

// GetSequencerRandomness is a free data retrieval call binding the contract method 0x700dcc1a.
//
// Solidity: function getSequencerRandomness(uint256 T) view returns(bytes32)
func (_SequencerRandomOracle *SequencerRandomOracleCaller) GetSequencerRandomness(opts *bind.CallOpts, T *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _SequencerRandomOracle.contract.Call(opts, &out, "getSequencerRandomness", T)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetSequencerRandomness is a free data retrieval call binding the contract method 0x700dcc1a.
//
// Solidity: function getSequencerRandomness(uint256 T) view returns(bytes32)
func (_SequencerRandomOracle *SequencerRandomOracleSession) GetSequencerRandomness(T *big.Int) ([32]byte, error) {
	return _SequencerRandomOracle.Contract.GetSequencerRandomness(&_SequencerRandomOracle.CallOpts, T)
}

// GetSequencerRandomness is a free data retrieval call binding the contract method 0x700dcc1a.
//
// Solidity: function getSequencerRandomness(uint256 T) view returns(bytes32)
func (_SequencerRandomOracle *SequencerRandomOracleCallerSession) GetSequencerRandomness(T *big.Int) ([32]byte, error) {
	return _SequencerRandomOracle.Contract.GetSequencerRandomness(&_SequencerRandomOracle.CallOpts, T)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SequencerRandomOracle *SequencerRandomOracleCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SequencerRandomOracle.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SequencerRandomOracle *SequencerRandomOracleSession) Owner() (common.Address, error) {
	return _SequencerRandomOracle.Contract.Owner(&_SequencerRandomOracle.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SequencerRandomOracle *SequencerRandomOracleCallerSession) Owner() (common.Address, error) {
	return _SequencerRandomOracle.Contract.Owner(&_SequencerRandomOracle.CallOpts)
}

// SequencerEntries is a free data retrieval call binding the contract method 0xd055bd04.
//
// Solidity: function sequencerEntries(uint256 ) view returns(bytes32 randomnessHash, bytes32 randomness, uint256 blockNumber, bool committed, bool revealed)
func (_SequencerRandomOracle *SequencerRandomOracleCaller) SequencerEntries(opts *bind.CallOpts, arg0 *big.Int) (struct {
	RandomnessHash [32]byte
	Randomness     [32]byte
	BlockNumber    *big.Int
	Committed      bool
	Revealed       bool
}, error) {
	var out []interface{}
	err := _SequencerRandomOracle.contract.Call(opts, &out, "sequencerEntries", arg0)

	outstruct := new(struct {
		RandomnessHash [32]byte
		Randomness     [32]byte
		BlockNumber    *big.Int
		Committed      bool
		Revealed       bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.RandomnessHash = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.Randomness = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.BlockNumber = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Committed = *abi.ConvertType(out[3], new(bool)).(*bool)
	outstruct.Revealed = *abi.ConvertType(out[4], new(bool)).(*bool)

	return *outstruct, err

}

// SequencerEntries is a free data retrieval call binding the contract method 0xd055bd04.
//
// Solidity: function sequencerEntries(uint256 ) view returns(bytes32 randomnessHash, bytes32 randomness, uint256 blockNumber, bool committed, bool revealed)
func (_SequencerRandomOracle *SequencerRandomOracleSession) SequencerEntries(arg0 *big.Int) (struct {
	RandomnessHash [32]byte
	Randomness     [32]byte
	BlockNumber    *big.Int
	Committed      bool
	Revealed       bool
}, error) {
	return _SequencerRandomOracle.Contract.SequencerEntries(&_SequencerRandomOracle.CallOpts, arg0)
}

// SequencerEntries is a free data retrieval call binding the contract method 0xd055bd04.
//
// Solidity: function sequencerEntries(uint256 ) view returns(bytes32 randomnessHash, bytes32 randomness, uint256 blockNumber, bool committed, bool revealed)
func (_SequencerRandomOracle *SequencerRandomOracleCallerSession) SequencerEntries(arg0 *big.Int) (struct {
	RandomnessHash [32]byte
	Randomness     [32]byte
	BlockNumber    *big.Int
	Committed      bool
	Revealed       bool
}, error) {
	return _SequencerRandomOracle.Contract.SequencerEntries(&_SequencerRandomOracle.CallOpts, arg0)
}

// PostRandomnessCommitment is a paid mutator transaction binding the contract method 0x47bdda48.
//
// Solidity: function postRandomnessCommitment(uint256 T, bytes32 randomnessHash) returns()
func (_SequencerRandomOracle *SequencerRandomOracleTransactor) PostRandomnessCommitment(opts *bind.TransactOpts, T *big.Int, randomnessHash [32]byte) (*types.Transaction, error) {
	return _SequencerRandomOracle.contract.Transact(opts, "postRandomnessCommitment", T, randomnessHash)
}

// PostRandomnessCommitment is a paid mutator transaction binding the contract method 0x47bdda48.
//
// Solidity: function postRandomnessCommitment(uint256 T, bytes32 randomnessHash) returns()
func (_SequencerRandomOracle *SequencerRandomOracleSession) PostRandomnessCommitment(T *big.Int, randomnessHash [32]byte) (*types.Transaction, error) {
	return _SequencerRandomOracle.Contract.PostRandomnessCommitment(&_SequencerRandomOracle.TransactOpts, T, randomnessHash)
}

// PostRandomnessCommitment is a paid mutator transaction binding the contract method 0x47bdda48.
//
// Solidity: function postRandomnessCommitment(uint256 T, bytes32 randomnessHash) returns()
func (_SequencerRandomOracle *SequencerRandomOracleTransactorSession) PostRandomnessCommitment(T *big.Int, randomnessHash [32]byte) (*types.Transaction, error) {
	return _SequencerRandomOracle.Contract.PostRandomnessCommitment(&_SequencerRandomOracle.TransactOpts, T, randomnessHash)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SequencerRandomOracle *SequencerRandomOracleTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SequencerRandomOracle.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SequencerRandomOracle *SequencerRandomOracleSession) RenounceOwnership() (*types.Transaction, error) {
	return _SequencerRandomOracle.Contract.RenounceOwnership(&_SequencerRandomOracle.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SequencerRandomOracle *SequencerRandomOracleTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _SequencerRandomOracle.Contract.RenounceOwnership(&_SequencerRandomOracle.TransactOpts)
}

// RevealSequencerRandomness is a paid mutator transaction binding the contract method 0x90095e9d.
//
// Solidity: function revealSequencerRandomness(uint256 T, bytes32 randomness) returns()
func (_SequencerRandomOracle *SequencerRandomOracleTransactor) RevealSequencerRandomness(opts *bind.TransactOpts, T *big.Int, randomness [32]byte) (*types.Transaction, error) {
	return _SequencerRandomOracle.contract.Transact(opts, "revealSequencerRandomness", T, randomness)
}

// RevealSequencerRandomness is a paid mutator transaction binding the contract method 0x90095e9d.
//
// Solidity: function revealSequencerRandomness(uint256 T, bytes32 randomness) returns()
func (_SequencerRandomOracle *SequencerRandomOracleSession) RevealSequencerRandomness(T *big.Int, randomness [32]byte) (*types.Transaction, error) {
	return _SequencerRandomOracle.Contract.RevealSequencerRandomness(&_SequencerRandomOracle.TransactOpts, T, randomness)
}

// RevealSequencerRandomness is a paid mutator transaction binding the contract method 0x90095e9d.
//
// Solidity: function revealSequencerRandomness(uint256 T, bytes32 randomness) returns()
func (_SequencerRandomOracle *SequencerRandomOracleTransactorSession) RevealSequencerRandomness(T *big.Int, randomness [32]byte) (*types.Transaction, error) {
	return _SequencerRandomOracle.Contract.RevealSequencerRandomness(&_SequencerRandomOracle.TransactOpts, T, randomness)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SequencerRandomOracle *SequencerRandomOracleTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _SequencerRandomOracle.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SequencerRandomOracle *SequencerRandomOracleSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SequencerRandomOracle.Contract.TransferOwnership(&_SequencerRandomOracle.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SequencerRandomOracle *SequencerRandomOracleTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SequencerRandomOracle.Contract.TransferOwnership(&_SequencerRandomOracle.TransactOpts, newOwner)
}

// SequencerRandomOracleOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the SequencerRandomOracle contract.
type SequencerRandomOracleOwnershipTransferredIterator struct {
	Event *SequencerRandomOracleOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SequencerRandomOracleOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SequencerRandomOracleOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SequencerRandomOracleOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SequencerRandomOracleOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SequencerRandomOracleOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SequencerRandomOracleOwnershipTransferred represents a OwnershipTransferred event raised by the SequencerRandomOracle contract.
type SequencerRandomOracleOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SequencerRandomOracle *SequencerRandomOracleFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SequencerRandomOracleOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SequencerRandomOracle.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SequencerRandomOracleOwnershipTransferredIterator{contract: _SequencerRandomOracle.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SequencerRandomOracle *SequencerRandomOracleFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SequencerRandomOracleOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SequencerRandomOracle.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SequencerRandomOracleOwnershipTransferred)
				if err := _SequencerRandomOracle.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SequencerRandomOracle *SequencerRandomOracleFilterer) ParseOwnershipTransferred(log types.Log) (*SequencerRandomOracleOwnershipTransferred, error) {
	event := new(SequencerRandomOracleOwnershipTransferred)
	if err := _SequencerRandomOracle.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SequencerRandomOracleSequencerRandomnessPostedIterator is returned from FilterSequencerRandomnessPosted and is used to iterate over the raw logs and unpacked data for SequencerRandomnessPosted events raised by the SequencerRandomOracle contract.
type SequencerRandomOracleSequencerRandomnessPostedIterator struct {
	Event *SequencerRandomOracleSequencerRandomnessPosted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SequencerRandomOracleSequencerRandomnessPostedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SequencerRandomOracleSequencerRandomnessPosted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SequencerRandomOracleSequencerRandomnessPosted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SequencerRandomOracleSequencerRandomnessPostedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SequencerRandomOracleSequencerRandomnessPostedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SequencerRandomOracleSequencerRandomnessPosted represents a SequencerRandomnessPosted event raised by the SequencerRandomOracle contract.
type SequencerRandomOracleSequencerRandomnessPosted struct {
	T              *big.Int
	RandomnessHash [32]byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterSequencerRandomnessPosted is a free log retrieval operation binding the contract event 0x7ee26e0e713ba7f9bd111f62267c6cac5984732e39b5a0350ddb50934dd18775.
//
// Solidity: event SequencerRandomnessPosted(uint256 indexed T, bytes32 randomnessHash)
func (_SequencerRandomOracle *SequencerRandomOracleFilterer) FilterSequencerRandomnessPosted(opts *bind.FilterOpts, T []*big.Int) (*SequencerRandomOracleSequencerRandomnessPostedIterator, error) {

	var TRule []interface{}
	for _, TItem := range T {
		TRule = append(TRule, TItem)
	}

	logs, sub, err := _SequencerRandomOracle.contract.FilterLogs(opts, "SequencerRandomnessPosted", TRule)
	if err != nil {
		return nil, err
	}
	return &SequencerRandomOracleSequencerRandomnessPostedIterator{contract: _SequencerRandomOracle.contract, event: "SequencerRandomnessPosted", logs: logs, sub: sub}, nil
}

// WatchSequencerRandomnessPosted is a free log subscription operation binding the contract event 0x7ee26e0e713ba7f9bd111f62267c6cac5984732e39b5a0350ddb50934dd18775.
//
// Solidity: event SequencerRandomnessPosted(uint256 indexed T, bytes32 randomnessHash)
func (_SequencerRandomOracle *SequencerRandomOracleFilterer) WatchSequencerRandomnessPosted(opts *bind.WatchOpts, sink chan<- *SequencerRandomOracleSequencerRandomnessPosted, T []*big.Int) (event.Subscription, error) {

	var TRule []interface{}
	for _, TItem := range T {
		TRule = append(TRule, TItem)
	}

	logs, sub, err := _SequencerRandomOracle.contract.WatchLogs(opts, "SequencerRandomnessPosted", TRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SequencerRandomOracleSequencerRandomnessPosted)
				if err := _SequencerRandomOracle.contract.UnpackLog(event, "SequencerRandomnessPosted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSequencerRandomnessPosted is a log parse operation binding the contract event 0x7ee26e0e713ba7f9bd111f62267c6cac5984732e39b5a0350ddb50934dd18775.
//
// Solidity: event SequencerRandomnessPosted(uint256 indexed T, bytes32 randomnessHash)
func (_SequencerRandomOracle *SequencerRandomOracleFilterer) ParseSequencerRandomnessPosted(log types.Log) (*SequencerRandomOracleSequencerRandomnessPosted, error) {
	event := new(SequencerRandomOracleSequencerRandomnessPosted)
	if err := _SequencerRandomOracle.contract.UnpackLog(event, "SequencerRandomnessPosted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SequencerRandomOracleSequencerRandomnessRevealedIterator is returned from FilterSequencerRandomnessRevealed and is used to iterate over the raw logs and unpacked data for SequencerRandomnessRevealed events raised by the SequencerRandomOracle contract.
type SequencerRandomOracleSequencerRandomnessRevealedIterator struct {
	Event *SequencerRandomOracleSequencerRandomnessRevealed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SequencerRandomOracleSequencerRandomnessRevealedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SequencerRandomOracleSequencerRandomnessRevealed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SequencerRandomOracleSequencerRandomnessRevealed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SequencerRandomOracleSequencerRandomnessRevealedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SequencerRandomOracleSequencerRandomnessRevealedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SequencerRandomOracleSequencerRandomnessRevealed represents a SequencerRandomnessRevealed event raised by the SequencerRandomOracle contract.
type SequencerRandomOracleSequencerRandomnessRevealed struct {
	T          *big.Int
	Randomness [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSequencerRandomnessRevealed is a free log retrieval operation binding the contract event 0xfaa891f32720f5e38fa109ba0262096d6555306fd75a9f6f0914d9f47b75128c.
//
// Solidity: event SequencerRandomnessRevealed(uint256 indexed T, bytes32 randomness)
func (_SequencerRandomOracle *SequencerRandomOracleFilterer) FilterSequencerRandomnessRevealed(opts *bind.FilterOpts, T []*big.Int) (*SequencerRandomOracleSequencerRandomnessRevealedIterator, error) {

	var TRule []interface{}
	for _, TItem := range T {
		TRule = append(TRule, TItem)
	}

	logs, sub, err := _SequencerRandomOracle.contract.FilterLogs(opts, "SequencerRandomnessRevealed", TRule)
	if err != nil {
		return nil, err
	}
	return &SequencerRandomOracleSequencerRandomnessRevealedIterator{contract: _SequencerRandomOracle.contract, event: "SequencerRandomnessRevealed", logs: logs, sub: sub}, nil
}

// WatchSequencerRandomnessRevealed is a free log subscription operation binding the contract event 0xfaa891f32720f5e38fa109ba0262096d6555306fd75a9f6f0914d9f47b75128c.
//
// Solidity: event SequencerRandomnessRevealed(uint256 indexed T, bytes32 randomness)
func (_SequencerRandomOracle *SequencerRandomOracleFilterer) WatchSequencerRandomnessRevealed(opts *bind.WatchOpts, sink chan<- *SequencerRandomOracleSequencerRandomnessRevealed, T []*big.Int) (event.Subscription, error) {

	var TRule []interface{}
	for _, TItem := range T {
		TRule = append(TRule, TItem)
	}

	logs, sub, err := _SequencerRandomOracle.contract.WatchLogs(opts, "SequencerRandomnessRevealed", TRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SequencerRandomOracleSequencerRandomnessRevealed)
				if err := _SequencerRandomOracle.contract.UnpackLog(event, "SequencerRandomnessRevealed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSequencerRandomnessRevealed is a log parse operation binding the contract event 0xfaa891f32720f5e38fa109ba0262096d6555306fd75a9f6f0914d9f47b75128c.
//
// Solidity: event SequencerRandomnessRevealed(uint256 indexed T, bytes32 randomness)
func (_SequencerRandomOracle *SequencerRandomOracleFilterer) ParseSequencerRandomnessRevealed(log types.Log) (*SequencerRandomOracleSequencerRandomnessRevealed, error) {
	event := new(SequencerRandomOracleSequencerRandomnessRevealed)
	if err := _SequencerRandomOracle.contract.UnpackLog(event, "SequencerRandomnessRevealed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
