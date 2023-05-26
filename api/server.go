package api

import (
	"context"
	"github.com/scroll-tech/go-ethereum/ethclient"
	"tool/accounts"
)

type solFunc func(ctx context.Context, accounts *accounts.Accounts, client *ethclient.Client) error

type SolType string

const (
	// native instance
	NativeName = SolType("native")

	// erc20 token contracts
	ERC20Name = SolType("erc20")
	WETHName  = SolType("weth")
	BTCName   = SolType("btc")

	// nft contract
	NftName = SolType("nft")

	// greeter contract
	GreeterName = SolType("greeter")

	// Sushi contract
	SushiName = SolType("sushi")
	ChefName  = SolType("chef")

	// Dao contract
	VoteName = SolType("vote")
	DaoName  = SolType("dao")

	// uniswapv2 contracts
	Uniswapv2Name = SolType("uniswapv2")
	FactoryName   = SolType("factory")
	RouterName    = SolType("router")
	MulticallName = SolType("multicall")

	// Scroll contracts
	ScrollName = SolType("scroll")
)
