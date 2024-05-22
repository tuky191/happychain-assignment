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

// DrandOracleMetaData contains all meta data concerning the DrandOracle contract.
var DrandOracleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"T\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"randomness\",\"type\":\"bytes32\"}],\"name\":\"DrandUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"drandEntries\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"randomness\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"filled\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"T\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"randomness\",\"type\":\"bytes32\"}],\"name\":\"postDrandRandomness\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"T\",\"type\":\"uint256\"}],\"name\":\"getDrand\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"T\",\"type\":\"uint256\"}],\"name\":\"isDrandAvailable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"T\",\"type\":\"uint256\"}],\"name\":\"hasUpdatePeriodExpired\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// DrandOracleABI is the input ABI used to generate the binding from.
// Deprecated: Use DrandOracleMetaData.ABI instead.
var DrandOracleABI = DrandOracleMetaData.ABI

// DrandOracle is an auto generated Go binding around an Ethereum contract.
type DrandOracle struct {
	DrandOracleCaller     // Read-only binding to the contract
	DrandOracleTransactor // Write-only binding to the contract
	DrandOracleFilterer   // Log filterer for contract events
}

// DrandOracleCaller is an auto generated read-only Go binding around an Ethereum contract.
type DrandOracleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DrandOracleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DrandOracleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DrandOracleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DrandOracleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DrandOracleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DrandOracleSession struct {
	Contract     *DrandOracle      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DrandOracleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DrandOracleCallerSession struct {
	Contract *DrandOracleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// DrandOracleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DrandOracleTransactorSession struct {
	Contract     *DrandOracleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// DrandOracleRaw is an auto generated low-level Go binding around an Ethereum contract.
type DrandOracleRaw struct {
	Contract *DrandOracle // Generic contract binding to access the raw methods on
}

// DrandOracleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DrandOracleCallerRaw struct {
	Contract *DrandOracleCaller // Generic read-only contract binding to access the raw methods on
}

// DrandOracleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DrandOracleTransactorRaw struct {
	Contract *DrandOracleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDrandOracle creates a new instance of DrandOracle, bound to a specific deployed contract.
func NewDrandOracle(address common.Address, backend bind.ContractBackend) (*DrandOracle, error) {
	contract, err := bindDrandOracle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DrandOracle{DrandOracleCaller: DrandOracleCaller{contract: contract}, DrandOracleTransactor: DrandOracleTransactor{contract: contract}, DrandOracleFilterer: DrandOracleFilterer{contract: contract}}, nil
}

// NewDrandOracleCaller creates a new read-only instance of DrandOracle, bound to a specific deployed contract.
func NewDrandOracleCaller(address common.Address, caller bind.ContractCaller) (*DrandOracleCaller, error) {
	contract, err := bindDrandOracle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DrandOracleCaller{contract: contract}, nil
}

// NewDrandOracleTransactor creates a new write-only instance of DrandOracle, bound to a specific deployed contract.
func NewDrandOracleTransactor(address common.Address, transactor bind.ContractTransactor) (*DrandOracleTransactor, error) {
	contract, err := bindDrandOracle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DrandOracleTransactor{contract: contract}, nil
}

// NewDrandOracleFilterer creates a new log filterer instance of DrandOracle, bound to a specific deployed contract.
func NewDrandOracleFilterer(address common.Address, filterer bind.ContractFilterer) (*DrandOracleFilterer, error) {
	contract, err := bindDrandOracle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DrandOracleFilterer{contract: contract}, nil
}

// bindDrandOracle binds a generic wrapper to an already deployed contract.
func bindDrandOracle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DrandOracleMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DrandOracle *DrandOracleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DrandOracle.Contract.DrandOracleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DrandOracle *DrandOracleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DrandOracle.Contract.DrandOracleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DrandOracle *DrandOracleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DrandOracle.Contract.DrandOracleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DrandOracle *DrandOracleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DrandOracle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DrandOracle *DrandOracleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DrandOracle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DrandOracle *DrandOracleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DrandOracle.Contract.contract.Transact(opts, method, params...)
}

// DrandEntries is a free data retrieval call binding the contract method 0x75a54743.
//
// Solidity: function drandEntries(uint256 ) view returns(bytes32 randomness, uint256 timestamp, bool filled)
func (_DrandOracle *DrandOracleCaller) DrandEntries(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Randomness [32]byte
	Timestamp  *big.Int
	Filled     bool
}, error) {
	var out []interface{}
	err := _DrandOracle.contract.Call(opts, &out, "drandEntries", arg0)

	outstruct := new(struct {
		Randomness [32]byte
		Timestamp  *big.Int
		Filled     bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Randomness = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.Timestamp = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Filled = *abi.ConvertType(out[2], new(bool)).(*bool)

	return *outstruct, err

}

// DrandEntries is a free data retrieval call binding the contract method 0x75a54743.
//
// Solidity: function drandEntries(uint256 ) view returns(bytes32 randomness, uint256 timestamp, bool filled)
func (_DrandOracle *DrandOracleSession) DrandEntries(arg0 *big.Int) (struct {
	Randomness [32]byte
	Timestamp  *big.Int
	Filled     bool
}, error) {
	return _DrandOracle.Contract.DrandEntries(&_DrandOracle.CallOpts, arg0)
}

// DrandEntries is a free data retrieval call binding the contract method 0x75a54743.
//
// Solidity: function drandEntries(uint256 ) view returns(bytes32 randomness, uint256 timestamp, bool filled)
func (_DrandOracle *DrandOracleCallerSession) DrandEntries(arg0 *big.Int) (struct {
	Randomness [32]byte
	Timestamp  *big.Int
	Filled     bool
}, error) {
	return _DrandOracle.Contract.DrandEntries(&_DrandOracle.CallOpts, arg0)
}

// GetDrand is a free data retrieval call binding the contract method 0x74dc0fc3.
//
// Solidity: function getDrand(uint256 T) view returns(bytes32)
func (_DrandOracle *DrandOracleCaller) GetDrand(opts *bind.CallOpts, T *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _DrandOracle.contract.Call(opts, &out, "getDrand", T)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetDrand is a free data retrieval call binding the contract method 0x74dc0fc3.
//
// Solidity: function getDrand(uint256 T) view returns(bytes32)
func (_DrandOracle *DrandOracleSession) GetDrand(T *big.Int) ([32]byte, error) {
	return _DrandOracle.Contract.GetDrand(&_DrandOracle.CallOpts, T)
}

// GetDrand is a free data retrieval call binding the contract method 0x74dc0fc3.
//
// Solidity: function getDrand(uint256 T) view returns(bytes32)
func (_DrandOracle *DrandOracleCallerSession) GetDrand(T *big.Int) ([32]byte, error) {
	return _DrandOracle.Contract.GetDrand(&_DrandOracle.CallOpts, T)
}

// HasUpdatePeriodExpired is a free data retrieval call binding the contract method 0xe0d81de0.
//
// Solidity: function hasUpdatePeriodExpired(uint256 T) view returns(bool)
func (_DrandOracle *DrandOracleCaller) HasUpdatePeriodExpired(opts *bind.CallOpts, T *big.Int) (bool, error) {
	var out []interface{}
	err := _DrandOracle.contract.Call(opts, &out, "hasUpdatePeriodExpired", T)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasUpdatePeriodExpired is a free data retrieval call binding the contract method 0xe0d81de0.
//
// Solidity: function hasUpdatePeriodExpired(uint256 T) view returns(bool)
func (_DrandOracle *DrandOracleSession) HasUpdatePeriodExpired(T *big.Int) (bool, error) {
	return _DrandOracle.Contract.HasUpdatePeriodExpired(&_DrandOracle.CallOpts, T)
}

// HasUpdatePeriodExpired is a free data retrieval call binding the contract method 0xe0d81de0.
//
// Solidity: function hasUpdatePeriodExpired(uint256 T) view returns(bool)
func (_DrandOracle *DrandOracleCallerSession) HasUpdatePeriodExpired(T *big.Int) (bool, error) {
	return _DrandOracle.Contract.HasUpdatePeriodExpired(&_DrandOracle.CallOpts, T)
}

// IsDrandAvailable is a free data retrieval call binding the contract method 0x079ea2ed.
//
// Solidity: function isDrandAvailable(uint256 T) view returns(bool)
func (_DrandOracle *DrandOracleCaller) IsDrandAvailable(opts *bind.CallOpts, T *big.Int) (bool, error) {
	var out []interface{}
	err := _DrandOracle.contract.Call(opts, &out, "isDrandAvailable", T)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsDrandAvailable is a free data retrieval call binding the contract method 0x079ea2ed.
//
// Solidity: function isDrandAvailable(uint256 T) view returns(bool)
func (_DrandOracle *DrandOracleSession) IsDrandAvailable(T *big.Int) (bool, error) {
	return _DrandOracle.Contract.IsDrandAvailable(&_DrandOracle.CallOpts, T)
}

// IsDrandAvailable is a free data retrieval call binding the contract method 0x079ea2ed.
//
// Solidity: function isDrandAvailable(uint256 T) view returns(bool)
func (_DrandOracle *DrandOracleCallerSession) IsDrandAvailable(T *big.Int) (bool, error) {
	return _DrandOracle.Contract.IsDrandAvailable(&_DrandOracle.CallOpts, T)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DrandOracle *DrandOracleCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DrandOracle.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DrandOracle *DrandOracleSession) Owner() (common.Address, error) {
	return _DrandOracle.Contract.Owner(&_DrandOracle.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DrandOracle *DrandOracleCallerSession) Owner() (common.Address, error) {
	return _DrandOracle.Contract.Owner(&_DrandOracle.CallOpts)
}

// PostDrandRandomness is a paid mutator transaction binding the contract method 0x1482882a.
//
// Solidity: function postDrandRandomness(uint256 T, bytes32 randomness) returns()
func (_DrandOracle *DrandOracleTransactor) PostDrandRandomness(opts *bind.TransactOpts, T *big.Int, randomness [32]byte) (*types.Transaction, error) {
	return _DrandOracle.contract.Transact(opts, "postDrandRandomness", T, randomness)
}

// PostDrandRandomness is a paid mutator transaction binding the contract method 0x1482882a.
//
// Solidity: function postDrandRandomness(uint256 T, bytes32 randomness) returns()
func (_DrandOracle *DrandOracleSession) PostDrandRandomness(T *big.Int, randomness [32]byte) (*types.Transaction, error) {
	return _DrandOracle.Contract.PostDrandRandomness(&_DrandOracle.TransactOpts, T, randomness)
}

// PostDrandRandomness is a paid mutator transaction binding the contract method 0x1482882a.
//
// Solidity: function postDrandRandomness(uint256 T, bytes32 randomness) returns()
func (_DrandOracle *DrandOracleTransactorSession) PostDrandRandomness(T *big.Int, randomness [32]byte) (*types.Transaction, error) {
	return _DrandOracle.Contract.PostDrandRandomness(&_DrandOracle.TransactOpts, T, randomness)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DrandOracle *DrandOracleTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DrandOracle.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DrandOracle *DrandOracleSession) RenounceOwnership() (*types.Transaction, error) {
	return _DrandOracle.Contract.RenounceOwnership(&_DrandOracle.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DrandOracle *DrandOracleTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _DrandOracle.Contract.RenounceOwnership(&_DrandOracle.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DrandOracle *DrandOracleTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _DrandOracle.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DrandOracle *DrandOracleSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DrandOracle.Contract.TransferOwnership(&_DrandOracle.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DrandOracle *DrandOracleTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DrandOracle.Contract.TransferOwnership(&_DrandOracle.TransactOpts, newOwner)
}

// DrandOracleDrandUpdatedIterator is returned from FilterDrandUpdated and is used to iterate over the raw logs and unpacked data for DrandUpdated events raised by the DrandOracle contract.
type DrandOracleDrandUpdatedIterator struct {
	Event *DrandOracleDrandUpdated // Event containing the contract specifics and raw log

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
func (it *DrandOracleDrandUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DrandOracleDrandUpdated)
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
		it.Event = new(DrandOracleDrandUpdated)
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
func (it *DrandOracleDrandUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DrandOracleDrandUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DrandOracleDrandUpdated represents a DrandUpdated event raised by the DrandOracle contract.
type DrandOracleDrandUpdated struct {
	T          *big.Int
	Randomness [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDrandUpdated is a free log retrieval operation binding the contract event 0x08fcfa9d5e57e534cb91c0d2628efd4da84db6c55fc26fca7b5fa7aca4102033.
//
// Solidity: event DrandUpdated(uint256 indexed T, bytes32 randomness)
func (_DrandOracle *DrandOracleFilterer) FilterDrandUpdated(opts *bind.FilterOpts, T []*big.Int) (*DrandOracleDrandUpdatedIterator, error) {

	var TRule []interface{}
	for _, TItem := range T {
		TRule = append(TRule, TItem)
	}

	logs, sub, err := _DrandOracle.contract.FilterLogs(opts, "DrandUpdated", TRule)
	if err != nil {
		return nil, err
	}
	return &DrandOracleDrandUpdatedIterator{contract: _DrandOracle.contract, event: "DrandUpdated", logs: logs, sub: sub}, nil
}

// WatchDrandUpdated is a free log subscription operation binding the contract event 0x08fcfa9d5e57e534cb91c0d2628efd4da84db6c55fc26fca7b5fa7aca4102033.
//
// Solidity: event DrandUpdated(uint256 indexed T, bytes32 randomness)
func (_DrandOracle *DrandOracleFilterer) WatchDrandUpdated(opts *bind.WatchOpts, sink chan<- *DrandOracleDrandUpdated, T []*big.Int) (event.Subscription, error) {

	var TRule []interface{}
	for _, TItem := range T {
		TRule = append(TRule, TItem)
	}

	logs, sub, err := _DrandOracle.contract.WatchLogs(opts, "DrandUpdated", TRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DrandOracleDrandUpdated)
				if err := _DrandOracle.contract.UnpackLog(event, "DrandUpdated", log); err != nil {
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

// ParseDrandUpdated is a log parse operation binding the contract event 0x08fcfa9d5e57e534cb91c0d2628efd4da84db6c55fc26fca7b5fa7aca4102033.
//
// Solidity: event DrandUpdated(uint256 indexed T, bytes32 randomness)
func (_DrandOracle *DrandOracleFilterer) ParseDrandUpdated(log types.Log) (*DrandOracleDrandUpdated, error) {
	event := new(DrandOracleDrandUpdated)
	if err := _DrandOracle.contract.UnpackLog(event, "DrandUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DrandOracleOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the DrandOracle contract.
type DrandOracleOwnershipTransferredIterator struct {
	Event *DrandOracleOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *DrandOracleOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DrandOracleOwnershipTransferred)
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
		it.Event = new(DrandOracleOwnershipTransferred)
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
func (it *DrandOracleOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DrandOracleOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DrandOracleOwnershipTransferred represents a OwnershipTransferred event raised by the DrandOracle contract.
type DrandOracleOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DrandOracle *DrandOracleFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*DrandOracleOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DrandOracle.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DrandOracleOwnershipTransferredIterator{contract: _DrandOracle.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DrandOracle *DrandOracleFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *DrandOracleOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DrandOracle.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DrandOracleOwnershipTransferred)
				if err := _DrandOracle.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_DrandOracle *DrandOracleFilterer) ParseOwnershipTransferred(log types.Log) (*DrandOracleOwnershipTransferred, error) {
	event := new(DrandOracleOwnershipTransferred)
	if err := _DrandOracle.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
