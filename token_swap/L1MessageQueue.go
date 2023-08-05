// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package token_swap

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

// L1MessageQueueMetaData contains all meta data concerning the L1MessageQueue contract.
var L1MessageQueueMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"queueIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"QueueTransaction\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_oldGasOracle\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_newGasOracle\",\"type\":\"address\"}],\"name\":\"UpdateGasOracle\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"appendCrossDomainMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"appendEnforcedTransaction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"}],\"name\":\"estimateCrossDomainMessageFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gasOracle\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_queueIndex\",\"type\":\"uint256\"}],\"name\":\"getCrossDomainMessage\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_messenger\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_gasOracle\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"messageQueue\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messenger\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextCrossDomainMessageIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newGasOracle\",\"type\":\"address\"}],\"name\":\"updateGasOracle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"9b159782": "appendCrossDomainMessage(address,uint256,bytes)",
		"bdc6f0a0": "appendEnforcedTransaction(address,address,uint256,uint256,bytes)",
		"9119a52b": "estimateCrossDomainMessageFee(address,address,bytes,uint256)",
		"5d62a8dd": "gasOracle()",
		"ae453cd5": "getCrossDomainMessage(uint256)",
		"485cc955": "initialize(address,address)",
		"29aa604b": "messageQueue(uint256)",
		"3cb747bf": "messenger()",
		"fd0ad31e": "nextCrossDomainMessageIndex()",
		"8da5cb5b": "owner()",
		"715018a6": "renounceOwnership()",
		"f2fde38b": "transferOwnership(address)",
		"70cee67f": "updateGasOracle(address)",
	},
	Bin: "0x608060405234801561001057600080fd5b50610bab806100206000396000f3fe608060405234801561001057600080fd5b50600436106100cf5760003560e01c80638da5cb5b1161008c578063ae453cd511610066578063ae453cd51461019f578063bdc6f0a0146101b2578063f2fde38b146101c8578063fd0ad31e146101db57600080fd5b80638da5cb5b146101685780639119a52b146101795780639b1597821461018c57600080fd5b806329aa604b146100d45780633cb747bf146100fa578063485cc955146101255780635d62a8dd1461013a57806370cee67f1461014d578063715018a614610160575b600080fd5b6100e76100e236600461071d565b6101e3565b6040519081526020015b60405180910390f35b60655461010d906001600160a01b031681565b6040516001600160a01b0390911681526020016100f1565b610138610133366004610752565b610204565b005b60665461010d906001600160a01b031681565b61013861015b366004610785565b6102fc565b610138610387565b6033546001600160a01b031661010d565b6100e76101873660046107bd565b6103bd565b61013861019a3660046108e0565b61045a565b6100e76101ad36600461071d565b6105aa565b6101386101c036600461093a565b505050505050565b6101386101d6366004610785565b6105d1565b6067546100e7565b606781815481106101f357600080fd5b600091825260209091200154905081565b600054610100900460ff1661021f5760005460ff1615610223565b303b155b61028b5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084015b60405180910390fd5b600054610100900460ff161580156102ad576000805461ffff19166101011790555b6102b561066c565b606580546001600160a01b038086166001600160a01b031992831617909255606680549285169290911691909117905580156102f7576000805461ff00191690555b505050565b6033546001600160a01b031633146103265760405162461bcd60e51b8152600401610282906109b2565b606680546001600160a01b038381166001600160a01b031983168117909355604080519190921680825260208201939093527f9ed5ec28f252b3e7f62f1ace8e54c5ebabf4c61cc2a7c33a806365b2ff7ecc5e910160405180910390a15050565b6033546001600160a01b031633146103b15760405162461bcd60e51b8152600401610282906109b2565b6103bb600061069b565b565b6066546000906001600160a01b0316806103db576000915050610452565b604051639119a52b60e01b81526001600160a01b03821690639119a52b9061040d9089908990899089906004016109e7565b602060405180830381865afa15801561042a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061044e9190610a58565b9150505b949350505050565b6065546001600160a01b031633146104c35760405162461bcd60e51b815260206004820152602660248201527f4f6e6c792063616c6c61626c6520627920746865204c315363726f6c6c4d657360448201526539b2b733b2b960d11b6064820152608401610282565b6040513373111100000000000000000000000000000000111101906000906104f990839088908490899089908990602001610a9a565b60405160208183030381529060405280519060200120905060006067805490509050866001600160a01b0316836001600160a01b03167fbdcc7517f8fe3db6506dfd910942d0bbecaf3d6a506dadea65b0d988e75b94396000848a8a8a604051610567959493929190610ae3565b60405180910390a350606780546001810182556000919091527f9787eeb91fe3101235e4a76063c7023ecb40f923f97916639c598592fa30d6ae01555050505050565b6000606782815481106105bf576105bf610b14565b90600052602060002001549050919050565b6033546001600160a01b031633146105fb5760405162461bcd60e51b8152600401610282906109b2565b6001600160a01b0381166106605760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610282565b6106698161069b565b50565b600054610100900460ff166106935760405162461bcd60e51b815260040161028290610b2a565b6103bb6106ed565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b600054610100900460ff166107145760405162461bcd60e51b815260040161028290610b2a565b6103bb3361069b565b60006020828403121561072f57600080fd5b5035919050565b80356001600160a01b038116811461074d57600080fd5b919050565b6000806040838503121561076557600080fd5b61076e83610736565b915061077c60208401610736565b90509250929050565b60006020828403121561079757600080fd5b6107a082610736565b9392505050565b634e487b7160e01b600052604160045260246000fd5b600080600080608085870312156107d357600080fd5b6107dc85610736565b93506107ea60208601610736565b9250604085013567ffffffffffffffff8082111561080757600080fd5b818701915087601f83011261081b57600080fd5b81358181111561082d5761082d6107a7565b604051601f8201601f19908116603f01168101908382118183101715610855576108556107a7565b816040528281528a602084870101111561086e57600080fd5b826020860160208301376000928101602001929092525095989497509495606001359450505050565b60008083601f8401126108a957600080fd5b50813567ffffffffffffffff8111156108c157600080fd5b6020830191508360208285010111156108d957600080fd5b9250929050565b600080600080606085870312156108f657600080fd5b6108ff85610736565b935060208501359250604085013567ffffffffffffffff81111561092257600080fd5b61092e87828801610897565b95989497509550505050565b60008060008060008060a0878903121561095357600080fd5b61095c87610736565b955061096a60208801610736565b94506040870135935060608701359250608087013567ffffffffffffffff81111561099457600080fd5b6109a089828a01610897565b979a9699509497509295939492505050565b6020808252818101527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604082015260600190565b600060018060a01b038087168352602081871681850152608060408501528551915081608085015260005b82811015610a2e5786810182015185820160a001528101610a12565b5050600060a0828501015260a0601f19601f83011684010191505082606083015295945050505050565b600060208284031215610a6a57600080fd5b5051919050565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b6001600160a01b0387811682528616602082015260ff851660408201526060810184905260a060808201819052600090610ad79083018486610a71565b98975050505050505050565b858152846020820152836040820152608060608201526000610b09608083018486610a71565b979650505050505050565b634e487b7160e01b600052603260045260246000fd5b6020808252602b908201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960408201526a6e697469616c697a696e6760a81b60608201526080019056fea26469706673582212207ef5bf54099dff0d915860851ec10cfda793e78ca09931766d371e499675fd4d64736f6c63430008110033",
}

// L1MessageQueueABI is the input ABI used to generate the binding from.
// Deprecated: Use L1MessageQueueMetaData.ABI instead.
var L1MessageQueueABI = L1MessageQueueMetaData.ABI

// Deprecated: Use L1MessageQueueMetaData.Sigs instead.
// L1MessageQueueFuncSigs maps the 4-byte function signature to its string representation.
var L1MessageQueueFuncSigs = L1MessageQueueMetaData.Sigs

// L1MessageQueueBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use L1MessageQueueMetaData.Bin instead.
var L1MessageQueueBin = L1MessageQueueMetaData.Bin

// DeployL1MessageQueue deploys a new Ethereum contract, binding an instance of L1MessageQueue to it.
func DeployL1MessageQueue(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *L1MessageQueue, error) {
	parsed, err := L1MessageQueueMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(L1MessageQueueBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &L1MessageQueue{L1MessageQueueCaller: L1MessageQueueCaller{contract: contract}, L1MessageQueueTransactor: L1MessageQueueTransactor{contract: contract}, L1MessageQueueFilterer: L1MessageQueueFilterer{contract: contract}}, nil
}

// L1MessageQueue is an auto generated Go binding around an Ethereum contract.
type L1MessageQueue struct {
	L1MessageQueueCaller     // Read-only binding to the contract
	L1MessageQueueTransactor // Write-only binding to the contract
	L1MessageQueueFilterer   // Log filterer for contract events
}

// L1MessageQueueCaller is an auto generated read-only Go binding around an Ethereum contract.
type L1MessageQueueCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L1MessageQueueTransactor is an auto generated write-only Go binding around an Ethereum contract.
type L1MessageQueueTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L1MessageQueueFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type L1MessageQueueFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L1MessageQueueSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type L1MessageQueueSession struct {
	Contract     *L1MessageQueue   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// L1MessageQueueCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type L1MessageQueueCallerSession struct {
	Contract *L1MessageQueueCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// L1MessageQueueTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type L1MessageQueueTransactorSession struct {
	Contract     *L1MessageQueueTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// L1MessageQueueRaw is an auto generated low-level Go binding around an Ethereum contract.
type L1MessageQueueRaw struct {
	Contract *L1MessageQueue // Generic contract binding to access the raw methods on
}

// L1MessageQueueCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type L1MessageQueueCallerRaw struct {
	Contract *L1MessageQueueCaller // Generic read-only contract binding to access the raw methods on
}

// L1MessageQueueTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type L1MessageQueueTransactorRaw struct {
	Contract *L1MessageQueueTransactor // Generic write-only contract binding to access the raw methods on
}

// NewL1MessageQueue creates a new instance of L1MessageQueue, bound to a specific deployed contract.
func NewL1MessageQueue(address common.Address, backend bind.ContractBackend) (*L1MessageQueue, error) {
	contract, err := bindL1MessageQueue(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &L1MessageQueue{L1MessageQueueCaller: L1MessageQueueCaller{contract: contract}, L1MessageQueueTransactor: L1MessageQueueTransactor{contract: contract}, L1MessageQueueFilterer: L1MessageQueueFilterer{contract: contract}}, nil
}

// NewL1MessageQueueCaller creates a new read-only instance of L1MessageQueue, bound to a specific deployed contract.
func NewL1MessageQueueCaller(address common.Address, caller bind.ContractCaller) (*L1MessageQueueCaller, error) {
	contract, err := bindL1MessageQueue(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &L1MessageQueueCaller{contract: contract}, nil
}

// NewL1MessageQueueTransactor creates a new write-only instance of L1MessageQueue, bound to a specific deployed contract.
func NewL1MessageQueueTransactor(address common.Address, transactor bind.ContractTransactor) (*L1MessageQueueTransactor, error) {
	contract, err := bindL1MessageQueue(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &L1MessageQueueTransactor{contract: contract}, nil
}

// NewL1MessageQueueFilterer creates a new log filterer instance of L1MessageQueue, bound to a specific deployed contract.
func NewL1MessageQueueFilterer(address common.Address, filterer bind.ContractFilterer) (*L1MessageQueueFilterer, error) {
	contract, err := bindL1MessageQueue(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &L1MessageQueueFilterer{contract: contract}, nil
}

// bindL1MessageQueue binds a generic wrapper to an already deployed contract.
func bindL1MessageQueue(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(L1MessageQueueABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L1MessageQueue *L1MessageQueueRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L1MessageQueue.Contract.L1MessageQueueCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L1MessageQueue *L1MessageQueueRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1MessageQueue.Contract.L1MessageQueueTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L1MessageQueue *L1MessageQueueRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L1MessageQueue.Contract.L1MessageQueueTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L1MessageQueue *L1MessageQueueCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L1MessageQueue.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L1MessageQueue *L1MessageQueueTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1MessageQueue.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L1MessageQueue *L1MessageQueueTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L1MessageQueue.Contract.contract.Transact(opts, method, params...)
}

// EstimateCrossDomainMessageFee is a free data retrieval call binding the contract method 0x9119a52b.
//
// Solidity: function estimateCrossDomainMessageFee(address _sender, address _target, bytes _message, uint256 _gasLimit) view returns(uint256)
func (_L1MessageQueue *L1MessageQueueCaller) EstimateCrossDomainMessageFee(opts *bind.CallOpts, _sender common.Address, _target common.Address, _message []byte, _gasLimit *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _L1MessageQueue.contract.Call(opts, &out, "estimateCrossDomainMessageFee", _sender, _target, _message, _gasLimit)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EstimateCrossDomainMessageFee is a free data retrieval call binding the contract method 0x9119a52b.
//
// Solidity: function estimateCrossDomainMessageFee(address _sender, address _target, bytes _message, uint256 _gasLimit) view returns(uint256)
func (_L1MessageQueue *L1MessageQueueSession) EstimateCrossDomainMessageFee(_sender common.Address, _target common.Address, _message []byte, _gasLimit *big.Int) (*big.Int, error) {
	return _L1MessageQueue.Contract.EstimateCrossDomainMessageFee(&_L1MessageQueue.CallOpts, _sender, _target, _message, _gasLimit)
}

// EstimateCrossDomainMessageFee is a free data retrieval call binding the contract method 0x9119a52b.
//
// Solidity: function estimateCrossDomainMessageFee(address _sender, address _target, bytes _message, uint256 _gasLimit) view returns(uint256)
func (_L1MessageQueue *L1MessageQueueCallerSession) EstimateCrossDomainMessageFee(_sender common.Address, _target common.Address, _message []byte, _gasLimit *big.Int) (*big.Int, error) {
	return _L1MessageQueue.Contract.EstimateCrossDomainMessageFee(&_L1MessageQueue.CallOpts, _sender, _target, _message, _gasLimit)
}

// GasOracle is a free data retrieval call binding the contract method 0x5d62a8dd.
//
// Solidity: function gasOracle() view returns(address)
func (_L1MessageQueue *L1MessageQueueCaller) GasOracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1MessageQueue.contract.Call(opts, &out, "gasOracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GasOracle is a free data retrieval call binding the contract method 0x5d62a8dd.
//
// Solidity: function gasOracle() view returns(address)
func (_L1MessageQueue *L1MessageQueueSession) GasOracle() (common.Address, error) {
	return _L1MessageQueue.Contract.GasOracle(&_L1MessageQueue.CallOpts)
}

// GasOracle is a free data retrieval call binding the contract method 0x5d62a8dd.
//
// Solidity: function gasOracle() view returns(address)
func (_L1MessageQueue *L1MessageQueueCallerSession) GasOracle() (common.Address, error) {
	return _L1MessageQueue.Contract.GasOracle(&_L1MessageQueue.CallOpts)
}

// GetCrossDomainMessage is a free data retrieval call binding the contract method 0xae453cd5.
//
// Solidity: function getCrossDomainMessage(uint256 _queueIndex) view returns(bytes32)
func (_L1MessageQueue *L1MessageQueueCaller) GetCrossDomainMessage(opts *bind.CallOpts, _queueIndex *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _L1MessageQueue.contract.Call(opts, &out, "getCrossDomainMessage", _queueIndex)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetCrossDomainMessage is a free data retrieval call binding the contract method 0xae453cd5.
//
// Solidity: function getCrossDomainMessage(uint256 _queueIndex) view returns(bytes32)
func (_L1MessageQueue *L1MessageQueueSession) GetCrossDomainMessage(_queueIndex *big.Int) ([32]byte, error) {
	return _L1MessageQueue.Contract.GetCrossDomainMessage(&_L1MessageQueue.CallOpts, _queueIndex)
}

// GetCrossDomainMessage is a free data retrieval call binding the contract method 0xae453cd5.
//
// Solidity: function getCrossDomainMessage(uint256 _queueIndex) view returns(bytes32)
func (_L1MessageQueue *L1MessageQueueCallerSession) GetCrossDomainMessage(_queueIndex *big.Int) ([32]byte, error) {
	return _L1MessageQueue.Contract.GetCrossDomainMessage(&_L1MessageQueue.CallOpts, _queueIndex)
}

// MessageQueue is a free data retrieval call binding the contract method 0x29aa604b.
//
// Solidity: function messageQueue(uint256 ) view returns(bytes32)
func (_L1MessageQueue *L1MessageQueueCaller) MessageQueue(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _L1MessageQueue.contract.Call(opts, &out, "messageQueue", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MessageQueue is a free data retrieval call binding the contract method 0x29aa604b.
//
// Solidity: function messageQueue(uint256 ) view returns(bytes32)
func (_L1MessageQueue *L1MessageQueueSession) MessageQueue(arg0 *big.Int) ([32]byte, error) {
	return _L1MessageQueue.Contract.MessageQueue(&_L1MessageQueue.CallOpts, arg0)
}

// MessageQueue is a free data retrieval call binding the contract method 0x29aa604b.
//
// Solidity: function messageQueue(uint256 ) view returns(bytes32)
func (_L1MessageQueue *L1MessageQueueCallerSession) MessageQueue(arg0 *big.Int) ([32]byte, error) {
	return _L1MessageQueue.Contract.MessageQueue(&_L1MessageQueue.CallOpts, arg0)
}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_L1MessageQueue *L1MessageQueueCaller) Messenger(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1MessageQueue.contract.Call(opts, &out, "messenger")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_L1MessageQueue *L1MessageQueueSession) Messenger() (common.Address, error) {
	return _L1MessageQueue.Contract.Messenger(&_L1MessageQueue.CallOpts)
}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_L1MessageQueue *L1MessageQueueCallerSession) Messenger() (common.Address, error) {
	return _L1MessageQueue.Contract.Messenger(&_L1MessageQueue.CallOpts)
}

// NextCrossDomainMessageIndex is a free data retrieval call binding the contract method 0xfd0ad31e.
//
// Solidity: function nextCrossDomainMessageIndex() view returns(uint256)
func (_L1MessageQueue *L1MessageQueueCaller) NextCrossDomainMessageIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _L1MessageQueue.contract.Call(opts, &out, "nextCrossDomainMessageIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextCrossDomainMessageIndex is a free data retrieval call binding the contract method 0xfd0ad31e.
//
// Solidity: function nextCrossDomainMessageIndex() view returns(uint256)
func (_L1MessageQueue *L1MessageQueueSession) NextCrossDomainMessageIndex() (*big.Int, error) {
	return _L1MessageQueue.Contract.NextCrossDomainMessageIndex(&_L1MessageQueue.CallOpts)
}

// NextCrossDomainMessageIndex is a free data retrieval call binding the contract method 0xfd0ad31e.
//
// Solidity: function nextCrossDomainMessageIndex() view returns(uint256)
func (_L1MessageQueue *L1MessageQueueCallerSession) NextCrossDomainMessageIndex() (*big.Int, error) {
	return _L1MessageQueue.Contract.NextCrossDomainMessageIndex(&_L1MessageQueue.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L1MessageQueue *L1MessageQueueCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1MessageQueue.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L1MessageQueue *L1MessageQueueSession) Owner() (common.Address, error) {
	return _L1MessageQueue.Contract.Owner(&_L1MessageQueue.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L1MessageQueue *L1MessageQueueCallerSession) Owner() (common.Address, error) {
	return _L1MessageQueue.Contract.Owner(&_L1MessageQueue.CallOpts)
}

// AppendCrossDomainMessage is a paid mutator transaction binding the contract method 0x9b159782.
//
// Solidity: function appendCrossDomainMessage(address _target, uint256 _gasLimit, bytes _data) returns()
func (_L1MessageQueue *L1MessageQueueTransactor) AppendCrossDomainMessage(opts *bind.TransactOpts, _target common.Address, _gasLimit *big.Int, _data []byte) (*types.Transaction, error) {
	return _L1MessageQueue.contract.Transact(opts, "appendCrossDomainMessage", _target, _gasLimit, _data)
}

// AppendCrossDomainMessage is a paid mutator transaction binding the contract method 0x9b159782.
//
// Solidity: function appendCrossDomainMessage(address _target, uint256 _gasLimit, bytes _data) returns()
func (_L1MessageQueue *L1MessageQueueSession) AppendCrossDomainMessage(_target common.Address, _gasLimit *big.Int, _data []byte) (*types.Transaction, error) {
	return _L1MessageQueue.Contract.AppendCrossDomainMessage(&_L1MessageQueue.TransactOpts, _target, _gasLimit, _data)
}

// AppendCrossDomainMessage is a paid mutator transaction binding the contract method 0x9b159782.
//
// Solidity: function appendCrossDomainMessage(address _target, uint256 _gasLimit, bytes _data) returns()
func (_L1MessageQueue *L1MessageQueueTransactorSession) AppendCrossDomainMessage(_target common.Address, _gasLimit *big.Int, _data []byte) (*types.Transaction, error) {
	return _L1MessageQueue.Contract.AppendCrossDomainMessage(&_L1MessageQueue.TransactOpts, _target, _gasLimit, _data)
}

// AppendEnforcedTransaction is a paid mutator transaction binding the contract method 0xbdc6f0a0.
//
// Solidity: function appendEnforcedTransaction(address , address , uint256 , uint256 , bytes ) returns()
func (_L1MessageQueue *L1MessageQueueTransactor) AppendEnforcedTransaction(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _L1MessageQueue.contract.Transact(opts, "appendEnforcedTransaction", arg0, arg1, arg2, arg3, arg4)
}

// AppendEnforcedTransaction is a paid mutator transaction binding the contract method 0xbdc6f0a0.
//
// Solidity: function appendEnforcedTransaction(address , address , uint256 , uint256 , bytes ) returns()
func (_L1MessageQueue *L1MessageQueueSession) AppendEnforcedTransaction(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _L1MessageQueue.Contract.AppendEnforcedTransaction(&_L1MessageQueue.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// AppendEnforcedTransaction is a paid mutator transaction binding the contract method 0xbdc6f0a0.
//
// Solidity: function appendEnforcedTransaction(address , address , uint256 , uint256 , bytes ) returns()
func (_L1MessageQueue *L1MessageQueueTransactorSession) AppendEnforcedTransaction(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _L1MessageQueue.Contract.AppendEnforcedTransaction(&_L1MessageQueue.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _messenger, address _gasOracle) returns()
func (_L1MessageQueue *L1MessageQueueTransactor) Initialize(opts *bind.TransactOpts, _messenger common.Address, _gasOracle common.Address) (*types.Transaction, error) {
	return _L1MessageQueue.contract.Transact(opts, "initialize", _messenger, _gasOracle)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _messenger, address _gasOracle) returns()
func (_L1MessageQueue *L1MessageQueueSession) Initialize(_messenger common.Address, _gasOracle common.Address) (*types.Transaction, error) {
	return _L1MessageQueue.Contract.Initialize(&_L1MessageQueue.TransactOpts, _messenger, _gasOracle)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _messenger, address _gasOracle) returns()
func (_L1MessageQueue *L1MessageQueueTransactorSession) Initialize(_messenger common.Address, _gasOracle common.Address) (*types.Transaction, error) {
	return _L1MessageQueue.Contract.Initialize(&_L1MessageQueue.TransactOpts, _messenger, _gasOracle)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_L1MessageQueue *L1MessageQueueTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1MessageQueue.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_L1MessageQueue *L1MessageQueueSession) RenounceOwnership() (*types.Transaction, error) {
	return _L1MessageQueue.Contract.RenounceOwnership(&_L1MessageQueue.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_L1MessageQueue *L1MessageQueueTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _L1MessageQueue.Contract.RenounceOwnership(&_L1MessageQueue.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_L1MessageQueue *L1MessageQueueTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _L1MessageQueue.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_L1MessageQueue *L1MessageQueueSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _L1MessageQueue.Contract.TransferOwnership(&_L1MessageQueue.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_L1MessageQueue *L1MessageQueueTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _L1MessageQueue.Contract.TransferOwnership(&_L1MessageQueue.TransactOpts, newOwner)
}

// UpdateGasOracle is a paid mutator transaction binding the contract method 0x70cee67f.
//
// Solidity: function updateGasOracle(address _newGasOracle) returns()
func (_L1MessageQueue *L1MessageQueueTransactor) UpdateGasOracle(opts *bind.TransactOpts, _newGasOracle common.Address) (*types.Transaction, error) {
	return _L1MessageQueue.contract.Transact(opts, "updateGasOracle", _newGasOracle)
}

// UpdateGasOracle is a paid mutator transaction binding the contract method 0x70cee67f.
//
// Solidity: function updateGasOracle(address _newGasOracle) returns()
func (_L1MessageQueue *L1MessageQueueSession) UpdateGasOracle(_newGasOracle common.Address) (*types.Transaction, error) {
	return _L1MessageQueue.Contract.UpdateGasOracle(&_L1MessageQueue.TransactOpts, _newGasOracle)
}

// UpdateGasOracle is a paid mutator transaction binding the contract method 0x70cee67f.
//
// Solidity: function updateGasOracle(address _newGasOracle) returns()
func (_L1MessageQueue *L1MessageQueueTransactorSession) UpdateGasOracle(_newGasOracle common.Address) (*types.Transaction, error) {
	return _L1MessageQueue.Contract.UpdateGasOracle(&_L1MessageQueue.TransactOpts, _newGasOracle)
}

// L1MessageQueueOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the L1MessageQueue contract.
type L1MessageQueueOwnershipTransferredIterator struct {
	Event *L1MessageQueueOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *L1MessageQueueOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1MessageQueueOwnershipTransferred)
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
		it.Event = new(L1MessageQueueOwnershipTransferred)
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
func (it *L1MessageQueueOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1MessageQueueOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1MessageQueueOwnershipTransferred represents a OwnershipTransferred event raised by the L1MessageQueue contract.
type L1MessageQueueOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_L1MessageQueue *L1MessageQueueFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*L1MessageQueueOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _L1MessageQueue.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &L1MessageQueueOwnershipTransferredIterator{contract: _L1MessageQueue.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_L1MessageQueue *L1MessageQueueFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *L1MessageQueueOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _L1MessageQueue.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1MessageQueueOwnershipTransferred)
				if err := _L1MessageQueue.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_L1MessageQueue *L1MessageQueueFilterer) ParseOwnershipTransferred(log types.Log) (*L1MessageQueueOwnershipTransferred, error) {
	event := new(L1MessageQueueOwnershipTransferred)
	if err := _L1MessageQueue.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1MessageQueueQueueTransactionIterator is returned from FilterQueueTransaction and is used to iterate over the raw logs and unpacked data for QueueTransaction events raised by the L1MessageQueue contract.
type L1MessageQueueQueueTransactionIterator struct {
	Event *L1MessageQueueQueueTransaction // Event containing the contract specifics and raw log

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
func (it *L1MessageQueueQueueTransactionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1MessageQueueQueueTransaction)
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
		it.Event = new(L1MessageQueueQueueTransaction)
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
func (it *L1MessageQueueQueueTransactionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1MessageQueueQueueTransactionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1MessageQueueQueueTransaction represents a QueueTransaction event raised by the L1MessageQueue contract.
type L1MessageQueueQueueTransaction struct {
	Sender     common.Address
	Target     common.Address
	Value      *big.Int
	QueueIndex *big.Int
	GasLimit   *big.Int
	Data       []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterQueueTransaction is a free log retrieval operation binding the contract event 0xbdcc7517f8fe3db6506dfd910942d0bbecaf3d6a506dadea65b0d988e75b9439.
//
// Solidity: event QueueTransaction(address indexed sender, address indexed target, uint256 value, uint256 queueIndex, uint256 gasLimit, bytes data)
func (_L1MessageQueue *L1MessageQueueFilterer) FilterQueueTransaction(opts *bind.FilterOpts, sender []common.Address, target []common.Address) (*L1MessageQueueQueueTransactionIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _L1MessageQueue.contract.FilterLogs(opts, "QueueTransaction", senderRule, targetRule)
	if err != nil {
		return nil, err
	}
	return &L1MessageQueueQueueTransactionIterator{contract: _L1MessageQueue.contract, event: "QueueTransaction", logs: logs, sub: sub}, nil
}

// WatchQueueTransaction is a free log subscription operation binding the contract event 0xbdcc7517f8fe3db6506dfd910942d0bbecaf3d6a506dadea65b0d988e75b9439.
//
// Solidity: event QueueTransaction(address indexed sender, address indexed target, uint256 value, uint256 queueIndex, uint256 gasLimit, bytes data)
func (_L1MessageQueue *L1MessageQueueFilterer) WatchQueueTransaction(opts *bind.WatchOpts, sink chan<- *L1MessageQueueQueueTransaction, sender []common.Address, target []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _L1MessageQueue.contract.WatchLogs(opts, "QueueTransaction", senderRule, targetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1MessageQueueQueueTransaction)
				if err := _L1MessageQueue.contract.UnpackLog(event, "QueueTransaction", log); err != nil {
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

// ParseQueueTransaction is a log parse operation binding the contract event 0xbdcc7517f8fe3db6506dfd910942d0bbecaf3d6a506dadea65b0d988e75b9439.
//
// Solidity: event QueueTransaction(address indexed sender, address indexed target, uint256 value, uint256 queueIndex, uint256 gasLimit, bytes data)
func (_L1MessageQueue *L1MessageQueueFilterer) ParseQueueTransaction(log types.Log) (*L1MessageQueueQueueTransaction, error) {
	event := new(L1MessageQueueQueueTransaction)
	if err := _L1MessageQueue.contract.UnpackLog(event, "QueueTransaction", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1MessageQueueUpdateGasOracleIterator is returned from FilterUpdateGasOracle and is used to iterate over the raw logs and unpacked data for UpdateGasOracle events raised by the L1MessageQueue contract.
type L1MessageQueueUpdateGasOracleIterator struct {
	Event *L1MessageQueueUpdateGasOracle // Event containing the contract specifics and raw log

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
func (it *L1MessageQueueUpdateGasOracleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1MessageQueueUpdateGasOracle)
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
		it.Event = new(L1MessageQueueUpdateGasOracle)
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
func (it *L1MessageQueueUpdateGasOracleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1MessageQueueUpdateGasOracleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1MessageQueueUpdateGasOracle represents a UpdateGasOracle event raised by the L1MessageQueue contract.
type L1MessageQueueUpdateGasOracle struct {
	OldGasOracle common.Address
	NewGasOracle common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUpdateGasOracle is a free log retrieval operation binding the contract event 0x9ed5ec28f252b3e7f62f1ace8e54c5ebabf4c61cc2a7c33a806365b2ff7ecc5e.
//
// Solidity: event UpdateGasOracle(address _oldGasOracle, address _newGasOracle)
func (_L1MessageQueue *L1MessageQueueFilterer) FilterUpdateGasOracle(opts *bind.FilterOpts) (*L1MessageQueueUpdateGasOracleIterator, error) {

	logs, sub, err := _L1MessageQueue.contract.FilterLogs(opts, "UpdateGasOracle")
	if err != nil {
		return nil, err
	}
	return &L1MessageQueueUpdateGasOracleIterator{contract: _L1MessageQueue.contract, event: "UpdateGasOracle", logs: logs, sub: sub}, nil
}

// WatchUpdateGasOracle is a free log subscription operation binding the contract event 0x9ed5ec28f252b3e7f62f1ace8e54c5ebabf4c61cc2a7c33a806365b2ff7ecc5e.
//
// Solidity: event UpdateGasOracle(address _oldGasOracle, address _newGasOracle)
func (_L1MessageQueue *L1MessageQueueFilterer) WatchUpdateGasOracle(opts *bind.WatchOpts, sink chan<- *L1MessageQueueUpdateGasOracle) (event.Subscription, error) {

	logs, sub, err := _L1MessageQueue.contract.WatchLogs(opts, "UpdateGasOracle")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1MessageQueueUpdateGasOracle)
				if err := _L1MessageQueue.contract.UnpackLog(event, "UpdateGasOracle", log); err != nil {
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

// ParseUpdateGasOracle is a log parse operation binding the contract event 0x9ed5ec28f252b3e7f62f1ace8e54c5ebabf4c61cc2a7c33a806365b2ff7ecc5e.
//
// Solidity: event UpdateGasOracle(address _oldGasOracle, address _newGasOracle)
func (_L1MessageQueue *L1MessageQueueFilterer) ParseUpdateGasOracle(log types.Log) (*L1MessageQueueUpdateGasOracle, error) {
	event := new(L1MessageQueueUpdateGasOracle)
	if err := _L1MessageQueue.contract.UnpackLog(event, "UpdateGasOracle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
