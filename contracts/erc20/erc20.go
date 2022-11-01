package erc20

import (
	"context"
	"errors"
	"math/big"
	"math/rand"
	"strings"
	"sync"
	"time"

	"github.com/scroll-tech/go-ethereum/accounts/abi/bind"
	"github.com/scroll-tech/go-ethereum/common"
	"github.com/scroll-tech/go-ethereum/ethclient"
	"github.com/scroll-tech/go-ethereum/event"
	"github.com/scroll-tech/go-ethereum/log"

	"tool/accounts"
	"tool/utils"
)

type ERC20Token struct {
	ctx     context.Context
	Address common.Address
	*ERC20TemplateSession

	accs   *accounts.Accounts
	client bind.ContractBackend

	batchCh chan struct{}
	txsFeed event.Feed
	stopCh  chan struct{}
}

func NewERC20Token(ctx context.Context, symbol string, accs *accounts.Accounts, client *ethclient.Client, mint bool) (*ERC20Token, error) {
	auth := accs.GetAccount()
	defer accs.SetAccount(auth)
	// Deploy or load contract
	tokenAddr, tx, erc20Token, err := DeployERC20Template(auth, client, auth.From, auth.From, symbol+" coin", symbol, 18)
	if err != nil {
		log.Crit("Failed to deploy erc20 contract", "err", err)
	}
	utils.WaitPendingTx(ctx, client, tx.Hash())

	// Mint balance for static accounts.
	if mint {
		for _, acc := range accounts.AddrPrivs {
			tx, err = erc20Token.Mint(auth, acc.Address, big.NewInt(1).Mul(big.NewInt(1e3), utils.Ether))
			if err != nil {
				return nil, err
			}
		}
	}

	// change owner
	tx, err = erc20Token.ChangeUser(auth, accs.Root.From, accs.Root.From)
	if err != nil {
		return nil, err
	}
	log.Info("Deploy "+strings.ToLower(symbol)+" contract successful", "address", tokenAddr.String())
	utils.WaitPendingTx(ctx, client, tx.Hash())

	return &ERC20Token{
		ctx:     ctx,
		accs:    accs,
		client:  client,
		batchCh: make(chan struct{}, 1),
		stopCh:  make(chan struct{}, 1),
		Address: tokenAddr,
		ERC20TemplateSession: &ERC20TemplateSession{
			Contract:     erc20Token,
			TransactOpts: *auth,
		},
	}, nil
}

func (e *ERC20Token) Pressure(count, batch, tps int) error {
	select {
	case e.batchCh <- struct{}{}:
		batch = utils.MinInt(utils.MinInt(utils.MinInt(batch, count), tps), len(accounts.AddrPrivs))
		sleep := time.Millisecond * 1000 * time.Duration(batch) / time.Duration(tps)
		go e.sendTxs(e.ctx, count, batch, sleep)
	default:
		return errors.New("current batch task is not finished")
	}
	return nil
}

// Batch send transactions.
func (e *ERC20Token) sendTxs(ctx context.Context, count, batch int, sleep time.Duration) {
	defer func() { <-e.batchCh }()
	start := time.Now()

	var wait sync.WaitGroup
	for j := 0; j < batch; j++ {
		wait.Add(1)
		go func() {
			defer wait.Done()
			loopCount := count / batch
			for i := 0; i < loopCount; i++ {
				to := common.BigToAddress(big.NewInt(int64(rand.Intn(1000000))))
				auth := e.accs.GetAccount()
				if _, err := e.Contract.Transfer(auth, to, big.NewInt(100)); err != nil {
					log.Error("Failed to send tx", "err", err)
				}
				e.accs.SetAccount(auth)
				if i < loopCount-1 {
					time.Sleep(sleep)
				}
			}
		}()
	}
	wait.Wait()

	log.Info("Send erc20 txs finished", "txs sum", count, "time used(ms)", time.Now().Sub(start).Milliseconds())
}
