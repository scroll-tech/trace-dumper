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

// IL2ETHGatewayMetaData contains all meta data concerning the IL2ETHGateway contract.
var IL2ETHGatewayMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"FinalizeDepositETH\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"WithdrawETH\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"finalizeDepositETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"}],\"name\":\"withdrawETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"}],\"name\":\"withdrawETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"}],\"name\":\"withdrawETHAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"232e8748": "finalizeDepositETH(address,address,uint256,bytes)",
		"2fcc29fa": "withdrawETH(address,uint256,uint256)",
		"c7cdea37": "withdrawETH(uint256,uint256)",
		"6dc24183": "withdrawETHAndCall(address,uint256,bytes,uint256)",
	},
}

// IL2ETHGatewayABI is the input ABI used to generate the binding from.
// Deprecated: Use IL2ETHGatewayMetaData.ABI instead.
var IL2ETHGatewayABI = IL2ETHGatewayMetaData.ABI

// Deprecated: Use IL2ETHGatewayMetaData.Sigs instead.
// IL2ETHGatewayFuncSigs maps the 4-byte function signature to its string representation.
var IL2ETHGatewayFuncSigs = IL2ETHGatewayMetaData.Sigs

// IL2ETHGateway is an auto generated Go binding around an Ethereum contract.
type IL2ETHGateway struct {
	IL2ETHGatewayCaller     // Read-only binding to the contract
	IL2ETHGatewayTransactor // Write-only binding to the contract
	IL2ETHGatewayFilterer   // Log filterer for contract events
}

// IL2ETHGatewayCaller is an auto generated read-only Go binding around an Ethereum contract.
type IL2ETHGatewayCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IL2ETHGatewayTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IL2ETHGatewayTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IL2ETHGatewayFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IL2ETHGatewayFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IL2ETHGatewaySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IL2ETHGatewaySession struct {
	Contract     *IL2ETHGateway    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IL2ETHGatewayCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IL2ETHGatewayCallerSession struct {
	Contract *IL2ETHGatewayCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// IL2ETHGatewayTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IL2ETHGatewayTransactorSession struct {
	Contract     *IL2ETHGatewayTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// IL2ETHGatewayRaw is an auto generated low-level Go binding around an Ethereum contract.
type IL2ETHGatewayRaw struct {
	Contract *IL2ETHGateway // Generic contract binding to access the raw methods on
}

// IL2ETHGatewayCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IL2ETHGatewayCallerRaw struct {
	Contract *IL2ETHGatewayCaller // Generic read-only contract binding to access the raw methods on
}

// IL2ETHGatewayTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IL2ETHGatewayTransactorRaw struct {
	Contract *IL2ETHGatewayTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIL2ETHGateway creates a new instance of IL2ETHGateway, bound to a specific deployed contract.
func NewIL2ETHGateway(address common.Address, backend bind.ContractBackend) (*IL2ETHGateway, error) {
	contract, err := bindIL2ETHGateway(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IL2ETHGateway{IL2ETHGatewayCaller: IL2ETHGatewayCaller{contract: contract}, IL2ETHGatewayTransactor: IL2ETHGatewayTransactor{contract: contract}, IL2ETHGatewayFilterer: IL2ETHGatewayFilterer{contract: contract}}, nil
}

// NewIL2ETHGatewayCaller creates a new read-only instance of IL2ETHGateway, bound to a specific deployed contract.
func NewIL2ETHGatewayCaller(address common.Address, caller bind.ContractCaller) (*IL2ETHGatewayCaller, error) {
	contract, err := bindIL2ETHGateway(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IL2ETHGatewayCaller{contract: contract}, nil
}

// NewIL2ETHGatewayTransactor creates a new write-only instance of IL2ETHGateway, bound to a specific deployed contract.
func NewIL2ETHGatewayTransactor(address common.Address, transactor bind.ContractTransactor) (*IL2ETHGatewayTransactor, error) {
	contract, err := bindIL2ETHGateway(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IL2ETHGatewayTransactor{contract: contract}, nil
}

// NewIL2ETHGatewayFilterer creates a new log filterer instance of IL2ETHGateway, bound to a specific deployed contract.
func NewIL2ETHGatewayFilterer(address common.Address, filterer bind.ContractFilterer) (*IL2ETHGatewayFilterer, error) {
	contract, err := bindIL2ETHGateway(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IL2ETHGatewayFilterer{contract: contract}, nil
}

// bindIL2ETHGateway binds a generic wrapper to an already deployed contract.
func bindIL2ETHGateway(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IL2ETHGatewayABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IL2ETHGateway *IL2ETHGatewayRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IL2ETHGateway.Contract.IL2ETHGatewayCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IL2ETHGateway *IL2ETHGatewayRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IL2ETHGateway.Contract.IL2ETHGatewayTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IL2ETHGateway *IL2ETHGatewayRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IL2ETHGateway.Contract.IL2ETHGatewayTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IL2ETHGateway *IL2ETHGatewayCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IL2ETHGateway.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IL2ETHGateway *IL2ETHGatewayTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IL2ETHGateway.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IL2ETHGateway *IL2ETHGatewayTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IL2ETHGateway.Contract.contract.Transact(opts, method, params...)
}

// FinalizeDepositETH is a paid mutator transaction binding the contract method 0x232e8748.
//
// Solidity: function finalizeDepositETH(address _from, address _to, uint256 _amount, bytes _data) payable returns()
func (_IL2ETHGateway *IL2ETHGatewayTransactor) FinalizeDepositETH(opts *bind.TransactOpts, _from common.Address, _to common.Address, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _IL2ETHGateway.contract.Transact(opts, "finalizeDepositETH", _from, _to, _amount, _data)
}

// FinalizeDepositETH is a paid mutator transaction binding the contract method 0x232e8748.
//
// Solidity: function finalizeDepositETH(address _from, address _to, uint256 _amount, bytes _data) payable returns()
func (_IL2ETHGateway *IL2ETHGatewaySession) FinalizeDepositETH(_from common.Address, _to common.Address, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _IL2ETHGateway.Contract.FinalizeDepositETH(&_IL2ETHGateway.TransactOpts, _from, _to, _amount, _data)
}

// FinalizeDepositETH is a paid mutator transaction binding the contract method 0x232e8748.
//
// Solidity: function finalizeDepositETH(address _from, address _to, uint256 _amount, bytes _data) payable returns()
func (_IL2ETHGateway *IL2ETHGatewayTransactorSession) FinalizeDepositETH(_from common.Address, _to common.Address, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _IL2ETHGateway.Contract.FinalizeDepositETH(&_IL2ETHGateway.TransactOpts, _from, _to, _amount, _data)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0x2fcc29fa.
//
// Solidity: function withdrawETH(address to, uint256 amount, uint256 gasLimit) payable returns()
func (_IL2ETHGateway *IL2ETHGatewayTransactor) WithdrawETH(opts *bind.TransactOpts, to common.Address, amount *big.Int, gasLimit *big.Int) (*types.Transaction, error) {
	return _IL2ETHGateway.contract.Transact(opts, "withdrawETH", to, amount, gasLimit)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0x2fcc29fa.
//
// Solidity: function withdrawETH(address to, uint256 amount, uint256 gasLimit) payable returns()
func (_IL2ETHGateway *IL2ETHGatewaySession) WithdrawETH(to common.Address, amount *big.Int, gasLimit *big.Int) (*types.Transaction, error) {
	return _IL2ETHGateway.Contract.WithdrawETH(&_IL2ETHGateway.TransactOpts, to, amount, gasLimit)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0x2fcc29fa.
//
// Solidity: function withdrawETH(address to, uint256 amount, uint256 gasLimit) payable returns()
func (_IL2ETHGateway *IL2ETHGatewayTransactorSession) WithdrawETH(to common.Address, amount *big.Int, gasLimit *big.Int) (*types.Transaction, error) {
	return _IL2ETHGateway.Contract.WithdrawETH(&_IL2ETHGateway.TransactOpts, to, amount, gasLimit)
}

// WithdrawETH0 is a paid mutator transaction binding the contract method 0xc7cdea37.
//
// Solidity: function withdrawETH(uint256 amount, uint256 gasLimit) payable returns()
func (_IL2ETHGateway *IL2ETHGatewayTransactor) WithdrawETH0(opts *bind.TransactOpts, amount *big.Int, gasLimit *big.Int) (*types.Transaction, error) {
	return _IL2ETHGateway.contract.Transact(opts, "withdrawETH0", amount, gasLimit)
}

// WithdrawETH0 is a paid mutator transaction binding the contract method 0xc7cdea37.
//
// Solidity: function withdrawETH(uint256 amount, uint256 gasLimit) payable returns()
func (_IL2ETHGateway *IL2ETHGatewaySession) WithdrawETH0(amount *big.Int, gasLimit *big.Int) (*types.Transaction, error) {
	return _IL2ETHGateway.Contract.WithdrawETH0(&_IL2ETHGateway.TransactOpts, amount, gasLimit)
}

// WithdrawETH0 is a paid mutator transaction binding the contract method 0xc7cdea37.
//
// Solidity: function withdrawETH(uint256 amount, uint256 gasLimit) payable returns()
func (_IL2ETHGateway *IL2ETHGatewayTransactorSession) WithdrawETH0(amount *big.Int, gasLimit *big.Int) (*types.Transaction, error) {
	return _IL2ETHGateway.Contract.WithdrawETH0(&_IL2ETHGateway.TransactOpts, amount, gasLimit)
}

// WithdrawETHAndCall is a paid mutator transaction binding the contract method 0x6dc24183.
//
// Solidity: function withdrawETHAndCall(address to, uint256 amount, bytes data, uint256 gasLimit) payable returns()
func (_IL2ETHGateway *IL2ETHGatewayTransactor) WithdrawETHAndCall(opts *bind.TransactOpts, to common.Address, amount *big.Int, data []byte, gasLimit *big.Int) (*types.Transaction, error) {
	return _IL2ETHGateway.contract.Transact(opts, "withdrawETHAndCall", to, amount, data, gasLimit)
}

// WithdrawETHAndCall is a paid mutator transaction binding the contract method 0x6dc24183.
//
// Solidity: function withdrawETHAndCall(address to, uint256 amount, bytes data, uint256 gasLimit) payable returns()
func (_IL2ETHGateway *IL2ETHGatewaySession) WithdrawETHAndCall(to common.Address, amount *big.Int, data []byte, gasLimit *big.Int) (*types.Transaction, error) {
	return _IL2ETHGateway.Contract.WithdrawETHAndCall(&_IL2ETHGateway.TransactOpts, to, amount, data, gasLimit)
}

// WithdrawETHAndCall is a paid mutator transaction binding the contract method 0x6dc24183.
//
// Solidity: function withdrawETHAndCall(address to, uint256 amount, bytes data, uint256 gasLimit) payable returns()
func (_IL2ETHGateway *IL2ETHGatewayTransactorSession) WithdrawETHAndCall(to common.Address, amount *big.Int, data []byte, gasLimit *big.Int) (*types.Transaction, error) {
	return _IL2ETHGateway.Contract.WithdrawETHAndCall(&_IL2ETHGateway.TransactOpts, to, amount, data, gasLimit)
}

// IL2ETHGatewayFinalizeDepositETHIterator is returned from FilterFinalizeDepositETH and is used to iterate over the raw logs and unpacked data for FinalizeDepositETH events raised by the IL2ETHGateway contract.
type IL2ETHGatewayFinalizeDepositETHIterator struct {
	Event *IL2ETHGatewayFinalizeDepositETH // Event containing the contract specifics and raw log

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
func (it *IL2ETHGatewayFinalizeDepositETHIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IL2ETHGatewayFinalizeDepositETH)
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
		it.Event = new(IL2ETHGatewayFinalizeDepositETH)
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
func (it *IL2ETHGatewayFinalizeDepositETHIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IL2ETHGatewayFinalizeDepositETHIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IL2ETHGatewayFinalizeDepositETH represents a FinalizeDepositETH event raised by the IL2ETHGateway contract.
type IL2ETHGatewayFinalizeDepositETH struct {
	From   common.Address
	To     common.Address
	Amount *big.Int
	Data   []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFinalizeDepositETH is a free log retrieval operation binding the contract event 0x9e86c356e14e24e26e3ce769bf8b87de38e0faa0ed0ca946fa09659aa606bd2d.
//
// Solidity: event FinalizeDepositETH(address indexed from, address indexed to, uint256 amount, bytes data)
func (_IL2ETHGateway *IL2ETHGatewayFilterer) FilterFinalizeDepositETH(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*IL2ETHGatewayFinalizeDepositETHIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IL2ETHGateway.contract.FilterLogs(opts, "FinalizeDepositETH", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IL2ETHGatewayFinalizeDepositETHIterator{contract: _IL2ETHGateway.contract, event: "FinalizeDepositETH", logs: logs, sub: sub}, nil
}

// WatchFinalizeDepositETH is a free log subscription operation binding the contract event 0x9e86c356e14e24e26e3ce769bf8b87de38e0faa0ed0ca946fa09659aa606bd2d.
//
// Solidity: event FinalizeDepositETH(address indexed from, address indexed to, uint256 amount, bytes data)
func (_IL2ETHGateway *IL2ETHGatewayFilterer) WatchFinalizeDepositETH(opts *bind.WatchOpts, sink chan<- *IL2ETHGatewayFinalizeDepositETH, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IL2ETHGateway.contract.WatchLogs(opts, "FinalizeDepositETH", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IL2ETHGatewayFinalizeDepositETH)
				if err := _IL2ETHGateway.contract.UnpackLog(event, "FinalizeDepositETH", log); err != nil {
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

// ParseFinalizeDepositETH is a log parse operation binding the contract event 0x9e86c356e14e24e26e3ce769bf8b87de38e0faa0ed0ca946fa09659aa606bd2d.
//
// Solidity: event FinalizeDepositETH(address indexed from, address indexed to, uint256 amount, bytes data)
func (_IL2ETHGateway *IL2ETHGatewayFilterer) ParseFinalizeDepositETH(log types.Log) (*IL2ETHGatewayFinalizeDepositETH, error) {
	event := new(IL2ETHGatewayFinalizeDepositETH)
	if err := _IL2ETHGateway.contract.UnpackLog(event, "FinalizeDepositETH", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IL2ETHGatewayWithdrawETHIterator is returned from FilterWithdrawETH and is used to iterate over the raw logs and unpacked data for WithdrawETH events raised by the IL2ETHGateway contract.
type IL2ETHGatewayWithdrawETHIterator struct {
	Event *IL2ETHGatewayWithdrawETH // Event containing the contract specifics and raw log

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
func (it *IL2ETHGatewayWithdrawETHIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IL2ETHGatewayWithdrawETH)
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
		it.Event = new(IL2ETHGatewayWithdrawETH)
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
func (it *IL2ETHGatewayWithdrawETHIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IL2ETHGatewayWithdrawETHIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IL2ETHGatewayWithdrawETH represents a WithdrawETH event raised by the IL2ETHGateway contract.
type IL2ETHGatewayWithdrawETH struct {
	From   common.Address
	To     common.Address
	Amount *big.Int
	Data   []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawETH is a free log retrieval operation binding the contract event 0xd8ed6eaa9a7a8980d7901e911fde6686810b989d3082182d1d3a3df6306ce20e.
//
// Solidity: event WithdrawETH(address indexed from, address indexed to, uint256 amount, bytes data)
func (_IL2ETHGateway *IL2ETHGatewayFilterer) FilterWithdrawETH(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*IL2ETHGatewayWithdrawETHIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IL2ETHGateway.contract.FilterLogs(opts, "WithdrawETH", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IL2ETHGatewayWithdrawETHIterator{contract: _IL2ETHGateway.contract, event: "WithdrawETH", logs: logs, sub: sub}, nil
}

// WatchWithdrawETH is a free log subscription operation binding the contract event 0xd8ed6eaa9a7a8980d7901e911fde6686810b989d3082182d1d3a3df6306ce20e.
//
// Solidity: event WithdrawETH(address indexed from, address indexed to, uint256 amount, bytes data)
func (_IL2ETHGateway *IL2ETHGatewayFilterer) WatchWithdrawETH(opts *bind.WatchOpts, sink chan<- *IL2ETHGatewayWithdrawETH, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IL2ETHGateway.contract.WatchLogs(opts, "WithdrawETH", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IL2ETHGatewayWithdrawETH)
				if err := _IL2ETHGateway.contract.UnpackLog(event, "WithdrawETH", log); err != nil {
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

// ParseWithdrawETH is a log parse operation binding the contract event 0xd8ed6eaa9a7a8980d7901e911fde6686810b989d3082182d1d3a3df6306ce20e.
//
// Solidity: event WithdrawETH(address indexed from, address indexed to, uint256 amount, bytes data)
func (_IL2ETHGateway *IL2ETHGatewayFilterer) ParseWithdrawETH(log types.Log) (*IL2ETHGatewayWithdrawETH, error) {
	event := new(IL2ETHGatewayWithdrawETH)
	if err := _IL2ETHGateway.contract.UnpackLog(event, "WithdrawETH", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
