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

// L1GasPriceOracleMetaData contains all meta data concerning the L1GasPriceOracle contract.
var L1GasPriceOracleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"l1BaseFee\",\"type\":\"uint256\"}],\"name\":\"L1BaseFeeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"overhead\",\"type\":\"uint256\"}],\"name\":\"OverheadUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_oldOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"scalar\",\"type\":\"uint256\"}],\"name\":\"ScalarUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_oldWhitelist\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_newWhitelist\",\"type\":\"address\"}],\"name\":\"UpdateWhitelist\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"getL1Fee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"getL1GasUsed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l1BaseFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"overhead\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"scalar\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_l1BaseFee\",\"type\":\"uint256\"}],\"name\":\"setL1BaseFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_overhead\",\"type\":\"uint256\"}],\"name\":\"setOverhead\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_scalar\",\"type\":\"uint256\"}],\"name\":\"setScalar\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newWhitelist\",\"type\":\"address\"}],\"name\":\"updateWhitelist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"whitelist\",\"outputs\":[{\"internalType\":\"contractIWhitelist\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"49948e0e": "getL1Fee(bytes)",
		"de26c4a1": "getL1GasUsed(bytes)",
		"519b4bd3": "l1BaseFee()",
		"0c18c162": "overhead()",
		"8da5cb5b": "owner()",
		"715018a6": "renounceOwnership()",
		"f45e65d8": "scalar()",
		"bede39b5": "setL1BaseFee(uint256)",
		"3577afc5": "setOverhead(uint256)",
		"70465597": "setScalar(uint256)",
		"f2fde38b": "transferOwnership(address)",
		"3d0f963e": "updateWhitelist(address)",
		"93e59dc1": "whitelist()",
	},
	Bin: "0x608060405234801561001057600080fd5b5060405161094038038061094083398101604081905261002f9161008e565b6100388161003e565b506100be565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6000602082840312156100a057600080fd5b81516001600160a01b03811681146100b757600080fd5b9392505050565b610873806100cd6000396000f3fe608060405234801561001057600080fd5b50600436106100cf5760003560e01c8063715018a61161008c578063bede39b511610066578063bede39b51461018d578063de26c4a1146101a0578063f2fde38b146101b3578063f45e65d8146101c657600080fd5b8063715018a6146101475780638da5cb5b1461014f57806393e59dc11461017a57600080fd5b80630c18c162146100d45780633577afc5146100f05780633d0f963e1461010557806349948e0e14610118578063519b4bd31461012b5780637046559714610134575b600080fd5b6100dd60025481565b6040519081526020015b60405180910390f35b6101036100fe366004610671565b6101cf565b005b61010361011336600461068a565b610291565b6100dd6101263660046106d0565b61031c565b6100dd60015481565b610103610142366004610671565b610361565b610103610416565b600054610162906001600160a01b031681565b6040516001600160a01b0390911681526020016100e7565b600454610162906001600160a01b031681565b61010361019b366004610671565b61044c565b6100dd6101ae3660046106d0565b610533565b6101036101c136600461068a565b610595565b6100dd60035481565b6000546001600160a01b031633146102025760405162461bcd60e51b81526004016101f990610781565b60405180910390fd5b621c9c388111156102555760405162461bcd60e51b815260206004820152601760248201527f657863656564206d6178696d756d206f7665726865616400000000000000000060448201526064016101f9565b60028190556040518181527f32740b35c0ea213650f60d44366b4fb211c9033b50714e4a1d34e65d5beb9bb4906020015b60405180910390a150565b6000546001600160a01b031633146102bb5760405162461bcd60e51b81526004016101f990610781565b600480546001600160a01b038381166001600160a01b031983168117909355604080519190921680825260208201939093527f22d1c35fe072d2e42c3c8f9bd4a0d34aa84a0101d020a62517b33fdb3174e5f7910160405180910390a15050565b60008061032883610533565b905060006001548261033a91906107b8565b9050633b9aca006003548261034f91906107b8565b61035991906107e3565b949350505050565b6000546001600160a01b0316331461038b5760405162461bcd60e51b81526004016101f990610781565b61039b633b9aca006103e86107b8565b8111156103e15760405162461bcd60e51b8152602060048201526014602482015273657863656564206d6178696d756d207363616c6560601b60448201526064016101f9565b60038190556040518181527f3336cd9708eaf2769a0f0dc0679f30e80f15dcd88d1921b5a16858e8b85c591a90602001610286565b6000546001600160a01b031633146104405760405162461bcd60e51b81526004016101f990610781565b61044a6000610621565b565b6004805460405163efc7840160e01b815233928101929092526001600160a01b03169063efc7840190602401602060405180830381865afa158015610495573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104b99190610805565b6104fe5760405162461bcd60e51b81526020600482015260166024820152752737ba103bb434ba32b634b9ba32b21039b2b73232b960511b60448201526064016101f9565b60018190556040518181527f351fb23757bb5ea0546c85b7996ddd7155f96b939ebaa5ff7bc49c75f27f2c4490602001610286565b80516000908190815b818110156105865784818151811061055657610556610827565b01602001516001600160f81b0319166000036105775760048301925061057e565b6010830192505b60010161053c565b50506002540160400192915050565b6000546001600160a01b031633146105bf5760405162461bcd60e51b81526004016101f990610781565b6001600160a01b0381166106155760405162461bcd60e51b815260206004820152601d60248201527f6e6577206f776e657220697320746865207a65726f206164647265737300000060448201526064016101f9565b61061e81610621565b50565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b60006020828403121561068357600080fd5b5035919050565b60006020828403121561069c57600080fd5b81356001600160a01b03811681146106b357600080fd5b9392505050565b634e487b7160e01b600052604160045260246000fd5b6000602082840312156106e257600080fd5b813567ffffffffffffffff808211156106fa57600080fd5b818401915084601f83011261070e57600080fd5b813581811115610720576107206106ba565b604051601f8201601f19908116603f01168101908382118183101715610748576107486106ba565b8160405282815287602084870101111561076157600080fd5b826020860160208301376000928101602001929092525095945050505050565b60208082526017908201527f63616c6c6572206973206e6f7420746865206f776e6572000000000000000000604082015260600190565b80820281158282048414176107dd57634e487b7160e01b600052601160045260246000fd5b92915050565b60008261080057634e487b7160e01b600052601260045260246000fd5b500490565b60006020828403121561081757600080fd5b815180151581146106b357600080fd5b634e487b7160e01b600052603260045260246000fdfea264697066735822122035edadd3df0d6b060c40e0515f07a3f9c388de98f0151c9e3c5db7b5b68dec6c64736f6c63430008110033",
}

// L1GasPriceOracleABI is the input ABI used to generate the binding from.
// Deprecated: Use L1GasPriceOracleMetaData.ABI instead.
var L1GasPriceOracleABI = L1GasPriceOracleMetaData.ABI

// Deprecated: Use L1GasPriceOracleMetaData.Sigs instead.
// L1GasPriceOracleFuncSigs maps the 4-byte function signature to its string representation.
var L1GasPriceOracleFuncSigs = L1GasPriceOracleMetaData.Sigs

// L1GasPriceOracleBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use L1GasPriceOracleMetaData.Bin instead.
var L1GasPriceOracleBin = L1GasPriceOracleMetaData.Bin

// DeployL1GasPriceOracle deploys a new Ethereum contract, binding an instance of L1GasPriceOracle to it.
func DeployL1GasPriceOracle(auth *bind.TransactOpts, backend bind.ContractBackend, _owner common.Address) (common.Address, *types.Transaction, *L1GasPriceOracle, error) {
	parsed, err := L1GasPriceOracleMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(L1GasPriceOracleBin), backend, _owner)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &L1GasPriceOracle{L1GasPriceOracleCaller: L1GasPriceOracleCaller{contract: contract}, L1GasPriceOracleTransactor: L1GasPriceOracleTransactor{contract: contract}, L1GasPriceOracleFilterer: L1GasPriceOracleFilterer{contract: contract}}, nil
}

// L1GasPriceOracle is an auto generated Go binding around an Ethereum contract.
type L1GasPriceOracle struct {
	L1GasPriceOracleCaller     // Read-only binding to the contract
	L1GasPriceOracleTransactor // Write-only binding to the contract
	L1GasPriceOracleFilterer   // Log filterer for contract events
}

// L1GasPriceOracleCaller is an auto generated read-only Go binding around an Ethereum contract.
type L1GasPriceOracleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L1GasPriceOracleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type L1GasPriceOracleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L1GasPriceOracleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type L1GasPriceOracleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L1GasPriceOracleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type L1GasPriceOracleSession struct {
	Contract     *L1GasPriceOracle // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// L1GasPriceOracleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type L1GasPriceOracleCallerSession struct {
	Contract *L1GasPriceOracleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// L1GasPriceOracleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type L1GasPriceOracleTransactorSession struct {
	Contract     *L1GasPriceOracleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// L1GasPriceOracleRaw is an auto generated low-level Go binding around an Ethereum contract.
type L1GasPriceOracleRaw struct {
	Contract *L1GasPriceOracle // Generic contract binding to access the raw methods on
}

// L1GasPriceOracleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type L1GasPriceOracleCallerRaw struct {
	Contract *L1GasPriceOracleCaller // Generic read-only contract binding to access the raw methods on
}

// L1GasPriceOracleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type L1GasPriceOracleTransactorRaw struct {
	Contract *L1GasPriceOracleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewL1GasPriceOracle creates a new instance of L1GasPriceOracle, bound to a specific deployed contract.
func NewL1GasPriceOracle(address common.Address, backend bind.ContractBackend) (*L1GasPriceOracle, error) {
	contract, err := bindL1GasPriceOracle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &L1GasPriceOracle{L1GasPriceOracleCaller: L1GasPriceOracleCaller{contract: contract}, L1GasPriceOracleTransactor: L1GasPriceOracleTransactor{contract: contract}, L1GasPriceOracleFilterer: L1GasPriceOracleFilterer{contract: contract}}, nil
}

// NewL1GasPriceOracleCaller creates a new read-only instance of L1GasPriceOracle, bound to a specific deployed contract.
func NewL1GasPriceOracleCaller(address common.Address, caller bind.ContractCaller) (*L1GasPriceOracleCaller, error) {
	contract, err := bindL1GasPriceOracle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &L1GasPriceOracleCaller{contract: contract}, nil
}

// NewL1GasPriceOracleTransactor creates a new write-only instance of L1GasPriceOracle, bound to a specific deployed contract.
func NewL1GasPriceOracleTransactor(address common.Address, transactor bind.ContractTransactor) (*L1GasPriceOracleTransactor, error) {
	contract, err := bindL1GasPriceOracle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &L1GasPriceOracleTransactor{contract: contract}, nil
}

// NewL1GasPriceOracleFilterer creates a new log filterer instance of L1GasPriceOracle, bound to a specific deployed contract.
func NewL1GasPriceOracleFilterer(address common.Address, filterer bind.ContractFilterer) (*L1GasPriceOracleFilterer, error) {
	contract, err := bindL1GasPriceOracle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &L1GasPriceOracleFilterer{contract: contract}, nil
}

// bindL1GasPriceOracle binds a generic wrapper to an already deployed contract.
func bindL1GasPriceOracle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(L1GasPriceOracleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L1GasPriceOracle *L1GasPriceOracleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L1GasPriceOracle.Contract.L1GasPriceOracleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L1GasPriceOracle *L1GasPriceOracleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1GasPriceOracle.Contract.L1GasPriceOracleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L1GasPriceOracle *L1GasPriceOracleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L1GasPriceOracle.Contract.L1GasPriceOracleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L1GasPriceOracle *L1GasPriceOracleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L1GasPriceOracle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L1GasPriceOracle *L1GasPriceOracleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1GasPriceOracle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L1GasPriceOracle *L1GasPriceOracleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L1GasPriceOracle.Contract.contract.Transact(opts, method, params...)
}

// GetL1Fee is a free data retrieval call binding the contract method 0x49948e0e.
//
// Solidity: function getL1Fee(bytes _data) view returns(uint256)
func (_L1GasPriceOracle *L1GasPriceOracleCaller) GetL1Fee(opts *bind.CallOpts, _data []byte) (*big.Int, error) {
	var out []interface{}
	err := _L1GasPriceOracle.contract.Call(opts, &out, "getL1Fee", _data)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetL1Fee is a free data retrieval call binding the contract method 0x49948e0e.
//
// Solidity: function getL1Fee(bytes _data) view returns(uint256)
func (_L1GasPriceOracle *L1GasPriceOracleSession) GetL1Fee(_data []byte) (*big.Int, error) {
	return _L1GasPriceOracle.Contract.GetL1Fee(&_L1GasPriceOracle.CallOpts, _data)
}

// GetL1Fee is a free data retrieval call binding the contract method 0x49948e0e.
//
// Solidity: function getL1Fee(bytes _data) view returns(uint256)
func (_L1GasPriceOracle *L1GasPriceOracleCallerSession) GetL1Fee(_data []byte) (*big.Int, error) {
	return _L1GasPriceOracle.Contract.GetL1Fee(&_L1GasPriceOracle.CallOpts, _data)
}

// GetL1GasUsed is a free data retrieval call binding the contract method 0xde26c4a1.
//
// Solidity: function getL1GasUsed(bytes _data) view returns(uint256)
func (_L1GasPriceOracle *L1GasPriceOracleCaller) GetL1GasUsed(opts *bind.CallOpts, _data []byte) (*big.Int, error) {
	var out []interface{}
	err := _L1GasPriceOracle.contract.Call(opts, &out, "getL1GasUsed", _data)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetL1GasUsed is a free data retrieval call binding the contract method 0xde26c4a1.
//
// Solidity: function getL1GasUsed(bytes _data) view returns(uint256)
func (_L1GasPriceOracle *L1GasPriceOracleSession) GetL1GasUsed(_data []byte) (*big.Int, error) {
	return _L1GasPriceOracle.Contract.GetL1GasUsed(&_L1GasPriceOracle.CallOpts, _data)
}

// GetL1GasUsed is a free data retrieval call binding the contract method 0xde26c4a1.
//
// Solidity: function getL1GasUsed(bytes _data) view returns(uint256)
func (_L1GasPriceOracle *L1GasPriceOracleCallerSession) GetL1GasUsed(_data []byte) (*big.Int, error) {
	return _L1GasPriceOracle.Contract.GetL1GasUsed(&_L1GasPriceOracle.CallOpts, _data)
}

// L1BaseFee is a free data retrieval call binding the contract method 0x519b4bd3.
//
// Solidity: function l1BaseFee() view returns(uint256)
func (_L1GasPriceOracle *L1GasPriceOracleCaller) L1BaseFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _L1GasPriceOracle.contract.Call(opts, &out, "l1BaseFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// L1BaseFee is a free data retrieval call binding the contract method 0x519b4bd3.
//
// Solidity: function l1BaseFee() view returns(uint256)
func (_L1GasPriceOracle *L1GasPriceOracleSession) L1BaseFee() (*big.Int, error) {
	return _L1GasPriceOracle.Contract.L1BaseFee(&_L1GasPriceOracle.CallOpts)
}

// L1BaseFee is a free data retrieval call binding the contract method 0x519b4bd3.
//
// Solidity: function l1BaseFee() view returns(uint256)
func (_L1GasPriceOracle *L1GasPriceOracleCallerSession) L1BaseFee() (*big.Int, error) {
	return _L1GasPriceOracle.Contract.L1BaseFee(&_L1GasPriceOracle.CallOpts)
}

// Overhead is a free data retrieval call binding the contract method 0x0c18c162.
//
// Solidity: function overhead() view returns(uint256)
func (_L1GasPriceOracle *L1GasPriceOracleCaller) Overhead(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _L1GasPriceOracle.contract.Call(opts, &out, "overhead")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Overhead is a free data retrieval call binding the contract method 0x0c18c162.
//
// Solidity: function overhead() view returns(uint256)
func (_L1GasPriceOracle *L1GasPriceOracleSession) Overhead() (*big.Int, error) {
	return _L1GasPriceOracle.Contract.Overhead(&_L1GasPriceOracle.CallOpts)
}

// Overhead is a free data retrieval call binding the contract method 0x0c18c162.
//
// Solidity: function overhead() view returns(uint256)
func (_L1GasPriceOracle *L1GasPriceOracleCallerSession) Overhead() (*big.Int, error) {
	return _L1GasPriceOracle.Contract.Overhead(&_L1GasPriceOracle.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L1GasPriceOracle *L1GasPriceOracleCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1GasPriceOracle.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L1GasPriceOracle *L1GasPriceOracleSession) Owner() (common.Address, error) {
	return _L1GasPriceOracle.Contract.Owner(&_L1GasPriceOracle.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L1GasPriceOracle *L1GasPriceOracleCallerSession) Owner() (common.Address, error) {
	return _L1GasPriceOracle.Contract.Owner(&_L1GasPriceOracle.CallOpts)
}

// Scalar is a free data retrieval call binding the contract method 0xf45e65d8.
//
// Solidity: function scalar() view returns(uint256)
func (_L1GasPriceOracle *L1GasPriceOracleCaller) Scalar(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _L1GasPriceOracle.contract.Call(opts, &out, "scalar")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Scalar is a free data retrieval call binding the contract method 0xf45e65d8.
//
// Solidity: function scalar() view returns(uint256)
func (_L1GasPriceOracle *L1GasPriceOracleSession) Scalar() (*big.Int, error) {
	return _L1GasPriceOracle.Contract.Scalar(&_L1GasPriceOracle.CallOpts)
}

// Scalar is a free data retrieval call binding the contract method 0xf45e65d8.
//
// Solidity: function scalar() view returns(uint256)
func (_L1GasPriceOracle *L1GasPriceOracleCallerSession) Scalar() (*big.Int, error) {
	return _L1GasPriceOracle.Contract.Scalar(&_L1GasPriceOracle.CallOpts)
}

// Whitelist is a free data retrieval call binding the contract method 0x93e59dc1.
//
// Solidity: function whitelist() view returns(address)
func (_L1GasPriceOracle *L1GasPriceOracleCaller) Whitelist(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1GasPriceOracle.contract.Call(opts, &out, "whitelist")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Whitelist is a free data retrieval call binding the contract method 0x93e59dc1.
//
// Solidity: function whitelist() view returns(address)
func (_L1GasPriceOracle *L1GasPriceOracleSession) Whitelist() (common.Address, error) {
	return _L1GasPriceOracle.Contract.Whitelist(&_L1GasPriceOracle.CallOpts)
}

// Whitelist is a free data retrieval call binding the contract method 0x93e59dc1.
//
// Solidity: function whitelist() view returns(address)
func (_L1GasPriceOracle *L1GasPriceOracleCallerSession) Whitelist() (common.Address, error) {
	return _L1GasPriceOracle.Contract.Whitelist(&_L1GasPriceOracle.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_L1GasPriceOracle *L1GasPriceOracleTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1GasPriceOracle.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_L1GasPriceOracle *L1GasPriceOracleSession) RenounceOwnership() (*types.Transaction, error) {
	return _L1GasPriceOracle.Contract.RenounceOwnership(&_L1GasPriceOracle.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_L1GasPriceOracle *L1GasPriceOracleTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _L1GasPriceOracle.Contract.RenounceOwnership(&_L1GasPriceOracle.TransactOpts)
}

// SetL1BaseFee is a paid mutator transaction binding the contract method 0xbede39b5.
//
// Solidity: function setL1BaseFee(uint256 _l1BaseFee) returns()
func (_L1GasPriceOracle *L1GasPriceOracleTransactor) SetL1BaseFee(opts *bind.TransactOpts, _l1BaseFee *big.Int) (*types.Transaction, error) {
	return _L1GasPriceOracle.contract.Transact(opts, "setL1BaseFee", _l1BaseFee)
}

// SetL1BaseFee is a paid mutator transaction binding the contract method 0xbede39b5.
//
// Solidity: function setL1BaseFee(uint256 _l1BaseFee) returns()
func (_L1GasPriceOracle *L1GasPriceOracleSession) SetL1BaseFee(_l1BaseFee *big.Int) (*types.Transaction, error) {
	return _L1GasPriceOracle.Contract.SetL1BaseFee(&_L1GasPriceOracle.TransactOpts, _l1BaseFee)
}

// SetL1BaseFee is a paid mutator transaction binding the contract method 0xbede39b5.
//
// Solidity: function setL1BaseFee(uint256 _l1BaseFee) returns()
func (_L1GasPriceOracle *L1GasPriceOracleTransactorSession) SetL1BaseFee(_l1BaseFee *big.Int) (*types.Transaction, error) {
	return _L1GasPriceOracle.Contract.SetL1BaseFee(&_L1GasPriceOracle.TransactOpts, _l1BaseFee)
}

// SetOverhead is a paid mutator transaction binding the contract method 0x3577afc5.
//
// Solidity: function setOverhead(uint256 _overhead) returns()
func (_L1GasPriceOracle *L1GasPriceOracleTransactor) SetOverhead(opts *bind.TransactOpts, _overhead *big.Int) (*types.Transaction, error) {
	return _L1GasPriceOracle.contract.Transact(opts, "setOverhead", _overhead)
}

// SetOverhead is a paid mutator transaction binding the contract method 0x3577afc5.
//
// Solidity: function setOverhead(uint256 _overhead) returns()
func (_L1GasPriceOracle *L1GasPriceOracleSession) SetOverhead(_overhead *big.Int) (*types.Transaction, error) {
	return _L1GasPriceOracle.Contract.SetOverhead(&_L1GasPriceOracle.TransactOpts, _overhead)
}

// SetOverhead is a paid mutator transaction binding the contract method 0x3577afc5.
//
// Solidity: function setOverhead(uint256 _overhead) returns()
func (_L1GasPriceOracle *L1GasPriceOracleTransactorSession) SetOverhead(_overhead *big.Int) (*types.Transaction, error) {
	return _L1GasPriceOracle.Contract.SetOverhead(&_L1GasPriceOracle.TransactOpts, _overhead)
}

// SetScalar is a paid mutator transaction binding the contract method 0x70465597.
//
// Solidity: function setScalar(uint256 _scalar) returns()
func (_L1GasPriceOracle *L1GasPriceOracleTransactor) SetScalar(opts *bind.TransactOpts, _scalar *big.Int) (*types.Transaction, error) {
	return _L1GasPriceOracle.contract.Transact(opts, "setScalar", _scalar)
}

// SetScalar is a paid mutator transaction binding the contract method 0x70465597.
//
// Solidity: function setScalar(uint256 _scalar) returns()
func (_L1GasPriceOracle *L1GasPriceOracleSession) SetScalar(_scalar *big.Int) (*types.Transaction, error) {
	return _L1GasPriceOracle.Contract.SetScalar(&_L1GasPriceOracle.TransactOpts, _scalar)
}

// SetScalar is a paid mutator transaction binding the contract method 0x70465597.
//
// Solidity: function setScalar(uint256 _scalar) returns()
func (_L1GasPriceOracle *L1GasPriceOracleTransactorSession) SetScalar(_scalar *big.Int) (*types.Transaction, error) {
	return _L1GasPriceOracle.Contract.SetScalar(&_L1GasPriceOracle.TransactOpts, _scalar)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_L1GasPriceOracle *L1GasPriceOracleTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _L1GasPriceOracle.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_L1GasPriceOracle *L1GasPriceOracleSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _L1GasPriceOracle.Contract.TransferOwnership(&_L1GasPriceOracle.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_L1GasPriceOracle *L1GasPriceOracleTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _L1GasPriceOracle.Contract.TransferOwnership(&_L1GasPriceOracle.TransactOpts, _newOwner)
}

// UpdateWhitelist is a paid mutator transaction binding the contract method 0x3d0f963e.
//
// Solidity: function updateWhitelist(address _newWhitelist) returns()
func (_L1GasPriceOracle *L1GasPriceOracleTransactor) UpdateWhitelist(opts *bind.TransactOpts, _newWhitelist common.Address) (*types.Transaction, error) {
	return _L1GasPriceOracle.contract.Transact(opts, "updateWhitelist", _newWhitelist)
}

// UpdateWhitelist is a paid mutator transaction binding the contract method 0x3d0f963e.
//
// Solidity: function updateWhitelist(address _newWhitelist) returns()
func (_L1GasPriceOracle *L1GasPriceOracleSession) UpdateWhitelist(_newWhitelist common.Address) (*types.Transaction, error) {
	return _L1GasPriceOracle.Contract.UpdateWhitelist(&_L1GasPriceOracle.TransactOpts, _newWhitelist)
}

// UpdateWhitelist is a paid mutator transaction binding the contract method 0x3d0f963e.
//
// Solidity: function updateWhitelist(address _newWhitelist) returns()
func (_L1GasPriceOracle *L1GasPriceOracleTransactorSession) UpdateWhitelist(_newWhitelist common.Address) (*types.Transaction, error) {
	return _L1GasPriceOracle.Contract.UpdateWhitelist(&_L1GasPriceOracle.TransactOpts, _newWhitelist)
}

// L1GasPriceOracleL1BaseFeeUpdatedIterator is returned from FilterL1BaseFeeUpdated and is used to iterate over the raw logs and unpacked data for L1BaseFeeUpdated events raised by the L1GasPriceOracle contract.
type L1GasPriceOracleL1BaseFeeUpdatedIterator struct {
	Event *L1GasPriceOracleL1BaseFeeUpdated // Event containing the contract specifics and raw log

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
func (it *L1GasPriceOracleL1BaseFeeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1GasPriceOracleL1BaseFeeUpdated)
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
		it.Event = new(L1GasPriceOracleL1BaseFeeUpdated)
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
func (it *L1GasPriceOracleL1BaseFeeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1GasPriceOracleL1BaseFeeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1GasPriceOracleL1BaseFeeUpdated represents a L1BaseFeeUpdated event raised by the L1GasPriceOracle contract.
type L1GasPriceOracleL1BaseFeeUpdated struct {
	L1BaseFee *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterL1BaseFeeUpdated is a free log retrieval operation binding the contract event 0x351fb23757bb5ea0546c85b7996ddd7155f96b939ebaa5ff7bc49c75f27f2c44.
//
// Solidity: event L1BaseFeeUpdated(uint256 l1BaseFee)
func (_L1GasPriceOracle *L1GasPriceOracleFilterer) FilterL1BaseFeeUpdated(opts *bind.FilterOpts) (*L1GasPriceOracleL1BaseFeeUpdatedIterator, error) {

	logs, sub, err := _L1GasPriceOracle.contract.FilterLogs(opts, "L1BaseFeeUpdated")
	if err != nil {
		return nil, err
	}
	return &L1GasPriceOracleL1BaseFeeUpdatedIterator{contract: _L1GasPriceOracle.contract, event: "L1BaseFeeUpdated", logs: logs, sub: sub}, nil
}

// WatchL1BaseFeeUpdated is a free log subscription operation binding the contract event 0x351fb23757bb5ea0546c85b7996ddd7155f96b939ebaa5ff7bc49c75f27f2c44.
//
// Solidity: event L1BaseFeeUpdated(uint256 l1BaseFee)
func (_L1GasPriceOracle *L1GasPriceOracleFilterer) WatchL1BaseFeeUpdated(opts *bind.WatchOpts, sink chan<- *L1GasPriceOracleL1BaseFeeUpdated) (event.Subscription, error) {

	logs, sub, err := _L1GasPriceOracle.contract.WatchLogs(opts, "L1BaseFeeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1GasPriceOracleL1BaseFeeUpdated)
				if err := _L1GasPriceOracle.contract.UnpackLog(event, "L1BaseFeeUpdated", log); err != nil {
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

// ParseL1BaseFeeUpdated is a log parse operation binding the contract event 0x351fb23757bb5ea0546c85b7996ddd7155f96b939ebaa5ff7bc49c75f27f2c44.
//
// Solidity: event L1BaseFeeUpdated(uint256 l1BaseFee)
func (_L1GasPriceOracle *L1GasPriceOracleFilterer) ParseL1BaseFeeUpdated(log types.Log) (*L1GasPriceOracleL1BaseFeeUpdated, error) {
	event := new(L1GasPriceOracleL1BaseFeeUpdated)
	if err := _L1GasPriceOracle.contract.UnpackLog(event, "L1BaseFeeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1GasPriceOracleOverheadUpdatedIterator is returned from FilterOverheadUpdated and is used to iterate over the raw logs and unpacked data for OverheadUpdated events raised by the L1GasPriceOracle contract.
type L1GasPriceOracleOverheadUpdatedIterator struct {
	Event *L1GasPriceOracleOverheadUpdated // Event containing the contract specifics and raw log

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
func (it *L1GasPriceOracleOverheadUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1GasPriceOracleOverheadUpdated)
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
		it.Event = new(L1GasPriceOracleOverheadUpdated)
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
func (it *L1GasPriceOracleOverheadUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1GasPriceOracleOverheadUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1GasPriceOracleOverheadUpdated represents a OverheadUpdated event raised by the L1GasPriceOracle contract.
type L1GasPriceOracleOverheadUpdated struct {
	Overhead *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOverheadUpdated is a free log retrieval operation binding the contract event 0x32740b35c0ea213650f60d44366b4fb211c9033b50714e4a1d34e65d5beb9bb4.
//
// Solidity: event OverheadUpdated(uint256 overhead)
func (_L1GasPriceOracle *L1GasPriceOracleFilterer) FilterOverheadUpdated(opts *bind.FilterOpts) (*L1GasPriceOracleOverheadUpdatedIterator, error) {

	logs, sub, err := _L1GasPriceOracle.contract.FilterLogs(opts, "OverheadUpdated")
	if err != nil {
		return nil, err
	}
	return &L1GasPriceOracleOverheadUpdatedIterator{contract: _L1GasPriceOracle.contract, event: "OverheadUpdated", logs: logs, sub: sub}, nil
}

// WatchOverheadUpdated is a free log subscription operation binding the contract event 0x32740b35c0ea213650f60d44366b4fb211c9033b50714e4a1d34e65d5beb9bb4.
//
// Solidity: event OverheadUpdated(uint256 overhead)
func (_L1GasPriceOracle *L1GasPriceOracleFilterer) WatchOverheadUpdated(opts *bind.WatchOpts, sink chan<- *L1GasPriceOracleOverheadUpdated) (event.Subscription, error) {

	logs, sub, err := _L1GasPriceOracle.contract.WatchLogs(opts, "OverheadUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1GasPriceOracleOverheadUpdated)
				if err := _L1GasPriceOracle.contract.UnpackLog(event, "OverheadUpdated", log); err != nil {
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

// ParseOverheadUpdated is a log parse operation binding the contract event 0x32740b35c0ea213650f60d44366b4fb211c9033b50714e4a1d34e65d5beb9bb4.
//
// Solidity: event OverheadUpdated(uint256 overhead)
func (_L1GasPriceOracle *L1GasPriceOracleFilterer) ParseOverheadUpdated(log types.Log) (*L1GasPriceOracleOverheadUpdated, error) {
	event := new(L1GasPriceOracleOverheadUpdated)
	if err := _L1GasPriceOracle.contract.UnpackLog(event, "OverheadUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1GasPriceOracleOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the L1GasPriceOracle contract.
type L1GasPriceOracleOwnershipTransferredIterator struct {
	Event *L1GasPriceOracleOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *L1GasPriceOracleOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1GasPriceOracleOwnershipTransferred)
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
		it.Event = new(L1GasPriceOracleOwnershipTransferred)
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
func (it *L1GasPriceOracleOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1GasPriceOracleOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1GasPriceOracleOwnershipTransferred represents a OwnershipTransferred event raised by the L1GasPriceOracle contract.
type L1GasPriceOracleOwnershipTransferred struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed _oldOwner, address indexed _newOwner)
func (_L1GasPriceOracle *L1GasPriceOracleFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, _oldOwner []common.Address, _newOwner []common.Address) (*L1GasPriceOracleOwnershipTransferredIterator, error) {

	var _oldOwnerRule []interface{}
	for _, _oldOwnerItem := range _oldOwner {
		_oldOwnerRule = append(_oldOwnerRule, _oldOwnerItem)
	}
	var _newOwnerRule []interface{}
	for _, _newOwnerItem := range _newOwner {
		_newOwnerRule = append(_newOwnerRule, _newOwnerItem)
	}

	logs, sub, err := _L1GasPriceOracle.contract.FilterLogs(opts, "OwnershipTransferred", _oldOwnerRule, _newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &L1GasPriceOracleOwnershipTransferredIterator{contract: _L1GasPriceOracle.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed _oldOwner, address indexed _newOwner)
func (_L1GasPriceOracle *L1GasPriceOracleFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *L1GasPriceOracleOwnershipTransferred, _oldOwner []common.Address, _newOwner []common.Address) (event.Subscription, error) {

	var _oldOwnerRule []interface{}
	for _, _oldOwnerItem := range _oldOwner {
		_oldOwnerRule = append(_oldOwnerRule, _oldOwnerItem)
	}
	var _newOwnerRule []interface{}
	for _, _newOwnerItem := range _newOwner {
		_newOwnerRule = append(_newOwnerRule, _newOwnerItem)
	}

	logs, sub, err := _L1GasPriceOracle.contract.WatchLogs(opts, "OwnershipTransferred", _oldOwnerRule, _newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1GasPriceOracleOwnershipTransferred)
				if err := _L1GasPriceOracle.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
// Solidity: event OwnershipTransferred(address indexed _oldOwner, address indexed _newOwner)
func (_L1GasPriceOracle *L1GasPriceOracleFilterer) ParseOwnershipTransferred(log types.Log) (*L1GasPriceOracleOwnershipTransferred, error) {
	event := new(L1GasPriceOracleOwnershipTransferred)
	if err := _L1GasPriceOracle.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1GasPriceOracleScalarUpdatedIterator is returned from FilterScalarUpdated and is used to iterate over the raw logs and unpacked data for ScalarUpdated events raised by the L1GasPriceOracle contract.
type L1GasPriceOracleScalarUpdatedIterator struct {
	Event *L1GasPriceOracleScalarUpdated // Event containing the contract specifics and raw log

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
func (it *L1GasPriceOracleScalarUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1GasPriceOracleScalarUpdated)
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
		it.Event = new(L1GasPriceOracleScalarUpdated)
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
func (it *L1GasPriceOracleScalarUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1GasPriceOracleScalarUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1GasPriceOracleScalarUpdated represents a ScalarUpdated event raised by the L1GasPriceOracle contract.
type L1GasPriceOracleScalarUpdated struct {
	Scalar *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterScalarUpdated is a free log retrieval operation binding the contract event 0x3336cd9708eaf2769a0f0dc0679f30e80f15dcd88d1921b5a16858e8b85c591a.
//
// Solidity: event ScalarUpdated(uint256 scalar)
func (_L1GasPriceOracle *L1GasPriceOracleFilterer) FilterScalarUpdated(opts *bind.FilterOpts) (*L1GasPriceOracleScalarUpdatedIterator, error) {

	logs, sub, err := _L1GasPriceOracle.contract.FilterLogs(opts, "ScalarUpdated")
	if err != nil {
		return nil, err
	}
	return &L1GasPriceOracleScalarUpdatedIterator{contract: _L1GasPriceOracle.contract, event: "ScalarUpdated", logs: logs, sub: sub}, nil
}

// WatchScalarUpdated is a free log subscription operation binding the contract event 0x3336cd9708eaf2769a0f0dc0679f30e80f15dcd88d1921b5a16858e8b85c591a.
//
// Solidity: event ScalarUpdated(uint256 scalar)
func (_L1GasPriceOracle *L1GasPriceOracleFilterer) WatchScalarUpdated(opts *bind.WatchOpts, sink chan<- *L1GasPriceOracleScalarUpdated) (event.Subscription, error) {

	logs, sub, err := _L1GasPriceOracle.contract.WatchLogs(opts, "ScalarUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1GasPriceOracleScalarUpdated)
				if err := _L1GasPriceOracle.contract.UnpackLog(event, "ScalarUpdated", log); err != nil {
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

// ParseScalarUpdated is a log parse operation binding the contract event 0x3336cd9708eaf2769a0f0dc0679f30e80f15dcd88d1921b5a16858e8b85c591a.
//
// Solidity: event ScalarUpdated(uint256 scalar)
func (_L1GasPriceOracle *L1GasPriceOracleFilterer) ParseScalarUpdated(log types.Log) (*L1GasPriceOracleScalarUpdated, error) {
	event := new(L1GasPriceOracleScalarUpdated)
	if err := _L1GasPriceOracle.contract.UnpackLog(event, "ScalarUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1GasPriceOracleUpdateWhitelistIterator is returned from FilterUpdateWhitelist and is used to iterate over the raw logs and unpacked data for UpdateWhitelist events raised by the L1GasPriceOracle contract.
type L1GasPriceOracleUpdateWhitelistIterator struct {
	Event *L1GasPriceOracleUpdateWhitelist // Event containing the contract specifics and raw log

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
func (it *L1GasPriceOracleUpdateWhitelistIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1GasPriceOracleUpdateWhitelist)
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
		it.Event = new(L1GasPriceOracleUpdateWhitelist)
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
func (it *L1GasPriceOracleUpdateWhitelistIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1GasPriceOracleUpdateWhitelistIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1GasPriceOracleUpdateWhitelist represents a UpdateWhitelist event raised by the L1GasPriceOracle contract.
type L1GasPriceOracleUpdateWhitelist struct {
	OldWhitelist common.Address
	NewWhitelist common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUpdateWhitelist is a free log retrieval operation binding the contract event 0x22d1c35fe072d2e42c3c8f9bd4a0d34aa84a0101d020a62517b33fdb3174e5f7.
//
// Solidity: event UpdateWhitelist(address _oldWhitelist, address _newWhitelist)
func (_L1GasPriceOracle *L1GasPriceOracleFilterer) FilterUpdateWhitelist(opts *bind.FilterOpts) (*L1GasPriceOracleUpdateWhitelistIterator, error) {

	logs, sub, err := _L1GasPriceOracle.contract.FilterLogs(opts, "UpdateWhitelist")
	if err != nil {
		return nil, err
	}
	return &L1GasPriceOracleUpdateWhitelistIterator{contract: _L1GasPriceOracle.contract, event: "UpdateWhitelist", logs: logs, sub: sub}, nil
}

// WatchUpdateWhitelist is a free log subscription operation binding the contract event 0x22d1c35fe072d2e42c3c8f9bd4a0d34aa84a0101d020a62517b33fdb3174e5f7.
//
// Solidity: event UpdateWhitelist(address _oldWhitelist, address _newWhitelist)
func (_L1GasPriceOracle *L1GasPriceOracleFilterer) WatchUpdateWhitelist(opts *bind.WatchOpts, sink chan<- *L1GasPriceOracleUpdateWhitelist) (event.Subscription, error) {

	logs, sub, err := _L1GasPriceOracle.contract.WatchLogs(opts, "UpdateWhitelist")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1GasPriceOracleUpdateWhitelist)
				if err := _L1GasPriceOracle.contract.UnpackLog(event, "UpdateWhitelist", log); err != nil {
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
func (_L1GasPriceOracle *L1GasPriceOracleFilterer) ParseUpdateWhitelist(log types.Log) (*L1GasPriceOracleUpdateWhitelist, error) {
	event := new(L1GasPriceOracleUpdateWhitelist)
	if err := _L1GasPriceOracle.contract.UnpackLog(event, "UpdateWhitelist", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
