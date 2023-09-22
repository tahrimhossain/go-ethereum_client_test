// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package new

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

// NewMetaData contains all meta data concerning the New contract.
var NewMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"val\",\"type\":\"uint256\"}],\"name\":\"DataStored\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"compute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// NewABI is the input ABI used to generate the binding from.
// Deprecated: Use NewMetaData.ABI instead.
var NewABI = NewMetaData.ABI

// New is an auto generated Go binding around an Ethereum contract.
type New struct {
	NewCaller     // Read-only binding to the contract
	NewTransactor // Write-only binding to the contract
	NewFilterer   // Log filterer for contract events
}

// NewCaller is an auto generated read-only Go binding around an Ethereum contract.
type NewCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NewTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NewTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NewFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NewFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NewSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NewSession struct {
	Contract     *New              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NewCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NewCallerSession struct {
	Contract *NewCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// NewTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NewTransactorSession struct {
	Contract     *NewTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NewRaw is an auto generated low-level Go binding around an Ethereum contract.
type NewRaw struct {
	Contract *New // Generic contract binding to access the raw methods on
}

// NewCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NewCallerRaw struct {
	Contract *NewCaller // Generic read-only contract binding to access the raw methods on
}

// NewTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NewTransactorRaw struct {
	Contract *NewTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNew creates a new instance of New, bound to a specific deployed contract.
func NewNew(address common.Address, backend bind.ContractBackend) (*New, error) {
	contract, err := bindNew(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &New{NewCaller: NewCaller{contract: contract}, NewTransactor: NewTransactor{contract: contract}, NewFilterer: NewFilterer{contract: contract}}, nil
}

// NewNewCaller creates a new read-only instance of New, bound to a specific deployed contract.
func NewNewCaller(address common.Address, caller bind.ContractCaller) (*NewCaller, error) {
	contract, err := bindNew(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NewCaller{contract: contract}, nil
}

// NewNewTransactor creates a new write-only instance of New, bound to a specific deployed contract.
func NewNewTransactor(address common.Address, transactor bind.ContractTransactor) (*NewTransactor, error) {
	contract, err := bindNew(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NewTransactor{contract: contract}, nil
}

// NewNewFilterer creates a new log filterer instance of New, bound to a specific deployed contract.
func NewNewFilterer(address common.Address, filterer bind.ContractFilterer) (*NewFilterer, error) {
	contract, err := bindNew(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NewFilterer{contract: contract}, nil
}

// bindNew binds a generic wrapper to an already deployed contract.
func bindNew(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := NewMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_New *NewRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _New.Contract.NewCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_New *NewRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _New.Contract.NewTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_New *NewRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _New.Contract.NewTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_New *NewCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _New.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_New *NewTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _New.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_New *NewTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _New.Contract.contract.Transact(opts, method, params...)
}

// Compute is a paid mutator transaction binding the contract method 0x1a43c338.
//
// Solidity: function compute() returns()
func (_New *NewTransactor) Compute(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _New.contract.Transact(opts, "compute")
}

// Compute is a paid mutator transaction binding the contract method 0x1a43c338.
//
// Solidity: function compute() returns()
func (_New *NewSession) Compute() (*types.Transaction, error) {
	return _New.Contract.Compute(&_New.TransactOpts)
}

// Compute is a paid mutator transaction binding the contract method 0x1a43c338.
//
// Solidity: function compute() returns()
func (_New *NewTransactorSession) Compute() (*types.Transaction, error) {
	return _New.Contract.Compute(&_New.TransactOpts)
}

// NewDataStoredIterator is returned from FilterDataStored and is used to iterate over the raw logs and unpacked data for DataStored events raised by the New contract.
type NewDataStoredIterator struct {
	Event *NewDataStored // Event containing the contract specifics and raw log

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
func (it *NewDataStoredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NewDataStored)
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
		it.Event = new(NewDataStored)
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
func (it *NewDataStoredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NewDataStoredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NewDataStored represents a DataStored event raised by the New contract.
type NewDataStored struct {
	Val *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDataStored is a free log retrieval operation binding the contract event 0x9455957c3b77d1d4ed071e2b469dd77e37fc5dfd3b4d44dc8a997cc97c7b3d49.
//
// Solidity: event DataStored(uint256 val)
func (_New *NewFilterer) FilterDataStored(opts *bind.FilterOpts) (*NewDataStoredIterator, error) {

	logs, sub, err := _New.contract.FilterLogs(opts, "DataStored")
	if err != nil {
		return nil, err
	}
	return &NewDataStoredIterator{contract: _New.contract, event: "DataStored", logs: logs, sub: sub}, nil
}

// WatchDataStored is a free log subscription operation binding the contract event 0x9455957c3b77d1d4ed071e2b469dd77e37fc5dfd3b4d44dc8a997cc97c7b3d49.
//
// Solidity: event DataStored(uint256 val)
func (_New *NewFilterer) WatchDataStored(opts *bind.WatchOpts, sink chan<- *NewDataStored) (event.Subscription, error) {

	logs, sub, err := _New.contract.WatchLogs(opts, "DataStored")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NewDataStored)
				if err := _New.contract.UnpackLog(event, "DataStored", log); err != nil {
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

// ParseDataStored is a log parse operation binding the contract event 0x9455957c3b77d1d4ed071e2b469dd77e37fc5dfd3b4d44dc8a997cc97c7b3d49.
//
// Solidity: event DataStored(uint256 val)
func (_New *NewFilterer) ParseDataStored(log types.Log) (*NewDataStored, error) {
	event := new(NewDataStored)
	if err := _New.contract.UnpackLog(event, "DataStored", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
