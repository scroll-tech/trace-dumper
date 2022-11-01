package api

import (
	"context"
	"encoding/json"
	"math/big"
	"os"
	"testing"

	"github.com/scroll-tech/go-ethereum/accounts/abi/bind"
	"github.com/scroll-tech/go-ethereum/common"
	"github.com/scroll-tech/go-ethereum/ethclient"
	"github.com/scroll-tech/go-ethereum/rpc"
	"github.com/stretchr/testify/assert"

	"tool/accounts"
	"tool/contracts/dao"
	"tool/contracts/nft"
	"tool/contracts/sushi"
	"tool/utils"
)

var (
	modules = map[SolType]common.Address{}
)

func loadModules() error {
	client, err := rpc.Dial("http://127.0.0.1:8190")
	if err != nil {
		return err
	}
	return client.Call(&modules, "rpc_addresses")
}

func TestAPI(t *testing.T) {
	srv := rpc.NewServer()
	session := &sushi.ERC20Session{
		Contract: nil,
	}
	err := srv.RegisterName("xx", session)
	assert.NoError(t, err)
}

func storeBlockResult(ctx context.Context, client *ethclient.Client, name string) error {
	header, err := client.HeaderByNumber(ctx, nil)
	if err != nil {
		return err
	}
	blockResult, err := client.GetBlockResultByHash(ctx, header.Hash())
	if err != nil {
		return err
	}
	data, err := json.Marshal(blockResult)
	if err != nil {
		return err
	}
	return os.WriteFile(name, data, os.ModePerm)
}

func TestStoreBlockResults(t *testing.T) {
	ctx := context.Background()
	assert := assert.New(t)

	assert.NoError(loadModules())

	client, err := ethclient.DialContext(ctx, "http://127.0.0.1:8545")
	assert.NoError(err)

	auth, err := bind.NewKeyedTransactorWithChainID(accounts.AddrPrivs[0].PrivateKey, big.NewInt(53077))
	assert.NoError(err)
	auth.GasPrice = big.NewInt(1108583800)
	auth.GasLimit = 11529940

	auth2, err := bind.NewKeyedTransactorWithChainID(accounts.AddrPrivs[1].PrivateKey, big.NewInt(53077))
	assert.NoError(err)

	// get dao trace
	daoToken, err := dao.NewGovernorMock(modules[DaoName], client)
	assert.NoError(err)
	_, err = daoToken.Propose(auth, []common.Address{auth2.From}, []*big.Int{big.NewInt(1)}, [][]byte{}, "dao propose test")
	assert.NoError(err)
	tx, err := daoToken.Cancel(auth, []common.Address{auth2.From}, []*big.Int{big.NewInt(1)}, [][]byte{}, common.Hash{})
	assert.NoError(err)
	utils.WaitPendingTx(ctx, client, tx.Hash())
	assert.NoError(storeBlockResult(ctx, client, "/tmp/dao.json"))

	// get nft trace
	nftToken, err := nft.NewERC721Mock(modules[NftName], client)
	assert.NoError(err)
	_, err = nftToken.Mint(auth, auth.From, big.NewInt(1))
	assert.NoError(err)
	_, err = nftToken.TransferFrom(auth, auth.From, auth2.From, big.NewInt(1))
	assert.NoError(err)
	tx, err = nftToken.Burn(auth, big.NewInt(1))
	utils.WaitPendingTx(ctx, client, tx.Hash())
	assert.NoError(storeBlockResult(ctx, client, "/tmp/nft.json"))

	// get sushi trace
	masterChef, err := sushi.NewMasterChef(modules[ChefName], client)
	assert.NoError(err)
	tx, err = masterChef.Add(auth, big.NewInt(1), modules[SushiName], true)
	assert.NoError(err)
	//utils.WaitPendingTx(ctx, client, tx.Hash())
	pid, err := masterChef.PoolLength(&bind.CallOpts{Pending: true, Context: ctx})
	assert.NoError(err)
	_, err = masterChef.Set(auth, pid, big.NewInt(1), true)
	assert.NoError(err)
	//masterChef.GetMultiplier(nil, auth.From, auth2.From)
	_, err = masterChef.Deposit(auth, pid, big.NewInt(1))
	assert.NoError(err)
	tx, err = masterChef.Withdraw(auth, pid, big.NewInt(1))
	utils.WaitPendingTx(ctx, client, tx.Hash())
	assert.NoError(storeBlockResult(ctx, client, "/tmp/sushi.json"))
}
