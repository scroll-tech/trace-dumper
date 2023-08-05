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

// IL2ERC20GatewayMetaData contains all meta data concerning the IL2ERC20Gateway contract.
var IL2ERC20GatewayMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l2Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"FinalizeDepositERC20\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l2Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"WithdrawERC20\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"l2Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"finalizeDepositERC20\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"l2Token\",\"type\":\"address\"}],\"name\":\"getL1ERC20Address\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\"}],\"name\":\"getL2ERC20Address\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"}],\"name\":\"withdrawERC20\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"}],\"name\":\"withdrawERC20\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"}],\"name\":\"withdrawERC20AndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"8431f5c1": "finalizeDepositERC20(address,address,address,address,uint256,bytes)",
		"54bbd59c": "getL1ERC20Address(address)",
		"c676ad29": "getL2ERC20Address(address)",
		"a93a4af9": "withdrawERC20(address,address,uint256,uint256)",
		"6c07ea43": "withdrawERC20(address,uint256,uint256)",
		"575361b6": "withdrawERC20AndCall(address,address,uint256,bytes,uint256)",
	},
}

// IL2ERC20GatewayABI is the input ABI used to generate the binding from.
// Deprecated: Use IL2ERC20GatewayMetaData.ABI instead.
var IL2ERC20GatewayABI = IL2ERC20GatewayMetaData.ABI

// Deprecated: Use IL2ERC20GatewayMetaData.Sigs instead.
// IL2ERC20GatewayFuncSigs maps the 4-byte function signature to its string representation.
var IL2ERC20GatewayFuncSigs = IL2ERC20GatewayMetaData.Sigs

// IL2ERC20Gateway is an auto generated Go binding around an Ethereum contract.
type IL2ERC20Gateway struct {
	IL2ERC20GatewayCaller     // Read-only binding to the contract
	IL2ERC20GatewayTransactor // Write-only binding to the contract
	IL2ERC20GatewayFilterer   // Log filterer for contract events
}

// IL2ERC20GatewayCaller is an auto generated read-only Go binding around an Ethereum contract.
type IL2ERC20GatewayCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IL2ERC20GatewayTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IL2ERC20GatewayTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IL2ERC20GatewayFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IL2ERC20GatewayFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IL2ERC20GatewaySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IL2ERC20GatewaySession struct {
	Contract     *IL2ERC20Gateway  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IL2ERC20GatewayCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IL2ERC20GatewayCallerSession struct {
	Contract *IL2ERC20GatewayCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// IL2ERC20GatewayTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IL2ERC20GatewayTransactorSession struct {
	Contract     *IL2ERC20GatewayTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// IL2ERC20GatewayRaw is an auto generated low-level Go binding around an Ethereum contract.
type IL2ERC20GatewayRaw struct {
	Contract *IL2ERC20Gateway // Generic contract binding to access the raw methods on
}

// IL2ERC20GatewayCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IL2ERC20GatewayCallerRaw struct {
	Contract *IL2ERC20GatewayCaller // Generic read-only contract binding to access the raw methods on
}

// IL2ERC20GatewayTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IL2ERC20GatewayTransactorRaw struct {
	Contract *IL2ERC20GatewayTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIL2ERC20Gateway creates a new instance of IL2ERC20Gateway, bound to a specific deployed contract.
func NewIL2ERC20Gateway(address common.Address, backend bind.ContractBackend) (*IL2ERC20Gateway, error) {
	contract, err := bindIL2ERC20Gateway(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IL2ERC20Gateway{IL2ERC20GatewayCaller: IL2ERC20GatewayCaller{contract: contract}, IL2ERC20GatewayTransactor: IL2ERC20GatewayTransactor{contract: contract}, IL2ERC20GatewayFilterer: IL2ERC20GatewayFilterer{contract: contract}}, nil
}

// NewIL2ERC20GatewayCaller creates a new read-only instance of IL2ERC20Gateway, bound to a specific deployed contract.
func NewIL2ERC20GatewayCaller(address common.Address, caller bind.ContractCaller) (*IL2ERC20GatewayCaller, error) {
	contract, err := bindIL2ERC20Gateway(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IL2ERC20GatewayCaller{contract: contract}, nil
}

// NewIL2ERC20GatewayTransactor creates a new write-only instance of IL2ERC20Gateway, bound to a specific deployed contract.
func NewIL2ERC20GatewayTransactor(address common.Address, transactor bind.ContractTransactor) (*IL2ERC20GatewayTransactor, error) {
	contract, err := bindIL2ERC20Gateway(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IL2ERC20GatewayTransactor{contract: contract}, nil
}

// NewIL2ERC20GatewayFilterer creates a new log filterer instance of IL2ERC20Gateway, bound to a specific deployed contract.
func NewIL2ERC20GatewayFilterer(address common.Address, filterer bind.ContractFilterer) (*IL2ERC20GatewayFilterer, error) {
	contract, err := bindIL2ERC20Gateway(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IL2ERC20GatewayFilterer{contract: contract}, nil
}

// bindIL2ERC20Gateway binds a generic wrapper to an already deployed contract.
func bindIL2ERC20Gateway(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IL2ERC20GatewayABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IL2ERC20Gateway *IL2ERC20GatewayRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IL2ERC20Gateway.Contract.IL2ERC20GatewayCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IL2ERC20Gateway *IL2ERC20GatewayRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IL2ERC20Gateway.Contract.IL2ERC20GatewayTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IL2ERC20Gateway *IL2ERC20GatewayRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IL2ERC20Gateway.Contract.IL2ERC20GatewayTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IL2ERC20Gateway *IL2ERC20GatewayCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IL2ERC20Gateway.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IL2ERC20Gateway *IL2ERC20GatewayTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IL2ERC20Gateway.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IL2ERC20Gateway *IL2ERC20GatewayTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IL2ERC20Gateway.Contract.contract.Transact(opts, method, params...)
}

// GetL1ERC20Address is a free data retrieval call binding the contract method 0x54bbd59c.
//
// Solidity: function getL1ERC20Address(address l2Token) view returns(address)
func (_IL2ERC20Gateway *IL2ERC20GatewayCaller) GetL1ERC20Address(opts *bind.CallOpts, l2Token common.Address) (common.Address, error) {
	var out []interface{}
	err := _IL2ERC20Gateway.contract.Call(opts, &out, "getL1ERC20Address", l2Token)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetL1ERC20Address is a free data retrieval call binding the contract method 0x54bbd59c.
//
// Solidity: function getL1ERC20Address(address l2Token) view returns(address)
func (_IL2ERC20Gateway *IL2ERC20GatewaySession) GetL1ERC20Address(l2Token common.Address) (common.Address, error) {
	return _IL2ERC20Gateway.Contract.GetL1ERC20Address(&_IL2ERC20Gateway.CallOpts, l2Token)
}

// GetL1ERC20Address is a free data retrieval call binding the contract method 0x54bbd59c.
//
// Solidity: function getL1ERC20Address(address l2Token) view returns(address)
func (_IL2ERC20Gateway *IL2ERC20GatewayCallerSession) GetL1ERC20Address(l2Token common.Address) (common.Address, error) {
	return _IL2ERC20Gateway.Contract.GetL1ERC20Address(&_IL2ERC20Gateway.CallOpts, l2Token)
}

// GetL2ERC20Address is a free data retrieval call binding the contract method 0xc676ad29.
//
// Solidity: function getL2ERC20Address(address l1Token) view returns(address)
func (_IL2ERC20Gateway *IL2ERC20GatewayCaller) GetL2ERC20Address(opts *bind.CallOpts, l1Token common.Address) (common.Address, error) {
	var out []interface{}
	err := _IL2ERC20Gateway.contract.Call(opts, &out, "getL2ERC20Address", l1Token)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetL2ERC20Address is a free data retrieval call binding the contract method 0xc676ad29.
//
// Solidity: function getL2ERC20Address(address l1Token) view returns(address)
func (_IL2ERC20Gateway *IL2ERC20GatewaySession) GetL2ERC20Address(l1Token common.Address) (common.Address, error) {
	return _IL2ERC20Gateway.Contract.GetL2ERC20Address(&_IL2ERC20Gateway.CallOpts, l1Token)
}

// GetL2ERC20Address is a free data retrieval call binding the contract method 0xc676ad29.
//
// Solidity: function getL2ERC20Address(address l1Token) view returns(address)
func (_IL2ERC20Gateway *IL2ERC20GatewayCallerSession) GetL2ERC20Address(l1Token common.Address) (common.Address, error) {
	return _IL2ERC20Gateway.Contract.GetL2ERC20Address(&_IL2ERC20Gateway.CallOpts, l1Token)
}

// FinalizeDepositERC20 is a paid mutator transaction binding the contract method 0x8431f5c1.
//
// Solidity: function finalizeDepositERC20(address l1Token, address l2Token, address from, address to, uint256 amount, bytes data) payable returns()
func (_IL2ERC20Gateway *IL2ERC20GatewayTransactor) FinalizeDepositERC20(opts *bind.TransactOpts, l1Token common.Address, l2Token common.Address, from common.Address, to common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _IL2ERC20Gateway.contract.Transact(opts, "finalizeDepositERC20", l1Token, l2Token, from, to, amount, data)
}

// FinalizeDepositERC20 is a paid mutator transaction binding the contract method 0x8431f5c1.
//
// Solidity: function finalizeDepositERC20(address l1Token, address l2Token, address from, address to, uint256 amount, bytes data) payable returns()
func (_IL2ERC20Gateway *IL2ERC20GatewaySession) FinalizeDepositERC20(l1Token common.Address, l2Token common.Address, from common.Address, to common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _IL2ERC20Gateway.Contract.FinalizeDepositERC20(&_IL2ERC20Gateway.TransactOpts, l1Token, l2Token, from, to, amount, data)
}

// FinalizeDepositERC20 is a paid mutator transaction binding the contract method 0x8431f5c1.
//
// Solidity: function finalizeDepositERC20(address l1Token, address l2Token, address from, address to, uint256 amount, bytes data) payable returns()
func (_IL2ERC20Gateway *IL2ERC20GatewayTransactorSession) FinalizeDepositERC20(l1Token common.Address, l2Token common.Address, from common.Address, to common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _IL2ERC20Gateway.Contract.FinalizeDepositERC20(&_IL2ERC20Gateway.TransactOpts, l1Token, l2Token, from, to, amount, data)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0x6c07ea43.
//
// Solidity: function withdrawERC20(address token, uint256 amount, uint256 gasLimit) payable returns()
func (_IL2ERC20Gateway *IL2ERC20GatewayTransactor) WithdrawERC20(opts *bind.TransactOpts, token common.Address, amount *big.Int, gasLimit *big.Int) (*types.Transaction, error) {
	return _IL2ERC20Gateway.contract.Transact(opts, "withdrawERC20", token, amount, gasLimit)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0x6c07ea43.
//
// Solidity: function withdrawERC20(address token, uint256 amount, uint256 gasLimit) payable returns()
func (_IL2ERC20Gateway *IL2ERC20GatewaySession) WithdrawERC20(token common.Address, amount *big.Int, gasLimit *big.Int) (*types.Transaction, error) {
	return _IL2ERC20Gateway.Contract.WithdrawERC20(&_IL2ERC20Gateway.TransactOpts, token, amount, gasLimit)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0x6c07ea43.
//
// Solidity: function withdrawERC20(address token, uint256 amount, uint256 gasLimit) payable returns()
func (_IL2ERC20Gateway *IL2ERC20GatewayTransactorSession) WithdrawERC20(token common.Address, amount *big.Int, gasLimit *big.Int) (*types.Transaction, error) {
	return _IL2ERC20Gateway.Contract.WithdrawERC20(&_IL2ERC20Gateway.TransactOpts, token, amount, gasLimit)
}

// WithdrawERC200 is a paid mutator transaction binding the contract method 0xa93a4af9.
//
// Solidity: function withdrawERC20(address token, address to, uint256 amount, uint256 gasLimit) payable returns()
func (_IL2ERC20Gateway *IL2ERC20GatewayTransactor) WithdrawERC200(opts *bind.TransactOpts, token common.Address, to common.Address, amount *big.Int, gasLimit *big.Int) (*types.Transaction, error) {
	return _IL2ERC20Gateway.contract.Transact(opts, "withdrawERC200", token, to, amount, gasLimit)
}

// WithdrawERC200 is a paid mutator transaction binding the contract method 0xa93a4af9.
//
// Solidity: function withdrawERC20(address token, address to, uint256 amount, uint256 gasLimit) payable returns()
func (_IL2ERC20Gateway *IL2ERC20GatewaySession) WithdrawERC200(token common.Address, to common.Address, amount *big.Int, gasLimit *big.Int) (*types.Transaction, error) {
	return _IL2ERC20Gateway.Contract.WithdrawERC200(&_IL2ERC20Gateway.TransactOpts, token, to, amount, gasLimit)
}

// WithdrawERC200 is a paid mutator transaction binding the contract method 0xa93a4af9.
//
// Solidity: function withdrawERC20(address token, address to, uint256 amount, uint256 gasLimit) payable returns()
func (_IL2ERC20Gateway *IL2ERC20GatewayTransactorSession) WithdrawERC200(token common.Address, to common.Address, amount *big.Int, gasLimit *big.Int) (*types.Transaction, error) {
	return _IL2ERC20Gateway.Contract.WithdrawERC200(&_IL2ERC20Gateway.TransactOpts, token, to, amount, gasLimit)
}

// WithdrawERC20AndCall is a paid mutator transaction binding the contract method 0x575361b6.
//
// Solidity: function withdrawERC20AndCall(address token, address to, uint256 amount, bytes data, uint256 gasLimit) payable returns()
func (_IL2ERC20Gateway *IL2ERC20GatewayTransactor) WithdrawERC20AndCall(opts *bind.TransactOpts, token common.Address, to common.Address, amount *big.Int, data []byte, gasLimit *big.Int) (*types.Transaction, error) {
	return _IL2ERC20Gateway.contract.Transact(opts, "withdrawERC20AndCall", token, to, amount, data, gasLimit)
}

// WithdrawERC20AndCall is a paid mutator transaction binding the contract method 0x575361b6.
//
// Solidity: function withdrawERC20AndCall(address token, address to, uint256 amount, bytes data, uint256 gasLimit) payable returns()
func (_IL2ERC20Gateway *IL2ERC20GatewaySession) WithdrawERC20AndCall(token common.Address, to common.Address, amount *big.Int, data []byte, gasLimit *big.Int) (*types.Transaction, error) {
	return _IL2ERC20Gateway.Contract.WithdrawERC20AndCall(&_IL2ERC20Gateway.TransactOpts, token, to, amount, data, gasLimit)
}

// WithdrawERC20AndCall is a paid mutator transaction binding the contract method 0x575361b6.
//
// Solidity: function withdrawERC20AndCall(address token, address to, uint256 amount, bytes data, uint256 gasLimit) payable returns()
func (_IL2ERC20Gateway *IL2ERC20GatewayTransactorSession) WithdrawERC20AndCall(token common.Address, to common.Address, amount *big.Int, data []byte, gasLimit *big.Int) (*types.Transaction, error) {
	return _IL2ERC20Gateway.Contract.WithdrawERC20AndCall(&_IL2ERC20Gateway.TransactOpts, token, to, amount, data, gasLimit)
}

// IL2ERC20GatewayFinalizeDepositERC20Iterator is returned from FilterFinalizeDepositERC20 and is used to iterate over the raw logs and unpacked data for FinalizeDepositERC20 events raised by the IL2ERC20Gateway contract.
type IL2ERC20GatewayFinalizeDepositERC20Iterator struct {
	Event *IL2ERC20GatewayFinalizeDepositERC20 // Event containing the contract specifics and raw log

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
func (it *IL2ERC20GatewayFinalizeDepositERC20Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IL2ERC20GatewayFinalizeDepositERC20)
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
		it.Event = new(IL2ERC20GatewayFinalizeDepositERC20)
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
func (it *IL2ERC20GatewayFinalizeDepositERC20Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IL2ERC20GatewayFinalizeDepositERC20Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IL2ERC20GatewayFinalizeDepositERC20 represents a FinalizeDepositERC20 event raised by the IL2ERC20Gateway contract.
type IL2ERC20GatewayFinalizeDepositERC20 struct {
	L1Token common.Address
	L2Token common.Address
	From    common.Address
	To      common.Address
	Amount  *big.Int
	Data    []byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterFinalizeDepositERC20 is a free log retrieval operation binding the contract event 0x165ba69f6ab40c50cade6f65431801e5f9c7d7830b7545391920db039133ba34.
//
// Solidity: event FinalizeDepositERC20(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256 amount, bytes data)
func (_IL2ERC20Gateway *IL2ERC20GatewayFilterer) FilterFinalizeDepositERC20(opts *bind.FilterOpts, l1Token []common.Address, l2Token []common.Address, from []common.Address) (*IL2ERC20GatewayFinalizeDepositERC20Iterator, error) {

	var l1TokenRule []interface{}
	for _, l1TokenItem := range l1Token {
		l1TokenRule = append(l1TokenRule, l1TokenItem)
	}
	var l2TokenRule []interface{}
	for _, l2TokenItem := range l2Token {
		l2TokenRule = append(l2TokenRule, l2TokenItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _IL2ERC20Gateway.contract.FilterLogs(opts, "FinalizeDepositERC20", l1TokenRule, l2TokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &IL2ERC20GatewayFinalizeDepositERC20Iterator{contract: _IL2ERC20Gateway.contract, event: "FinalizeDepositERC20", logs: logs, sub: sub}, nil
}

// WatchFinalizeDepositERC20 is a free log subscription operation binding the contract event 0x165ba69f6ab40c50cade6f65431801e5f9c7d7830b7545391920db039133ba34.
//
// Solidity: event FinalizeDepositERC20(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256 amount, bytes data)
func (_IL2ERC20Gateway *IL2ERC20GatewayFilterer) WatchFinalizeDepositERC20(opts *bind.WatchOpts, sink chan<- *IL2ERC20GatewayFinalizeDepositERC20, l1Token []common.Address, l2Token []common.Address, from []common.Address) (event.Subscription, error) {

	var l1TokenRule []interface{}
	for _, l1TokenItem := range l1Token {
		l1TokenRule = append(l1TokenRule, l1TokenItem)
	}
	var l2TokenRule []interface{}
	for _, l2TokenItem := range l2Token {
		l2TokenRule = append(l2TokenRule, l2TokenItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _IL2ERC20Gateway.contract.WatchLogs(opts, "FinalizeDepositERC20", l1TokenRule, l2TokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IL2ERC20GatewayFinalizeDepositERC20)
				if err := _IL2ERC20Gateway.contract.UnpackLog(event, "FinalizeDepositERC20", log); err != nil {
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

// ParseFinalizeDepositERC20 is a log parse operation binding the contract event 0x165ba69f6ab40c50cade6f65431801e5f9c7d7830b7545391920db039133ba34.
//
// Solidity: event FinalizeDepositERC20(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256 amount, bytes data)
func (_IL2ERC20Gateway *IL2ERC20GatewayFilterer) ParseFinalizeDepositERC20(log types.Log) (*IL2ERC20GatewayFinalizeDepositERC20, error) {
	event := new(IL2ERC20GatewayFinalizeDepositERC20)
	if err := _IL2ERC20Gateway.contract.UnpackLog(event, "FinalizeDepositERC20", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IL2ERC20GatewayWithdrawERC20Iterator is returned from FilterWithdrawERC20 and is used to iterate over the raw logs and unpacked data for WithdrawERC20 events raised by the IL2ERC20Gateway contract.
type IL2ERC20GatewayWithdrawERC20Iterator struct {
	Event *IL2ERC20GatewayWithdrawERC20 // Event containing the contract specifics and raw log

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
func (it *IL2ERC20GatewayWithdrawERC20Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IL2ERC20GatewayWithdrawERC20)
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
		it.Event = new(IL2ERC20GatewayWithdrawERC20)
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
func (it *IL2ERC20GatewayWithdrawERC20Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IL2ERC20GatewayWithdrawERC20Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IL2ERC20GatewayWithdrawERC20 represents a WithdrawERC20 event raised by the IL2ERC20Gateway contract.
type IL2ERC20GatewayWithdrawERC20 struct {
	L1Token common.Address
	L2Token common.Address
	From    common.Address
	To      common.Address
	Amount  *big.Int
	Data    []byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWithdrawERC20 is a free log retrieval operation binding the contract event 0xd8d3a3f4ab95694bef40475997598bcf8acd3ed9617a4c1013795429414c27e8.
//
// Solidity: event WithdrawERC20(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256 amount, bytes data)
func (_IL2ERC20Gateway *IL2ERC20GatewayFilterer) FilterWithdrawERC20(opts *bind.FilterOpts, l1Token []common.Address, l2Token []common.Address, from []common.Address) (*IL2ERC20GatewayWithdrawERC20Iterator, error) {

	var l1TokenRule []interface{}
	for _, l1TokenItem := range l1Token {
		l1TokenRule = append(l1TokenRule, l1TokenItem)
	}
	var l2TokenRule []interface{}
	for _, l2TokenItem := range l2Token {
		l2TokenRule = append(l2TokenRule, l2TokenItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _IL2ERC20Gateway.contract.FilterLogs(opts, "WithdrawERC20", l1TokenRule, l2TokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &IL2ERC20GatewayWithdrawERC20Iterator{contract: _IL2ERC20Gateway.contract, event: "WithdrawERC20", logs: logs, sub: sub}, nil
}

// WatchWithdrawERC20 is a free log subscription operation binding the contract event 0xd8d3a3f4ab95694bef40475997598bcf8acd3ed9617a4c1013795429414c27e8.
//
// Solidity: event WithdrawERC20(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256 amount, bytes data)
func (_IL2ERC20Gateway *IL2ERC20GatewayFilterer) WatchWithdrawERC20(opts *bind.WatchOpts, sink chan<- *IL2ERC20GatewayWithdrawERC20, l1Token []common.Address, l2Token []common.Address, from []common.Address) (event.Subscription, error) {

	var l1TokenRule []interface{}
	for _, l1TokenItem := range l1Token {
		l1TokenRule = append(l1TokenRule, l1TokenItem)
	}
	var l2TokenRule []interface{}
	for _, l2TokenItem := range l2Token {
		l2TokenRule = append(l2TokenRule, l2TokenItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _IL2ERC20Gateway.contract.WatchLogs(opts, "WithdrawERC20", l1TokenRule, l2TokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IL2ERC20GatewayWithdrawERC20)
				if err := _IL2ERC20Gateway.contract.UnpackLog(event, "WithdrawERC20", log); err != nil {
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

// ParseWithdrawERC20 is a log parse operation binding the contract event 0xd8d3a3f4ab95694bef40475997598bcf8acd3ed9617a4c1013795429414c27e8.
//
// Solidity: event WithdrawERC20(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256 amount, bytes data)
func (_IL2ERC20Gateway *IL2ERC20GatewayFilterer) ParseWithdrawERC20(log types.Log) (*IL2ERC20GatewayWithdrawERC20, error) {
	event := new(IL2ERC20GatewayWithdrawERC20)
	if err := _IL2ERC20Gateway.contract.UnpackLog(event, "WithdrawERC20", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
