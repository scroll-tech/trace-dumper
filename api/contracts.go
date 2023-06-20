package api

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"math/big"
	"math/rand"
	"os"

	"github.com/scroll-tech/go-ethereum/accounts/abi/bind"
	"github.com/scroll-tech/go-ethereum/common"
	"github.com/scroll-tech/go-ethereum/core/types"
	"github.com/scroll-tech/go-ethereum/crypto"
	"github.com/scroll-tech/go-ethereum/ethclient"
	"github.com/scroll-tech/go-ethereum/log"

	"tool/accounts"
	"tool/contracts/dao"
	"tool/contracts/erc20"
	"tool/contracts/greeter"
	"tool/contracts/nft"
	"tool/contracts/sushi"
	"tool/contracts/uniswap/factory"
	"tool/contracts/uniswap/router"
	"tool/contracts/uniswap/weth9"
	"tool/contracts/vote"
	"tool/utils"
)

var (
	WrapJson bool
)

type jsonrpcMessage struct {
	Version string          `json:"jsonrpc,omitempty"`
	ID      int             `json:"id,omitempty"`
	Result  json.RawMessage `json:"result,omitempty"`
}

const TRACEDATA_DIR_PREFIX = "./tracedata/"

func storeBlockResultsForTxs(ctx context.Context, client *ethclient.Client, path, file string, txs ...*types.Transaction) error {
	numberList, err := getTxsBlockNumbers(ctx, client, file, txs...)
	if err != nil {
		return err
	}

	return storeBlockResultsForBlocks(ctx, client, path, file, numberList)
}

func getTxsBlockNumbers(ctx context.Context, client *ethclient.Client, file string, txs ...*types.Transaction) ([]*big.Int, error) {
	// Wait tx mined.
	var (
		preNumber  *big.Int
		numberList []*big.Int
	)
	for _, tx := range txs {
		receipt, err := bind.WaitMined(ctx, client, tx)
		if err != nil {
			return nil, err
		}
		if receipt.Status != types.ReceiptStatusSuccessful {
			return nil, errors.New("receipt status is fail")
		}
		if preNumber != nil && preNumber.Uint64() == receipt.BlockNumber.Uint64() {
			continue
		}
		preNumber = receipt.BlockNumber
		numberList = append(numberList, receipt.BlockNumber)

		log.Info(file, "number", receipt.BlockNumber.Uint64(), "txHash", tx.Hash().String())
	}
	return numberList, nil
}

func storeBlockResultsForBlocks(ctx context.Context, client *ethclient.Client, path, file string, numberList []*big.Int) error {
	for _, number := range numberList {
		trace, err := client.GetBlockTraceByNumber(ctx, number)
		if err != nil {
			return err
		}
		data, err := json.MarshalIndent(trace, "", "    ")
		if err != nil {
			return err
		}

		if WrapJson {
			wrapData := json.RawMessage(data)
			var wrapJson = &jsonrpcMessage{
				Version: "2.0",
				ID:      1,
				Result:  wrapData,
			}
			data, err = json.MarshalIndent(wrapJson, "", "    ")
			if err != nil {
				return err
			}
		}

		// Check dir exist or not.
		if exist, _ := utils.PathExists(path); !exist {
			if err = os.Mkdir(path, 0755); err != nil {
				return err
			}
		}

		name := fmt.Sprintf("%s/%s_%d.json", path, file, number)
		if len(numberList) == 1 {
			name = fmt.Sprintf("%s/%s.json", path, file)
		}
		// Write file
		if err = os.WriteFile(name, data, 0600); err != nil {
			return err
		}
	}

	return nil
}

func Native(ctx context.Context, client *ethclient.Client, root *bind.TransactOpts, to common.Address, value *big.Int) error {
	tx, err := accounts.CreateSignedTx(client, root, &to, value, nil)
	if err != nil {
		return err
	}
	if err = client.SendTransaction(ctx, tx); err != nil {
		return err
	}
	return storeBlockResultsForTxs(ctx, client, TRACEDATA_DIR_PREFIX+"native/", "transfer", tx)
}

func NewERC20(ctx context.Context, client *ethclient.Client, root, auth *bind.TransactOpts) error {
	_, tx, erc20Token, err := erc20.DeployERC20Template(root, client, root.From, root.From, "WETH coin", "WETH", 18)
	if err != nil {
		return err
	}

	path := TRACEDATA_DIR_PREFIX + "erc20/"
	if err = storeBlockResultsForTxs(ctx, client, path, "deploy", tx); err != nil {
		return err
	}

	originVal := big.NewInt(1).Mul(big.NewInt(3e3), utils.Ether)
	tx, err = erc20Token.Mint(root, root.From, originVal)
	if err != nil {
		return err
	}
	if err = storeBlockResultsForTxs(ctx, client, path, "mint", tx); err != nil {
		return err
	}

	// erc20 transfer
	tx, err = erc20Token.Transfer(root, auth.From, big.NewInt(1000))
	if err != nil {
		return err
	}
	if err = storeBlockResultsForTxs(ctx, client, path, "1_transfer", tx); err != nil {
		return err
	}

	var txs = make([]*types.Transaction, 0, 10)
	for i := 0; i < 10; i++ {
		// erc20 transfer
		tx, err = erc20Token.Transfer(root, auth.From, big.NewInt(1000))
		if err != nil {
			return err
		}
		txs = append(txs, tx)
	}

	return storeBlockResultsForTxs(ctx, client, path, "10_transfer", txs...)
}

func NewGreeter(ctx context.Context, client *ethclient.Client, root *bind.TransactOpts) error {
	_, tx, token, err := greeter.DeployGreeter(root, client, big.NewInt(10))
	if err != nil {
		return err
	}

	path := TRACEDATA_DIR_PREFIX + "greeter/"
	if err = storeBlockResultsForTxs(ctx, client, path, "deploy", tx); err != nil {
		return err
	}

	tx, err = token.SetValue(root, big.NewInt(10))
	if err != nil {
		return err
	}
	return storeBlockResultsForTxs(ctx, client, path, "setValue", tx)
}

func NewNft(ctx context.Context, client *ethclient.Client, root, auth *bind.TransactOpts) error {
	_, tx, token, err := nft.DeployERC721Mock(root, client, "ERC721 coin", "ERC721")
	if err != nil {
		return err
	}
	path := TRACEDATA_DIR_PREFIX + "nft/"
	if err = storeBlockResultsForTxs(ctx, client, path, "deploy", tx); err != nil {
		return err
	}

	tokenId := big.NewInt(rand.Int63())
	tx, err = token.Mint(root, root.From, tokenId)
	if err != nil {
		return err
	}
	if err = storeBlockResultsForTxs(ctx, client, path, "mint", tx); err != nil {
		return err
	}

	tx, err = token.TransferFrom(root, root.From, auth.From, tokenId)
	if err != nil {
		return err
	}
	if err = storeBlockResultsForTxs(ctx, client, path, "transferFrom", tx); err != nil {
		return err
	}

	tx, err = token.Burn(auth, tokenId)
	if err != nil {
		return err
	}
	return storeBlockResultsForTxs(ctx, client, path, "burn", tx)
}

func NewSushi(ctx context.Context, client *ethclient.Client, root *bind.TransactOpts) error {
	sushiAddr, tx, sushiToken, err := sushi.DeploySushiToken(root, client)
	if err != nil {
		return err
	}

	path := TRACEDATA_DIR_PREFIX + "sushi/"
	if err = storeBlockResultsForTxs(ctx, client, path, "deploy-sushi", tx); err != nil {
		return err
	}

	chefAddr, tx, chefToken, err := sushi.DeployMasterChef(root, client, sushiAddr, root.From, big.NewInt(1), big.NewInt(1), big.NewInt(math.MaxInt))
	if err != nil {
		return err
	}
	if err = storeBlockResultsForTxs(ctx, client, path, "deploy-chef", tx); err != nil {
		return err
	}

	amount := big.NewInt(1e18)
	tx, err = sushiToken.Mint(root, root.From, amount)
	if err != nil {
		return err
	}
	if err = storeBlockResultsForTxs(ctx, client, path, "sushi-mint", tx); err != nil {
		return err
	}

	allocPoint := utils.Ether
	tx, err = chefToken.Add(root, allocPoint, sushiAddr, true)
	if err != nil {
		return err
	}
	if err = storeBlockResultsForTxs(ctx, client, path, "chef-add", tx); err != nil {
		return err
	}

	pid, err := chefToken.PoolLength(&bind.CallOpts{Pending: true})
	if err != nil {
		return err
	}
	pid.Sub(pid, big.NewInt(1))
	tx, err = chefToken.Set(root, pid, allocPoint, true)
	if err != nil {
		return err
	}
	if err = storeBlockResultsForTxs(ctx, client, path, "chef-set", tx); err != nil {
		return err
	}

	tx, err = sushiToken.Approve(root, chefAddr, amount)
	if err = storeBlockResultsForTxs(ctx, client, path, "sushi-approve", tx); err != nil {
		return err
	}

	// deposit amount to chef
	tx, err = chefToken.Deposit(root, pid, amount)
	if err = storeBlockResultsForTxs(ctx, client, path, "chef-deposit", tx); err != nil {
		return err
	}

	// change sushiToken's owner to masterChef.
	tx, err = sushiToken.TransferOwnership(root, chefAddr)
	if err = storeBlockResultsForTxs(ctx, client, path, "sushi-transferOwnership", tx); err != nil {
		return err
	}

	res, err := chefToken.UserInfo(nil, pid, root.From)
	if err != nil {
		return err
	}

	// withdraw amount from chef
	tx, err = chefToken.Withdraw(root, pid, res.Amount)
	if err = storeBlockResultsForTxs(ctx, client, path, "chef-withdraw", tx); err != nil {
		return err
	}
	return nil
}

func NewDao(ctx context.Context, client *ethclient.Client, root, auth *bind.TransactOpts) error {
	voteAddr, tx, _, err := vote.DeployVotesMock(root, client, "vote v2")
	if err != nil {
		return err
	}
	path := TRACEDATA_DIR_PREFIX + "/dao"
	if err = storeBlockResultsForTxs(ctx, client, path, "deploy-votes", tx); err != nil {
		return err
	}

	_, tx, daoToken, err := dao.DeployGovernorMock(root, client, "governor mock", voteAddr, big.NewInt(1), big.NewInt(1), big.NewInt(100))
	if err != nil {
		return err
	}
	if err = storeBlockResultsForTxs(ctx, client, path, "deploy-dao", tx); err != nil {
		return err
	}

	callData := [][]byte{big.NewInt(1).Bytes()}
	target := common.BigToAddress(big.NewInt(1))
	value := big.NewInt(1)
	description := "dao propose test"
	tx, err = daoToken.Propose(root, []common.Address{target}, []*big.Int{value}, callData, description)
	if err != nil {
		return err
	}
	if err = storeBlockResultsForTxs(ctx, client, path, "dao-propose", tx); err != nil {
		return err
	}

	salt := crypto.Keccak256Hash([]byte(description))
	tx, err = daoToken.Cancel(root, []common.Address{target}, []*big.Int{value}, callData, salt)
	if err != nil {
		return err
	}
	if err = storeBlockResultsForTxs(ctx, client, path, "dao-cancel", tx); err != nil {
		return err
	}

	return nil
}

func NewUniswapv2(ctx context.Context, client *ethclient.Client, root, auth *bind.TransactOpts) error {
	wethAddr, tx, wethToken, err := weth9.DeployWETH9(root, client)
	if err != nil {
		return err
	}

	path := TRACEDATA_DIR_PREFIX + "uniswapv2"
	if err = storeBlockResultsForTxs(ctx, client, path, "deploy-weth9", tx); err != nil {
		return err
	}

	// deploy factory
	fAddr, tx, fToken, err := factory.DeployUniswapV2Factory(root, client, root.From)
	if err != nil {
		return err
	}
	if err = storeBlockResultsForTxs(ctx, client, path, "deploy-factory", tx); err != nil {
		return err
	}

	// deploy router
	rAddr, tx, rToken, err := router.DeployUniswapV2Router02(root, client, fAddr, wethAddr)
	if err != nil {
		return err
	}
	if err = storeBlockResultsForTxs(ctx, client, path, "deploy-router", tx); err != nil {
		return err
	}

	btcAddr, tx, btcToken, err := erc20.DeployERC20Template(root, client, root.From, root.From, "BTC coin", "BTC", 18)
	if err != nil {
		return err
	}
	if err = storeBlockResultsForTxs(ctx, client, path, "deploy-btc", tx); err != nil {
		return err
	}

	// init balance
	originVal := big.NewInt(1).Mul(big.NewInt(3e3), utils.Ether)
	auth.Value = originVal
	tx0, err := wethToken.Deposit(auth)
	if err != nil {
		return err
	}
	auth.Value = nil
	tx1, err := btcToken.Mint(root, auth.From, originVal)
	if err != nil {
		return err
	}
	tx2, err := wethToken.Approve(auth, rAddr, originVal)
	if err != nil {
		return err
	}
	tx3, err := btcToken.Approve(auth, rAddr, originVal)
	if err != nil {
		return err
	}
	if err = storeBlockResultsForTxs(ctx, client, path, "token-initBalance", []*types.Transaction{tx0, tx1, tx2, tx3}...); err != nil {
		return err
	}
	bls, _ := wethToken.BalanceOf(nil, auth.From)
	log.Info("weth balance", "balance", bls.String())
	bls, _ = btcToken.BalanceOf(nil, auth.From)
	log.Info("btc balance", "balance", bls.String())

	// create pair
	tx, err = fToken.CreatePair(root, wethAddr, btcAddr)
	if err = storeBlockResultsForTxs(ctx, client, path, "factory-createPair", tx); err != nil {
		return err
	}

	// add liquidity, pool is 1:1
	liqVal := big.NewInt(1).Mul(big.NewInt(1e3), utils.Ether)
	tx, err = rToken.AddLiquidity(
		auth,
		wethAddr,
		btcAddr,
		liqVal,
		liqVal,
		big.NewInt(0),
		big.NewInt(0),
		auth.From,
		big.NewInt(2e9),
	)
	if err != nil {
		return err
	}
	if err = storeBlockResultsForTxs(ctx, client, path, "router-AddLiquidity", tx); err != nil {
		return err
	}

	header, err := client.HeaderByNumber(ctx, nil)
	if err != nil {
		return err
	}

	// swap weth => btc
	// swapVal := utils.Ether
	swapVal := big.NewInt(1e15)
	txs := make([]*types.Transaction, 0, 100)
	auth.GasLimit = 5000000
	for i := 0; i < 100; i++ {
		tx, err = rToken.SwapExactTokensForTokens(
			auth,
			swapVal,
			big.NewInt(0),
			[]common.Address{wethAddr, btcAddr},
			auth.From,
			big.NewInt(int64(header.Time)*2),
		)
		if err != nil {
			return err
		}
		txs = append(txs, tx)
	}
	for i, tx := range txs {
		receipt, err := bind.WaitMined(ctx, client, tx)
		if err != nil {
			return err
		}
		if receipt.Status != types.ReceiptStatusSuccessful {
			log.Error("tx status is not right", "index", i, "txHash", tx.Hash().String())
		}
	}

	return storeBlockResultsForTxs(ctx, client, path, "router-swapExactTokensForTokens", tx)
}
