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

// L2StandardERC20GatewayMetaData contains all meta data concerning the L2StandardERC20Gateway contract.
var L2StandardERC20GatewayMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l2Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"FinalizeDepositERC20\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l2Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"WithdrawERC20\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"counterpart\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"finalizeDepositERC20\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"}],\"name\":\"getL1ERC20Address\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"}],\"name\":\"getL2ERC20Address\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_counterpart\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_router\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_messenger\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tokenFactory\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messenger\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"router\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenFactory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"}],\"name\":\"withdrawERC20\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"}],\"name\":\"withdrawERC20\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"}],\"name\":\"withdrawERC20AndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"797594b0": "counterpart()",
		"8431f5c1": "finalizeDepositERC20(address,address,address,address,uint256,bytes)",
		"54bbd59c": "getL1ERC20Address(address)",
		"c676ad29": "getL2ERC20Address(address)",
		"f8c8765e": "initialize(address,address,address,address)",
		"3cb747bf": "messenger()",
		"f887ea40": "router()",
		"e77772fe": "tokenFactory()",
		"a93a4af9": "withdrawERC20(address,address,uint256,uint256)",
		"6c07ea43": "withdrawERC20(address,uint256,uint256)",
		"575361b6": "withdrawERC20AndCall(address,address,uint256,bytes,uint256)",
	},
	Bin: "0x608060405234801561001057600080fd5b50611342806100206000396000f3fe60806040526004361061009c5760003560e01c80638431f5c1116100645780638431f5c114610164578063a93a4af914610177578063c676ad291461018a578063e77772fe146101aa578063f887ea40146101ca578063f8c8765e146101ea57600080fd5b80633cb747bf146100a157806354bbd59c146100dd578063575361b6146101165780636c07ea431461012b578063797594b01461013e575b600080fd5b3480156100ad57600080fd5b506002546100c1906001600160a01b031681565b6040516001600160a01b03909116815260200160405180910390f35b3480156100e957600080fd5b506100c16100f8366004610cef565b6001600160a01b039081166000908152600460205260409020541690565b610129610124366004610d5c565b61020a565b005b610129610139366004610dd7565b610256565b34801561014a57600080fd5b506000546100c1906201000090046001600160a01b031681565b610129610172366004610e0c565b610295565b610129610185366004610ea4565b61064a565b34801561019657600080fd5b506100c16101a5366004610cef565b61065d565b3480156101b657600080fd5b506005546100c1906001600160a01b031681565b3480156101d657600080fd5b506001546100c1906001600160a01b031681565b3480156101f657600080fd5b50610129610205366004610eea565b6106d8565b61024e86868686868080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250889250610852915050565b505050505050565b61029083338460005b6040519080825280601f01601f191660200182016040528015610289576020820181803683370190505b5085610852565b505050565b6002546001600160a01b03163381146102f55760405162461bcd60e51b815260206004820152601760248201527f6f6e6c79206d657373656e6765722063616e2063616c6c00000000000000000060448201526064015b60405180910390fd5b806001600160a01b0316636e296e456040518163ffffffff1660e01b8152600401602060405180830381865afa158015610333573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103579190610f5c565b6000546201000090046001600160a01b039081169116146103ba5760405162461bcd60e51b815260206004820152601760248201527f6f6e6c792063616c6c20627920636f6e7465727061727400000000000000000060448201526064016102ec565b34156103fc5760405162461bcd60e51b81526020600482015260116024820152706e6f6e7a65726f206d73672e76616c756560781b60448201526064016102ec565b6005546040516361e98ca160e01b81523060048201526001600160a01b038a8116602483015260009216906361e98ca190604401602060405180830381865afa15801561044d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104719190610f5c565b9050806001600160a01b0316886001600160a01b0316146104c85760405162461bcd60e51b81526020600482015260116024820152700d86440e8ded6cadc40dad2e6dac2e8c6d607b1b60448201526064016102ec565b506001600160a01b03878116600090815260046020526040902054606091829116610530576001600160a01b03898116600090815260046020526040902080546001600160a01b031916918c1691909117905561052785850186611028565b9250905061056a565b84848080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509293505050505b6001600160a01b0389163b61058357610583828b610ac1565b6040516340c10f1960e01b81526001600160a01b038881166004830152602482018890528a16906340c10f1990604401600060405180830381600087803b1580156105cd57600080fd5b505af11580156105e1573d6000803e3d6000fd5b50505050876001600160a01b0316896001600160a01b03168b6001600160a01b03167f165ba69f6ab40c50cade6f65431801e5f9c7d7830b7545391920db039133ba348a8a86604051610636939291906110dc565b60405180910390a450505050505050505050565b610657848484600061025f565b50505050565b6005546040516361e98ca160e01b81523060048201526001600160a01b03838116602483015260009216906361e98ca190604401602060405180830381865afa1580156106ae573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106d29190610f5c565b92915050565b600054610100900460ff166106f35760005460ff16156106f7565b303b155b61075a5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084016102ec565b600054610100900460ff1615801561077c576000805461ffff19166101011790555b6001600160a01b0384166107c85760405162461bcd60e51b81526020600482015260136024820152727a65726f20726f75746572206164647265737360681b60448201526064016102ec565b6107d3858585610bc7565b6001600160a01b03821661081e5760405162461bcd60e51b81526020600482015260126024820152717a65726f20746f6b656e20666163746f727960701b60448201526064016102ec565b600580546001600160a01b0319166001600160a01b038416179055801561084b576000805461ff00191690555b5050505050565b600083116108995760405162461bcd60e51b81526020600482015260146024820152731dda5d1a191c985dc81e995c9bc8185b5bdd5b9d60621b60448201526064016102ec565b60015433906001600160a01b03168190036108c757828060200190518101906108c2919061113c565b935090505b6001600160a01b03808716600090815260046020526040902054168061092f5760405162461bcd60e51b815260206004820152601960248201527f6e6f20636f72726573706f6e64696e67206c3120746f6b656e0000000000000060448201526064016102ec565b604051632770a7eb60e21b81526001600160a01b03838116600483015260248201879052881690639dc29fac90604401600060405180830381600087803b15801561097957600080fd5b505af115801561098d573d6000803e3d6000fd5b5050505060006384bd13b060e01b8289858a8a8a6040516024016109b696959493929190611197565b60408051601f198184030181529181526020820180516001600160e01b03166001600160e01b03199094169390931790925260025460008054935163b2267a7b60e01b81529294506001600160a01b039182169363b2267a7b933493610a2c936201000090930416919087908b906004016111e6565b6000604051808303818588803b158015610a4557600080fd5b505af1158015610a59573d6000803e3d6000fd5b5050505050826001600160a01b0316886001600160a01b0316836001600160a01b03167fd8d3a3f4ab95694bef40475997598bcf8acd3ed9617a4c1013795429414c27e88a8a8a604051610aaf939291906110dc565b60405180910390a45050505050505050565b600554604051637bdbcbbf60e01b81523060048201526001600160a01b0383811660248301526000921690637bdbcbbf906044016020604051808303816000875af1158015610b14573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b389190610f5c565b9050600080600085806020019051810190610b53919061123e565b925092509250836001600160a01b031663c820f146838584308a6040518663ffffffff1660e01b8152600401610b8d9594939291906112bc565b600060405180830381600087803b158015610ba757600080fd5b505af1158015610bbb573d6000803e3d6000fd5b50505050505050505050565b6001600160a01b038316610c1d5760405162461bcd60e51b815260206004820152601860248201527f7a65726f20636f756e746572706172742061646472657373000000000000000060448201526064016102ec565b6001600160a01b038116610c6c5760405162461bcd60e51b81526020600482015260166024820152757a65726f206d657373656e676572206164647265737360501b60448201526064016102ec565b6000805462010000600160b01b031916620100006001600160a01b038681169190910291909117909155600280546001600160a01b031916838316179055821615610ccd57600180546001600160a01b0319166001600160a01b0384161790555b5050600160035550565b6001600160a01b0381168114610cec57600080fd5b50565b600060208284031215610d0157600080fd5b8135610d0c81610cd7565b9392505050565b60008083601f840112610d2557600080fd5b50813567ffffffffffffffff811115610d3d57600080fd5b602083019150836020828501011115610d5557600080fd5b9250929050565b60008060008060008060a08789031215610d7557600080fd5b8635610d8081610cd7565b95506020870135610d9081610cd7565b945060408701359350606087013567ffffffffffffffff811115610db357600080fd5b610dbf89828a01610d13565b979a9699509497949695608090950135949350505050565b600080600060608486031215610dec57600080fd5b8335610df781610cd7565b95602085013595506040909401359392505050565b600080600080600080600060c0888a031215610e2757600080fd5b8735610e3281610cd7565b96506020880135610e4281610cd7565b95506040880135610e5281610cd7565b94506060880135610e6281610cd7565b93506080880135925060a088013567ffffffffffffffff811115610e8557600080fd5b610e918a828b01610d13565b989b979a50959850939692959293505050565b60008060008060808587031215610eba57600080fd5b8435610ec581610cd7565b93506020850135610ed581610cd7565b93969395505050506040820135916060013590565b60008060008060808587031215610f0057600080fd5b8435610f0b81610cd7565b93506020850135610f1b81610cd7565b92506040850135610f2b81610cd7565b91506060850135610f3b81610cd7565b939692955090935050565b634e487b7160e01b600052604160045260246000fd5b600060208284031215610f6e57600080fd5b8151610d0c81610cd7565b604051601f8201601f1916810167ffffffffffffffff81118282101715610fa257610fa2610f46565b604052919050565b600067ffffffffffffffff821115610fc457610fc4610f46565b50601f01601f191660200190565b600082601f830112610fe357600080fd5b8135610ff6610ff182610faa565b610f79565b81815284602083860101111561100b57600080fd5b816020850160208301376000918101602001919091529392505050565b6000806040838503121561103b57600080fd5b823567ffffffffffffffff8082111561105357600080fd5b61105f86838701610fd2565b9350602085013591508082111561107557600080fd5b5061108285828601610fd2565b9150509250929050565b60005b838110156110a757818101518382015260200161108f565b50506000910152565b600081518084526110c881602086016020860161108c565b601f01601f19169290920160200192915050565b60018060a01b038416815282602082015260606040820152600061110360608301846110b0565b95945050505050565b600061111a610ff184610faa565b905082815283838301111561112e57600080fd5b610d0c83602083018461108c565b6000806040838503121561114f57600080fd5b825161115a81610cd7565b602084015190925067ffffffffffffffff81111561117757600080fd5b8301601f8101851361118857600080fd5b6110828582516020840161110c565b6001600160a01b03878116825286811660208301528581166040830152841660608201526080810183905260c060a082018190526000906111da908301846110b0565b98975050505050505050565b60018060a01b038516815283602082015260806040820152600061120d60808301856110b0565b905082606083015295945050505050565b600082601f83011261122f57600080fd5b610d0c8383516020850161110c565b60008060006060848603121561125357600080fd5b835167ffffffffffffffff8082111561126b57600080fd5b6112778783880161121e565b9450602086015191508082111561128d57600080fd5b5061129a8682870161121e565b925050604084015160ff811681146112b157600080fd5b809150509250925092565b60a0815260006112cf60a08301886110b0565b82810360208401526112e181886110b0565b60ff96909616604084015250506001600160a01b03928316606082015291166080909101529291505056fea26469706673582212205a20d41ef7a648468a1f9a4253516e80fa745056c1af4c1dbfda05743fba446264736f6c63430008110033",
}

// L2StandardERC20GatewayABI is the input ABI used to generate the binding from.
// Deprecated: Use L2StandardERC20GatewayMetaData.ABI instead.
var L2StandardERC20GatewayABI = L2StandardERC20GatewayMetaData.ABI

// Deprecated: Use L2StandardERC20GatewayMetaData.Sigs instead.
// L2StandardERC20GatewayFuncSigs maps the 4-byte function signature to its string representation.
var L2StandardERC20GatewayFuncSigs = L2StandardERC20GatewayMetaData.Sigs

// L2StandardERC20GatewayBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use L2StandardERC20GatewayMetaData.Bin instead.
var L2StandardERC20GatewayBin = L2StandardERC20GatewayMetaData.Bin

// DeployL2StandardERC20Gateway deploys a new Ethereum contract, binding an instance of L2StandardERC20Gateway to it.
func DeployL2StandardERC20Gateway(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *L2StandardERC20Gateway, error) {
	parsed, err := L2StandardERC20GatewayMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(L2StandardERC20GatewayBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &L2StandardERC20Gateway{L2StandardERC20GatewayCaller: L2StandardERC20GatewayCaller{contract: contract}, L2StandardERC20GatewayTransactor: L2StandardERC20GatewayTransactor{contract: contract}, L2StandardERC20GatewayFilterer: L2StandardERC20GatewayFilterer{contract: contract}}, nil
}

// L2StandardERC20Gateway is an auto generated Go binding around an Ethereum contract.
type L2StandardERC20Gateway struct {
	L2StandardERC20GatewayCaller     // Read-only binding to the contract
	L2StandardERC20GatewayTransactor // Write-only binding to the contract
	L2StandardERC20GatewayFilterer   // Log filterer for contract events
}

// L2StandardERC20GatewayCaller is an auto generated read-only Go binding around an Ethereum contract.
type L2StandardERC20GatewayCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2StandardERC20GatewayTransactor is an auto generated write-only Go binding around an Ethereum contract.
type L2StandardERC20GatewayTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2StandardERC20GatewayFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type L2StandardERC20GatewayFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2StandardERC20GatewaySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type L2StandardERC20GatewaySession struct {
	Contract     *L2StandardERC20Gateway // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// L2StandardERC20GatewayCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type L2StandardERC20GatewayCallerSession struct {
	Contract *L2StandardERC20GatewayCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// L2StandardERC20GatewayTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type L2StandardERC20GatewayTransactorSession struct {
	Contract     *L2StandardERC20GatewayTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// L2StandardERC20GatewayRaw is an auto generated low-level Go binding around an Ethereum contract.
type L2StandardERC20GatewayRaw struct {
	Contract *L2StandardERC20Gateway // Generic contract binding to access the raw methods on
}

// L2StandardERC20GatewayCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type L2StandardERC20GatewayCallerRaw struct {
	Contract *L2StandardERC20GatewayCaller // Generic read-only contract binding to access the raw methods on
}

// L2StandardERC20GatewayTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type L2StandardERC20GatewayTransactorRaw struct {
	Contract *L2StandardERC20GatewayTransactor // Generic write-only contract binding to access the raw methods on
}

// NewL2StandardERC20Gateway creates a new instance of L2StandardERC20Gateway, bound to a specific deployed contract.
func NewL2StandardERC20Gateway(address common.Address, backend bind.ContractBackend) (*L2StandardERC20Gateway, error) {
	contract, err := bindL2StandardERC20Gateway(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &L2StandardERC20Gateway{L2StandardERC20GatewayCaller: L2StandardERC20GatewayCaller{contract: contract}, L2StandardERC20GatewayTransactor: L2StandardERC20GatewayTransactor{contract: contract}, L2StandardERC20GatewayFilterer: L2StandardERC20GatewayFilterer{contract: contract}}, nil
}

// NewL2StandardERC20GatewayCaller creates a new read-only instance of L2StandardERC20Gateway, bound to a specific deployed contract.
func NewL2StandardERC20GatewayCaller(address common.Address, caller bind.ContractCaller) (*L2StandardERC20GatewayCaller, error) {
	contract, err := bindL2StandardERC20Gateway(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &L2StandardERC20GatewayCaller{contract: contract}, nil
}

// NewL2StandardERC20GatewayTransactor creates a new write-only instance of L2StandardERC20Gateway, bound to a specific deployed contract.
func NewL2StandardERC20GatewayTransactor(address common.Address, transactor bind.ContractTransactor) (*L2StandardERC20GatewayTransactor, error) {
	contract, err := bindL2StandardERC20Gateway(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &L2StandardERC20GatewayTransactor{contract: contract}, nil
}

// NewL2StandardERC20GatewayFilterer creates a new log filterer instance of L2StandardERC20Gateway, bound to a specific deployed contract.
func NewL2StandardERC20GatewayFilterer(address common.Address, filterer bind.ContractFilterer) (*L2StandardERC20GatewayFilterer, error) {
	contract, err := bindL2StandardERC20Gateway(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &L2StandardERC20GatewayFilterer{contract: contract}, nil
}

// bindL2StandardERC20Gateway binds a generic wrapper to an already deployed contract.
func bindL2StandardERC20Gateway(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(L2StandardERC20GatewayABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L2StandardERC20Gateway *L2StandardERC20GatewayRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L2StandardERC20Gateway.Contract.L2StandardERC20GatewayCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L2StandardERC20Gateway *L2StandardERC20GatewayRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2StandardERC20Gateway.Contract.L2StandardERC20GatewayTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L2StandardERC20Gateway *L2StandardERC20GatewayRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L2StandardERC20Gateway.Contract.L2StandardERC20GatewayTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L2StandardERC20Gateway *L2StandardERC20GatewayCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L2StandardERC20Gateway.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L2StandardERC20Gateway *L2StandardERC20GatewayTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2StandardERC20Gateway.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L2StandardERC20Gateway *L2StandardERC20GatewayTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L2StandardERC20Gateway.Contract.contract.Transact(opts, method, params...)
}

// Counterpart is a free data retrieval call binding the contract method 0x797594b0.
//
// Solidity: function counterpart() view returns(address)
func (_L2StandardERC20Gateway *L2StandardERC20GatewayCaller) Counterpart(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2StandardERC20Gateway.contract.Call(opts, &out, "counterpart")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Counterpart is a free data retrieval call binding the contract method 0x797594b0.
//
// Solidity: function counterpart() view returns(address)
func (_L2StandardERC20Gateway *L2StandardERC20GatewaySession) Counterpart() (common.Address, error) {
	return _L2StandardERC20Gateway.Contract.Counterpart(&_L2StandardERC20Gateway.CallOpts)
}

// Counterpart is a free data retrieval call binding the contract method 0x797594b0.
//
// Solidity: function counterpart() view returns(address)
func (_L2StandardERC20Gateway *L2StandardERC20GatewayCallerSession) Counterpart() (common.Address, error) {
	return _L2StandardERC20Gateway.Contract.Counterpart(&_L2StandardERC20Gateway.CallOpts)
}

// GetL1ERC20Address is a free data retrieval call binding the contract method 0x54bbd59c.
//
// Solidity: function getL1ERC20Address(address _l2Token) view returns(address)
func (_L2StandardERC20Gateway *L2StandardERC20GatewayCaller) GetL1ERC20Address(opts *bind.CallOpts, _l2Token common.Address) (common.Address, error) {
	var out []interface{}
	err := _L2StandardERC20Gateway.contract.Call(opts, &out, "getL1ERC20Address", _l2Token)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetL1ERC20Address is a free data retrieval call binding the contract method 0x54bbd59c.
//
// Solidity: function getL1ERC20Address(address _l2Token) view returns(address)
func (_L2StandardERC20Gateway *L2StandardERC20GatewaySession) GetL1ERC20Address(_l2Token common.Address) (common.Address, error) {
	return _L2StandardERC20Gateway.Contract.GetL1ERC20Address(&_L2StandardERC20Gateway.CallOpts, _l2Token)
}

// GetL1ERC20Address is a free data retrieval call binding the contract method 0x54bbd59c.
//
// Solidity: function getL1ERC20Address(address _l2Token) view returns(address)
func (_L2StandardERC20Gateway *L2StandardERC20GatewayCallerSession) GetL1ERC20Address(_l2Token common.Address) (common.Address, error) {
	return _L2StandardERC20Gateway.Contract.GetL1ERC20Address(&_L2StandardERC20Gateway.CallOpts, _l2Token)
}

// GetL2ERC20Address is a free data retrieval call binding the contract method 0xc676ad29.
//
// Solidity: function getL2ERC20Address(address _l1Token) view returns(address)
func (_L2StandardERC20Gateway *L2StandardERC20GatewayCaller) GetL2ERC20Address(opts *bind.CallOpts, _l1Token common.Address) (common.Address, error) {
	var out []interface{}
	err := _L2StandardERC20Gateway.contract.Call(opts, &out, "getL2ERC20Address", _l1Token)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetL2ERC20Address is a free data retrieval call binding the contract method 0xc676ad29.
//
// Solidity: function getL2ERC20Address(address _l1Token) view returns(address)
func (_L2StandardERC20Gateway *L2StandardERC20GatewaySession) GetL2ERC20Address(_l1Token common.Address) (common.Address, error) {
	return _L2StandardERC20Gateway.Contract.GetL2ERC20Address(&_L2StandardERC20Gateway.CallOpts, _l1Token)
}

// GetL2ERC20Address is a free data retrieval call binding the contract method 0xc676ad29.
//
// Solidity: function getL2ERC20Address(address _l1Token) view returns(address)
func (_L2StandardERC20Gateway *L2StandardERC20GatewayCallerSession) GetL2ERC20Address(_l1Token common.Address) (common.Address, error) {
	return _L2StandardERC20Gateway.Contract.GetL2ERC20Address(&_L2StandardERC20Gateway.CallOpts, _l1Token)
}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_L2StandardERC20Gateway *L2StandardERC20GatewayCaller) Messenger(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2StandardERC20Gateway.contract.Call(opts, &out, "messenger")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_L2StandardERC20Gateway *L2StandardERC20GatewaySession) Messenger() (common.Address, error) {
	return _L2StandardERC20Gateway.Contract.Messenger(&_L2StandardERC20Gateway.CallOpts)
}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_L2StandardERC20Gateway *L2StandardERC20GatewayCallerSession) Messenger() (common.Address, error) {
	return _L2StandardERC20Gateway.Contract.Messenger(&_L2StandardERC20Gateway.CallOpts)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_L2StandardERC20Gateway *L2StandardERC20GatewayCaller) Router(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2StandardERC20Gateway.contract.Call(opts, &out, "router")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_L2StandardERC20Gateway *L2StandardERC20GatewaySession) Router() (common.Address, error) {
	return _L2StandardERC20Gateway.Contract.Router(&_L2StandardERC20Gateway.CallOpts)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_L2StandardERC20Gateway *L2StandardERC20GatewayCallerSession) Router() (common.Address, error) {
	return _L2StandardERC20Gateway.Contract.Router(&_L2StandardERC20Gateway.CallOpts)
}

// TokenFactory is a free data retrieval call binding the contract method 0xe77772fe.
//
// Solidity: function tokenFactory() view returns(address)
func (_L2StandardERC20Gateway *L2StandardERC20GatewayCaller) TokenFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2StandardERC20Gateway.contract.Call(opts, &out, "tokenFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenFactory is a free data retrieval call binding the contract method 0xe77772fe.
//
// Solidity: function tokenFactory() view returns(address)
func (_L2StandardERC20Gateway *L2StandardERC20GatewaySession) TokenFactory() (common.Address, error) {
	return _L2StandardERC20Gateway.Contract.TokenFactory(&_L2StandardERC20Gateway.CallOpts)
}

// TokenFactory is a free data retrieval call binding the contract method 0xe77772fe.
//
// Solidity: function tokenFactory() view returns(address)
func (_L2StandardERC20Gateway *L2StandardERC20GatewayCallerSession) TokenFactory() (common.Address, error) {
	return _L2StandardERC20Gateway.Contract.TokenFactory(&_L2StandardERC20Gateway.CallOpts)
}

// FinalizeDepositERC20 is a paid mutator transaction binding the contract method 0x8431f5c1.
//
// Solidity: function finalizeDepositERC20(address _l1Token, address _l2Token, address _from, address _to, uint256 _amount, bytes _data) payable returns()
func (_L2StandardERC20Gateway *L2StandardERC20GatewayTransactor) FinalizeDepositERC20(opts *bind.TransactOpts, _l1Token common.Address, _l2Token common.Address, _from common.Address, _to common.Address, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _L2StandardERC20Gateway.contract.Transact(opts, "finalizeDepositERC20", _l1Token, _l2Token, _from, _to, _amount, _data)
}

// FinalizeDepositERC20 is a paid mutator transaction binding the contract method 0x8431f5c1.
//
// Solidity: function finalizeDepositERC20(address _l1Token, address _l2Token, address _from, address _to, uint256 _amount, bytes _data) payable returns()
func (_L2StandardERC20Gateway *L2StandardERC20GatewaySession) FinalizeDepositERC20(_l1Token common.Address, _l2Token common.Address, _from common.Address, _to common.Address, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _L2StandardERC20Gateway.Contract.FinalizeDepositERC20(&_L2StandardERC20Gateway.TransactOpts, _l1Token, _l2Token, _from, _to, _amount, _data)
}

// FinalizeDepositERC20 is a paid mutator transaction binding the contract method 0x8431f5c1.
//
// Solidity: function finalizeDepositERC20(address _l1Token, address _l2Token, address _from, address _to, uint256 _amount, bytes _data) payable returns()
func (_L2StandardERC20Gateway *L2StandardERC20GatewayTransactorSession) FinalizeDepositERC20(_l1Token common.Address, _l2Token common.Address, _from common.Address, _to common.Address, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _L2StandardERC20Gateway.Contract.FinalizeDepositERC20(&_L2StandardERC20Gateway.TransactOpts, _l1Token, _l2Token, _from, _to, _amount, _data)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _counterpart, address _router, address _messenger, address _tokenFactory) returns()
func (_L2StandardERC20Gateway *L2StandardERC20GatewayTransactor) Initialize(opts *bind.TransactOpts, _counterpart common.Address, _router common.Address, _messenger common.Address, _tokenFactory common.Address) (*types.Transaction, error) {
	return _L2StandardERC20Gateway.contract.Transact(opts, "initialize", _counterpart, _router, _messenger, _tokenFactory)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _counterpart, address _router, address _messenger, address _tokenFactory) returns()
func (_L2StandardERC20Gateway *L2StandardERC20GatewaySession) Initialize(_counterpart common.Address, _router common.Address, _messenger common.Address, _tokenFactory common.Address) (*types.Transaction, error) {
	return _L2StandardERC20Gateway.Contract.Initialize(&_L2StandardERC20Gateway.TransactOpts, _counterpart, _router, _messenger, _tokenFactory)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _counterpart, address _router, address _messenger, address _tokenFactory) returns()
func (_L2StandardERC20Gateway *L2StandardERC20GatewayTransactorSession) Initialize(_counterpart common.Address, _router common.Address, _messenger common.Address, _tokenFactory common.Address) (*types.Transaction, error) {
	return _L2StandardERC20Gateway.Contract.Initialize(&_L2StandardERC20Gateway.TransactOpts, _counterpart, _router, _messenger, _tokenFactory)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0x6c07ea43.
//
// Solidity: function withdrawERC20(address _token, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L2StandardERC20Gateway *L2StandardERC20GatewayTransactor) WithdrawERC20(opts *bind.TransactOpts, _token common.Address, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2StandardERC20Gateway.contract.Transact(opts, "withdrawERC20", _token, _amount, _gasLimit)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0x6c07ea43.
//
// Solidity: function withdrawERC20(address _token, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L2StandardERC20Gateway *L2StandardERC20GatewaySession) WithdrawERC20(_token common.Address, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2StandardERC20Gateway.Contract.WithdrawERC20(&_L2StandardERC20Gateway.TransactOpts, _token, _amount, _gasLimit)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0x6c07ea43.
//
// Solidity: function withdrawERC20(address _token, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L2StandardERC20Gateway *L2StandardERC20GatewayTransactorSession) WithdrawERC20(_token common.Address, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2StandardERC20Gateway.Contract.WithdrawERC20(&_L2StandardERC20Gateway.TransactOpts, _token, _amount, _gasLimit)
}

// WithdrawERC200 is a paid mutator transaction binding the contract method 0xa93a4af9.
//
// Solidity: function withdrawERC20(address _token, address _to, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L2StandardERC20Gateway *L2StandardERC20GatewayTransactor) WithdrawERC200(opts *bind.TransactOpts, _token common.Address, _to common.Address, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2StandardERC20Gateway.contract.Transact(opts, "withdrawERC200", _token, _to, _amount, _gasLimit)
}

// WithdrawERC200 is a paid mutator transaction binding the contract method 0xa93a4af9.
//
// Solidity: function withdrawERC20(address _token, address _to, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L2StandardERC20Gateway *L2StandardERC20GatewaySession) WithdrawERC200(_token common.Address, _to common.Address, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2StandardERC20Gateway.Contract.WithdrawERC200(&_L2StandardERC20Gateway.TransactOpts, _token, _to, _amount, _gasLimit)
}

// WithdrawERC200 is a paid mutator transaction binding the contract method 0xa93a4af9.
//
// Solidity: function withdrawERC20(address _token, address _to, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L2StandardERC20Gateway *L2StandardERC20GatewayTransactorSession) WithdrawERC200(_token common.Address, _to common.Address, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2StandardERC20Gateway.Contract.WithdrawERC200(&_L2StandardERC20Gateway.TransactOpts, _token, _to, _amount, _gasLimit)
}

// WithdrawERC20AndCall is a paid mutator transaction binding the contract method 0x575361b6.
//
// Solidity: function withdrawERC20AndCall(address _token, address _to, uint256 _amount, bytes _data, uint256 _gasLimit) payable returns()
func (_L2StandardERC20Gateway *L2StandardERC20GatewayTransactor) WithdrawERC20AndCall(opts *bind.TransactOpts, _token common.Address, _to common.Address, _amount *big.Int, _data []byte, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2StandardERC20Gateway.contract.Transact(opts, "withdrawERC20AndCall", _token, _to, _amount, _data, _gasLimit)
}

// WithdrawERC20AndCall is a paid mutator transaction binding the contract method 0x575361b6.
//
// Solidity: function withdrawERC20AndCall(address _token, address _to, uint256 _amount, bytes _data, uint256 _gasLimit) payable returns()
func (_L2StandardERC20Gateway *L2StandardERC20GatewaySession) WithdrawERC20AndCall(_token common.Address, _to common.Address, _amount *big.Int, _data []byte, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2StandardERC20Gateway.Contract.WithdrawERC20AndCall(&_L2StandardERC20Gateway.TransactOpts, _token, _to, _amount, _data, _gasLimit)
}

// WithdrawERC20AndCall is a paid mutator transaction binding the contract method 0x575361b6.
//
// Solidity: function withdrawERC20AndCall(address _token, address _to, uint256 _amount, bytes _data, uint256 _gasLimit) payable returns()
func (_L2StandardERC20Gateway *L2StandardERC20GatewayTransactorSession) WithdrawERC20AndCall(_token common.Address, _to common.Address, _amount *big.Int, _data []byte, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2StandardERC20Gateway.Contract.WithdrawERC20AndCall(&_L2StandardERC20Gateway.TransactOpts, _token, _to, _amount, _data, _gasLimit)
}

// L2StandardERC20GatewayFinalizeDepositERC20Iterator is returned from FilterFinalizeDepositERC20 and is used to iterate over the raw logs and unpacked data for FinalizeDepositERC20 events raised by the L2StandardERC20Gateway contract.
type L2StandardERC20GatewayFinalizeDepositERC20Iterator struct {
	Event *L2StandardERC20GatewayFinalizeDepositERC20 // Event containing the contract specifics and raw log

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
func (it *L2StandardERC20GatewayFinalizeDepositERC20Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2StandardERC20GatewayFinalizeDepositERC20)
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
		it.Event = new(L2StandardERC20GatewayFinalizeDepositERC20)
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
func (it *L2StandardERC20GatewayFinalizeDepositERC20Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2StandardERC20GatewayFinalizeDepositERC20Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2StandardERC20GatewayFinalizeDepositERC20 represents a FinalizeDepositERC20 event raised by the L2StandardERC20Gateway contract.
type L2StandardERC20GatewayFinalizeDepositERC20 struct {
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
func (_L2StandardERC20Gateway *L2StandardERC20GatewayFilterer) FilterFinalizeDepositERC20(opts *bind.FilterOpts, l1Token []common.Address, l2Token []common.Address, from []common.Address) (*L2StandardERC20GatewayFinalizeDepositERC20Iterator, error) {

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

	logs, sub, err := _L2StandardERC20Gateway.contract.FilterLogs(opts, "FinalizeDepositERC20", l1TokenRule, l2TokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &L2StandardERC20GatewayFinalizeDepositERC20Iterator{contract: _L2StandardERC20Gateway.contract, event: "FinalizeDepositERC20", logs: logs, sub: sub}, nil
}

// WatchFinalizeDepositERC20 is a free log subscription operation binding the contract event 0x165ba69f6ab40c50cade6f65431801e5f9c7d7830b7545391920db039133ba34.
//
// Solidity: event FinalizeDepositERC20(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256 amount, bytes data)
func (_L2StandardERC20Gateway *L2StandardERC20GatewayFilterer) WatchFinalizeDepositERC20(opts *bind.WatchOpts, sink chan<- *L2StandardERC20GatewayFinalizeDepositERC20, l1Token []common.Address, l2Token []common.Address, from []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _L2StandardERC20Gateway.contract.WatchLogs(opts, "FinalizeDepositERC20", l1TokenRule, l2TokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2StandardERC20GatewayFinalizeDepositERC20)
				if err := _L2StandardERC20Gateway.contract.UnpackLog(event, "FinalizeDepositERC20", log); err != nil {
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
func (_L2StandardERC20Gateway *L2StandardERC20GatewayFilterer) ParseFinalizeDepositERC20(log types.Log) (*L2StandardERC20GatewayFinalizeDepositERC20, error) {
	event := new(L2StandardERC20GatewayFinalizeDepositERC20)
	if err := _L2StandardERC20Gateway.contract.UnpackLog(event, "FinalizeDepositERC20", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2StandardERC20GatewayWithdrawERC20Iterator is returned from FilterWithdrawERC20 and is used to iterate over the raw logs and unpacked data for WithdrawERC20 events raised by the L2StandardERC20Gateway contract.
type L2StandardERC20GatewayWithdrawERC20Iterator struct {
	Event *L2StandardERC20GatewayWithdrawERC20 // Event containing the contract specifics and raw log

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
func (it *L2StandardERC20GatewayWithdrawERC20Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2StandardERC20GatewayWithdrawERC20)
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
		it.Event = new(L2StandardERC20GatewayWithdrawERC20)
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
func (it *L2StandardERC20GatewayWithdrawERC20Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2StandardERC20GatewayWithdrawERC20Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2StandardERC20GatewayWithdrawERC20 represents a WithdrawERC20 event raised by the L2StandardERC20Gateway contract.
type L2StandardERC20GatewayWithdrawERC20 struct {
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
func (_L2StandardERC20Gateway *L2StandardERC20GatewayFilterer) FilterWithdrawERC20(opts *bind.FilterOpts, l1Token []common.Address, l2Token []common.Address, from []common.Address) (*L2StandardERC20GatewayWithdrawERC20Iterator, error) {

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

	logs, sub, err := _L2StandardERC20Gateway.contract.FilterLogs(opts, "WithdrawERC20", l1TokenRule, l2TokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &L2StandardERC20GatewayWithdrawERC20Iterator{contract: _L2StandardERC20Gateway.contract, event: "WithdrawERC20", logs: logs, sub: sub}, nil
}

// WatchWithdrawERC20 is a free log subscription operation binding the contract event 0xd8d3a3f4ab95694bef40475997598bcf8acd3ed9617a4c1013795429414c27e8.
//
// Solidity: event WithdrawERC20(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256 amount, bytes data)
func (_L2StandardERC20Gateway *L2StandardERC20GatewayFilterer) WatchWithdrawERC20(opts *bind.WatchOpts, sink chan<- *L2StandardERC20GatewayWithdrawERC20, l1Token []common.Address, l2Token []common.Address, from []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _L2StandardERC20Gateway.contract.WatchLogs(opts, "WithdrawERC20", l1TokenRule, l2TokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2StandardERC20GatewayWithdrawERC20)
				if err := _L2StandardERC20Gateway.contract.UnpackLog(event, "WithdrawERC20", log); err != nil {
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
func (_L2StandardERC20Gateway *L2StandardERC20GatewayFilterer) ParseWithdrawERC20(log types.Log) (*L2StandardERC20GatewayWithdrawERC20, error) {
	event := new(L2StandardERC20GatewayWithdrawERC20)
	if err := _L2StandardERC20Gateway.contract.UnpackLog(event, "WithdrawERC20", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
