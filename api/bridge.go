package api

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/scroll-tech/go-ethereum/core/types"
	"math/big"
	"os"
	"strings"
	"tool/contracts/scroll/L1"
	"tool/contracts/scroll/L2"

	"github.com/scroll-tech/go-ethereum/accounts/abi/bind"
	"github.com/scroll-tech/go-ethereum/common"
	"github.com/scroll-tech/go-ethereum/ethclient"
	"github.com/scroll-tech/go-ethereum/log"

	l1gateway "tool/contracts/scroll/L1/gateways"
	l2gateway "tool/contracts/scroll/L2/gateways"
)

func NewBridge(ctx context.Context, sk *ecdsa.PrivateKey, contractsFile string, l1Cli, l2Cli *ethclient.Client) error {
	contracts, err := NewContracts(ctx, sk, l1Cli, l2Cli, contractsFile)
	if err != nil {
		return err
	}
	// get 01~20 number traces.
	if err = getInitTraces(ctx, l2Cli, 1, 20); err != nil {
		return err
	}

	// get deposit trace.
	if err = ethDeposit(ctx, contracts); err != nil {
		return err
	}
	// get withdraw trace.
	if err = ethWithdraw(ctx, contracts); err != nil {
		return err
	}

	// get deposit, approve, withdraw traces.
	if err = erc20DepositAndWithdraw(ctx, contracts); err != nil {
		return err
	}

	return nil
}

// get batch traces by numbers.
func getInitTraces(ctx context.Context, l2Cli *ethclient.Client, start, end int) error {
	path := TRACEDATA_DIR_PREFIX + "bridge/"
	for i := start; i <= end; i++ {
		if err := storeBlockResultsForBlocks(ctx, l2Cli, path, fmt.Sprintf("%2d", i), []*big.Int{big.NewInt(int64(i))}); err != nil {
			return err
		}
	}
	return nil
}

// deposit eth and get trace on l2chain.
func ethDeposit(ctx context.Context, contracts *Contracts) error {

	// get l2chain number
	l2Number, err := contracts.L2CLi.BlockNumber(ctx)
	if err != nil {
		return err
	}

	l1Auth := contracts.L1Auth
	l1Auth.Value = big.NewInt(1e15)
	defer func() { l1Auth.Value = nil }()

	// make a deposit transaction
	tx, err := contracts.IL1ETHGateway.DepositETH(l1Auth, big.NewInt(1e14), big.NewInt(1e6))
	if err != nil {
		return err
	}
	receipt, err := bind.WaitMined(ctx, contracts.L1CLi, tx)
	if err != nil {
		return err
	}
	if receipt.Status != types.ReceiptStatusSuccessful {
		return fmt.Errorf("deposit tx(%s) status is fail", tx.Hash().String())
	}

	iter, err := contracts.IL2ETHGateway.FilterFinalizeDepositETH(
		&bind.FilterOpts{
			Start:   l2Number,
			Context: ctx,
		},
		[]common.Address{l1Auth.From},
		nil,
	)
	if err != nil {
		return err
	}

	path := TRACEDATA_DIR_PREFIX + "bridge/"
	for iter.Next() {
		return storeBlockResultsForBlocks(
			ctx,
			contracts.L2CLi,
			path,
			"depositETH",
			[]*big.Int{big.NewInt(0).SetUint64(iter.Event.Raw.BlockNumber)},
		)
	}
	return nil
}

func ethWithdraw(ctx context.Context, contracts *Contracts) error {
	// get l2chain number
	l2Number, err := contracts.L2CLi.BlockNumber(ctx)
	if err != nil {
		return err
	}

	l2Auth := contracts.L2Auth
	l2Auth.Value = big.NewInt(1e16)
	defer func() { l2Auth.Value = nil }()
	tx, err := contracts.IL2ETHGateway.WithdrawETH0(l2Auth, big.NewInt(1e15), big.NewInt(1e6))
	if err != nil {
		return err
	}
	receipt, err := bind.WaitMined(ctx, contracts.L2CLi, tx)
	if err != nil {
		return err
	}
	if receipt.Status != types.ReceiptStatusSuccessful {
		return fmt.Errorf("withdraw tx(%s) status is fail", tx.Hash().String())
	}

	iter, err := contracts.IL2ETHGateway.FilterWithdrawETH(
		&bind.FilterOpts{
			Start:   l2Number,
			Context: ctx,
		},
		[]common.Address{l2Auth.From},
		nil,
	)
	if err != nil {
		return err
	}

	path := TRACEDATA_DIR_PREFIX + "bridge/"
	for iter.Next() {
		return storeBlockResultsForBlocks(
			ctx,
			contracts.L2CLi,
			path,
			"withdrawETH",
			[]*big.Int{big.NewInt(0).SetUint64(iter.Event.Raw.BlockNumber)},
		)
	}
	return nil
}

func erc20DepositAndWithdraw(ctx context.Context, contracts *Contracts) error {
	return nil
}

type Contracts struct {
	sk *ecdsa.PrivateKey

	L1CLi *ethclient.Client
	L2CLi *ethclient.Client

	L1ChainID *big.Int
	L2ChainID *big.Int

	L1Auth *bind.TransactOpts
	L2Auth *bind.TransactOpts

	IL1ETHGateway *l1gateway.IL1ETHGateway
	IL2ETHGateway *l2gateway.IL2ETHGateway

	IL1StandardERC20Gateway *l1gateway.IL1ERC20Gateway
	IL2StandardERC20Gateway *l2gateway.IL2ERC20Gateway

	IL1ScrollMessenger *L1.IL1ScrollMessenger
	IL2ScrollMessenger *L2.IL2ScrollMessenger
}

func NewContracts(ctx context.Context, sk *ecdsa.PrivateKey, l1Cli, l2Cli *ethclient.Client, file string) (*Contracts, error) {
	cts := &Contracts{
		sk:    sk,
		L1CLi: l1Cli,
		L2CLi: l2Cli,
	}
	if err := cts.getL1AndL1Auth(ctx); err != nil {
		return nil, err
	}
	if err := cts.parseContracts(file); err != nil {
		return nil, err
	}

	return cts, nil
}

func (c *Contracts) getL1AndL1Auth(ctx context.Context) (err error) {
	c.L1ChainID, err = c.L1CLi.ChainID(ctx)
	if err != nil {
		return err
	}
	c.L1Auth, err = bind.NewKeyedTransactorWithChainID(c.sk, c.L1ChainID)
	if err != nil {
		return err
	}
	c.L2ChainID, err = c.L2CLi.ChainID(ctx)
	if err != nil {
		return err
	}
	c.L2Auth, err = bind.NewKeyedTransactorWithChainID(c.sk, c.L2ChainID)
	if err != nil {
		return err
	}
	return nil
}

func (c *Contracts) parseContracts(file string) error {
	data, err := os.ReadFile(file)
	if err != nil {
		return err
	}
	lines := strings.Split(string(data), "\n")
	for _, val := range lines {
		kv := strings.Split(strings.TrimSpace(val), "=")
		if len(kv) != 2 {
			log.Warn("contract key value is abnormal", "msg", val)
			continue
		}
		switch kv[0] {
		case "L1_ETH_GATEWAY_PROXY_ADDR":
			c.IL1ETHGateway, err = l1gateway.NewIL1ETHGateway(common.HexToAddress(kv[1]), c.L1CLi)
		case "L2_ETH_GATEWAY_PROXY_ADDR":
			c.IL2ETHGateway, err = l2gateway.NewIL2ETHGateway(common.HexToAddress(kv[1]), c.L2CLi)
		case "L1_STANDARD_ERC20_GATEWAY_PROXY_ADDR":
			c.IL1StandardERC20Gateway, err = l1gateway.NewIL1ERC20Gateway(common.HexToAddress(kv[1]), c.L1CLi)
		case "L2_STANDARD_ERC20_GATEWAY_PROXY_ADDR":
			c.IL2StandardERC20Gateway, err = l2gateway.NewIL2ERC20Gateway(common.HexToAddress(kv[1]), c.L2CLi)
		case "L1_SCROLL_MESSENGER_PROXY_ADDR":
			c.IL1ScrollMessenger, err = L1.NewIL1ScrollMessenger(common.HexToAddress(kv[1]), c.L1CLi)
		case "L2_SCROLL_MESSENGER_PROXY_ADDR":
			c.IL2ScrollMessenger, err = L2.NewIL2ScrollMessenger(common.HexToAddress(kv[1]), c.L2CLi)
		}
		if err != nil {
			return err
		}
	}
	return nil
}
