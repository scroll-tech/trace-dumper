package api

import (
	"context"
	"errors"

	"github.com/scroll-tech/go-ethereum/common"
	"github.com/scroll-tech/go-ethereum/ethclient"
	ethrpc "github.com/scroll-tech/go-ethereum/rpc"

	"tool/accounts"
	"tool/rpc"
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
)

type SolFactory struct {
	accounts *accounts.Accounts
	client   *ethclient.Client

	solNodes map[SolType]common.Address
}

func NewFactory(accounts *accounts.Accounts, client *ethclient.Client) error {
	return rpc.Register([]ethrpc.API{
		{
			Namespace: "rpc",
			Service: &SolFactory{
				accounts: accounts,
				client:   client,
				solNodes: make(map[SolType]common.Address),
			},
			Public: true,
		},
	})
}

func (s *SolFactory) Addresses() (map[string]string, error) {
	res := make(map[string]string)
	for name, addr := range s.solNodes {
		res[string(name)] = addr.String()
	}
	return res, nil
}

func (s *SolFactory) DeployContract(name string) (err error) {
	tpName := SolType(name)

	var (
		apis []ethrpc.API
		ctx  = context.Background()
	)
	switch tpName {
	case NativeName:
		api, err := s.newNativectx(ctx)
		if err != nil {
			return err
		}
		apis = append(apis, *api)
	case ERC20Name:
		api, err := s.erc20Instance(ctx, ERC20Name)
		if err != nil {
			return err
		}
		apis = append(apis, *api)
	case BTCName:
		api, err := s.erc20Instance(ctx, BTCName)
		if err != nil {
			return err
		}
		apis = append(apis, *api)
	case NftName:
		api, err := s.newNFT(ctx)
		if err != nil {
			return err
		}
		apis = append(apis, *api)
	case GreeterName:
		api, err := s.newGreeter(ctx)
		if err != nil {
			return err
		}
		apis = append(apis, *api)
	case DaoName:
		vApi, err := s.newVote(ctx)
		if err != nil {
			return err
		}
		dApi, err := s.newDao(ctx)
		if err != nil {
			return err
		}
		apis = append(apis, []ethrpc.API{*vApi, *dApi}...)
	case SushiName:
		sApi, err := s.newSushiToken(ctx)
		if err != nil {
			return err
		}
		mApi, err := s.newMasterChef(ctx)
		if err != nil {
			return err
		}
		apis = append(apis, []ethrpc.API{*sApi, *mApi}...)
	case Uniswapv2Name:
		uApis, err := s.newUniswap(ctx)
		if err != nil {
			return err
		}
		apis = append(apis, uApis...)
	default:
		return errors.New("unsupported contract")
	}

	return rpc.Register(apis)
}
