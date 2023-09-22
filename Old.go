// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package old

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

// OldMetaData contains all meta data concerning the Old contract.
var OldMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"val\",\"type\":\"uint256\"}],\"name\":\"DataStored\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"compute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getResult\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b5060055f8190555060036001819055505f600281905550610180806100335f395ff3fe608060405234801561000f575f80fd5b5060043610610034575f3560e01c80631a43c33814610038578063de29278914610042575b5f80fd5b610040610060565b005b61004a6100b0565b60405161005791906100d1565b60405180910390f35b6001545f5461006f9190610117565b6002819055507f9455957c3b77d1d4ed071e2b469dd77e37fc5dfd3b4d44dc8a997cc97c7b3d496002546040516100a691906100d1565b60405180910390a1565b5f600254905090565b5f819050919050565b6100cb816100b9565b82525050565b5f6020820190506100e45f8301846100c2565b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f610121826100b9565b915061012c836100b9565b9250828201905080821115610144576101436100ea565b5b9291505056fea26469706673582212200344822f69fc14a4ca4bf0dd443f9bb99d04b89035ebee1184320a8fbeba1f6164736f6c63430008140033",
}

// OldABI is the input ABI used to generate the binding from.
// Deprecated: Use OldMetaData.ABI instead.
var OldABI = OldMetaData.ABI

// OldBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use OldMetaData.Bin instead.
var OldBin = OldMetaData.Bin

// DeployOld deploys a new Ethereum contract, binding an instance of Old to it.
func DeployOld(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Old, error) {
	parsed, err := OldMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OldBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Old{OldCaller: OldCaller{contract: contract}, OldTransactor: OldTransactor{contract: contract}, OldFilterer: OldFilterer{contract: contract}}, nil
}

// Old is an auto generated Go binding around an Ethereum contract.
type Old struct {
	OldCaller     // Read-only binding to the contract
	OldTransactor // Write-only binding to the contract
	OldFilterer   // Log filterer for contract events
}

// OldCaller is an auto generated read-only Go binding around an Ethereum contract.
type OldCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OldTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OldTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OldFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OldFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OldSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OldSession struct {
	Contract     *Old              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OldCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OldCallerSession struct {
	Contract *OldCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// OldTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OldTransactorSession struct {
	Contract     *OldTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OldRaw is an auto generated low-level Go binding around an Ethereum contract.
type OldRaw struct {
	Contract *Old // Generic contract binding to access the raw methods on
}

// OldCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OldCallerRaw struct {
	Contract *OldCaller // Generic read-only contract binding to access the raw methods on
}

// OldTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OldTransactorRaw struct {
	Contract *OldTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOld creates a new instance of Old, bound to a specific deployed contract.
func NewOld(address common.Address, backend bind.ContractBackend) (*Old, error) {
	contract, err := bindOld(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Old{OldCaller: OldCaller{contract: contract}, OldTransactor: OldTransactor{contract: contract}, OldFilterer: OldFilterer{contract: contract}}, nil
}

// NewOldCaller creates a new read-only instance of Old, bound to a specific deployed contract.
func NewOldCaller(address common.Address, caller bind.ContractCaller) (*OldCaller, error) {
	contract, err := bindOld(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OldCaller{contract: contract}, nil
}

// NewOldTransactor creates a new write-only instance of Old, bound to a specific deployed contract.
func NewOldTransactor(address common.Address, transactor bind.ContractTransactor) (*OldTransactor, error) {
	contract, err := bindOld(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OldTransactor{contract: contract}, nil
}

// NewOldFilterer creates a new log filterer instance of Old, bound to a specific deployed contract.
func NewOldFilterer(address common.Address, filterer bind.ContractFilterer) (*OldFilterer, error) {
	contract, err := bindOld(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OldFilterer{contract: contract}, nil
}

// bindOld binds a generic wrapper to an already deployed contract.
func bindOld(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OldMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Old *OldRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Old.Contract.OldCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Old *OldRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Old.Contract.OldTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Old *OldRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Old.Contract.OldTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Old *OldCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Old.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Old *OldTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Old.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Old *OldTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Old.Contract.contract.Transact(opts, method, params...)
}

// GetResult is a free data retrieval call binding the contract method 0xde292789.
//
// Solidity: function getResult() view returns(uint256)
func (_Old *OldCaller) GetResult(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Old.contract.Call(opts, &out, "getResult")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetResult is a free data retrieval call binding the contract method 0xde292789.
//
// Solidity: function getResult() view returns(uint256)
func (_Old *OldSession) GetResult() (*big.Int, error) {
	return _Old.Contract.GetResult(&_Old.CallOpts)
}

// GetResult is a free data retrieval call binding the contract method 0xde292789.
//
// Solidity: function getResult() view returns(uint256)
func (_Old *OldCallerSession) GetResult() (*big.Int, error) {
	return _Old.Contract.GetResult(&_Old.CallOpts)
}

// Compute is a paid mutator transaction binding the contract method 0x1a43c338.
//
// Solidity: function compute() returns()
func (_Old *OldTransactor) Compute(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Old.contract.Transact(opts, "compute")
}

// Compute is a paid mutator transaction binding the contract method 0x1a43c338.
//
// Solidity: function compute() returns()
func (_Old *OldSession) Compute() (*types.Transaction, error) {
	return _Old.Contract.Compute(&_Old.TransactOpts)
}

// Compute is a paid mutator transaction binding the contract method 0x1a43c338.
//
// Solidity: function compute() returns()
func (_Old *OldTransactorSession) Compute() (*types.Transaction, error) {
	return _Old.Contract.Compute(&_Old.TransactOpts)
}

// OldDataStoredIterator is returned from FilterDataStored and is used to iterate over the raw logs and unpacked data for DataStored events raised by the Old contract.
type OldDataStoredIterator struct {
	Event *OldDataStored // Event containing the contract specifics and raw log

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
func (it *OldDataStoredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OldDataStored)
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
		it.Event = new(OldDataStored)
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
func (it *OldDataStoredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OldDataStoredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OldDataStored represents a DataStored event raised by the Old contract.
type OldDataStored struct {
	Val *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDataStored is a free log retrieval operation binding the contract event 0x9455957c3b77d1d4ed071e2b469dd77e37fc5dfd3b4d44dc8a997cc97c7b3d49.
//
// Solidity: event DataStored(uint256 val)
func (_Old *OldFilterer) FilterDataStored(opts *bind.FilterOpts) (*OldDataStoredIterator, error) {

	logs, sub, err := _Old.contract.FilterLogs(opts, "DataStored")
	if err != nil {
		return nil, err
	}
	return &OldDataStoredIterator{contract: _Old.contract, event: "DataStored", logs: logs, sub: sub}, nil
}

// WatchDataStored is a free log subscription operation binding the contract event 0x9455957c3b77d1d4ed071e2b469dd77e37fc5dfd3b4d44dc8a997cc97c7b3d49.
//
// Solidity: event DataStored(uint256 val)
func (_Old *OldFilterer) WatchDataStored(opts *bind.WatchOpts, sink chan<- *OldDataStored) (event.Subscription, error) {

	logs, sub, err := _Old.contract.WatchLogs(opts, "DataStored")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OldDataStored)
				if err := _Old.contract.UnpackLog(event, "DataStored", log); err != nil {
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
func (_Old *OldFilterer) ParseDataStored(log types.Log) (*OldDataStored, error) {
	event := new(OldDataStored)
	if err := _Old.contract.UnpackLog(event, "DataStored", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
