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

// L1StandardERC20GatewayMetaData contains all meta data concerning the L1StandardERC20Gateway contract.
var L1StandardERC20GatewayMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l2Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"DepositERC20\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l2Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"FinalizeWithdrawERC20\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"counterpart\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"}],\"name\":\"depositERC20\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"}],\"name\":\"depositERC20\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"}],\"name\":\"depositERC20AndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"finalizeWithdrawERC20\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"}],\"name\":\"getL2ERC20Address\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_counterpart\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_router\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_messenger\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l2TokenImplementation\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l2TokenFactory\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l2TokenFactory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l2TokenImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messenger\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"router\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"797594b0": "counterpart()",
		"f219fa66": "depositERC20(address,address,uint256,uint256)",
		"21425ee0": "depositERC20(address,uint256,uint256)",
		"0aea8c26": "depositERC20AndCall(address,address,uint256,bytes,uint256)",
		"84bd13b0": "finalizeWithdrawERC20(address,address,address,address,uint256,bytes)",
		"c676ad29": "getL2ERC20Address(address)",
		"1459457a": "initialize(address,address,address,address,address)",
		"eddd5e82": "l2TokenFactory()",
		"0e28c1f2": "l2TokenImplementation()",
		"3cb747bf": "messenger()",
		"f887ea40": "router()",
	},
	Bin: "0x608060405234801561001057600080fd5b506116da806100206000396000f3fe60806040526004361061009c5760003560e01c8063797594b011610064578063797594b01461014557806384bd13b01461016b578063c676ad291461017e578063eddd5e821461019e578063f219fa66146101be578063f887ea40146101d157600080fd5b80630aea8c26146100a15780630e28c1f2146100b65780631459457a146100f257806321425ee0146101125780633cb747bf14610125575b600080fd5b6100b46100af3660046110c1565b6101f1565b005b3480156100c257600080fd5b506004546100d6906001600160a01b031681565b6040516001600160a01b03909116815260200160405180910390f35b3480156100fe57600080fd5b506100b461010d366004611178565b610205565b6100b46101203660046111e9565b6103f2565b34801561013157600080fd5b506002546100d6906001600160a01b031681565b34801561015157600080fd5b506000546100d6906201000090046001600160a01b031681565b6100b461017936600461121e565b610431565b34801561018a57600080fd5b506100d66101993660046112e5565b61060c565b3480156101aa57600080fd5b506005546100d6906001600160a01b031681565b6100b46101cc366004611302565b61070e565b3480156101dd57600080fd5b506001546100d6906001600160a01b031681565b6101fe8585858585610721565b5050505050565b600054610100900460ff166102205760005460ff1615610224565b303b155b61028c5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084015b60405180910390fd5b600054610100900460ff161580156102ae576000805461ffff19166101011790555b6001600160a01b0385166102fa5760405162461bcd60e51b81526020600482015260136024820152727a65726f20726f75746572206164647265737360681b6044820152606401610283565b610305868686610c3c565b6001600160a01b03831661035b5760405162461bcd60e51b815260206004820152601860248201527f7a65726f20696d706c656d656e746174696f6e206861736800000000000000006044820152606401610283565b6001600160a01b0382166103a85760405162461bcd60e51b81526020600482015260146024820152737a65726f20666163746f7279206164647265737360601b6044820152606401610283565b600480546001600160a01b038086166001600160a01b031992831617909255600580549285169290911691909117905580156103ea576000805461ff00191690555b505050505050565b61042c83338460005b6040519080825280601f01601f191660200182016040528015610425576020820181803683370190505b5085610721565b505050565b6002546001600160a01b031633811461048c5760405162461bcd60e51b815260206004820152601760248201527f6f6e6c79206d657373656e6765722063616e2063616c6c0000000000000000006044820152606401610283565b806001600160a01b0316636e296e456040518163ffffffff1660e01b8152600401602060405180830381865afa1580156104ca573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104ee9190611348565b6000546201000090046001600160a01b039081169116146105515760405162461bcd60e51b815260206004820152601760248201527f6f6e6c792063616c6c20627920636f6e746572706172740000000000000000006044820152606401610283565b34156105935760405162461bcd60e51b81526020600482015260116024820152706e6f6e7a65726f206d73672e76616c756560781b6044820152606401610283565b6105a76001600160a01b0389168686610d4c565b856001600160a01b0316876001600160a01b0316896001600160a01b03167fc6f985873b37805705f6bce756dce3d1ff4b603e298d506288cce499926846a7888888886040516105fa9493929190611365565b60405180910390a45050505050505050565b600080546040516001600160601b0319606085901b16602082015282916201000090046001600160a01b0316906034016040516020818303038152906040528051906020012060405160200161067b92919060609290921b6001600160601b0319168252601482015260340190565b60408051808303601f19018152908290528051602090910120600454600554733d602d80600a3d3981f3363d3d373d3d3d363d7360601b84526001600160601b0319606092831b811660148601526f5af43d82803e903d91602b57fd5bf3ff60801b6028860152911b166038830152604c82018190526037808320606c84015260559201919091209091505b9392505050565b61071b84848460006103fb565b50505050565b6002600354036107735760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c006044820152606401610283565b6002600355826107bb5760405162461bcd60e51b815260206004820152601360248201527219195c1bdcda5d081e995c9bc8185b5bdd5b9d606a1b6044820152606401610283565b60015433906001600160a01b03168190036107e957828060200190518101906107e49190611401565b935090505b6040516370a0823160e01b81523060048201526000906001600160a01b038816906370a0823190602401602060405180830381865afa158015610830573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108549190611466565b905061086b6001600160a01b038816833088610daf565b6040516370a0823160e01b81523060048201526000906001600160a01b038916906370a0823190602401602060405180830381865afa1580156108b2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108d69190611466565b90506108e2828261147f565b95506000861161092a5760405162461bcd60e51b815260206004820152601360248201527219195c1bdcda5d081e995c9bc8185b5bdd5b9d606a1b6044820152606401610283565b50506001600160a01b03808716600090815260066020526040902054168381610b06576109568861060c565b6001600160a01b0389811660008181526006602052604080822080546001600160a01b0319169486169490941790935582516395d89b4160e01b815292519395509290916395d89b4191600480830192869291908290030181865afa1580156109c3573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526109eb91908101906114a6565b90506000896001600160a01b03166306fdde036040518163ffffffff1660e01b8152600401600060405180830381865afa158015610a2d573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610a5591908101906114a6565b905060008a6001600160a01b031663313ce5676040518163ffffffff1660e01b8152600401602060405180830381865afa158015610a97573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610abb91906114ef565b905087838383604051602001610ad39392919061153e565b60408051601f1981840301815290829052610af19291602001611577565b60405160208183030381529060405293505050505b6000638431f5c160e01b8984868b8b87604051602401610b2b969594939291906115a5565b60408051601f198184030181529181526020820180516001600160e01b03166001600160e01b03199094169390931790925260025460008054935163b2267a7b60e01b81529294506001600160a01b039182169363b2267a7b933493610ba1936201000090930416919087908c906004016115f4565b6000604051808303818588803b158015610bba57600080fd5b505af1158015610bce573d6000803e3d6000fd5b5050505050836001600160a01b0316836001600160a01b03168a6001600160a01b03167f31cd3b976e4d654022bf95c68a2ce53f1d5d94afabe0454d2832208eeb40af258b8b8b604051610c249392919061162c565b60405180910390a45050600160035550505050505050565b6001600160a01b038316610c925760405162461bcd60e51b815260206004820152601860248201527f7a65726f20636f756e74657270617274206164647265737300000000000000006044820152606401610283565b6001600160a01b038116610ce15760405162461bcd60e51b81526020600482015260166024820152757a65726f206d657373656e676572206164647265737360501b6044820152606401610283565b6000805462010000600160b01b031916620100006001600160a01b038681169190910291909117909155600280546001600160a01b031916838316179055821615610d4257600180546001600160a01b0319166001600160a01b0384161790555b5050600160035550565b6040516001600160a01b03831660248201526044810182905261042c90849063a9059cbb60e01b906064015b60408051601f198184030181529190526020810180516001600160e01b03166001600160e01b031990931692909217909152610de7565b6040516001600160a01b038085166024830152831660448201526064810182905261071b9085906323b872dd60e01b90608401610d78565b6000610e3c826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b0316610eb99092919063ffffffff16565b80519091501561042c5780806020019051810190610e5a9190611653565b61042c5760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b6064820152608401610283565b6060610ec88484600085610ed0565b949350505050565b606082471015610f315760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f6044820152651c8818d85b1b60d21b6064820152608401610283565b6001600160a01b0385163b610f885760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610283565b600080866001600160a01b03168587604051610fa49190611675565b60006040518083038185875af1925050503d8060008114610fe1576040519150601f19603f3d011682016040523d82523d6000602084013e610fe6565b606091505b5091509150610ff6828286611001565b979650505050505050565b60608315611010575081610707565b8251156110205782518084602001fd5b8160405162461bcd60e51b81526004016102839190611691565b6001600160a01b038116811461104f57600080fd5b50565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f1916810167ffffffffffffffff8111828210171561109157611091611052565b604052919050565b600067ffffffffffffffff8211156110b3576110b3611052565b50601f01601f191660200190565b600080600080600060a086880312156110d957600080fd5b85356110e48161103a565b945060208601356110f48161103a565b935060408601359250606086013567ffffffffffffffff81111561111757600080fd5b8601601f8101881361112857600080fd5b803561113b61113682611099565b611068565b81815289602083850101111561115057600080fd5b8160208401602083013760009181016020019190915295989497509295608001359392505050565b600080600080600060a0868803121561119057600080fd5b853561119b8161103a565b945060208601356111ab8161103a565b935060408601356111bb8161103a565b925060608601356111cb8161103a565b915060808601356111db8161103a565b809150509295509295909350565b6000806000606084860312156111fe57600080fd5b83356112098161103a565b95602085013595506040909401359392505050565b600080600080600080600060c0888a03121561123957600080fd5b87356112448161103a565b965060208801356112548161103a565b955060408801356112648161103a565b945060608801356112748161103a565b93506080880135925060a088013567ffffffffffffffff8082111561129857600080fd5b818a0191508a601f8301126112ac57600080fd5b8135818111156112bb57600080fd5b8b60208285010111156112cd57600080fd5b60208301945080935050505092959891949750929550565b6000602082840312156112f757600080fd5b81356107078161103a565b6000806000806080858703121561131857600080fd5b84356113238161103a565b935060208501356113338161103a565b93969395505050506040820135916060013590565b60006020828403121561135a57600080fd5b81516107078161103a565b6001600160a01b0385168152602081018490526060604082018190528101829052818360808301376000818301608090810191909152601f909201601f191601019392505050565b60005b838110156113c85781810151838201526020016113b0565b50506000910152565b60006113df61113684611099565b90508281528383830111156113f357600080fd5b6107078360208301846113ad565b6000806040838503121561141457600080fd5b825161141f8161103a565b602084015190925067ffffffffffffffff81111561143c57600080fd5b8301601f8101851361144d57600080fd5b61145c858251602084016113d1565b9150509250929050565b60006020828403121561147857600080fd5b5051919050565b818103818111156114a057634e487b7160e01b600052601160045260246000fd5b92915050565b6000602082840312156114b857600080fd5b815167ffffffffffffffff8111156114cf57600080fd5b8201601f810184136114e057600080fd5b610ec8848251602084016113d1565b60006020828403121561150157600080fd5b815160ff8116811461070757600080fd5b6000815180845261152a8160208601602086016113ad565b601f01601f19169290920160200192915050565b6060815260006115516060830186611512565b82810360208401526115638186611512565b91505060ff83166040830152949350505050565b60408152600061158a6040830185611512565b828103602084015261159c8185611512565b95945050505050565b6001600160a01b03878116825286811660208301528581166040830152841660608201526080810183905260c060a082018190526000906115e890830184611512565b98975050505050505050565b60018060a01b038516815283602082015260806040820152600061161b6080830185611512565b905082606083015295945050505050565b60018060a01b038416815282602082015260606040820152600061159c6060830184611512565b60006020828403121561166557600080fd5b8151801515811461070757600080fd5b600082516116878184602087016113ad565b9190910192915050565b602081526000610707602083018461151256fea2646970667358221220013b1640de503a48ac4f40802b5b9b9334526ebc30671f52cfa1832dfda8efec64736f6c63430008110033",
}

// L1StandardERC20GatewayABI is the input ABI used to generate the binding from.
// Deprecated: Use L1StandardERC20GatewayMetaData.ABI instead.
var L1StandardERC20GatewayABI = L1StandardERC20GatewayMetaData.ABI

// Deprecated: Use L1StandardERC20GatewayMetaData.Sigs instead.
// L1StandardERC20GatewayFuncSigs maps the 4-byte function signature to its string representation.
var L1StandardERC20GatewayFuncSigs = L1StandardERC20GatewayMetaData.Sigs

// L1StandardERC20GatewayBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use L1StandardERC20GatewayMetaData.Bin instead.
var L1StandardERC20GatewayBin = L1StandardERC20GatewayMetaData.Bin

// DeployL1StandardERC20Gateway deploys a new Ethereum contract, binding an instance of L1StandardERC20Gateway to it.
func DeployL1StandardERC20Gateway(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *L1StandardERC20Gateway, error) {
	parsed, err := L1StandardERC20GatewayMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(L1StandardERC20GatewayBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &L1StandardERC20Gateway{L1StandardERC20GatewayCaller: L1StandardERC20GatewayCaller{contract: contract}, L1StandardERC20GatewayTransactor: L1StandardERC20GatewayTransactor{contract: contract}, L1StandardERC20GatewayFilterer: L1StandardERC20GatewayFilterer{contract: contract}}, nil
}

// L1StandardERC20Gateway is an auto generated Go binding around an Ethereum contract.
type L1StandardERC20Gateway struct {
	L1StandardERC20GatewayCaller     // Read-only binding to the contract
	L1StandardERC20GatewayTransactor // Write-only binding to the contract
	L1StandardERC20GatewayFilterer   // Log filterer for contract events
}

// L1StandardERC20GatewayCaller is an auto generated read-only Go binding around an Ethereum contract.
type L1StandardERC20GatewayCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L1StandardERC20GatewayTransactor is an auto generated write-only Go binding around an Ethereum contract.
type L1StandardERC20GatewayTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L1StandardERC20GatewayFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type L1StandardERC20GatewayFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L1StandardERC20GatewaySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type L1StandardERC20GatewaySession struct {
	Contract     *L1StandardERC20Gateway // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// L1StandardERC20GatewayCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type L1StandardERC20GatewayCallerSession struct {
	Contract *L1StandardERC20GatewayCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// L1StandardERC20GatewayTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type L1StandardERC20GatewayTransactorSession struct {
	Contract     *L1StandardERC20GatewayTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// L1StandardERC20GatewayRaw is an auto generated low-level Go binding around an Ethereum contract.
type L1StandardERC20GatewayRaw struct {
	Contract *L1StandardERC20Gateway // Generic contract binding to access the raw methods on
}

// L1StandardERC20GatewayCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type L1StandardERC20GatewayCallerRaw struct {
	Contract *L1StandardERC20GatewayCaller // Generic read-only contract binding to access the raw methods on
}

// L1StandardERC20GatewayTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type L1StandardERC20GatewayTransactorRaw struct {
	Contract *L1StandardERC20GatewayTransactor // Generic write-only contract binding to access the raw methods on
}

// NewL1StandardERC20Gateway creates a new instance of L1StandardERC20Gateway, bound to a specific deployed contract.
func NewL1StandardERC20Gateway(address common.Address, backend bind.ContractBackend) (*L1StandardERC20Gateway, error) {
	contract, err := bindL1StandardERC20Gateway(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &L1StandardERC20Gateway{L1StandardERC20GatewayCaller: L1StandardERC20GatewayCaller{contract: contract}, L1StandardERC20GatewayTransactor: L1StandardERC20GatewayTransactor{contract: contract}, L1StandardERC20GatewayFilterer: L1StandardERC20GatewayFilterer{contract: contract}}, nil
}

// NewL1StandardERC20GatewayCaller creates a new read-only instance of L1StandardERC20Gateway, bound to a specific deployed contract.
func NewL1StandardERC20GatewayCaller(address common.Address, caller bind.ContractCaller) (*L1StandardERC20GatewayCaller, error) {
	contract, err := bindL1StandardERC20Gateway(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &L1StandardERC20GatewayCaller{contract: contract}, nil
}

// NewL1StandardERC20GatewayTransactor creates a new write-only instance of L1StandardERC20Gateway, bound to a specific deployed contract.
func NewL1StandardERC20GatewayTransactor(address common.Address, transactor bind.ContractTransactor) (*L1StandardERC20GatewayTransactor, error) {
	contract, err := bindL1StandardERC20Gateway(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &L1StandardERC20GatewayTransactor{contract: contract}, nil
}

// NewL1StandardERC20GatewayFilterer creates a new log filterer instance of L1StandardERC20Gateway, bound to a specific deployed contract.
func NewL1StandardERC20GatewayFilterer(address common.Address, filterer bind.ContractFilterer) (*L1StandardERC20GatewayFilterer, error) {
	contract, err := bindL1StandardERC20Gateway(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &L1StandardERC20GatewayFilterer{contract: contract}, nil
}

// bindL1StandardERC20Gateway binds a generic wrapper to an already deployed contract.
func bindL1StandardERC20Gateway(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(L1StandardERC20GatewayABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L1StandardERC20Gateway *L1StandardERC20GatewayRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L1StandardERC20Gateway.Contract.L1StandardERC20GatewayCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L1StandardERC20Gateway *L1StandardERC20GatewayRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1StandardERC20Gateway.Contract.L1StandardERC20GatewayTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L1StandardERC20Gateway *L1StandardERC20GatewayRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L1StandardERC20Gateway.Contract.L1StandardERC20GatewayTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L1StandardERC20Gateway *L1StandardERC20GatewayCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L1StandardERC20Gateway.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L1StandardERC20Gateway *L1StandardERC20GatewayTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1StandardERC20Gateway.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L1StandardERC20Gateway *L1StandardERC20GatewayTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L1StandardERC20Gateway.Contract.contract.Transact(opts, method, params...)
}

// Counterpart is a free data retrieval call binding the contract method 0x797594b0.
//
// Solidity: function counterpart() view returns(address)
func (_L1StandardERC20Gateway *L1StandardERC20GatewayCaller) Counterpart(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1StandardERC20Gateway.contract.Call(opts, &out, "counterpart")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Counterpart is a free data retrieval call binding the contract method 0x797594b0.
//
// Solidity: function counterpart() view returns(address)
func (_L1StandardERC20Gateway *L1StandardERC20GatewaySession) Counterpart() (common.Address, error) {
	return _L1StandardERC20Gateway.Contract.Counterpart(&_L1StandardERC20Gateway.CallOpts)
}

// Counterpart is a free data retrieval call binding the contract method 0x797594b0.
//
// Solidity: function counterpart() view returns(address)
func (_L1StandardERC20Gateway *L1StandardERC20GatewayCallerSession) Counterpart() (common.Address, error) {
	return _L1StandardERC20Gateway.Contract.Counterpart(&_L1StandardERC20Gateway.CallOpts)
}

// GetL2ERC20Address is a free data retrieval call binding the contract method 0xc676ad29.
//
// Solidity: function getL2ERC20Address(address _l1Token) view returns(address)
func (_L1StandardERC20Gateway *L1StandardERC20GatewayCaller) GetL2ERC20Address(opts *bind.CallOpts, _l1Token common.Address) (common.Address, error) {
	var out []interface{}
	err := _L1StandardERC20Gateway.contract.Call(opts, &out, "getL2ERC20Address", _l1Token)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetL2ERC20Address is a free data retrieval call binding the contract method 0xc676ad29.
//
// Solidity: function getL2ERC20Address(address _l1Token) view returns(address)
func (_L1StandardERC20Gateway *L1StandardERC20GatewaySession) GetL2ERC20Address(_l1Token common.Address) (common.Address, error) {
	return _L1StandardERC20Gateway.Contract.GetL2ERC20Address(&_L1StandardERC20Gateway.CallOpts, _l1Token)
}

// GetL2ERC20Address is a free data retrieval call binding the contract method 0xc676ad29.
//
// Solidity: function getL2ERC20Address(address _l1Token) view returns(address)
func (_L1StandardERC20Gateway *L1StandardERC20GatewayCallerSession) GetL2ERC20Address(_l1Token common.Address) (common.Address, error) {
	return _L1StandardERC20Gateway.Contract.GetL2ERC20Address(&_L1StandardERC20Gateway.CallOpts, _l1Token)
}

// L2TokenFactory is a free data retrieval call binding the contract method 0xeddd5e82.
//
// Solidity: function l2TokenFactory() view returns(address)
func (_L1StandardERC20Gateway *L1StandardERC20GatewayCaller) L2TokenFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1StandardERC20Gateway.contract.Call(opts, &out, "l2TokenFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L2TokenFactory is a free data retrieval call binding the contract method 0xeddd5e82.
//
// Solidity: function l2TokenFactory() view returns(address)
func (_L1StandardERC20Gateway *L1StandardERC20GatewaySession) L2TokenFactory() (common.Address, error) {
	return _L1StandardERC20Gateway.Contract.L2TokenFactory(&_L1StandardERC20Gateway.CallOpts)
}

// L2TokenFactory is a free data retrieval call binding the contract method 0xeddd5e82.
//
// Solidity: function l2TokenFactory() view returns(address)
func (_L1StandardERC20Gateway *L1StandardERC20GatewayCallerSession) L2TokenFactory() (common.Address, error) {
	return _L1StandardERC20Gateway.Contract.L2TokenFactory(&_L1StandardERC20Gateway.CallOpts)
}

// L2TokenImplementation is a free data retrieval call binding the contract method 0x0e28c1f2.
//
// Solidity: function l2TokenImplementation() view returns(address)
func (_L1StandardERC20Gateway *L1StandardERC20GatewayCaller) L2TokenImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1StandardERC20Gateway.contract.Call(opts, &out, "l2TokenImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L2TokenImplementation is a free data retrieval call binding the contract method 0x0e28c1f2.
//
// Solidity: function l2TokenImplementation() view returns(address)
func (_L1StandardERC20Gateway *L1StandardERC20GatewaySession) L2TokenImplementation() (common.Address, error) {
	return _L1StandardERC20Gateway.Contract.L2TokenImplementation(&_L1StandardERC20Gateway.CallOpts)
}

// L2TokenImplementation is a free data retrieval call binding the contract method 0x0e28c1f2.
//
// Solidity: function l2TokenImplementation() view returns(address)
func (_L1StandardERC20Gateway *L1StandardERC20GatewayCallerSession) L2TokenImplementation() (common.Address, error) {
	return _L1StandardERC20Gateway.Contract.L2TokenImplementation(&_L1StandardERC20Gateway.CallOpts)
}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_L1StandardERC20Gateway *L1StandardERC20GatewayCaller) Messenger(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1StandardERC20Gateway.contract.Call(opts, &out, "messenger")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_L1StandardERC20Gateway *L1StandardERC20GatewaySession) Messenger() (common.Address, error) {
	return _L1StandardERC20Gateway.Contract.Messenger(&_L1StandardERC20Gateway.CallOpts)
}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_L1StandardERC20Gateway *L1StandardERC20GatewayCallerSession) Messenger() (common.Address, error) {
	return _L1StandardERC20Gateway.Contract.Messenger(&_L1StandardERC20Gateway.CallOpts)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_L1StandardERC20Gateway *L1StandardERC20GatewayCaller) Router(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1StandardERC20Gateway.contract.Call(opts, &out, "router")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_L1StandardERC20Gateway *L1StandardERC20GatewaySession) Router() (common.Address, error) {
	return _L1StandardERC20Gateway.Contract.Router(&_L1StandardERC20Gateway.CallOpts)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_L1StandardERC20Gateway *L1StandardERC20GatewayCallerSession) Router() (common.Address, error) {
	return _L1StandardERC20Gateway.Contract.Router(&_L1StandardERC20Gateway.CallOpts)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0x21425ee0.
//
// Solidity: function depositERC20(address _token, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L1StandardERC20Gateway *L1StandardERC20GatewayTransactor) DepositERC20(opts *bind.TransactOpts, _token common.Address, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L1StandardERC20Gateway.contract.Transact(opts, "depositERC20", _token, _amount, _gasLimit)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0x21425ee0.
//
// Solidity: function depositERC20(address _token, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L1StandardERC20Gateway *L1StandardERC20GatewaySession) DepositERC20(_token common.Address, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L1StandardERC20Gateway.Contract.DepositERC20(&_L1StandardERC20Gateway.TransactOpts, _token, _amount, _gasLimit)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0x21425ee0.
//
// Solidity: function depositERC20(address _token, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L1StandardERC20Gateway *L1StandardERC20GatewayTransactorSession) DepositERC20(_token common.Address, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L1StandardERC20Gateway.Contract.DepositERC20(&_L1StandardERC20Gateway.TransactOpts, _token, _amount, _gasLimit)
}

// DepositERC200 is a paid mutator transaction binding the contract method 0xf219fa66.
//
// Solidity: function depositERC20(address _token, address _to, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L1StandardERC20Gateway *L1StandardERC20GatewayTransactor) DepositERC200(opts *bind.TransactOpts, _token common.Address, _to common.Address, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L1StandardERC20Gateway.contract.Transact(opts, "depositERC200", _token, _to, _amount, _gasLimit)
}

// DepositERC200 is a paid mutator transaction binding the contract method 0xf219fa66.
//
// Solidity: function depositERC20(address _token, address _to, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L1StandardERC20Gateway *L1StandardERC20GatewaySession) DepositERC200(_token common.Address, _to common.Address, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L1StandardERC20Gateway.Contract.DepositERC200(&_L1StandardERC20Gateway.TransactOpts, _token, _to, _amount, _gasLimit)
}

// DepositERC200 is a paid mutator transaction binding the contract method 0xf219fa66.
//
// Solidity: function depositERC20(address _token, address _to, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L1StandardERC20Gateway *L1StandardERC20GatewayTransactorSession) DepositERC200(_token common.Address, _to common.Address, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L1StandardERC20Gateway.Contract.DepositERC200(&_L1StandardERC20Gateway.TransactOpts, _token, _to, _amount, _gasLimit)
}

// DepositERC20AndCall is a paid mutator transaction binding the contract method 0x0aea8c26.
//
// Solidity: function depositERC20AndCall(address _token, address _to, uint256 _amount, bytes _data, uint256 _gasLimit) payable returns()
func (_L1StandardERC20Gateway *L1StandardERC20GatewayTransactor) DepositERC20AndCall(opts *bind.TransactOpts, _token common.Address, _to common.Address, _amount *big.Int, _data []byte, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L1StandardERC20Gateway.contract.Transact(opts, "depositERC20AndCall", _token, _to, _amount, _data, _gasLimit)
}

// DepositERC20AndCall is a paid mutator transaction binding the contract method 0x0aea8c26.
//
// Solidity: function depositERC20AndCall(address _token, address _to, uint256 _amount, bytes _data, uint256 _gasLimit) payable returns()
func (_L1StandardERC20Gateway *L1StandardERC20GatewaySession) DepositERC20AndCall(_token common.Address, _to common.Address, _amount *big.Int, _data []byte, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L1StandardERC20Gateway.Contract.DepositERC20AndCall(&_L1StandardERC20Gateway.TransactOpts, _token, _to, _amount, _data, _gasLimit)
}

// DepositERC20AndCall is a paid mutator transaction binding the contract method 0x0aea8c26.
//
// Solidity: function depositERC20AndCall(address _token, address _to, uint256 _amount, bytes _data, uint256 _gasLimit) payable returns()
func (_L1StandardERC20Gateway *L1StandardERC20GatewayTransactorSession) DepositERC20AndCall(_token common.Address, _to common.Address, _amount *big.Int, _data []byte, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L1StandardERC20Gateway.Contract.DepositERC20AndCall(&_L1StandardERC20Gateway.TransactOpts, _token, _to, _amount, _data, _gasLimit)
}

// FinalizeWithdrawERC20 is a paid mutator transaction binding the contract method 0x84bd13b0.
//
// Solidity: function finalizeWithdrawERC20(address _l1Token, address _l2Token, address _from, address _to, uint256 _amount, bytes _data) payable returns()
func (_L1StandardERC20Gateway *L1StandardERC20GatewayTransactor) FinalizeWithdrawERC20(opts *bind.TransactOpts, _l1Token common.Address, _l2Token common.Address, _from common.Address, _to common.Address, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _L1StandardERC20Gateway.contract.Transact(opts, "finalizeWithdrawERC20", _l1Token, _l2Token, _from, _to, _amount, _data)
}

// FinalizeWithdrawERC20 is a paid mutator transaction binding the contract method 0x84bd13b0.
//
// Solidity: function finalizeWithdrawERC20(address _l1Token, address _l2Token, address _from, address _to, uint256 _amount, bytes _data) payable returns()
func (_L1StandardERC20Gateway *L1StandardERC20GatewaySession) FinalizeWithdrawERC20(_l1Token common.Address, _l2Token common.Address, _from common.Address, _to common.Address, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _L1StandardERC20Gateway.Contract.FinalizeWithdrawERC20(&_L1StandardERC20Gateway.TransactOpts, _l1Token, _l2Token, _from, _to, _amount, _data)
}

// FinalizeWithdrawERC20 is a paid mutator transaction binding the contract method 0x84bd13b0.
//
// Solidity: function finalizeWithdrawERC20(address _l1Token, address _l2Token, address _from, address _to, uint256 _amount, bytes _data) payable returns()
func (_L1StandardERC20Gateway *L1StandardERC20GatewayTransactorSession) FinalizeWithdrawERC20(_l1Token common.Address, _l2Token common.Address, _from common.Address, _to common.Address, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _L1StandardERC20Gateway.Contract.FinalizeWithdrawERC20(&_L1StandardERC20Gateway.TransactOpts, _l1Token, _l2Token, _from, _to, _amount, _data)
}

// Initialize is a paid mutator transaction binding the contract method 0x1459457a.
//
// Solidity: function initialize(address _counterpart, address _router, address _messenger, address _l2TokenImplementation, address _l2TokenFactory) returns()
func (_L1StandardERC20Gateway *L1StandardERC20GatewayTransactor) Initialize(opts *bind.TransactOpts, _counterpart common.Address, _router common.Address, _messenger common.Address, _l2TokenImplementation common.Address, _l2TokenFactory common.Address) (*types.Transaction, error) {
	return _L1StandardERC20Gateway.contract.Transact(opts, "initialize", _counterpart, _router, _messenger, _l2TokenImplementation, _l2TokenFactory)
}

// Initialize is a paid mutator transaction binding the contract method 0x1459457a.
//
// Solidity: function initialize(address _counterpart, address _router, address _messenger, address _l2TokenImplementation, address _l2TokenFactory) returns()
func (_L1StandardERC20Gateway *L1StandardERC20GatewaySession) Initialize(_counterpart common.Address, _router common.Address, _messenger common.Address, _l2TokenImplementation common.Address, _l2TokenFactory common.Address) (*types.Transaction, error) {
	return _L1StandardERC20Gateway.Contract.Initialize(&_L1StandardERC20Gateway.TransactOpts, _counterpart, _router, _messenger, _l2TokenImplementation, _l2TokenFactory)
}

// Initialize is a paid mutator transaction binding the contract method 0x1459457a.
//
// Solidity: function initialize(address _counterpart, address _router, address _messenger, address _l2TokenImplementation, address _l2TokenFactory) returns()
func (_L1StandardERC20Gateway *L1StandardERC20GatewayTransactorSession) Initialize(_counterpart common.Address, _router common.Address, _messenger common.Address, _l2TokenImplementation common.Address, _l2TokenFactory common.Address) (*types.Transaction, error) {
	return _L1StandardERC20Gateway.Contract.Initialize(&_L1StandardERC20Gateway.TransactOpts, _counterpart, _router, _messenger, _l2TokenImplementation, _l2TokenFactory)
}

// L1StandardERC20GatewayDepositERC20Iterator is returned from FilterDepositERC20 and is used to iterate over the raw logs and unpacked data for DepositERC20 events raised by the L1StandardERC20Gateway contract.
type L1StandardERC20GatewayDepositERC20Iterator struct {
	Event *L1StandardERC20GatewayDepositERC20 // Event containing the contract specifics and raw log

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
func (it *L1StandardERC20GatewayDepositERC20Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1StandardERC20GatewayDepositERC20)
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
		it.Event = new(L1StandardERC20GatewayDepositERC20)
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
func (it *L1StandardERC20GatewayDepositERC20Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1StandardERC20GatewayDepositERC20Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1StandardERC20GatewayDepositERC20 represents a DepositERC20 event raised by the L1StandardERC20Gateway contract.
type L1StandardERC20GatewayDepositERC20 struct {
	L1Token common.Address
	L2Token common.Address
	From    common.Address
	To      common.Address
	Amount  *big.Int
	Data    []byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDepositERC20 is a free log retrieval operation binding the contract event 0x31cd3b976e4d654022bf95c68a2ce53f1d5d94afabe0454d2832208eeb40af25.
//
// Solidity: event DepositERC20(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256 amount, bytes data)
func (_L1StandardERC20Gateway *L1StandardERC20GatewayFilterer) FilterDepositERC20(opts *bind.FilterOpts, l1Token []common.Address, l2Token []common.Address, from []common.Address) (*L1StandardERC20GatewayDepositERC20Iterator, error) {

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

	logs, sub, err := _L1StandardERC20Gateway.contract.FilterLogs(opts, "DepositERC20", l1TokenRule, l2TokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &L1StandardERC20GatewayDepositERC20Iterator{contract: _L1StandardERC20Gateway.contract, event: "DepositERC20", logs: logs, sub: sub}, nil
}

// WatchDepositERC20 is a free log subscription operation binding the contract event 0x31cd3b976e4d654022bf95c68a2ce53f1d5d94afabe0454d2832208eeb40af25.
//
// Solidity: event DepositERC20(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256 amount, bytes data)
func (_L1StandardERC20Gateway *L1StandardERC20GatewayFilterer) WatchDepositERC20(opts *bind.WatchOpts, sink chan<- *L1StandardERC20GatewayDepositERC20, l1Token []common.Address, l2Token []common.Address, from []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _L1StandardERC20Gateway.contract.WatchLogs(opts, "DepositERC20", l1TokenRule, l2TokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1StandardERC20GatewayDepositERC20)
				if err := _L1StandardERC20Gateway.contract.UnpackLog(event, "DepositERC20", log); err != nil {
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

// ParseDepositERC20 is a log parse operation binding the contract event 0x31cd3b976e4d654022bf95c68a2ce53f1d5d94afabe0454d2832208eeb40af25.
//
// Solidity: event DepositERC20(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256 amount, bytes data)
func (_L1StandardERC20Gateway *L1StandardERC20GatewayFilterer) ParseDepositERC20(log types.Log) (*L1StandardERC20GatewayDepositERC20, error) {
	event := new(L1StandardERC20GatewayDepositERC20)
	if err := _L1StandardERC20Gateway.contract.UnpackLog(event, "DepositERC20", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1StandardERC20GatewayFinalizeWithdrawERC20Iterator is returned from FilterFinalizeWithdrawERC20 and is used to iterate over the raw logs and unpacked data for FinalizeWithdrawERC20 events raised by the L1StandardERC20Gateway contract.
type L1StandardERC20GatewayFinalizeWithdrawERC20Iterator struct {
	Event *L1StandardERC20GatewayFinalizeWithdrawERC20 // Event containing the contract specifics and raw log

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
func (it *L1StandardERC20GatewayFinalizeWithdrawERC20Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1StandardERC20GatewayFinalizeWithdrawERC20)
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
		it.Event = new(L1StandardERC20GatewayFinalizeWithdrawERC20)
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
func (it *L1StandardERC20GatewayFinalizeWithdrawERC20Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1StandardERC20GatewayFinalizeWithdrawERC20Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1StandardERC20GatewayFinalizeWithdrawERC20 represents a FinalizeWithdrawERC20 event raised by the L1StandardERC20Gateway contract.
type L1StandardERC20GatewayFinalizeWithdrawERC20 struct {
	L1Token common.Address
	L2Token common.Address
	From    common.Address
	To      common.Address
	Amount  *big.Int
	Data    []byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterFinalizeWithdrawERC20 is a free log retrieval operation binding the contract event 0xc6f985873b37805705f6bce756dce3d1ff4b603e298d506288cce499926846a7.
//
// Solidity: event FinalizeWithdrawERC20(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256 amount, bytes data)
func (_L1StandardERC20Gateway *L1StandardERC20GatewayFilterer) FilterFinalizeWithdrawERC20(opts *bind.FilterOpts, l1Token []common.Address, l2Token []common.Address, from []common.Address) (*L1StandardERC20GatewayFinalizeWithdrawERC20Iterator, error) {

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

	logs, sub, err := _L1StandardERC20Gateway.contract.FilterLogs(opts, "FinalizeWithdrawERC20", l1TokenRule, l2TokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &L1StandardERC20GatewayFinalizeWithdrawERC20Iterator{contract: _L1StandardERC20Gateway.contract, event: "FinalizeWithdrawERC20", logs: logs, sub: sub}, nil
}

// WatchFinalizeWithdrawERC20 is a free log subscription operation binding the contract event 0xc6f985873b37805705f6bce756dce3d1ff4b603e298d506288cce499926846a7.
//
// Solidity: event FinalizeWithdrawERC20(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256 amount, bytes data)
func (_L1StandardERC20Gateway *L1StandardERC20GatewayFilterer) WatchFinalizeWithdrawERC20(opts *bind.WatchOpts, sink chan<- *L1StandardERC20GatewayFinalizeWithdrawERC20, l1Token []common.Address, l2Token []common.Address, from []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _L1StandardERC20Gateway.contract.WatchLogs(opts, "FinalizeWithdrawERC20", l1TokenRule, l2TokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1StandardERC20GatewayFinalizeWithdrawERC20)
				if err := _L1StandardERC20Gateway.contract.UnpackLog(event, "FinalizeWithdrawERC20", log); err != nil {
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

// ParseFinalizeWithdrawERC20 is a log parse operation binding the contract event 0xc6f985873b37805705f6bce756dce3d1ff4b603e298d506288cce499926846a7.
//
// Solidity: event FinalizeWithdrawERC20(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256 amount, bytes data)
func (_L1StandardERC20Gateway *L1StandardERC20GatewayFilterer) ParseFinalizeWithdrawERC20(log types.Log) (*L1StandardERC20GatewayFinalizeWithdrawERC20, error) {
	event := new(L1StandardERC20GatewayFinalizeWithdrawERC20)
	if err := _L1StandardERC20Gateway.contract.UnpackLog(event, "FinalizeWithdrawERC20", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
