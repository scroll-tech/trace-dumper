// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package L2

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

// IL2ScrollMessengerMetaData contains all meta data concerning the IL2ScrollMessenger contract.
var IL2ScrollMessengerMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"}],\"name\":\"FailedRelayedMessage\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"}],\"name\":\"RelayedMessage\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"messageNonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"SentMessage\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"relayMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"refundAddress\",\"type\":\"address\"}],\"name\":\"sendMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"}],\"name\":\"sendMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"xDomainMessageSender\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"8ef1332e": "relayMessage(address,address,uint256,uint256,bytes)",
		"b2267a7b": "sendMessage(address,uint256,bytes,uint256)",
		"5f7b1577": "sendMessage(address,uint256,bytes,uint256,address)",
		"6e296e45": "xDomainMessageSender()",
	},
}

// IL2ScrollMessengerABI is the input ABI used to generate the binding from.
// Deprecated: Use IL2ScrollMessengerMetaData.ABI instead.
var IL2ScrollMessengerABI = IL2ScrollMessengerMetaData.ABI

// Deprecated: Use IL2ScrollMessengerMetaData.Sigs instead.
// IL2ScrollMessengerFuncSigs maps the 4-byte function signature to its string representation.
var IL2ScrollMessengerFuncSigs = IL2ScrollMessengerMetaData.Sigs

// IL2ScrollMessenger is an auto generated Go binding around an Ethereum contract.
type IL2ScrollMessenger struct {
	IL2ScrollMessengerCaller     // Read-only binding to the contract
	IL2ScrollMessengerTransactor // Write-only binding to the contract
	IL2ScrollMessengerFilterer   // Log filterer for contract events
}

// IL2ScrollMessengerCaller is an auto generated read-only Go binding around an Ethereum contract.
type IL2ScrollMessengerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IL2ScrollMessengerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IL2ScrollMessengerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IL2ScrollMessengerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IL2ScrollMessengerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IL2ScrollMessengerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IL2ScrollMessengerSession struct {
	Contract     *IL2ScrollMessenger // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// IL2ScrollMessengerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IL2ScrollMessengerCallerSession struct {
	Contract *IL2ScrollMessengerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// IL2ScrollMessengerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IL2ScrollMessengerTransactorSession struct {
	Contract     *IL2ScrollMessengerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// IL2ScrollMessengerRaw is an auto generated low-level Go binding around an Ethereum contract.
type IL2ScrollMessengerRaw struct {
	Contract *IL2ScrollMessenger // Generic contract binding to access the raw methods on
}

// IL2ScrollMessengerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IL2ScrollMessengerCallerRaw struct {
	Contract *IL2ScrollMessengerCaller // Generic read-only contract binding to access the raw methods on
}

// IL2ScrollMessengerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IL2ScrollMessengerTransactorRaw struct {
	Contract *IL2ScrollMessengerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIL2ScrollMessenger creates a new instance of IL2ScrollMessenger, bound to a specific deployed contract.
func NewIL2ScrollMessenger(address common.Address, backend bind.ContractBackend) (*IL2ScrollMessenger, error) {
	contract, err := bindIL2ScrollMessenger(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IL2ScrollMessenger{IL2ScrollMessengerCaller: IL2ScrollMessengerCaller{contract: contract}, IL2ScrollMessengerTransactor: IL2ScrollMessengerTransactor{contract: contract}, IL2ScrollMessengerFilterer: IL2ScrollMessengerFilterer{contract: contract}}, nil
}

// NewIL2ScrollMessengerCaller creates a new read-only instance of IL2ScrollMessenger, bound to a specific deployed contract.
func NewIL2ScrollMessengerCaller(address common.Address, caller bind.ContractCaller) (*IL2ScrollMessengerCaller, error) {
	contract, err := bindIL2ScrollMessenger(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IL2ScrollMessengerCaller{contract: contract}, nil
}

// NewIL2ScrollMessengerTransactor creates a new write-only instance of IL2ScrollMessenger, bound to a specific deployed contract.
func NewIL2ScrollMessengerTransactor(address common.Address, transactor bind.ContractTransactor) (*IL2ScrollMessengerTransactor, error) {
	contract, err := bindIL2ScrollMessenger(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IL2ScrollMessengerTransactor{contract: contract}, nil
}

// NewIL2ScrollMessengerFilterer creates a new log filterer instance of IL2ScrollMessenger, bound to a specific deployed contract.
func NewIL2ScrollMessengerFilterer(address common.Address, filterer bind.ContractFilterer) (*IL2ScrollMessengerFilterer, error) {
	contract, err := bindIL2ScrollMessenger(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IL2ScrollMessengerFilterer{contract: contract}, nil
}

// bindIL2ScrollMessenger binds a generic wrapper to an already deployed contract.
func bindIL2ScrollMessenger(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IL2ScrollMessengerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IL2ScrollMessenger *IL2ScrollMessengerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IL2ScrollMessenger.Contract.IL2ScrollMessengerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IL2ScrollMessenger *IL2ScrollMessengerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IL2ScrollMessenger.Contract.IL2ScrollMessengerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IL2ScrollMessenger *IL2ScrollMessengerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IL2ScrollMessenger.Contract.IL2ScrollMessengerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IL2ScrollMessenger *IL2ScrollMessengerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IL2ScrollMessenger.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IL2ScrollMessenger *IL2ScrollMessengerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IL2ScrollMessenger.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IL2ScrollMessenger *IL2ScrollMessengerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IL2ScrollMessenger.Contract.contract.Transact(opts, method, params...)
}

// XDomainMessageSender is a free data retrieval call binding the contract method 0x6e296e45.
//
// Solidity: function xDomainMessageSender() view returns(address)
func (_IL2ScrollMessenger *IL2ScrollMessengerCaller) XDomainMessageSender(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IL2ScrollMessenger.contract.Call(opts, &out, "xDomainMessageSender")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// XDomainMessageSender is a free data retrieval call binding the contract method 0x6e296e45.
//
// Solidity: function xDomainMessageSender() view returns(address)
func (_IL2ScrollMessenger *IL2ScrollMessengerSession) XDomainMessageSender() (common.Address, error) {
	return _IL2ScrollMessenger.Contract.XDomainMessageSender(&_IL2ScrollMessenger.CallOpts)
}

// XDomainMessageSender is a free data retrieval call binding the contract method 0x6e296e45.
//
// Solidity: function xDomainMessageSender() view returns(address)
func (_IL2ScrollMessenger *IL2ScrollMessengerCallerSession) XDomainMessageSender() (common.Address, error) {
	return _IL2ScrollMessenger.Contract.XDomainMessageSender(&_IL2ScrollMessenger.CallOpts)
}

// RelayMessage is a paid mutator transaction binding the contract method 0x8ef1332e.
//
// Solidity: function relayMessage(address from, address to, uint256 value, uint256 nonce, bytes message) returns()
func (_IL2ScrollMessenger *IL2ScrollMessengerTransactor) RelayMessage(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int, nonce *big.Int, message []byte) (*types.Transaction, error) {
	return _IL2ScrollMessenger.contract.Transact(opts, "relayMessage", from, to, value, nonce, message)
}

// RelayMessage is a paid mutator transaction binding the contract method 0x8ef1332e.
//
// Solidity: function relayMessage(address from, address to, uint256 value, uint256 nonce, bytes message) returns()
func (_IL2ScrollMessenger *IL2ScrollMessengerSession) RelayMessage(from common.Address, to common.Address, value *big.Int, nonce *big.Int, message []byte) (*types.Transaction, error) {
	return _IL2ScrollMessenger.Contract.RelayMessage(&_IL2ScrollMessenger.TransactOpts, from, to, value, nonce, message)
}

// RelayMessage is a paid mutator transaction binding the contract method 0x8ef1332e.
//
// Solidity: function relayMessage(address from, address to, uint256 value, uint256 nonce, bytes message) returns()
func (_IL2ScrollMessenger *IL2ScrollMessengerTransactorSession) RelayMessage(from common.Address, to common.Address, value *big.Int, nonce *big.Int, message []byte) (*types.Transaction, error) {
	return _IL2ScrollMessenger.Contract.RelayMessage(&_IL2ScrollMessenger.TransactOpts, from, to, value, nonce, message)
}

// SendMessage is a paid mutator transaction binding the contract method 0x5f7b1577.
//
// Solidity: function sendMessage(address target, uint256 value, bytes message, uint256 gasLimit, address refundAddress) payable returns()
func (_IL2ScrollMessenger *IL2ScrollMessengerTransactor) SendMessage(opts *bind.TransactOpts, target common.Address, value *big.Int, message []byte, gasLimit *big.Int, refundAddress common.Address) (*types.Transaction, error) {
	return _IL2ScrollMessenger.contract.Transact(opts, "sendMessage", target, value, message, gasLimit, refundAddress)
}

// SendMessage is a paid mutator transaction binding the contract method 0x5f7b1577.
//
// Solidity: function sendMessage(address target, uint256 value, bytes message, uint256 gasLimit, address refundAddress) payable returns()
func (_IL2ScrollMessenger *IL2ScrollMessengerSession) SendMessage(target common.Address, value *big.Int, message []byte, gasLimit *big.Int, refundAddress common.Address) (*types.Transaction, error) {
	return _IL2ScrollMessenger.Contract.SendMessage(&_IL2ScrollMessenger.TransactOpts, target, value, message, gasLimit, refundAddress)
}

// SendMessage is a paid mutator transaction binding the contract method 0x5f7b1577.
//
// Solidity: function sendMessage(address target, uint256 value, bytes message, uint256 gasLimit, address refundAddress) payable returns()
func (_IL2ScrollMessenger *IL2ScrollMessengerTransactorSession) SendMessage(target common.Address, value *big.Int, message []byte, gasLimit *big.Int, refundAddress common.Address) (*types.Transaction, error) {
	return _IL2ScrollMessenger.Contract.SendMessage(&_IL2ScrollMessenger.TransactOpts, target, value, message, gasLimit, refundAddress)
}

// SendMessage0 is a paid mutator transaction binding the contract method 0xb2267a7b.
//
// Solidity: function sendMessage(address target, uint256 value, bytes message, uint256 gasLimit) payable returns()
func (_IL2ScrollMessenger *IL2ScrollMessengerTransactor) SendMessage0(opts *bind.TransactOpts, target common.Address, value *big.Int, message []byte, gasLimit *big.Int) (*types.Transaction, error) {
	return _IL2ScrollMessenger.contract.Transact(opts, "sendMessage0", target, value, message, gasLimit)
}

// SendMessage0 is a paid mutator transaction binding the contract method 0xb2267a7b.
//
// Solidity: function sendMessage(address target, uint256 value, bytes message, uint256 gasLimit) payable returns()
func (_IL2ScrollMessenger *IL2ScrollMessengerSession) SendMessage0(target common.Address, value *big.Int, message []byte, gasLimit *big.Int) (*types.Transaction, error) {
	return _IL2ScrollMessenger.Contract.SendMessage0(&_IL2ScrollMessenger.TransactOpts, target, value, message, gasLimit)
}

// SendMessage0 is a paid mutator transaction binding the contract method 0xb2267a7b.
//
// Solidity: function sendMessage(address target, uint256 value, bytes message, uint256 gasLimit) payable returns()
func (_IL2ScrollMessenger *IL2ScrollMessengerTransactorSession) SendMessage0(target common.Address, value *big.Int, message []byte, gasLimit *big.Int) (*types.Transaction, error) {
	return _IL2ScrollMessenger.Contract.SendMessage0(&_IL2ScrollMessenger.TransactOpts, target, value, message, gasLimit)
}

// IL2ScrollMessengerFailedRelayedMessageIterator is returned from FilterFailedRelayedMessage and is used to iterate over the raw logs and unpacked data for FailedRelayedMessage events raised by the IL2ScrollMessenger contract.
type IL2ScrollMessengerFailedRelayedMessageIterator struct {
	Event *IL2ScrollMessengerFailedRelayedMessage // Event containing the contract specifics and raw log

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
func (it *IL2ScrollMessengerFailedRelayedMessageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IL2ScrollMessengerFailedRelayedMessage)
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
		it.Event = new(IL2ScrollMessengerFailedRelayedMessage)
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
func (it *IL2ScrollMessengerFailedRelayedMessageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IL2ScrollMessengerFailedRelayedMessageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IL2ScrollMessengerFailedRelayedMessage represents a FailedRelayedMessage event raised by the IL2ScrollMessenger contract.
type IL2ScrollMessengerFailedRelayedMessage struct {
	MessageHash [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterFailedRelayedMessage is a free log retrieval operation binding the contract event 0x99d0e048484baa1b1540b1367cb128acd7ab2946d1ed91ec10e3c85e4bf51b8f.
//
// Solidity: event FailedRelayedMessage(bytes32 indexed messageHash)
func (_IL2ScrollMessenger *IL2ScrollMessengerFilterer) FilterFailedRelayedMessage(opts *bind.FilterOpts, messageHash [][32]byte) (*IL2ScrollMessengerFailedRelayedMessageIterator, error) {

	var messageHashRule []interface{}
	for _, messageHashItem := range messageHash {
		messageHashRule = append(messageHashRule, messageHashItem)
	}

	logs, sub, err := _IL2ScrollMessenger.contract.FilterLogs(opts, "FailedRelayedMessage", messageHashRule)
	if err != nil {
		return nil, err
	}
	return &IL2ScrollMessengerFailedRelayedMessageIterator{contract: _IL2ScrollMessenger.contract, event: "FailedRelayedMessage", logs: logs, sub: sub}, nil
}

// WatchFailedRelayedMessage is a free log subscription operation binding the contract event 0x99d0e048484baa1b1540b1367cb128acd7ab2946d1ed91ec10e3c85e4bf51b8f.
//
// Solidity: event FailedRelayedMessage(bytes32 indexed messageHash)
func (_IL2ScrollMessenger *IL2ScrollMessengerFilterer) WatchFailedRelayedMessage(opts *bind.WatchOpts, sink chan<- *IL2ScrollMessengerFailedRelayedMessage, messageHash [][32]byte) (event.Subscription, error) {

	var messageHashRule []interface{}
	for _, messageHashItem := range messageHash {
		messageHashRule = append(messageHashRule, messageHashItem)
	}

	logs, sub, err := _IL2ScrollMessenger.contract.WatchLogs(opts, "FailedRelayedMessage", messageHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IL2ScrollMessengerFailedRelayedMessage)
				if err := _IL2ScrollMessenger.contract.UnpackLog(event, "FailedRelayedMessage", log); err != nil {
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

// ParseFailedRelayedMessage is a log parse operation binding the contract event 0x99d0e048484baa1b1540b1367cb128acd7ab2946d1ed91ec10e3c85e4bf51b8f.
//
// Solidity: event FailedRelayedMessage(bytes32 indexed messageHash)
func (_IL2ScrollMessenger *IL2ScrollMessengerFilterer) ParseFailedRelayedMessage(log types.Log) (*IL2ScrollMessengerFailedRelayedMessage, error) {
	event := new(IL2ScrollMessengerFailedRelayedMessage)
	if err := _IL2ScrollMessenger.contract.UnpackLog(event, "FailedRelayedMessage", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IL2ScrollMessengerRelayedMessageIterator is returned from FilterRelayedMessage and is used to iterate over the raw logs and unpacked data for RelayedMessage events raised by the IL2ScrollMessenger contract.
type IL2ScrollMessengerRelayedMessageIterator struct {
	Event *IL2ScrollMessengerRelayedMessage // Event containing the contract specifics and raw log

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
func (it *IL2ScrollMessengerRelayedMessageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IL2ScrollMessengerRelayedMessage)
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
		it.Event = new(IL2ScrollMessengerRelayedMessage)
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
func (it *IL2ScrollMessengerRelayedMessageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IL2ScrollMessengerRelayedMessageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IL2ScrollMessengerRelayedMessage represents a RelayedMessage event raised by the IL2ScrollMessenger contract.
type IL2ScrollMessengerRelayedMessage struct {
	MessageHash [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRelayedMessage is a free log retrieval operation binding the contract event 0x4641df4a962071e12719d8c8c8e5ac7fc4d97b927346a3d7a335b1f7517e133c.
//
// Solidity: event RelayedMessage(bytes32 indexed messageHash)
func (_IL2ScrollMessenger *IL2ScrollMessengerFilterer) FilterRelayedMessage(opts *bind.FilterOpts, messageHash [][32]byte) (*IL2ScrollMessengerRelayedMessageIterator, error) {

	var messageHashRule []interface{}
	for _, messageHashItem := range messageHash {
		messageHashRule = append(messageHashRule, messageHashItem)
	}

	logs, sub, err := _IL2ScrollMessenger.contract.FilterLogs(opts, "RelayedMessage", messageHashRule)
	if err != nil {
		return nil, err
	}
	return &IL2ScrollMessengerRelayedMessageIterator{contract: _IL2ScrollMessenger.contract, event: "RelayedMessage", logs: logs, sub: sub}, nil
}

// WatchRelayedMessage is a free log subscription operation binding the contract event 0x4641df4a962071e12719d8c8c8e5ac7fc4d97b927346a3d7a335b1f7517e133c.
//
// Solidity: event RelayedMessage(bytes32 indexed messageHash)
func (_IL2ScrollMessenger *IL2ScrollMessengerFilterer) WatchRelayedMessage(opts *bind.WatchOpts, sink chan<- *IL2ScrollMessengerRelayedMessage, messageHash [][32]byte) (event.Subscription, error) {

	var messageHashRule []interface{}
	for _, messageHashItem := range messageHash {
		messageHashRule = append(messageHashRule, messageHashItem)
	}

	logs, sub, err := _IL2ScrollMessenger.contract.WatchLogs(opts, "RelayedMessage", messageHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IL2ScrollMessengerRelayedMessage)
				if err := _IL2ScrollMessenger.contract.UnpackLog(event, "RelayedMessage", log); err != nil {
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

// ParseRelayedMessage is a log parse operation binding the contract event 0x4641df4a962071e12719d8c8c8e5ac7fc4d97b927346a3d7a335b1f7517e133c.
//
// Solidity: event RelayedMessage(bytes32 indexed messageHash)
func (_IL2ScrollMessenger *IL2ScrollMessengerFilterer) ParseRelayedMessage(log types.Log) (*IL2ScrollMessengerRelayedMessage, error) {
	event := new(IL2ScrollMessengerRelayedMessage)
	if err := _IL2ScrollMessenger.contract.UnpackLog(event, "RelayedMessage", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IL2ScrollMessengerSentMessageIterator is returned from FilterSentMessage and is used to iterate over the raw logs and unpacked data for SentMessage events raised by the IL2ScrollMessenger contract.
type IL2ScrollMessengerSentMessageIterator struct {
	Event *IL2ScrollMessengerSentMessage // Event containing the contract specifics and raw log

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
func (it *IL2ScrollMessengerSentMessageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IL2ScrollMessengerSentMessage)
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
		it.Event = new(IL2ScrollMessengerSentMessage)
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
func (it *IL2ScrollMessengerSentMessageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IL2ScrollMessengerSentMessageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IL2ScrollMessengerSentMessage represents a SentMessage event raised by the IL2ScrollMessenger contract.
type IL2ScrollMessengerSentMessage struct {
	Sender       common.Address
	Target       common.Address
	Value        *big.Int
	MessageNonce *big.Int
	GasLimit     *big.Int
	Message      []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSentMessage is a free log retrieval operation binding the contract event 0x104371f3b442861a2a7b82a070afbbaab748bb13757bf47769e170e37809ec1e.
//
// Solidity: event SentMessage(address indexed sender, address indexed target, uint256 value, uint256 messageNonce, uint256 gasLimit, bytes message)
func (_IL2ScrollMessenger *IL2ScrollMessengerFilterer) FilterSentMessage(opts *bind.FilterOpts, sender []common.Address, target []common.Address) (*IL2ScrollMessengerSentMessageIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _IL2ScrollMessenger.contract.FilterLogs(opts, "SentMessage", senderRule, targetRule)
	if err != nil {
		return nil, err
	}
	return &IL2ScrollMessengerSentMessageIterator{contract: _IL2ScrollMessenger.contract, event: "SentMessage", logs: logs, sub: sub}, nil
}

// WatchSentMessage is a free log subscription operation binding the contract event 0x104371f3b442861a2a7b82a070afbbaab748bb13757bf47769e170e37809ec1e.
//
// Solidity: event SentMessage(address indexed sender, address indexed target, uint256 value, uint256 messageNonce, uint256 gasLimit, bytes message)
func (_IL2ScrollMessenger *IL2ScrollMessengerFilterer) WatchSentMessage(opts *bind.WatchOpts, sink chan<- *IL2ScrollMessengerSentMessage, sender []common.Address, target []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _IL2ScrollMessenger.contract.WatchLogs(opts, "SentMessage", senderRule, targetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IL2ScrollMessengerSentMessage)
				if err := _IL2ScrollMessenger.contract.UnpackLog(event, "SentMessage", log); err != nil {
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

// ParseSentMessage is a log parse operation binding the contract event 0x104371f3b442861a2a7b82a070afbbaab748bb13757bf47769e170e37809ec1e.
//
// Solidity: event SentMessage(address indexed sender, address indexed target, uint256 value, uint256 messageNonce, uint256 gasLimit, bytes message)
func (_IL2ScrollMessenger *IL2ScrollMessengerFilterer) ParseSentMessage(log types.Log) (*IL2ScrollMessengerSentMessage, error) {
	event := new(IL2ScrollMessengerSentMessage)
	if err := _IL2ScrollMessenger.contract.UnpackLog(event, "SentMessage", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
