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

// L2GasPriceOracleMetaData contains all meta data concerning the L2GasPriceOracle contract.
var L2GasPriceOracleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"txGas\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"txGasContractCreation\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"zeroGas\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonZeroGas\",\"type\":\"uint256\"}],\"name\":\"IntrinsicParamsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"l2BaseFee\",\"type\":\"uint256\"}],\"name\":\"L2BaseFeeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_oldWhitelist\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_newWhitelist\",\"type\":\"address\"}],\"name\":\"UpdateWhitelist\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"calculateIntrinsicGasFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"}],\"name\":\"estimateCrossDomainMessageFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_txGas\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_txGasContractCreation\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_zeroGas\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_nonZeroGas\",\"type\":\"uint64\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"intrinsicParams\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"txGas\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"txGasContractCreation\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"zeroGas\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nonZeroGas\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l2BaseFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_txGas\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_txGasContractCreation\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_zeroGas\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_nonZeroGas\",\"type\":\"uint64\"}],\"name\":\"setIntrinsicParams\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_l2BaseFee\",\"type\":\"uint256\"}],\"name\":\"setL2BaseFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newWhitelist\",\"type\":\"address\"}],\"name\":\"updateWhitelist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"whitelist\",\"outputs\":[{\"internalType\":\"contractIWhitelist\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"e172d3a1": "calculateIntrinsicGasFee(bytes)",
		"d7704bae": "estimateCrossDomainMessageFee(uint256)",
		"3366ff72": "initialize(uint64,uint64,uint64,uint64)",
		"64431a27": "intrinsicParams()",
		"e3176bd5": "l2BaseFee()",
		"8da5cb5b": "owner()",
		"715018a6": "renounceOwnership()",
		"accf9a60": "setIntrinsicParams(uint64,uint64,uint64,uint64)",
		"d99bc80e": "setL2BaseFee(uint256)",
		"f2fde38b": "transferOwnership(address)",
		"3d0f963e": "updateWhitelist(address)",
		"93e59dc1": "whitelist()",
	},
	Bin: "0x608060405234801561001057600080fd5b5061001961001e565b6100dd565b600054610100900460ff161561008a5760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b60005460ff908116146100db576000805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b610b7c806100ec6000396000f3fe608060405234801561001057600080fd5b50600436106100b45760003560e01c8063accf9a6011610071578063accf9a601461018d578063d7704bae146101a0578063d99bc80e146101c1578063e172d3a1146101d4578063e3176bd5146101e7578063f2fde38b146101f057600080fd5b80633366ff72146100b95780633d0f963e146100ce57806364431a27146100e1578063715018a61461014d5780638da5cb5b1461015557806393e59dc11461017a575b600080fd5b6100cc6100c73660046108f3565b610203565b005b6100cc6100dc366004610947565b61039a565b6067546101149067ffffffffffffffff80821691600160401b8104821691600160801b8204811691600160c01b90041684565b6040805167ffffffffffffffff958616815293851660208501529184169183019190915290911660608201526080015b60405180910390f35b6100cc610403565b6033546001600160a01b03165b6040516001600160a01b039091168152602001610144565b606654610162906001600160a01b031681565b6100cc61019b3660046108f3565b610417565b6101b36101ae366004610977565b610588565b604051908152602001610144565b6100cc6101cf366004610977565b61059e565b6101b36101e23660046109a6565b61068a565b6101b360655481565b6100cc6101fe366004610947565b610752565b600054610100900460ff16158080156102235750600054600160ff909116105b8061023d5750303b15801561023d575060005460ff166001145b6102a55760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084015b60405180910390fd5b6000805460ff1916600117905580156102c8576000805461ff0019166101001790555b6102d06107cb565b6040805160808101825267ffffffffffffffff878116808352878216602084018190528783169484018590529186166060909301839052606780546001600160801b031916909117600160401b909202919091176001600160801b0316600160801b9093026001600160c01b031692909217600160c01b9091021790558015610393576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b5050505050565b6103a26107fa565b606680546001600160a01b038381166001600160a01b031983168117909355604080519190921680825260208201939093527f22d1c35fe072d2e42c3c8f9bd4a0d34aa84a0101d020a62517b33fdb3174e5f7910160405180910390a15050565b61040b6107fa565b6104156000610854565b565b60665460405163efc7840160e01b81523360048201526001600160a01b039091169063efc7840190602401602060405180830381865afa15801561045f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104839190610a57565b6104c85760405162461bcd60e51b81526020600482015260166024820152752737ba103bb434ba32b634b9ba32b21039b2b73232b960511b604482015260640161029c565b604080516080808201835267ffffffffffffffff87811680845287821660208086018290528884168688018190529388166060968701819052606780546001600160801b0319168517600160401b8502176001600160801b0316600160801b87026001600160c01b031617600160c01b830217905587519384529083019190915294810191909152918201929092527f92d8a3003262a4b8ea0d2818ec49eb874ebb871df18bdaf071a0c577fdbd6854910160405180910390a150505050565b6000606554826105989190610a8f565b92915050565b60665460405163efc7840160e01b81523360048201526001600160a01b039091169063efc7840190602401602060405180830381865afa1580156105e6573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061060a9190610a57565b61064f5760405162461bcd60e51b81526020600482015260166024820152752737ba103bb434ba32b634b9ba32b21039b2b73232b960511b604482015260640161029c565b60658190556040518181527f12857a66e1699c61a21781cf161da52147ad9a840339be06baab2ff9a667c8429060200160405180910390a150565b606754815160009167ffffffffffffffff80821692600160801b8304821692600160c01b900490911690839015610749576000805b875181101561070e578781815181106106da576106da610aa6565b01602001516001600160f81b031916156106fc57816106f881610abc565b9250505b8061070681610abc565b9150506106bf565b508381885161071d9190610ad5565b6107279190610a8f565b6107318483610a8f565b61073b9190610ae8565b6107459083610ae8565b9150505b95945050505050565b61075a6107fa565b6001600160a01b0381166107bf5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b606482015260840161029c565b6107c881610854565b50565b600054610100900460ff166107f25760405162461bcd60e51b815260040161029c90610afb565b6104156108a6565b6033546001600160a01b031633146104155760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161029c565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b600054610100900460ff166108cd5760405162461bcd60e51b815260040161029c90610afb565b61041533610854565b803567ffffffffffffffff811681146108ee57600080fd5b919050565b6000806000806080858703121561090957600080fd5b610912856108d6565b9350610920602086016108d6565b925061092e604086016108d6565b915061093c606086016108d6565b905092959194509250565b60006020828403121561095957600080fd5b81356001600160a01b038116811461097057600080fd5b9392505050565b60006020828403121561098957600080fd5b5035919050565b634e487b7160e01b600052604160045260246000fd5b6000602082840312156109b857600080fd5b813567ffffffffffffffff808211156109d057600080fd5b818401915084601f8301126109e457600080fd5b8135818111156109f6576109f6610990565b604051601f8201601f19908116603f01168101908382118183101715610a1e57610a1e610990565b81604052828152876020848701011115610a3757600080fd5b826020860160208301376000928101602001929092525095945050505050565b600060208284031215610a6957600080fd5b8151801515811461097057600080fd5b634e487b7160e01b600052601160045260246000fd5b808202811582820484141761059857610598610a79565b634e487b7160e01b600052603260045260246000fd5b600060018201610ace57610ace610a79565b5060010190565b8181038181111561059857610598610a79565b8082018082111561059857610598610a79565b6020808252602b908201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960408201526a6e697469616c697a696e6760a81b60608201526080019056fea264697066735822122076e560db6994eaf18ac8a22792a3651dc609937d2a6fdfd92619788cb982eb3864736f6c63430008110033",
}

// L2GasPriceOracleABI is the input ABI used to generate the binding from.
// Deprecated: Use L2GasPriceOracleMetaData.ABI instead.
var L2GasPriceOracleABI = L2GasPriceOracleMetaData.ABI

// Deprecated: Use L2GasPriceOracleMetaData.Sigs instead.
// L2GasPriceOracleFuncSigs maps the 4-byte function signature to its string representation.
var L2GasPriceOracleFuncSigs = L2GasPriceOracleMetaData.Sigs

// L2GasPriceOracleBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use L2GasPriceOracleMetaData.Bin instead.
var L2GasPriceOracleBin = L2GasPriceOracleMetaData.Bin

// DeployL2GasPriceOracle deploys a new Ethereum contract, binding an instance of L2GasPriceOracle to it.
func DeployL2GasPriceOracle(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *L2GasPriceOracle, error) {
	parsed, err := L2GasPriceOracleMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(L2GasPriceOracleBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &L2GasPriceOracle{L2GasPriceOracleCaller: L2GasPriceOracleCaller{contract: contract}, L2GasPriceOracleTransactor: L2GasPriceOracleTransactor{contract: contract}, L2GasPriceOracleFilterer: L2GasPriceOracleFilterer{contract: contract}}, nil
}

// L2GasPriceOracle is an auto generated Go binding around an Ethereum contract.
type L2GasPriceOracle struct {
	L2GasPriceOracleCaller     // Read-only binding to the contract
	L2GasPriceOracleTransactor // Write-only binding to the contract
	L2GasPriceOracleFilterer   // Log filterer for contract events
}

// L2GasPriceOracleCaller is an auto generated read-only Go binding around an Ethereum contract.
type L2GasPriceOracleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2GasPriceOracleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type L2GasPriceOracleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2GasPriceOracleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type L2GasPriceOracleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2GasPriceOracleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type L2GasPriceOracleSession struct {
	Contract     *L2GasPriceOracle // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// L2GasPriceOracleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type L2GasPriceOracleCallerSession struct {
	Contract *L2GasPriceOracleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// L2GasPriceOracleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type L2GasPriceOracleTransactorSession struct {
	Contract     *L2GasPriceOracleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// L2GasPriceOracleRaw is an auto generated low-level Go binding around an Ethereum contract.
type L2GasPriceOracleRaw struct {
	Contract *L2GasPriceOracle // Generic contract binding to access the raw methods on
}

// L2GasPriceOracleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type L2GasPriceOracleCallerRaw struct {
	Contract *L2GasPriceOracleCaller // Generic read-only contract binding to access the raw methods on
}

// L2GasPriceOracleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type L2GasPriceOracleTransactorRaw struct {
	Contract *L2GasPriceOracleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewL2GasPriceOracle creates a new instance of L2GasPriceOracle, bound to a specific deployed contract.
func NewL2GasPriceOracle(address common.Address, backend bind.ContractBackend) (*L2GasPriceOracle, error) {
	contract, err := bindL2GasPriceOracle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &L2GasPriceOracle{L2GasPriceOracleCaller: L2GasPriceOracleCaller{contract: contract}, L2GasPriceOracleTransactor: L2GasPriceOracleTransactor{contract: contract}, L2GasPriceOracleFilterer: L2GasPriceOracleFilterer{contract: contract}}, nil
}

// NewL2GasPriceOracleCaller creates a new read-only instance of L2GasPriceOracle, bound to a specific deployed contract.
func NewL2GasPriceOracleCaller(address common.Address, caller bind.ContractCaller) (*L2GasPriceOracleCaller, error) {
	contract, err := bindL2GasPriceOracle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &L2GasPriceOracleCaller{contract: contract}, nil
}

// NewL2GasPriceOracleTransactor creates a new write-only instance of L2GasPriceOracle, bound to a specific deployed contract.
func NewL2GasPriceOracleTransactor(address common.Address, transactor bind.ContractTransactor) (*L2GasPriceOracleTransactor, error) {
	contract, err := bindL2GasPriceOracle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &L2GasPriceOracleTransactor{contract: contract}, nil
}

// NewL2GasPriceOracleFilterer creates a new log filterer instance of L2GasPriceOracle, bound to a specific deployed contract.
func NewL2GasPriceOracleFilterer(address common.Address, filterer bind.ContractFilterer) (*L2GasPriceOracleFilterer, error) {
	contract, err := bindL2GasPriceOracle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &L2GasPriceOracleFilterer{contract: contract}, nil
}

// bindL2GasPriceOracle binds a generic wrapper to an already deployed contract.
func bindL2GasPriceOracle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(L2GasPriceOracleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L2GasPriceOracle *L2GasPriceOracleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L2GasPriceOracle.Contract.L2GasPriceOracleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L2GasPriceOracle *L2GasPriceOracleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2GasPriceOracle.Contract.L2GasPriceOracleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L2GasPriceOracle *L2GasPriceOracleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L2GasPriceOracle.Contract.L2GasPriceOracleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L2GasPriceOracle *L2GasPriceOracleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L2GasPriceOracle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L2GasPriceOracle *L2GasPriceOracleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2GasPriceOracle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L2GasPriceOracle *L2GasPriceOracleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L2GasPriceOracle.Contract.contract.Transact(opts, method, params...)
}

// CalculateIntrinsicGasFee is a free data retrieval call binding the contract method 0xe172d3a1.
//
// Solidity: function calculateIntrinsicGasFee(bytes _message) view returns(uint256)
func (_L2GasPriceOracle *L2GasPriceOracleCaller) CalculateIntrinsicGasFee(opts *bind.CallOpts, _message []byte) (*big.Int, error) {
	var out []interface{}
	err := _L2GasPriceOracle.contract.Call(opts, &out, "calculateIntrinsicGasFee", _message)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateIntrinsicGasFee is a free data retrieval call binding the contract method 0xe172d3a1.
//
// Solidity: function calculateIntrinsicGasFee(bytes _message) view returns(uint256)
func (_L2GasPriceOracle *L2GasPriceOracleSession) CalculateIntrinsicGasFee(_message []byte) (*big.Int, error) {
	return _L2GasPriceOracle.Contract.CalculateIntrinsicGasFee(&_L2GasPriceOracle.CallOpts, _message)
}

// CalculateIntrinsicGasFee is a free data retrieval call binding the contract method 0xe172d3a1.
//
// Solidity: function calculateIntrinsicGasFee(bytes _message) view returns(uint256)
func (_L2GasPriceOracle *L2GasPriceOracleCallerSession) CalculateIntrinsicGasFee(_message []byte) (*big.Int, error) {
	return _L2GasPriceOracle.Contract.CalculateIntrinsicGasFee(&_L2GasPriceOracle.CallOpts, _message)
}

// EstimateCrossDomainMessageFee is a free data retrieval call binding the contract method 0xd7704bae.
//
// Solidity: function estimateCrossDomainMessageFee(uint256 _gasLimit) view returns(uint256)
func (_L2GasPriceOracle *L2GasPriceOracleCaller) EstimateCrossDomainMessageFee(opts *bind.CallOpts, _gasLimit *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _L2GasPriceOracle.contract.Call(opts, &out, "estimateCrossDomainMessageFee", _gasLimit)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EstimateCrossDomainMessageFee is a free data retrieval call binding the contract method 0xd7704bae.
//
// Solidity: function estimateCrossDomainMessageFee(uint256 _gasLimit) view returns(uint256)
func (_L2GasPriceOracle *L2GasPriceOracleSession) EstimateCrossDomainMessageFee(_gasLimit *big.Int) (*big.Int, error) {
	return _L2GasPriceOracle.Contract.EstimateCrossDomainMessageFee(&_L2GasPriceOracle.CallOpts, _gasLimit)
}

// EstimateCrossDomainMessageFee is a free data retrieval call binding the contract method 0xd7704bae.
//
// Solidity: function estimateCrossDomainMessageFee(uint256 _gasLimit) view returns(uint256)
func (_L2GasPriceOracle *L2GasPriceOracleCallerSession) EstimateCrossDomainMessageFee(_gasLimit *big.Int) (*big.Int, error) {
	return _L2GasPriceOracle.Contract.EstimateCrossDomainMessageFee(&_L2GasPriceOracle.CallOpts, _gasLimit)
}

// IntrinsicParams is a free data retrieval call binding the contract method 0x64431a27.
//
// Solidity: function intrinsicParams() view returns(uint64 txGas, uint64 txGasContractCreation, uint64 zeroGas, uint64 nonZeroGas)
func (_L2GasPriceOracle *L2GasPriceOracleCaller) IntrinsicParams(opts *bind.CallOpts) (struct {
	TxGas                 uint64
	TxGasContractCreation uint64
	ZeroGas               uint64
	NonZeroGas            uint64
}, error) {
	var out []interface{}
	err := _L2GasPriceOracle.contract.Call(opts, &out, "intrinsicParams")

	outstruct := new(struct {
		TxGas                 uint64
		TxGasContractCreation uint64
		ZeroGas               uint64
		NonZeroGas            uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TxGas = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.TxGasContractCreation = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	outstruct.ZeroGas = *abi.ConvertType(out[2], new(uint64)).(*uint64)
	outstruct.NonZeroGas = *abi.ConvertType(out[3], new(uint64)).(*uint64)

	return *outstruct, err

}

// IntrinsicParams is a free data retrieval call binding the contract method 0x64431a27.
//
// Solidity: function intrinsicParams() view returns(uint64 txGas, uint64 txGasContractCreation, uint64 zeroGas, uint64 nonZeroGas)
func (_L2GasPriceOracle *L2GasPriceOracleSession) IntrinsicParams() (struct {
	TxGas                 uint64
	TxGasContractCreation uint64
	ZeroGas               uint64
	NonZeroGas            uint64
}, error) {
	return _L2GasPriceOracle.Contract.IntrinsicParams(&_L2GasPriceOracle.CallOpts)
}

// IntrinsicParams is a free data retrieval call binding the contract method 0x64431a27.
//
// Solidity: function intrinsicParams() view returns(uint64 txGas, uint64 txGasContractCreation, uint64 zeroGas, uint64 nonZeroGas)
func (_L2GasPriceOracle *L2GasPriceOracleCallerSession) IntrinsicParams() (struct {
	TxGas                 uint64
	TxGasContractCreation uint64
	ZeroGas               uint64
	NonZeroGas            uint64
}, error) {
	return _L2GasPriceOracle.Contract.IntrinsicParams(&_L2GasPriceOracle.CallOpts)
}

// L2BaseFee is a free data retrieval call binding the contract method 0xe3176bd5.
//
// Solidity: function l2BaseFee() view returns(uint256)
func (_L2GasPriceOracle *L2GasPriceOracleCaller) L2BaseFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _L2GasPriceOracle.contract.Call(opts, &out, "l2BaseFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// L2BaseFee is a free data retrieval call binding the contract method 0xe3176bd5.
//
// Solidity: function l2BaseFee() view returns(uint256)
func (_L2GasPriceOracle *L2GasPriceOracleSession) L2BaseFee() (*big.Int, error) {
	return _L2GasPriceOracle.Contract.L2BaseFee(&_L2GasPriceOracle.CallOpts)
}

// L2BaseFee is a free data retrieval call binding the contract method 0xe3176bd5.
//
// Solidity: function l2BaseFee() view returns(uint256)
func (_L2GasPriceOracle *L2GasPriceOracleCallerSession) L2BaseFee() (*big.Int, error) {
	return _L2GasPriceOracle.Contract.L2BaseFee(&_L2GasPriceOracle.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L2GasPriceOracle *L2GasPriceOracleCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2GasPriceOracle.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L2GasPriceOracle *L2GasPriceOracleSession) Owner() (common.Address, error) {
	return _L2GasPriceOracle.Contract.Owner(&_L2GasPriceOracle.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L2GasPriceOracle *L2GasPriceOracleCallerSession) Owner() (common.Address, error) {
	return _L2GasPriceOracle.Contract.Owner(&_L2GasPriceOracle.CallOpts)
}

// Whitelist is a free data retrieval call binding the contract method 0x93e59dc1.
//
// Solidity: function whitelist() view returns(address)
func (_L2GasPriceOracle *L2GasPriceOracleCaller) Whitelist(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2GasPriceOracle.contract.Call(opts, &out, "whitelist")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Whitelist is a free data retrieval call binding the contract method 0x93e59dc1.
//
// Solidity: function whitelist() view returns(address)
func (_L2GasPriceOracle *L2GasPriceOracleSession) Whitelist() (common.Address, error) {
	return _L2GasPriceOracle.Contract.Whitelist(&_L2GasPriceOracle.CallOpts)
}

// Whitelist is a free data retrieval call binding the contract method 0x93e59dc1.
//
// Solidity: function whitelist() view returns(address)
func (_L2GasPriceOracle *L2GasPriceOracleCallerSession) Whitelist() (common.Address, error) {
	return _L2GasPriceOracle.Contract.Whitelist(&_L2GasPriceOracle.CallOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x3366ff72.
//
// Solidity: function initialize(uint64 _txGas, uint64 _txGasContractCreation, uint64 _zeroGas, uint64 _nonZeroGas) returns()
func (_L2GasPriceOracle *L2GasPriceOracleTransactor) Initialize(opts *bind.TransactOpts, _txGas uint64, _txGasContractCreation uint64, _zeroGas uint64, _nonZeroGas uint64) (*types.Transaction, error) {
	return _L2GasPriceOracle.contract.Transact(opts, "initialize", _txGas, _txGasContractCreation, _zeroGas, _nonZeroGas)
}

// Initialize is a paid mutator transaction binding the contract method 0x3366ff72.
//
// Solidity: function initialize(uint64 _txGas, uint64 _txGasContractCreation, uint64 _zeroGas, uint64 _nonZeroGas) returns()
func (_L2GasPriceOracle *L2GasPriceOracleSession) Initialize(_txGas uint64, _txGasContractCreation uint64, _zeroGas uint64, _nonZeroGas uint64) (*types.Transaction, error) {
	return _L2GasPriceOracle.Contract.Initialize(&_L2GasPriceOracle.TransactOpts, _txGas, _txGasContractCreation, _zeroGas, _nonZeroGas)
}

// Initialize is a paid mutator transaction binding the contract method 0x3366ff72.
//
// Solidity: function initialize(uint64 _txGas, uint64 _txGasContractCreation, uint64 _zeroGas, uint64 _nonZeroGas) returns()
func (_L2GasPriceOracle *L2GasPriceOracleTransactorSession) Initialize(_txGas uint64, _txGasContractCreation uint64, _zeroGas uint64, _nonZeroGas uint64) (*types.Transaction, error) {
	return _L2GasPriceOracle.Contract.Initialize(&_L2GasPriceOracle.TransactOpts, _txGas, _txGasContractCreation, _zeroGas, _nonZeroGas)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_L2GasPriceOracle *L2GasPriceOracleTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2GasPriceOracle.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_L2GasPriceOracle *L2GasPriceOracleSession) RenounceOwnership() (*types.Transaction, error) {
	return _L2GasPriceOracle.Contract.RenounceOwnership(&_L2GasPriceOracle.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_L2GasPriceOracle *L2GasPriceOracleTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _L2GasPriceOracle.Contract.RenounceOwnership(&_L2GasPriceOracle.TransactOpts)
}

// SetIntrinsicParams is a paid mutator transaction binding the contract method 0xaccf9a60.
//
// Solidity: function setIntrinsicParams(uint64 _txGas, uint64 _txGasContractCreation, uint64 _zeroGas, uint64 _nonZeroGas) returns()
func (_L2GasPriceOracle *L2GasPriceOracleTransactor) SetIntrinsicParams(opts *bind.TransactOpts, _txGas uint64, _txGasContractCreation uint64, _zeroGas uint64, _nonZeroGas uint64) (*types.Transaction, error) {
	return _L2GasPriceOracle.contract.Transact(opts, "setIntrinsicParams", _txGas, _txGasContractCreation, _zeroGas, _nonZeroGas)
}

// SetIntrinsicParams is a paid mutator transaction binding the contract method 0xaccf9a60.
//
// Solidity: function setIntrinsicParams(uint64 _txGas, uint64 _txGasContractCreation, uint64 _zeroGas, uint64 _nonZeroGas) returns()
func (_L2GasPriceOracle *L2GasPriceOracleSession) SetIntrinsicParams(_txGas uint64, _txGasContractCreation uint64, _zeroGas uint64, _nonZeroGas uint64) (*types.Transaction, error) {
	return _L2GasPriceOracle.Contract.SetIntrinsicParams(&_L2GasPriceOracle.TransactOpts, _txGas, _txGasContractCreation, _zeroGas, _nonZeroGas)
}

// SetIntrinsicParams is a paid mutator transaction binding the contract method 0xaccf9a60.
//
// Solidity: function setIntrinsicParams(uint64 _txGas, uint64 _txGasContractCreation, uint64 _zeroGas, uint64 _nonZeroGas) returns()
func (_L2GasPriceOracle *L2GasPriceOracleTransactorSession) SetIntrinsicParams(_txGas uint64, _txGasContractCreation uint64, _zeroGas uint64, _nonZeroGas uint64) (*types.Transaction, error) {
	return _L2GasPriceOracle.Contract.SetIntrinsicParams(&_L2GasPriceOracle.TransactOpts, _txGas, _txGasContractCreation, _zeroGas, _nonZeroGas)
}

// SetL2BaseFee is a paid mutator transaction binding the contract method 0xd99bc80e.
//
// Solidity: function setL2BaseFee(uint256 _l2BaseFee) returns()
func (_L2GasPriceOracle *L2GasPriceOracleTransactor) SetL2BaseFee(opts *bind.TransactOpts, _l2BaseFee *big.Int) (*types.Transaction, error) {
	return _L2GasPriceOracle.contract.Transact(opts, "setL2BaseFee", _l2BaseFee)
}

// SetL2BaseFee is a paid mutator transaction binding the contract method 0xd99bc80e.
//
// Solidity: function setL2BaseFee(uint256 _l2BaseFee) returns()
func (_L2GasPriceOracle *L2GasPriceOracleSession) SetL2BaseFee(_l2BaseFee *big.Int) (*types.Transaction, error) {
	return _L2GasPriceOracle.Contract.SetL2BaseFee(&_L2GasPriceOracle.TransactOpts, _l2BaseFee)
}

// SetL2BaseFee is a paid mutator transaction binding the contract method 0xd99bc80e.
//
// Solidity: function setL2BaseFee(uint256 _l2BaseFee) returns()
func (_L2GasPriceOracle *L2GasPriceOracleTransactorSession) SetL2BaseFee(_l2BaseFee *big.Int) (*types.Transaction, error) {
	return _L2GasPriceOracle.Contract.SetL2BaseFee(&_L2GasPriceOracle.TransactOpts, _l2BaseFee)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_L2GasPriceOracle *L2GasPriceOracleTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _L2GasPriceOracle.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_L2GasPriceOracle *L2GasPriceOracleSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _L2GasPriceOracle.Contract.TransferOwnership(&_L2GasPriceOracle.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_L2GasPriceOracle *L2GasPriceOracleTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _L2GasPriceOracle.Contract.TransferOwnership(&_L2GasPriceOracle.TransactOpts, newOwner)
}

// UpdateWhitelist is a paid mutator transaction binding the contract method 0x3d0f963e.
//
// Solidity: function updateWhitelist(address _newWhitelist) returns()
func (_L2GasPriceOracle *L2GasPriceOracleTransactor) UpdateWhitelist(opts *bind.TransactOpts, _newWhitelist common.Address) (*types.Transaction, error) {
	return _L2GasPriceOracle.contract.Transact(opts, "updateWhitelist", _newWhitelist)
}

// UpdateWhitelist is a paid mutator transaction binding the contract method 0x3d0f963e.
//
// Solidity: function updateWhitelist(address _newWhitelist) returns()
func (_L2GasPriceOracle *L2GasPriceOracleSession) UpdateWhitelist(_newWhitelist common.Address) (*types.Transaction, error) {
	return _L2GasPriceOracle.Contract.UpdateWhitelist(&_L2GasPriceOracle.TransactOpts, _newWhitelist)
}

// UpdateWhitelist is a paid mutator transaction binding the contract method 0x3d0f963e.
//
// Solidity: function updateWhitelist(address _newWhitelist) returns()
func (_L2GasPriceOracle *L2GasPriceOracleTransactorSession) UpdateWhitelist(_newWhitelist common.Address) (*types.Transaction, error) {
	return _L2GasPriceOracle.Contract.UpdateWhitelist(&_L2GasPriceOracle.TransactOpts, _newWhitelist)
}

// L2GasPriceOracleInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the L2GasPriceOracle contract.
type L2GasPriceOracleInitializedIterator struct {
	Event *L2GasPriceOracleInitialized // Event containing the contract specifics and raw log

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
func (it *L2GasPriceOracleInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2GasPriceOracleInitialized)
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
		it.Event = new(L2GasPriceOracleInitialized)
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
func (it *L2GasPriceOracleInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2GasPriceOracleInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2GasPriceOracleInitialized represents a Initialized event raised by the L2GasPriceOracle contract.
type L2GasPriceOracleInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_L2GasPriceOracle *L2GasPriceOracleFilterer) FilterInitialized(opts *bind.FilterOpts) (*L2GasPriceOracleInitializedIterator, error) {

	logs, sub, err := _L2GasPriceOracle.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &L2GasPriceOracleInitializedIterator{contract: _L2GasPriceOracle.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_L2GasPriceOracle *L2GasPriceOracleFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *L2GasPriceOracleInitialized) (event.Subscription, error) {

	logs, sub, err := _L2GasPriceOracle.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2GasPriceOracleInitialized)
				if err := _L2GasPriceOracle.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_L2GasPriceOracle *L2GasPriceOracleFilterer) ParseInitialized(log types.Log) (*L2GasPriceOracleInitialized, error) {
	event := new(L2GasPriceOracleInitialized)
	if err := _L2GasPriceOracle.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2GasPriceOracleIntrinsicParamsUpdatedIterator is returned from FilterIntrinsicParamsUpdated and is used to iterate over the raw logs and unpacked data for IntrinsicParamsUpdated events raised by the L2GasPriceOracle contract.
type L2GasPriceOracleIntrinsicParamsUpdatedIterator struct {
	Event *L2GasPriceOracleIntrinsicParamsUpdated // Event containing the contract specifics and raw log

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
func (it *L2GasPriceOracleIntrinsicParamsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2GasPriceOracleIntrinsicParamsUpdated)
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
		it.Event = new(L2GasPriceOracleIntrinsicParamsUpdated)
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
func (it *L2GasPriceOracleIntrinsicParamsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2GasPriceOracleIntrinsicParamsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2GasPriceOracleIntrinsicParamsUpdated represents a IntrinsicParamsUpdated event raised by the L2GasPriceOracle contract.
type L2GasPriceOracleIntrinsicParamsUpdated struct {
	TxGas                 *big.Int
	TxGasContractCreation *big.Int
	ZeroGas               *big.Int
	NonZeroGas            *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterIntrinsicParamsUpdated is a free log retrieval operation binding the contract event 0x92d8a3003262a4b8ea0d2818ec49eb874ebb871df18bdaf071a0c577fdbd6854.
//
// Solidity: event IntrinsicParamsUpdated(uint256 txGas, uint256 txGasContractCreation, uint256 zeroGas, uint256 nonZeroGas)
func (_L2GasPriceOracle *L2GasPriceOracleFilterer) FilterIntrinsicParamsUpdated(opts *bind.FilterOpts) (*L2GasPriceOracleIntrinsicParamsUpdatedIterator, error) {

	logs, sub, err := _L2GasPriceOracle.contract.FilterLogs(opts, "IntrinsicParamsUpdated")
	if err != nil {
		return nil, err
	}
	return &L2GasPriceOracleIntrinsicParamsUpdatedIterator{contract: _L2GasPriceOracle.contract, event: "IntrinsicParamsUpdated", logs: logs, sub: sub}, nil
}

// WatchIntrinsicParamsUpdated is a free log subscription operation binding the contract event 0x92d8a3003262a4b8ea0d2818ec49eb874ebb871df18bdaf071a0c577fdbd6854.
//
// Solidity: event IntrinsicParamsUpdated(uint256 txGas, uint256 txGasContractCreation, uint256 zeroGas, uint256 nonZeroGas)
func (_L2GasPriceOracle *L2GasPriceOracleFilterer) WatchIntrinsicParamsUpdated(opts *bind.WatchOpts, sink chan<- *L2GasPriceOracleIntrinsicParamsUpdated) (event.Subscription, error) {

	logs, sub, err := _L2GasPriceOracle.contract.WatchLogs(opts, "IntrinsicParamsUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2GasPriceOracleIntrinsicParamsUpdated)
				if err := _L2GasPriceOracle.contract.UnpackLog(event, "IntrinsicParamsUpdated", log); err != nil {
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

// ParseIntrinsicParamsUpdated is a log parse operation binding the contract event 0x92d8a3003262a4b8ea0d2818ec49eb874ebb871df18bdaf071a0c577fdbd6854.
//
// Solidity: event IntrinsicParamsUpdated(uint256 txGas, uint256 txGasContractCreation, uint256 zeroGas, uint256 nonZeroGas)
func (_L2GasPriceOracle *L2GasPriceOracleFilterer) ParseIntrinsicParamsUpdated(log types.Log) (*L2GasPriceOracleIntrinsicParamsUpdated, error) {
	event := new(L2GasPriceOracleIntrinsicParamsUpdated)
	if err := _L2GasPriceOracle.contract.UnpackLog(event, "IntrinsicParamsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2GasPriceOracleL2BaseFeeUpdatedIterator is returned from FilterL2BaseFeeUpdated and is used to iterate over the raw logs and unpacked data for L2BaseFeeUpdated events raised by the L2GasPriceOracle contract.
type L2GasPriceOracleL2BaseFeeUpdatedIterator struct {
	Event *L2GasPriceOracleL2BaseFeeUpdated // Event containing the contract specifics and raw log

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
func (it *L2GasPriceOracleL2BaseFeeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2GasPriceOracleL2BaseFeeUpdated)
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
		it.Event = new(L2GasPriceOracleL2BaseFeeUpdated)
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
func (it *L2GasPriceOracleL2BaseFeeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2GasPriceOracleL2BaseFeeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2GasPriceOracleL2BaseFeeUpdated represents a L2BaseFeeUpdated event raised by the L2GasPriceOracle contract.
type L2GasPriceOracleL2BaseFeeUpdated struct {
	L2BaseFee *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterL2BaseFeeUpdated is a free log retrieval operation binding the contract event 0x12857a66e1699c61a21781cf161da52147ad9a840339be06baab2ff9a667c842.
//
// Solidity: event L2BaseFeeUpdated(uint256 l2BaseFee)
func (_L2GasPriceOracle *L2GasPriceOracleFilterer) FilterL2BaseFeeUpdated(opts *bind.FilterOpts) (*L2GasPriceOracleL2BaseFeeUpdatedIterator, error) {

	logs, sub, err := _L2GasPriceOracle.contract.FilterLogs(opts, "L2BaseFeeUpdated")
	if err != nil {
		return nil, err
	}
	return &L2GasPriceOracleL2BaseFeeUpdatedIterator{contract: _L2GasPriceOracle.contract, event: "L2BaseFeeUpdated", logs: logs, sub: sub}, nil
}

// WatchL2BaseFeeUpdated is a free log subscription operation binding the contract event 0x12857a66e1699c61a21781cf161da52147ad9a840339be06baab2ff9a667c842.
//
// Solidity: event L2BaseFeeUpdated(uint256 l2BaseFee)
func (_L2GasPriceOracle *L2GasPriceOracleFilterer) WatchL2BaseFeeUpdated(opts *bind.WatchOpts, sink chan<- *L2GasPriceOracleL2BaseFeeUpdated) (event.Subscription, error) {

	logs, sub, err := _L2GasPriceOracle.contract.WatchLogs(opts, "L2BaseFeeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2GasPriceOracleL2BaseFeeUpdated)
				if err := _L2GasPriceOracle.contract.UnpackLog(event, "L2BaseFeeUpdated", log); err != nil {
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

// ParseL2BaseFeeUpdated is a log parse operation binding the contract event 0x12857a66e1699c61a21781cf161da52147ad9a840339be06baab2ff9a667c842.
//
// Solidity: event L2BaseFeeUpdated(uint256 l2BaseFee)
func (_L2GasPriceOracle *L2GasPriceOracleFilterer) ParseL2BaseFeeUpdated(log types.Log) (*L2GasPriceOracleL2BaseFeeUpdated, error) {
	event := new(L2GasPriceOracleL2BaseFeeUpdated)
	if err := _L2GasPriceOracle.contract.UnpackLog(event, "L2BaseFeeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2GasPriceOracleOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the L2GasPriceOracle contract.
type L2GasPriceOracleOwnershipTransferredIterator struct {
	Event *L2GasPriceOracleOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *L2GasPriceOracleOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2GasPriceOracleOwnershipTransferred)
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
		it.Event = new(L2GasPriceOracleOwnershipTransferred)
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
func (it *L2GasPriceOracleOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2GasPriceOracleOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2GasPriceOracleOwnershipTransferred represents a OwnershipTransferred event raised by the L2GasPriceOracle contract.
type L2GasPriceOracleOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_L2GasPriceOracle *L2GasPriceOracleFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*L2GasPriceOracleOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _L2GasPriceOracle.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &L2GasPriceOracleOwnershipTransferredIterator{contract: _L2GasPriceOracle.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_L2GasPriceOracle *L2GasPriceOracleFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *L2GasPriceOracleOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _L2GasPriceOracle.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2GasPriceOracleOwnershipTransferred)
				if err := _L2GasPriceOracle.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_L2GasPriceOracle *L2GasPriceOracleFilterer) ParseOwnershipTransferred(log types.Log) (*L2GasPriceOracleOwnershipTransferred, error) {
	event := new(L2GasPriceOracleOwnershipTransferred)
	if err := _L2GasPriceOracle.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2GasPriceOracleUpdateWhitelistIterator is returned from FilterUpdateWhitelist and is used to iterate over the raw logs and unpacked data for UpdateWhitelist events raised by the L2GasPriceOracle contract.
type L2GasPriceOracleUpdateWhitelistIterator struct {
	Event *L2GasPriceOracleUpdateWhitelist // Event containing the contract specifics and raw log

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
func (it *L2GasPriceOracleUpdateWhitelistIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2GasPriceOracleUpdateWhitelist)
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
		it.Event = new(L2GasPriceOracleUpdateWhitelist)
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
func (it *L2GasPriceOracleUpdateWhitelistIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2GasPriceOracleUpdateWhitelistIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2GasPriceOracleUpdateWhitelist represents a UpdateWhitelist event raised by the L2GasPriceOracle contract.
type L2GasPriceOracleUpdateWhitelist struct {
	OldWhitelist common.Address
	NewWhitelist common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUpdateWhitelist is a free log retrieval operation binding the contract event 0x22d1c35fe072d2e42c3c8f9bd4a0d34aa84a0101d020a62517b33fdb3174e5f7.
//
// Solidity: event UpdateWhitelist(address _oldWhitelist, address _newWhitelist)
func (_L2GasPriceOracle *L2GasPriceOracleFilterer) FilterUpdateWhitelist(opts *bind.FilterOpts) (*L2GasPriceOracleUpdateWhitelistIterator, error) {

	logs, sub, err := _L2GasPriceOracle.contract.FilterLogs(opts, "UpdateWhitelist")
	if err != nil {
		return nil, err
	}
	return &L2GasPriceOracleUpdateWhitelistIterator{contract: _L2GasPriceOracle.contract, event: "UpdateWhitelist", logs: logs, sub: sub}, nil
}

// WatchUpdateWhitelist is a free log subscription operation binding the contract event 0x22d1c35fe072d2e42c3c8f9bd4a0d34aa84a0101d020a62517b33fdb3174e5f7.
//
// Solidity: event UpdateWhitelist(address _oldWhitelist, address _newWhitelist)
func (_L2GasPriceOracle *L2GasPriceOracleFilterer) WatchUpdateWhitelist(opts *bind.WatchOpts, sink chan<- *L2GasPriceOracleUpdateWhitelist) (event.Subscription, error) {

	logs, sub, err := _L2GasPriceOracle.contract.WatchLogs(opts, "UpdateWhitelist")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2GasPriceOracleUpdateWhitelist)
				if err := _L2GasPriceOracle.contract.UnpackLog(event, "UpdateWhitelist", log); err != nil {
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

// ParseUpdateWhitelist is a log parse operation binding the contract event 0x22d1c35fe072d2e42c3c8f9bd4a0d34aa84a0101d020a62517b33fdb3174e5f7.
//
// Solidity: event UpdateWhitelist(address _oldWhitelist, address _newWhitelist)
func (_L2GasPriceOracle *L2GasPriceOracleFilterer) ParseUpdateWhitelist(log types.Log) (*L2GasPriceOracleUpdateWhitelist, error) {
	event := new(L2GasPriceOracleUpdateWhitelist)
	if err := _L2GasPriceOracle.contract.UnpackLog(event, "UpdateWhitelist", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
