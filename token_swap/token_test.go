package token_swap

import (
	"context"
	"encoding/json"
	"github.com/scroll-tech/go-ethereum/accounts/abi/bind"
	"github.com/scroll-tech/go-ethereum/common"
	"github.com/scroll-tech/go-ethereum/core/types"
	"github.com/scroll-tech/go-ethereum/crypto"
	"github.com/scroll-tech/go-ethereum/ethclient"
	"github.com/stretchr/testify/assert"
	"math/big"
	"os"
	"testing"
)

var (
	owner = common.HexToAddress("0xBC01a540724B544DA8DaB1c1Da507bE2e6821d46")

	l1URL = "https://sepolia.infura.io/v3/d869f8d95f1e467b97e9865d8decff12" //"http://sepolia-l1geth.scroll.tech:8545/l1"
	l2URL = "http://sepolia-l2geth-rpc-internal0.scroll.tech:8545"

	l1GasPriceOracleAddr = common.HexToAddress("0x5300000000000000000000000000000000000002")

	l1LinkAddr         = common.HexToAddress("0x779877A7B0D9E8603169DdbD7836e478b4624789")
	l1ERC20GatewayAddr = common.HexToAddress("0x362675E7bBb60b0fcAe52F91cFe33087f33D512A")

	l2LinkAddr         = common.HexToAddress("0x535E8a3E95ADe8d2199562192EFbb060071A55d7")
	l2ERC20GatewayAddr = common.HexToAddress("0x88A59756752854C7A494d76726Df64230Dcf7Ac5")
)

func TestDeposit(t *testing.T) {
	var (
		ctx                  = context.Background()
		standardERC20Addr    = common.HexToAddress("0x362675E7bBb60b0fcAe52F91cFe33087f33D512A")
		l2GasPriceOracleAddr = common.HexToAddress("0x9256C88C677497Cd42780Ce1D5c962a94cbF4eB3")
		//messgeQueueAddr   = common.HexToAddress("0xe7D7cB724c70B7413d4dc5494bEA51cFe927B403")
	)
	l1Cli, err := ethclient.Dial(l1URL)
	assert.NoError(t, err)

	l2GasPriceOracle, err := NewL2GasPriceOracle(l2GasPriceOracleAddr, l1Cli)
	assert.NoError(t, err)
	l2BaseFee, err := l2GasPriceOracle.L2BaseFee(nil)
	assert.NoError(t, err)
	t.Log("l2 baseFee: ", l2BaseFee.String())

	chainID, err := l1Cli.ChainID(ctx)
	assert.NoError(t, err)
	privk, err := crypto.HexToECDSA("e66f2fe3bf37ba70cf9c1ce6d88eac18ae5907f55663f92afcd19101fb5c964c")
	assert.NoError(t, err)
	auth, err := bind.NewKeyedTransactorWithChainID(privk, chainID)
	assert.NoError(t, err)

	bls, err := l1Cli.BalanceAt(ctx, owner, nil)
	assert.NoError(t, err)
	t.Log("balance: ", bls.String())

	t.Log("chainID: ", chainID.String())

	l1Link, err := NewERC20(l1LinkAddr, l1Cli)
	assert.NoError(t, err)
	l1Token, err := l1Link.BalanceOf(nil, owner)
	assert.NoError(t, err)
	t.Log("l1 chain link balance: ", l1Token.String())

	linkBls := big.NewInt(0).SetBytes(common.FromHex("8ac7230489e80000")) // 10 eth
	tx, err := l1Link.Approve(auth, standardERC20Addr, linkBls)
	assert.NoError(t, err)
	receipt, err := bind.WaitMined(ctx, l1Cli, tx)
	assert.NoError(t, err)
	assert.Equal(t, types.ReceiptStatusSuccessful, receipt.Status)

	standardERC20, err := NewL1StandardERC20Gateway(standardERC20Addr, l1Cli)
	assert.NoError(t, err)

	l2TokenAddr, err := standardERC20.GetL2ERC20Address(nil, l1LinkAddr)
	assert.NoError(t, err)
	t.Log("l2 link address: ", l2TokenAddr.String())

	auth.Value = big.NewInt(1e16)
	tx, err = standardERC20.DepositERC20(auth, l1LinkAddr, linkBls, big.NewInt(1e6))
	assert.NoError(t, err)
	t.Log("l1 deposit tx hash: ", tx.Hash().String())
	receipt, err = bind.WaitMined(ctx, l1Cli, tx)
	assert.NoError(t, err)
	assert.Equal(t, types.ReceiptStatusSuccessful, receipt.Status)
}

func TestGetTrace(t *testing.T) {
	var (
		ctx   = context.Background()
		l2URL = "http://sepolia-l2geth-rpc-internal0.scroll.tech:8545"
	)
	l2Cli, err := ethclient.Dial(l2URL)
	assert.NoError(t, err)

	// Get deposit ERC20 trace.
	trace, err := l2Cli.GetBlockTraceByNumber(ctx, big.NewInt(13174))
	assert.NoError(t, err)
	buf, err := json.Marshal(trace)
	assert.NoError(t, err)
	err = os.WriteFile("/Users/maskpp/projects/scroll-tech/test-traces/bridge/depositERC20.json", buf, 0644)
	assert.NoError(t, err)

	// Get approve ERC20 trace.
	trace, err = l2Cli.GetBlockTraceByNumber(ctx, big.NewInt(13201))
	assert.NoError(t, err)
	buf, err = json.Marshal(trace)
	assert.NoError(t, err)
	err = os.WriteFile("/Users/maskpp/projects/scroll-tech/test-traces/bridge/approveERC20.json", buf, 0644)
	assert.NoError(t, err)

	// Get withdraw ERC20 trace.
	trace, err = l2Cli.GetBlockTraceByNumber(ctx, big.NewInt(13205))
	assert.NoError(t, err)
	buf, err = json.Marshal(trace)
	assert.NoError(t, err)
	err = os.WriteFile("/Users/maskpp/projects/scroll-tech/test-traces/bridge/withdrawERC20.json", buf, 0644)
	assert.NoError(t, err)
}

func TestReceiptCheck(t *testing.T) {
	var (
		ctx    = context.Background()
		txHash = common.HexToHash("0xd8affc45f31a7dc7c1b5cfc1d54e1b5fbfbd9951bf7d4263f97eb7b855154359")
	)
	l2Cli, err := ethclient.Dial(l2URL)
	assert.NoError(t, err)

	l1GasPriceOracle, err := NewL1GasPriceOracle(l1GasPriceOracleAddr, l2Cli)
	assert.NoError(t, err)

	l1BaseFee, err := l1GasPriceOracle.L1BaseFee(nil)
	assert.NoError(t, err)
	t.Log("l1 baseFee: ", l1BaseFee.String())

	receipt, err := l2Cli.TransactionReceipt(ctx, txHash)
	assert.NoError(t, err)
	t.Log("l1 fee: ", receipt.L1Fee.String())
}

func TestWithdraw(t *testing.T) {
	var (
		ctx               = context.Background()
		standardERC20Addr = common.HexToAddress("0x88A59756752854C7A494d76726Df64230Dcf7Ac5")
		linkBls           = big.NewInt(0).SetBytes(common.FromHex("de0b6b3a7640000"))
	)
	l2Cli, err := ethclient.Dial(l2URL)
	assert.NoError(t, err)
	linkToken, err := NewERC20(l2LinkAddr, l2Cli)
	assert.NoError(t, err)

	chainID, err := l2Cli.ChainID(ctx)
	assert.NoError(t, err)
	privk, err := crypto.HexToECDSA("e66f2fe3bf37ba70cf9c1ce6d88eac18ae5907f55663f92afcd19101fb5c964c")
	assert.NoError(t, err)
	auth, err := bind.NewKeyedTransactorWithChainID(privk, chainID)
	assert.NoError(t, err)

	allowBls, _ := linkToken.Allowance(nil, owner, standardERC20Addr)
	if allowBls != nil && allowBls.Cmp(linkBls) < 0 {
		tx, err := linkToken.Approve(auth, standardERC20Addr, linkBls)
		assert.NoError(t, err)
		receipt, err := bind.WaitMined(ctx, l2Cli, tx)
		assert.NoError(t, err)
		assert.Equal(t, types.ReceiptStatusSuccessful, receipt.Status)
		t.Log("l2 chain ERC20 approve tx hash: ", tx.Hash().String())
	}

	standardERC20, err := NewL2StandardERC20Gateway(standardERC20Addr, l2Cli)
	assert.NoError(t, err)
	l1TokenAddr, err := standardERC20.GetL1ERC20Address(nil, l2LinkAddr)
	assert.NoError(t, err)
	assert.Equal(t, l1LinkAddr, l1TokenAddr)
	t.Log("actual l1 link contract address: ", l1TokenAddr.String())

	// withdraw erc20 token.
	auth.Value = big.NewInt(0)
	tx, err := standardERC20.WithdrawERC20(auth, l2LinkAddr, linkBls, big.NewInt(1e6))
	assert.NoError(t, err)
	receipt, err := bind.WaitMined(ctx, l2Cli, tx)
	assert.NoError(t, err)
	assert.Equal(t, types.ReceiptStatusSuccessful, receipt.Status)
}

func TestCC(t *testing.T) {
	var (
		ctx   = context.Background()
		l2URL = "http://sepolia-l2geth-rpc-internal0.scroll.tech:8545" //"https://sepolia-rpc.scroll.io/l2"
		owner = common.HexToAddress("0xBC01a540724B544DA8DaB1c1Da507bE2e6821d46")
	)

	l2Cli, err := ethclient.Dial(l2URL)
	assert.NoError(t, err)
	l2Bls, err := l2Cli.BalanceAt(ctx, owner, nil)
	assert.NoError(t, err)
	t.Log("l2 chain owner balance: ", l2Bls)
}
