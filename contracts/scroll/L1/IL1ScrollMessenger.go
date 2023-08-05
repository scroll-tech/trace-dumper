// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package L1

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

// IL1ScrollMessengerL2MessageProof is an auto generated low-level Go binding around an user-defined struct.
type IL1ScrollMessengerL2MessageProof struct {
	BatchIndex  *big.Int
	MerkleProof []byte
}

// IL1ScrollMessengerMetaData contains all meta data concerning the IL1ScrollMessenger contract.
var IL1ScrollMessengerMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"}],\"name\":\"FailedRelayedMessage\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"}],\"name\":\"RelayedMessage\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"messageNonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"SentMessage\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxReplayTimes\",\"type\":\"uint256\"}],\"name\":\"UpdateMaxReplayTimes\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"messageNonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"dropMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"batchIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"merkleProof\",\"type\":\"bytes\"}],\"internalType\":\"structIL1ScrollMessenger.L2MessageProof\",\"name\":\"proof\",\"type\":\"tuple\"}],\"name\":\"relayMessageWithProof\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"messageNonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"uint32\",\"name\":\"newGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"refundAddress\",\"type\":\"address\"}],\"name\":\"replayMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"refundAddress\",\"type\":\"address\"}],\"name\":\"sendMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"}],\"name\":\"sendMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"xDomainMessageSender\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"29907acd": "dropMessage(address,address,uint256,uint256,bytes)",
		"c311b6fc": "relayMessageWithProof(address,address,uint256,uint256,bytes,(uint256,bytes))",
		"55004105": "replayMessage(address,address,uint256,uint256,bytes,uint32,address)",
		"b2267a7b": "sendMessage(address,uint256,bytes,uint256)",
		"5f7b1577": "sendMessage(address,uint256,bytes,uint256,address)",
		"6e296e45": "xDomainMessageSender()",
	},
}

// IL1ScrollMessengerABI is the input ABI used to generate the binding from.
// Deprecated: Use IL1ScrollMessengerMetaData.ABI instead.
var IL1ScrollMessengerABI = IL1ScrollMessengerMetaData.ABI

// Deprecated: Use IL1ScrollMessengerMetaData.Sigs instead.
// IL1ScrollMessengerFuncSigs maps the 4-byte function signature to its string representation.
var IL1ScrollMessengerFuncSigs = IL1ScrollMessengerMetaData.Sigs

// IL1ScrollMessenger is an auto generated Go binding around an Ethereum contract.
type IL1ScrollMessenger struct {
	IL1ScrollMessengerCaller     // Read-only binding to the contract
	IL1ScrollMessengerTransactor // Write-only binding to the contract
	IL1ScrollMessengerFilterer   // Log filterer for contract events
}

// IL1ScrollMessengerCaller is an auto generated read-only Go binding around an Ethereum contract.
type IL1ScrollMessengerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IL1ScrollMessengerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IL1ScrollMessengerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IL1ScrollMessengerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IL1ScrollMessengerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IL1ScrollMessengerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IL1ScrollMessengerSession struct {
	Contract     *IL1ScrollMessenger // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// IL1ScrollMessengerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IL1ScrollMessengerCallerSession struct {
	Contract *IL1ScrollMessengerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// IL1ScrollMessengerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IL1ScrollMessengerTransactorSession struct {
	Contract     *IL1ScrollMessengerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// IL1ScrollMessengerRaw is an auto generated low-level Go binding around an Ethereum contract.
type IL1ScrollMessengerRaw struct {
	Contract *IL1ScrollMessenger // Generic contract binding to access the raw methods on
}

// IL1ScrollMessengerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IL1ScrollMessengerCallerRaw struct {
	Contract *IL1ScrollMessengerCaller // Generic read-only contract binding to access the raw methods on
}

// IL1ScrollMessengerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IL1ScrollMessengerTransactorRaw struct {
	Contract *IL1ScrollMessengerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIL1ScrollMessenger creates a new instance of IL1ScrollMessenger, bound to a specific deployed contract.
func NewIL1ScrollMessenger(address common.Address, backend bind.ContractBackend) (*IL1ScrollMessenger, error) {
	contract, err := bindIL1ScrollMessenger(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IL1ScrollMessenger{IL1ScrollMessengerCaller: IL1ScrollMessengerCaller{contract: contract}, IL1ScrollMessengerTransactor: IL1ScrollMessengerTransactor{contract: contract}, IL1ScrollMessengerFilterer: IL1ScrollMessengerFilterer{contract: contract}}, nil
}

// NewIL1ScrollMessengerCaller creates a new read-only instance of IL1ScrollMessenger, bound to a specific deployed contract.
func NewIL1ScrollMessengerCaller(address common.Address, caller bind.ContractCaller) (*IL1ScrollMessengerCaller, error) {
	contract, err := bindIL1ScrollMessenger(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IL1ScrollMessengerCaller{contract: contract}, nil
}

// NewIL1ScrollMessengerTransactor creates a new write-only instance of IL1ScrollMessenger, bound to a specific deployed contract.
func NewIL1ScrollMessengerTransactor(address common.Address, transactor bind.ContractTransactor) (*IL1ScrollMessengerTransactor, error) {
	contract, err := bindIL1ScrollMessenger(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IL1ScrollMessengerTransactor{contract: contract}, nil
}

// NewIL1ScrollMessengerFilterer creates a new log filterer instance of IL1ScrollMessenger, bound to a specific deployed contract.
func NewIL1ScrollMessengerFilterer(address common.Address, filterer bind.ContractFilterer) (*IL1ScrollMessengerFilterer, error) {
	contract, err := bindIL1ScrollMessenger(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IL1ScrollMessengerFilterer{contract: contract}, nil
}

// bindIL1ScrollMessenger binds a generic wrapper to an already deployed contract.
func bindIL1ScrollMessenger(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IL1ScrollMessengerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IL1ScrollMessenger *IL1ScrollMessengerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IL1ScrollMessenger.Contract.IL1ScrollMessengerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IL1ScrollMessenger *IL1ScrollMessengerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IL1ScrollMessenger.Contract.IL1ScrollMessengerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IL1ScrollMessenger *IL1ScrollMessengerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IL1ScrollMessenger.Contract.IL1ScrollMessengerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IL1ScrollMessenger *IL1ScrollMessengerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IL1ScrollMessenger.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IL1ScrollMessenger *IL1ScrollMessengerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IL1ScrollMessenger.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IL1ScrollMessenger *IL1ScrollMessengerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IL1ScrollMessenger.Contract.contract.Transact(opts, method, params...)
}

// XDomainMessageSender is a free data retrieval call binding the contract method 0x6e296e45.
//
// Solidity: function xDomainMessageSender() view returns(address)
func (_IL1ScrollMessenger *IL1ScrollMessengerCaller) XDomainMessageSender(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IL1ScrollMessenger.contract.Call(opts, &out, "xDomainMessageSender")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// XDomainMessageSender is a free data retrieval call binding the contract method 0x6e296e45.
//
// Solidity: function xDomainMessageSender() view returns(address)
func (_IL1ScrollMessenger *IL1ScrollMessengerSession) XDomainMessageSender() (common.Address, error) {
	return _IL1ScrollMessenger.Contract.XDomainMessageSender(&_IL1ScrollMessenger.CallOpts)
}

// XDomainMessageSender is a free data retrieval call binding the contract method 0x6e296e45.
//
// Solidity: function xDomainMessageSender() view returns(address)
func (_IL1ScrollMessenger *IL1ScrollMessengerCallerSession) XDomainMessageSender() (common.Address, error) {
	return _IL1ScrollMessenger.Contract.XDomainMessageSender(&_IL1ScrollMessenger.CallOpts)
}

// DropMessage is a paid mutator transaction binding the contract method 0x29907acd.
//
// Solidity: function dropMessage(address from, address to, uint256 value, uint256 messageNonce, bytes message) returns()
func (_IL1ScrollMessenger *IL1ScrollMessengerTransactor) DropMessage(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int, messageNonce *big.Int, message []byte) (*types.Transaction, error) {
	return _IL1ScrollMessenger.contract.Transact(opts, "dropMessage", from, to, value, messageNonce, message)
}

// DropMessage is a paid mutator transaction binding the contract method 0x29907acd.
//
// Solidity: function dropMessage(address from, address to, uint256 value, uint256 messageNonce, bytes message) returns()
func (_IL1ScrollMessenger *IL1ScrollMessengerSession) DropMessage(from common.Address, to common.Address, value *big.Int, messageNonce *big.Int, message []byte) (*types.Transaction, error) {
	return _IL1ScrollMessenger.Contract.DropMessage(&_IL1ScrollMessenger.TransactOpts, from, to, value, messageNonce, message)
}

// DropMessage is a paid mutator transaction binding the contract method 0x29907acd.
//
// Solidity: function dropMessage(address from, address to, uint256 value, uint256 messageNonce, bytes message) returns()
func (_IL1ScrollMessenger *IL1ScrollMessengerTransactorSession) DropMessage(from common.Address, to common.Address, value *big.Int, messageNonce *big.Int, message []byte) (*types.Transaction, error) {
	return _IL1ScrollMessenger.Contract.DropMessage(&_IL1ScrollMessenger.TransactOpts, from, to, value, messageNonce, message)
}

// RelayMessageWithProof is a paid mutator transaction binding the contract method 0xc311b6fc.
//
// Solidity: function relayMessageWithProof(address from, address to, uint256 value, uint256 nonce, bytes message, (uint256,bytes) proof) returns()
func (_IL1ScrollMessenger *IL1ScrollMessengerTransactor) RelayMessageWithProof(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int, nonce *big.Int, message []byte, proof IL1ScrollMessengerL2MessageProof) (*types.Transaction, error) {
	return _IL1ScrollMessenger.contract.Transact(opts, "relayMessageWithProof", from, to, value, nonce, message, proof)
}

// RelayMessageWithProof is a paid mutator transaction binding the contract method 0xc311b6fc.
//
// Solidity: function relayMessageWithProof(address from, address to, uint256 value, uint256 nonce, bytes message, (uint256,bytes) proof) returns()
func (_IL1ScrollMessenger *IL1ScrollMessengerSession) RelayMessageWithProof(from common.Address, to common.Address, value *big.Int, nonce *big.Int, message []byte, proof IL1ScrollMessengerL2MessageProof) (*types.Transaction, error) {
	return _IL1ScrollMessenger.Contract.RelayMessageWithProof(&_IL1ScrollMessenger.TransactOpts, from, to, value, nonce, message, proof)
}

// RelayMessageWithProof is a paid mutator transaction binding the contract method 0xc311b6fc.
//
// Solidity: function relayMessageWithProof(address from, address to, uint256 value, uint256 nonce, bytes message, (uint256,bytes) proof) returns()
func (_IL1ScrollMessenger *IL1ScrollMessengerTransactorSession) RelayMessageWithProof(from common.Address, to common.Address, value *big.Int, nonce *big.Int, message []byte, proof IL1ScrollMessengerL2MessageProof) (*types.Transaction, error) {
	return _IL1ScrollMessenger.Contract.RelayMessageWithProof(&_IL1ScrollMessenger.TransactOpts, from, to, value, nonce, message, proof)
}

// ReplayMessage is a paid mutator transaction binding the contract method 0x55004105.
//
// Solidity: function replayMessage(address from, address to, uint256 value, uint256 messageNonce, bytes message, uint32 newGasLimit, address refundAddress) payable returns()
func (_IL1ScrollMessenger *IL1ScrollMessengerTransactor) ReplayMessage(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int, messageNonce *big.Int, message []byte, newGasLimit uint32, refundAddress common.Address) (*types.Transaction, error) {
	return _IL1ScrollMessenger.contract.Transact(opts, "replayMessage", from, to, value, messageNonce, message, newGasLimit, refundAddress)
}

// ReplayMessage is a paid mutator transaction binding the contract method 0x55004105.
//
// Solidity: function replayMessage(address from, address to, uint256 value, uint256 messageNonce, bytes message, uint32 newGasLimit, address refundAddress) payable returns()
func (_IL1ScrollMessenger *IL1ScrollMessengerSession) ReplayMessage(from common.Address, to common.Address, value *big.Int, messageNonce *big.Int, message []byte, newGasLimit uint32, refundAddress common.Address) (*types.Transaction, error) {
	return _IL1ScrollMessenger.Contract.ReplayMessage(&_IL1ScrollMessenger.TransactOpts, from, to, value, messageNonce, message, newGasLimit, refundAddress)
}

// ReplayMessage is a paid mutator transaction binding the contract method 0x55004105.
//
// Solidity: function replayMessage(address from, address to, uint256 value, uint256 messageNonce, bytes message, uint32 newGasLimit, address refundAddress) payable returns()
func (_IL1ScrollMessenger *IL1ScrollMessengerTransactorSession) ReplayMessage(from common.Address, to common.Address, value *big.Int, messageNonce *big.Int, message []byte, newGasLimit uint32, refundAddress common.Address) (*types.Transaction, error) {
	return _IL1ScrollMessenger.Contract.ReplayMessage(&_IL1ScrollMessenger.TransactOpts, from, to, value, messageNonce, message, newGasLimit, refundAddress)
}

// SendMessage is a paid mutator transaction binding the contract method 0x5f7b1577.
//
// Solidity: function sendMessage(address target, uint256 value, bytes message, uint256 gasLimit, address refundAddress) payable returns()
func (_IL1ScrollMessenger *IL1ScrollMessengerTransactor) SendMessage(opts *bind.TransactOpts, target common.Address, value *big.Int, message []byte, gasLimit *big.Int, refundAddress common.Address) (*types.Transaction, error) {
	return _IL1ScrollMessenger.contract.Transact(opts, "sendMessage", target, value, message, gasLimit, refundAddress)
}

// SendMessage is a paid mutator transaction binding the contract method 0x5f7b1577.
//
// Solidity: function sendMessage(address target, uint256 value, bytes message, uint256 gasLimit, address refundAddress) payable returns()
func (_IL1ScrollMessenger *IL1ScrollMessengerSession) SendMessage(target common.Address, value *big.Int, message []byte, gasLimit *big.Int, refundAddress common.Address) (*types.Transaction, error) {
	return _IL1ScrollMessenger.Contract.SendMessage(&_IL1ScrollMessenger.TransactOpts, target, value, message, gasLimit, refundAddress)
}

// SendMessage is a paid mutator transaction binding the contract method 0x5f7b1577.
//
// Solidity: function sendMessage(address target, uint256 value, bytes message, uint256 gasLimit, address refundAddress) payable returns()
func (_IL1ScrollMessenger *IL1ScrollMessengerTransactorSession) SendMessage(target common.Address, value *big.Int, message []byte, gasLimit *big.Int, refundAddress common.Address) (*types.Transaction, error) {
	return _IL1ScrollMessenger.Contract.SendMessage(&_IL1ScrollMessenger.TransactOpts, target, value, message, gasLimit, refundAddress)
}

// SendMessage0 is a paid mutator transaction binding the contract method 0xb2267a7b.
//
// Solidity: function sendMessage(address target, uint256 value, bytes message, uint256 gasLimit) payable returns()
func (_IL1ScrollMessenger *IL1ScrollMessengerTransactor) SendMessage0(opts *bind.TransactOpts, target common.Address, value *big.Int, message []byte, gasLimit *big.Int) (*types.Transaction, error) {
	return _IL1ScrollMessenger.contract.Transact(opts, "sendMessage0", target, value, message, gasLimit)
}

// SendMessage0 is a paid mutator transaction binding the contract method 0xb2267a7b.
//
// Solidity: function sendMessage(address target, uint256 value, bytes message, uint256 gasLimit) payable returns()
func (_IL1ScrollMessenger *IL1ScrollMessengerSession) SendMessage0(target common.Address, value *big.Int, message []byte, gasLimit *big.Int) (*types.Transaction, error) {
	return _IL1ScrollMessenger.Contract.SendMessage0(&_IL1ScrollMessenger.TransactOpts, target, value, message, gasLimit)
}

// SendMessage0 is a paid mutator transaction binding the contract method 0xb2267a7b.
//
// Solidity: function sendMessage(address target, uint256 value, bytes message, uint256 gasLimit) payable returns()
func (_IL1ScrollMessenger *IL1ScrollMessengerTransactorSession) SendMessage0(target common.Address, value *big.Int, message []byte, gasLimit *big.Int) (*types.Transaction, error) {
	return _IL1ScrollMessenger.Contract.SendMessage0(&_IL1ScrollMessenger.TransactOpts, target, value, message, gasLimit)
}

// IL1ScrollMessengerFailedRelayedMessageIterator is returned from FilterFailedRelayedMessage and is used to iterate over the raw logs and unpacked data for FailedRelayedMessage events raised by the IL1ScrollMessenger contract.
type IL1ScrollMessengerFailedRelayedMessageIterator struct {
	Event *IL1ScrollMessengerFailedRelayedMessage // Event containing the contract specifics and raw log

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
func (it *IL1ScrollMessengerFailedRelayedMessageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IL1ScrollMessengerFailedRelayedMessage)
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
		it.Event = new(IL1ScrollMessengerFailedRelayedMessage)
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
func (it *IL1ScrollMessengerFailedRelayedMessageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IL1ScrollMessengerFailedRelayedMessageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IL1ScrollMessengerFailedRelayedMessage represents a FailedRelayedMessage event raised by the IL1ScrollMessenger contract.
type IL1ScrollMessengerFailedRelayedMessage struct {
	MessageHash [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterFailedRelayedMessage is a free log retrieval operation binding the contract event 0x99d0e048484baa1b1540b1367cb128acd7ab2946d1ed91ec10e3c85e4bf51b8f.
//
// Solidity: event FailedRelayedMessage(bytes32 indexed messageHash)
func (_IL1ScrollMessenger *IL1ScrollMessengerFilterer) FilterFailedRelayedMessage(opts *bind.FilterOpts, messageHash [][32]byte) (*IL1ScrollMessengerFailedRelayedMessageIterator, error) {

	var messageHashRule []interface{}
	for _, messageHashItem := range messageHash {
		messageHashRule = append(messageHashRule, messageHashItem)
	}

	logs, sub, err := _IL1ScrollMessenger.contract.FilterLogs(opts, "FailedRelayedMessage", messageHashRule)
	if err != nil {
		return nil, err
	}
	return &IL1ScrollMessengerFailedRelayedMessageIterator{contract: _IL1ScrollMessenger.contract, event: "FailedRelayedMessage", logs: logs, sub: sub}, nil
}

// WatchFailedRelayedMessage is a free log subscription operation binding the contract event 0x99d0e048484baa1b1540b1367cb128acd7ab2946d1ed91ec10e3c85e4bf51b8f.
//
// Solidity: event FailedRelayedMessage(bytes32 indexed messageHash)
func (_IL1ScrollMessenger *IL1ScrollMessengerFilterer) WatchFailedRelayedMessage(opts *bind.WatchOpts, sink chan<- *IL1ScrollMessengerFailedRelayedMessage, messageHash [][32]byte) (event.Subscription, error) {

	var messageHashRule []interface{}
	for _, messageHashItem := range messageHash {
		messageHashRule = append(messageHashRule, messageHashItem)
	}

	logs, sub, err := _IL1ScrollMessenger.contract.WatchLogs(opts, "FailedRelayedMessage", messageHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IL1ScrollMessengerFailedRelayedMessage)
				if err := _IL1ScrollMessenger.contract.UnpackLog(event, "FailedRelayedMessage", log); err != nil {
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
func (_IL1ScrollMessenger *IL1ScrollMessengerFilterer) ParseFailedRelayedMessage(log types.Log) (*IL1ScrollMessengerFailedRelayedMessage, error) {
	event := new(IL1ScrollMessengerFailedRelayedMessage)
	if err := _IL1ScrollMessenger.contract.UnpackLog(event, "FailedRelayedMessage", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IL1ScrollMessengerRelayedMessageIterator is returned from FilterRelayedMessage and is used to iterate over the raw logs and unpacked data for RelayedMessage events raised by the IL1ScrollMessenger contract.
type IL1ScrollMessengerRelayedMessageIterator struct {
	Event *IL1ScrollMessengerRelayedMessage // Event containing the contract specifics and raw log

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
func (it *IL1ScrollMessengerRelayedMessageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IL1ScrollMessengerRelayedMessage)
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
		it.Event = new(IL1ScrollMessengerRelayedMessage)
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
func (it *IL1ScrollMessengerRelayedMessageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IL1ScrollMessengerRelayedMessageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IL1ScrollMessengerRelayedMessage represents a RelayedMessage event raised by the IL1ScrollMessenger contract.
type IL1ScrollMessengerRelayedMessage struct {
	MessageHash [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRelayedMessage is a free log retrieval operation binding the contract event 0x4641df4a962071e12719d8c8c8e5ac7fc4d97b927346a3d7a335b1f7517e133c.
//
// Solidity: event RelayedMessage(bytes32 indexed messageHash)
func (_IL1ScrollMessenger *IL1ScrollMessengerFilterer) FilterRelayedMessage(opts *bind.FilterOpts, messageHash [][32]byte) (*IL1ScrollMessengerRelayedMessageIterator, error) {

	var messageHashRule []interface{}
	for _, messageHashItem := range messageHash {
		messageHashRule = append(messageHashRule, messageHashItem)
	}

	logs, sub, err := _IL1ScrollMessenger.contract.FilterLogs(opts, "RelayedMessage", messageHashRule)
	if err != nil {
		return nil, err
	}
	return &IL1ScrollMessengerRelayedMessageIterator{contract: _IL1ScrollMessenger.contract, event: "RelayedMessage", logs: logs, sub: sub}, nil
}

// WatchRelayedMessage is a free log subscription operation binding the contract event 0x4641df4a962071e12719d8c8c8e5ac7fc4d97b927346a3d7a335b1f7517e133c.
//
// Solidity: event RelayedMessage(bytes32 indexed messageHash)
func (_IL1ScrollMessenger *IL1ScrollMessengerFilterer) WatchRelayedMessage(opts *bind.WatchOpts, sink chan<- *IL1ScrollMessengerRelayedMessage, messageHash [][32]byte) (event.Subscription, error) {

	var messageHashRule []interface{}
	for _, messageHashItem := range messageHash {
		messageHashRule = append(messageHashRule, messageHashItem)
	}

	logs, sub, err := _IL1ScrollMessenger.contract.WatchLogs(opts, "RelayedMessage", messageHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IL1ScrollMessengerRelayedMessage)
				if err := _IL1ScrollMessenger.contract.UnpackLog(event, "RelayedMessage", log); err != nil {
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
func (_IL1ScrollMessenger *IL1ScrollMessengerFilterer) ParseRelayedMessage(log types.Log) (*IL1ScrollMessengerRelayedMessage, error) {
	event := new(IL1ScrollMessengerRelayedMessage)
	if err := _IL1ScrollMessenger.contract.UnpackLog(event, "RelayedMessage", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IL1ScrollMessengerSentMessageIterator is returned from FilterSentMessage and is used to iterate over the raw logs and unpacked data for SentMessage events raised by the IL1ScrollMessenger contract.
type IL1ScrollMessengerSentMessageIterator struct {
	Event *IL1ScrollMessengerSentMessage // Event containing the contract specifics and raw log

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
func (it *IL1ScrollMessengerSentMessageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IL1ScrollMessengerSentMessage)
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
		it.Event = new(IL1ScrollMessengerSentMessage)
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
func (it *IL1ScrollMessengerSentMessageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IL1ScrollMessengerSentMessageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IL1ScrollMessengerSentMessage represents a SentMessage event raised by the IL1ScrollMessenger contract.
type IL1ScrollMessengerSentMessage struct {
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
func (_IL1ScrollMessenger *IL1ScrollMessengerFilterer) FilterSentMessage(opts *bind.FilterOpts, sender []common.Address, target []common.Address) (*IL1ScrollMessengerSentMessageIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _IL1ScrollMessenger.contract.FilterLogs(opts, "SentMessage", senderRule, targetRule)
	if err != nil {
		return nil, err
	}
	return &IL1ScrollMessengerSentMessageIterator{contract: _IL1ScrollMessenger.contract, event: "SentMessage", logs: logs, sub: sub}, nil
}

// WatchSentMessage is a free log subscription operation binding the contract event 0x104371f3b442861a2a7b82a070afbbaab748bb13757bf47769e170e37809ec1e.
//
// Solidity: event SentMessage(address indexed sender, address indexed target, uint256 value, uint256 messageNonce, uint256 gasLimit, bytes message)
func (_IL1ScrollMessenger *IL1ScrollMessengerFilterer) WatchSentMessage(opts *bind.WatchOpts, sink chan<- *IL1ScrollMessengerSentMessage, sender []common.Address, target []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _IL1ScrollMessenger.contract.WatchLogs(opts, "SentMessage", senderRule, targetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IL1ScrollMessengerSentMessage)
				if err := _IL1ScrollMessenger.contract.UnpackLog(event, "SentMessage", log); err != nil {
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
func (_IL1ScrollMessenger *IL1ScrollMessengerFilterer) ParseSentMessage(log types.Log) (*IL1ScrollMessengerSentMessage, error) {
	event := new(IL1ScrollMessengerSentMessage)
	if err := _IL1ScrollMessenger.contract.UnpackLog(event, "SentMessage", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IL1ScrollMessengerUpdateMaxReplayTimesIterator is returned from FilterUpdateMaxReplayTimes and is used to iterate over the raw logs and unpacked data for UpdateMaxReplayTimes events raised by the IL1ScrollMessenger contract.
type IL1ScrollMessengerUpdateMaxReplayTimesIterator struct {
	Event *IL1ScrollMessengerUpdateMaxReplayTimes // Event containing the contract specifics and raw log

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
func (it *IL1ScrollMessengerUpdateMaxReplayTimesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IL1ScrollMessengerUpdateMaxReplayTimes)
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
		it.Event = new(IL1ScrollMessengerUpdateMaxReplayTimes)
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
func (it *IL1ScrollMessengerUpdateMaxReplayTimesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IL1ScrollMessengerUpdateMaxReplayTimesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IL1ScrollMessengerUpdateMaxReplayTimes represents a UpdateMaxReplayTimes event raised by the IL1ScrollMessenger contract.
type IL1ScrollMessengerUpdateMaxReplayTimes struct {
	MaxReplayTimes *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpdateMaxReplayTimes is a free log retrieval operation binding the contract event 0x87447df30e533f82a48d0b72640868bb678dbec0d35314564ad5590a0191d90a.
//
// Solidity: event UpdateMaxReplayTimes(uint256 maxReplayTimes)
func (_IL1ScrollMessenger *IL1ScrollMessengerFilterer) FilterUpdateMaxReplayTimes(opts *bind.FilterOpts) (*IL1ScrollMessengerUpdateMaxReplayTimesIterator, error) {

	logs, sub, err := _IL1ScrollMessenger.contract.FilterLogs(opts, "UpdateMaxReplayTimes")
	if err != nil {
		return nil, err
	}
	return &IL1ScrollMessengerUpdateMaxReplayTimesIterator{contract: _IL1ScrollMessenger.contract, event: "UpdateMaxReplayTimes", logs: logs, sub: sub}, nil
}

// WatchUpdateMaxReplayTimes is a free log subscription operation binding the contract event 0x87447df30e533f82a48d0b72640868bb678dbec0d35314564ad5590a0191d90a.
//
// Solidity: event UpdateMaxReplayTimes(uint256 maxReplayTimes)
func (_IL1ScrollMessenger *IL1ScrollMessengerFilterer) WatchUpdateMaxReplayTimes(opts *bind.WatchOpts, sink chan<- *IL1ScrollMessengerUpdateMaxReplayTimes) (event.Subscription, error) {

	logs, sub, err := _IL1ScrollMessenger.contract.WatchLogs(opts, "UpdateMaxReplayTimes")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IL1ScrollMessengerUpdateMaxReplayTimes)
				if err := _IL1ScrollMessenger.contract.UnpackLog(event, "UpdateMaxReplayTimes", log); err != nil {
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

// ParseUpdateMaxReplayTimes is a log parse operation binding the contract event 0x87447df30e533f82a48d0b72640868bb678dbec0d35314564ad5590a0191d90a.
//
// Solidity: event UpdateMaxReplayTimes(uint256 maxReplayTimes)
func (_IL1ScrollMessenger *IL1ScrollMessengerFilterer) ParseUpdateMaxReplayTimes(log types.Log) (*IL1ScrollMessengerUpdateMaxReplayTimes, error) {
	event := new(IL1ScrollMessengerUpdateMaxReplayTimes)
	if err := _IL1ScrollMessenger.contract.UnpackLog(event, "UpdateMaxReplayTimes", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
