// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package gateway

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/scroll-tech/go-ethereum"
	"github.com/scroll-tech/go-ethereum/accounts/abi"
	"github.com/scroll-tech/go-ethereum/accounts/abi/bind"
	"github.com/scroll-tech/go-ethereum/common"
	"github.com/scroll-tech/go-ethereum/core/types"
	"github.com/scroll-tech/go-ethereum/event"
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
)

// IL1ETHGatewayMetaData contains all meta data concerning the IL1ETHGateway contract.
var IL1ETHGatewayMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"DepositETH\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"FinalizeWithdrawETH\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RefundETH\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"}],\"name\":\"depositETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"}],\"name\":\"depositETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"}],\"name\":\"depositETHAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"finalizeWithdrawETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"ce0b63ce": "depositETH(address,uint256,uint256)",
		"9f8420b3": "depositETH(uint256,uint256)",
		"aac476f8": "depositETHAndCall(address,uint256,bytes,uint256)",
		"8eaac8a3": "finalizeWithdrawETH(address,address,uint256,bytes)",
	},
}

// IL1ETHGatewayABI is the input ABI used to generate the binding from.
// Deprecated: Use IL1ETHGatewayMetaData.ABI instead.
var IL1ETHGatewayABI = IL1ETHGatewayMetaData.ABI

// Deprecated: Use IL1ETHGatewayMetaData.Sigs instead.
// IL1ETHGatewayFuncSigs maps the 4-byte function signature to its string representation.
var IL1ETHGatewayFuncSigs = IL1ETHGatewayMetaData.Sigs

// IL1ETHGateway is an auto generated Go binding around an Ethereum contract.
type IL1ETHGateway struct {
	IL1ETHGatewayCaller     // Read-only binding to the contract
	IL1ETHGatewayTransactor // Write-only binding to the contract
	IL1ETHGatewayFilterer   // Log filterer for contract events
}

// IL1ETHGatewayCaller is an auto generated read-only Go binding around an Ethereum contract.
type IL1ETHGatewayCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IL1ETHGatewayTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IL1ETHGatewayTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IL1ETHGatewayFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IL1ETHGatewayFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IL1ETHGatewaySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IL1ETHGatewaySession struct {
	Contract     *IL1ETHGateway    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IL1ETHGatewayCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IL1ETHGatewayCallerSession struct {
	Contract *IL1ETHGatewayCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// IL1ETHGatewayTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IL1ETHGatewayTransactorSession struct {
	Contract     *IL1ETHGatewayTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// IL1ETHGatewayRaw is an auto generated low-level Go binding around an Ethereum contract.
type IL1ETHGatewayRaw struct {
	Contract *IL1ETHGateway // Generic contract binding to access the raw methods on
}

// IL1ETHGatewayCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IL1ETHGatewayCallerRaw struct {
	Contract *IL1ETHGatewayCaller // Generic read-only contract binding to access the raw methods on
}

// IL1ETHGatewayTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IL1ETHGatewayTransactorRaw struct {
	Contract *IL1ETHGatewayTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIL1ETHGateway creates a new instance of IL1ETHGateway, bound to a specific deployed contract.
func NewIL1ETHGateway(address common.Address, backend bind.ContractBackend) (*IL1ETHGateway, error) {
	contract, err := bindIL1ETHGateway(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IL1ETHGateway{IL1ETHGatewayCaller: IL1ETHGatewayCaller{contract: contract}, IL1ETHGatewayTransactor: IL1ETHGatewayTransactor{contract: contract}, IL1ETHGatewayFilterer: IL1ETHGatewayFilterer{contract: contract}}, nil
}

// NewIL1ETHGatewayCaller creates a new read-only instance of IL1ETHGateway, bound to a specific deployed contract.
func NewIL1ETHGatewayCaller(address common.Address, caller bind.ContractCaller) (*IL1ETHGatewayCaller, error) {
	contract, err := bindIL1ETHGateway(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IL1ETHGatewayCaller{contract: contract}, nil
}

// NewIL1ETHGatewayTransactor creates a new write-only instance of IL1ETHGateway, bound to a specific deployed contract.
func NewIL1ETHGatewayTransactor(address common.Address, transactor bind.ContractTransactor) (*IL1ETHGatewayTransactor, error) {
	contract, err := bindIL1ETHGateway(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IL1ETHGatewayTransactor{contract: contract}, nil
}

// NewIL1ETHGatewayFilterer creates a new log filterer instance of IL1ETHGateway, bound to a specific deployed contract.
func NewIL1ETHGatewayFilterer(address common.Address, filterer bind.ContractFilterer) (*IL1ETHGatewayFilterer, error) {
	contract, err := bindIL1ETHGateway(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IL1ETHGatewayFilterer{contract: contract}, nil
}

// bindIL1ETHGateway binds a generic wrapper to an already deployed contract.
func bindIL1ETHGateway(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IL1ETHGatewayABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IL1ETHGateway *IL1ETHGatewayRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IL1ETHGateway.Contract.IL1ETHGatewayCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IL1ETHGateway *IL1ETHGatewayRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IL1ETHGateway.Contract.IL1ETHGatewayTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IL1ETHGateway *IL1ETHGatewayRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IL1ETHGateway.Contract.IL1ETHGatewayTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IL1ETHGateway *IL1ETHGatewayCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IL1ETHGateway.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IL1ETHGateway *IL1ETHGatewayTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IL1ETHGateway.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IL1ETHGateway *IL1ETHGatewayTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IL1ETHGateway.Contract.contract.Transact(opts, method, params...)
}

// DepositETH is a paid mutator transaction binding the contract method 0x9f8420b3.
//
// Solidity: function depositETH(uint256 amount, uint256 gasLimit) payable returns()
func (_IL1ETHGateway *IL1ETHGatewayTransactor) DepositETH(opts *bind.TransactOpts, amount *big.Int, gasLimit *big.Int) (*types.Transaction, error) {
	return _IL1ETHGateway.contract.Transact(opts, "depositETH", amount, gasLimit)
}

// DepositETH is a paid mutator transaction binding the contract method 0x9f8420b3.
//
// Solidity: function depositETH(uint256 amount, uint256 gasLimit) payable returns()
func (_IL1ETHGateway *IL1ETHGatewaySession) DepositETH(amount *big.Int, gasLimit *big.Int) (*types.Transaction, error) {
	return _IL1ETHGateway.Contract.DepositETH(&_IL1ETHGateway.TransactOpts, amount, gasLimit)
}

// DepositETH is a paid mutator transaction binding the contract method 0x9f8420b3.
//
// Solidity: function depositETH(uint256 amount, uint256 gasLimit) payable returns()
func (_IL1ETHGateway *IL1ETHGatewayTransactorSession) DepositETH(amount *big.Int, gasLimit *big.Int) (*types.Transaction, error) {
	return _IL1ETHGateway.Contract.DepositETH(&_IL1ETHGateway.TransactOpts, amount, gasLimit)
}

// DepositETH0 is a paid mutator transaction binding the contract method 0xce0b63ce.
//
// Solidity: function depositETH(address to, uint256 amount, uint256 gasLimit) payable returns()
func (_IL1ETHGateway *IL1ETHGatewayTransactor) DepositETH0(opts *bind.TransactOpts, to common.Address, amount *big.Int, gasLimit *big.Int) (*types.Transaction, error) {
	return _IL1ETHGateway.contract.Transact(opts, "depositETH0", to, amount, gasLimit)
}

// DepositETH0 is a paid mutator transaction binding the contract method 0xce0b63ce.
//
// Solidity: function depositETH(address to, uint256 amount, uint256 gasLimit) payable returns()
func (_IL1ETHGateway *IL1ETHGatewaySession) DepositETH0(to common.Address, amount *big.Int, gasLimit *big.Int) (*types.Transaction, error) {
	return _IL1ETHGateway.Contract.DepositETH0(&_IL1ETHGateway.TransactOpts, to, amount, gasLimit)
}

// DepositETH0 is a paid mutator transaction binding the contract method 0xce0b63ce.
//
// Solidity: function depositETH(address to, uint256 amount, uint256 gasLimit) payable returns()
func (_IL1ETHGateway *IL1ETHGatewayTransactorSession) DepositETH0(to common.Address, amount *big.Int, gasLimit *big.Int) (*types.Transaction, error) {
	return _IL1ETHGateway.Contract.DepositETH0(&_IL1ETHGateway.TransactOpts, to, amount, gasLimit)
}

// DepositETHAndCall is a paid mutator transaction binding the contract method 0xaac476f8.
//
// Solidity: function depositETHAndCall(address to, uint256 amount, bytes data, uint256 gasLimit) payable returns()
func (_IL1ETHGateway *IL1ETHGatewayTransactor) DepositETHAndCall(opts *bind.TransactOpts, to common.Address, amount *big.Int, data []byte, gasLimit *big.Int) (*types.Transaction, error) {
	return _IL1ETHGateway.contract.Transact(opts, "depositETHAndCall", to, amount, data, gasLimit)
}

// DepositETHAndCall is a paid mutator transaction binding the contract method 0xaac476f8.
//
// Solidity: function depositETHAndCall(address to, uint256 amount, bytes data, uint256 gasLimit) payable returns()
func (_IL1ETHGateway *IL1ETHGatewaySession) DepositETHAndCall(to common.Address, amount *big.Int, data []byte, gasLimit *big.Int) (*types.Transaction, error) {
	return _IL1ETHGateway.Contract.DepositETHAndCall(&_IL1ETHGateway.TransactOpts, to, amount, data, gasLimit)
}

// DepositETHAndCall is a paid mutator transaction binding the contract method 0xaac476f8.
//
// Solidity: function depositETHAndCall(address to, uint256 amount, bytes data, uint256 gasLimit) payable returns()
func (_IL1ETHGateway *IL1ETHGatewayTransactorSession) DepositETHAndCall(to common.Address, amount *big.Int, data []byte, gasLimit *big.Int) (*types.Transaction, error) {
	return _IL1ETHGateway.Contract.DepositETHAndCall(&_IL1ETHGateway.TransactOpts, to, amount, data, gasLimit)
}

// FinalizeWithdrawETH is a paid mutator transaction binding the contract method 0x8eaac8a3.
//
// Solidity: function finalizeWithdrawETH(address from, address to, uint256 amount, bytes data) payable returns()
func (_IL1ETHGateway *IL1ETHGatewayTransactor) FinalizeWithdrawETH(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _IL1ETHGateway.contract.Transact(opts, "finalizeWithdrawETH", from, to, amount, data)
}

// FinalizeWithdrawETH is a paid mutator transaction binding the contract method 0x8eaac8a3.
//
// Solidity: function finalizeWithdrawETH(address from, address to, uint256 amount, bytes data) payable returns()
func (_IL1ETHGateway *IL1ETHGatewaySession) FinalizeWithdrawETH(from common.Address, to common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _IL1ETHGateway.Contract.FinalizeWithdrawETH(&_IL1ETHGateway.TransactOpts, from, to, amount, data)
}

// FinalizeWithdrawETH is a paid mutator transaction binding the contract method 0x8eaac8a3.
//
// Solidity: function finalizeWithdrawETH(address from, address to, uint256 amount, bytes data) payable returns()
func (_IL1ETHGateway *IL1ETHGatewayTransactorSession) FinalizeWithdrawETH(from common.Address, to common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _IL1ETHGateway.Contract.FinalizeWithdrawETH(&_IL1ETHGateway.TransactOpts, from, to, amount, data)
}

// IL1ETHGatewayDepositETHIterator is returned from FilterDepositETH and is used to iterate over the raw logs and unpacked data for DepositETH events raised by the IL1ETHGateway contract.
type IL1ETHGatewayDepositETHIterator struct {
	Event *IL1ETHGatewayDepositETH // Event containing the contract specifics and raw log

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
func (it *IL1ETHGatewayDepositETHIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IL1ETHGatewayDepositETH)
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
		it.Event = new(IL1ETHGatewayDepositETH)
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
func (it *IL1ETHGatewayDepositETHIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IL1ETHGatewayDepositETHIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IL1ETHGatewayDepositETH represents a DepositETH event raised by the IL1ETHGateway contract.
type IL1ETHGatewayDepositETH struct {
	From   common.Address
	To     common.Address
	Amount *big.Int
	Data   []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDepositETH is a free log retrieval operation binding the contract event 0x6670de856ec8bf5cb2b7e957c5dc24759716056f79d97ea5e7c939ca0ba5a675.
//
// Solidity: event DepositETH(address indexed from, address indexed to, uint256 amount, bytes data)
func (_IL1ETHGateway *IL1ETHGatewayFilterer) FilterDepositETH(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*IL1ETHGatewayDepositETHIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IL1ETHGateway.contract.FilterLogs(opts, "DepositETH", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IL1ETHGatewayDepositETHIterator{contract: _IL1ETHGateway.contract, event: "DepositETH", logs: logs, sub: sub}, nil
}

// WatchDepositETH is a free log subscription operation binding the contract event 0x6670de856ec8bf5cb2b7e957c5dc24759716056f79d97ea5e7c939ca0ba5a675.
//
// Solidity: event DepositETH(address indexed from, address indexed to, uint256 amount, bytes data)
func (_IL1ETHGateway *IL1ETHGatewayFilterer) WatchDepositETH(opts *bind.WatchOpts, sink chan<- *IL1ETHGatewayDepositETH, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IL1ETHGateway.contract.WatchLogs(opts, "DepositETH", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IL1ETHGatewayDepositETH)
				if err := _IL1ETHGateway.contract.UnpackLog(event, "DepositETH", log); err != nil {
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

// ParseDepositETH is a log parse operation binding the contract event 0x6670de856ec8bf5cb2b7e957c5dc24759716056f79d97ea5e7c939ca0ba5a675.
//
// Solidity: event DepositETH(address indexed from, address indexed to, uint256 amount, bytes data)
func (_IL1ETHGateway *IL1ETHGatewayFilterer) ParseDepositETH(log types.Log) (*IL1ETHGatewayDepositETH, error) {
	event := new(IL1ETHGatewayDepositETH)
	if err := _IL1ETHGateway.contract.UnpackLog(event, "DepositETH", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IL1ETHGatewayFinalizeWithdrawETHIterator is returned from FilterFinalizeWithdrawETH and is used to iterate over the raw logs and unpacked data for FinalizeWithdrawETH events raised by the IL1ETHGateway contract.
type IL1ETHGatewayFinalizeWithdrawETHIterator struct {
	Event *IL1ETHGatewayFinalizeWithdrawETH // Event containing the contract specifics and raw log

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
func (it *IL1ETHGatewayFinalizeWithdrawETHIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IL1ETHGatewayFinalizeWithdrawETH)
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
		it.Event = new(IL1ETHGatewayFinalizeWithdrawETH)
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
func (it *IL1ETHGatewayFinalizeWithdrawETHIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IL1ETHGatewayFinalizeWithdrawETHIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IL1ETHGatewayFinalizeWithdrawETH represents a FinalizeWithdrawETH event raised by the IL1ETHGateway contract.
type IL1ETHGatewayFinalizeWithdrawETH struct {
	From   common.Address
	To     common.Address
	Amount *big.Int
	Data   []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFinalizeWithdrawETH is a free log retrieval operation binding the contract event 0x96db5d1cee1dd2760826bb56fabd9c9f6e978083e0a8b88559c741a29e9746e7.
//
// Solidity: event FinalizeWithdrawETH(address indexed from, address indexed to, uint256 amount, bytes data)
func (_IL1ETHGateway *IL1ETHGatewayFilterer) FilterFinalizeWithdrawETH(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*IL1ETHGatewayFinalizeWithdrawETHIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IL1ETHGateway.contract.FilterLogs(opts, "FinalizeWithdrawETH", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IL1ETHGatewayFinalizeWithdrawETHIterator{contract: _IL1ETHGateway.contract, event: "FinalizeWithdrawETH", logs: logs, sub: sub}, nil
}

// WatchFinalizeWithdrawETH is a free log subscription operation binding the contract event 0x96db5d1cee1dd2760826bb56fabd9c9f6e978083e0a8b88559c741a29e9746e7.
//
// Solidity: event FinalizeWithdrawETH(address indexed from, address indexed to, uint256 amount, bytes data)
func (_IL1ETHGateway *IL1ETHGatewayFilterer) WatchFinalizeWithdrawETH(opts *bind.WatchOpts, sink chan<- *IL1ETHGatewayFinalizeWithdrawETH, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IL1ETHGateway.contract.WatchLogs(opts, "FinalizeWithdrawETH", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IL1ETHGatewayFinalizeWithdrawETH)
				if err := _IL1ETHGateway.contract.UnpackLog(event, "FinalizeWithdrawETH", log); err != nil {
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

// ParseFinalizeWithdrawETH is a log parse operation binding the contract event 0x96db5d1cee1dd2760826bb56fabd9c9f6e978083e0a8b88559c741a29e9746e7.
//
// Solidity: event FinalizeWithdrawETH(address indexed from, address indexed to, uint256 amount, bytes data)
func (_IL1ETHGateway *IL1ETHGatewayFilterer) ParseFinalizeWithdrawETH(log types.Log) (*IL1ETHGatewayFinalizeWithdrawETH, error) {
	event := new(IL1ETHGatewayFinalizeWithdrawETH)
	if err := _IL1ETHGateway.contract.UnpackLog(event, "FinalizeWithdrawETH", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IL1ETHGatewayRefundETHIterator is returned from FilterRefundETH and is used to iterate over the raw logs and unpacked data for RefundETH events raised by the IL1ETHGateway contract.
type IL1ETHGatewayRefundETHIterator struct {
	Event *IL1ETHGatewayRefundETH // Event containing the contract specifics and raw log

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
func (it *IL1ETHGatewayRefundETHIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IL1ETHGatewayRefundETH)
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
		it.Event = new(IL1ETHGatewayRefundETH)
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
func (it *IL1ETHGatewayRefundETHIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IL1ETHGatewayRefundETHIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IL1ETHGatewayRefundETH represents a RefundETH event raised by the IL1ETHGateway contract.
type IL1ETHGatewayRefundETH struct {
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRefundETH is a free log retrieval operation binding the contract event 0x289360176646a5f99cb4b6300628426dca46b723f40db3c04449d6ed1745a0e7.
//
// Solidity: event RefundETH(address indexed recipient, uint256 amount)
func (_IL1ETHGateway *IL1ETHGatewayFilterer) FilterRefundETH(opts *bind.FilterOpts, recipient []common.Address) (*IL1ETHGatewayRefundETHIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _IL1ETHGateway.contract.FilterLogs(opts, "RefundETH", recipientRule)
	if err != nil {
		return nil, err
	}
	return &IL1ETHGatewayRefundETHIterator{contract: _IL1ETHGateway.contract, event: "RefundETH", logs: logs, sub: sub}, nil
}

// WatchRefundETH is a free log subscription operation binding the contract event 0x289360176646a5f99cb4b6300628426dca46b723f40db3c04449d6ed1745a0e7.
//
// Solidity: event RefundETH(address indexed recipient, uint256 amount)
func (_IL1ETHGateway *IL1ETHGatewayFilterer) WatchRefundETH(opts *bind.WatchOpts, sink chan<- *IL1ETHGatewayRefundETH, recipient []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _IL1ETHGateway.contract.WatchLogs(opts, "RefundETH", recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IL1ETHGatewayRefundETH)
				if err := _IL1ETHGateway.contract.UnpackLog(event, "RefundETH", log); err != nil {
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

// ParseRefundETH is a log parse operation binding the contract event 0x289360176646a5f99cb4b6300628426dca46b723f40db3c04449d6ed1745a0e7.
//
// Solidity: event RefundETH(address indexed recipient, uint256 amount)
func (_IL1ETHGateway *IL1ETHGatewayFilterer) ParseRefundETH(log types.Log) (*IL1ETHGatewayRefundETH, error) {
	event := new(IL1ETHGatewayRefundETH)
	if err := _IL1ETHGateway.contract.UnpackLog(event, "RefundETH", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
