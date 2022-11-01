package api

import (
	"context"
	"math"
	"math/big"
	"strings"

	"github.com/scroll-tech/go-ethereum/accounts/abi/bind"
	"github.com/scroll-tech/go-ethereum/common"
	"github.com/scroll-tech/go-ethereum/log"
	"github.com/scroll-tech/go-ethereum/rpc"

	"tool/contracts/dao"
	"tool/contracts/erc20"
	"tool/contracts/greeter"
	"tool/contracts/native"
	"tool/contracts/nft"
	"tool/contracts/sushi"
	"tool/contracts/uniswap/factory"
	"tool/contracts/uniswap/multicall"
	"tool/contracts/uniswap/router"
	"tool/contracts/uniswap/weth9"
	"tool/contracts/vote"
	"tool/utils"
)

func (s *SolFactory) newNativectx(ctx context.Context) (*rpc.API, error) {
	s.solNodes[NativeName] = common.Address{}
	return &rpc.API{
		Namespace: string(NativeName),
		Service:   native.NewNative(ctx, s.accounts, s.client),
		Public:    true,
	}, nil
}

func (s *SolFactory) erc20Instance(ctx context.Context, solName SolType) (*rpc.API, error) {
	token, err := erc20.NewERC20Token(ctx, strings.ToUpper(string(solName)), s.accounts, s.client, solName == ERC20Name)
	if err != nil {
		return nil, err
	}
	s.solNodes[solName] = token.Address
	return &rpc.API{
		Namespace: string(solName),
		Service:   token,
		Public:    true,
	}, nil
}

func (s *SolFactory) newNFT(ctx context.Context) (*rpc.API, error) {
	accs := s.accounts
	client := s.client
	auth := accs.GetAccount()
	defer accs.SetAccount(auth)
	addr, tx, token, err := nft.DeployERC721Mock(auth, client, "ERC721 coin", "ERC721")
	if err != nil {
		return nil, err
	}
	// Wait pending tx until it's deployed.
	utils.WaitPendingTx(ctx, client, tx.Hash())
	log.Info("Deploy nft successful", "address", addr.String())

	s.solNodes[NftName] = addr
	return &rpc.API{
		Namespace: string(NftName),
		Service: &nft.ERC721MockSession{
			Contract:     token,
			TransactOpts: *auth,
			CallOpts: bind.CallOpts{
				Pending: true,
				Context: ctx,
			},
		},
		Public: true,
	}, nil
}

func (s *SolFactory) newVote(ctx context.Context) (*rpc.API, error) {
	auth := s.accounts.GetAccount()
	defer s.accounts.SetAccount(auth)
	// deploy vote
	addr, voteTx, token, err := vote.DeployVotesMock(auth, s.client, "vote v2")
	if err != nil {
		return nil, err
	}
	utils.WaitPendingTx(ctx, s.client, voteTx.Hash())
	log.Info("Deploy vote successful", "address", addr.String())

	s.solNodes[VoteName] = addr
	return &rpc.API{
		Namespace: string(VoteName),
		Service: &vote.VotesMockSession{
			Contract:     token,
			TransactOpts: *auth,
			CallOpts: bind.CallOpts{
				Pending: true,
				Context: ctx,
			},
		},
		Public: true,
	}, nil
}

func (s *SolFactory) newDao(ctx context.Context) (*rpc.API, error) {
	auth := s.accounts.GetAccount()
	defer s.accounts.SetAccount(auth)
	// deploy vote
	addr := s.solNodes[VoteName]

	// deploy governor vote
	addr, tx, token, err := dao.DeployGovernorMock(auth, s.client, "governor mock", addr, big.NewInt(1), big.NewInt(1), big.NewInt(100))
	if err != nil {
		return nil, err
	}
	utils.WaitPendingTx(ctx, s.client, tx.Hash())
	log.Info("Deploy dao successful", "address", addr.String())

	s.solNodes[DaoName] = addr
	return &rpc.API{
		Namespace: string(DaoName),
		Service: &dao.GovernorMockSession{
			Contract:     token,
			TransactOpts: *auth,
			CallOpts: bind.CallOpts{
				Pending: true,
				Context: ctx,
			},
		},
		Public: true,
	}, nil
}

func (s *SolFactory) newGreeter(ctx context.Context) (*rpc.API, error) {
	auth := s.accounts.GetAccount()
	defer s.accounts.SetAccount(auth)
	// deploy greeter
	addr, tx, token, err := greeter.DeployGreeter(auth, s.client, big.NewInt(10))
	if err != nil {
		return nil, err
	}
	// Wait pending tx until it's deployed.
	utils.WaitPendingTx(ctx, s.client, tx.Hash())
	log.Info("Deploy greeter successful", "address", addr.String())

	s.solNodes[GreeterName] = addr
	// return session
	return &rpc.API{
		Namespace: string(GreeterName),
		Service: &greeter.GreeterSession{
			Contract:     token,
			TransactOpts: *auth,
			CallOpts: bind.CallOpts{
				Pending: true,
				Context: ctx,
			},
		},
		Public: true,
	}, nil
}

func (s *SolFactory) newSushiToken(ctx context.Context) (*rpc.API, error) {
	auth := s.accounts.GetAccount()
	defer s.accounts.SetAccount(auth)
	// deploy sushi
	sushiAddr, tx, token, err := sushi.DeploySushiToken(auth, s.client)
	if err != nil {
		return nil, err
	}
	utils.WaitPendingTx(ctx, s.client, tx.Hash())

	// change owner
	tx, err = token.TransferOwnership(auth, s.accounts.Root.From)
	if err != nil {
		return nil, err
	}
	utils.WaitPendingTx(ctx, s.client, tx.Hash())
	log.Debug("Change sushiToken owner to root account", "address", s.accounts.Root.From.String())

	log.Info("Deploy sushiToken successful", "address", sushiAddr.String())

	s.solNodes[SushiName] = sushiAddr
	return &rpc.API{
		Namespace: string(SushiName),
		Service: &sushi.SushiTokenSession{
			Contract:     token,
			TransactOpts: *auth,
			CallOpts: bind.CallOpts{
				Pending: true,
				Context: ctx,
			},
		},
		Public: true,
	}, nil
}

func (s *SolFactory) newMasterChef(ctx context.Context) (*rpc.API, error) {
	accs := s.accounts
	client := s.client

	auth := accs.GetAccount()
	defer accs.SetAccount(auth)

	sushiAddr := s.solNodes[SushiName]
	addr, tx, chefToken, err := sushi.DeployMasterChef(auth, s.client, sushiAddr, auth.From, big.NewInt(1), big.NewInt(1), big.NewInt(math.MaxInt))
	if err != nil {
		return nil, err
	}
	// Wait pending tx until it's deployed.
	utils.WaitPendingTx(ctx, s.client, tx.Hash())
	log.Info("Deploy masterChef successful", "address", addr.String())

	tx, err = chefToken.TransferOwnership(auth, accs.Root.From)
	if err != nil {
		return nil, err
	}
	utils.WaitPendingTx(ctx, client, tx.Hash())
	log.Debug("Change chef owner to root account", "address", accs.Root.From.String())

	s.solNodes[ChefName] = addr
	// return session
	return &rpc.API{
		Namespace: string(ChefName),
		Service: &sushi.MasterChefSession{
			Contract:     chefToken,
			TransactOpts: *auth,
			CallOpts: bind.CallOpts{
				Pending: true,
				Context: ctx,
			},
		},
		Public: true,
	}, nil
}

func (s *SolFactory) newUniswap(ctx context.Context) (apis []rpc.API, err error) {
	auth := s.accounts.GetAccount()
	defer s.accounts.SetAccount(auth)

	wethAddr, _, wToken, err := weth9.DeployWETH9(auth, s.client)
	if err != nil {
		return nil, err
	}

	// deploy factory
	fAddr, _, fToken, err := factory.DeployUniswapV2Factory(auth, s.client, auth.From)
	if err != nil {
		return nil, err
	}

	// deploy router
	rAddr, _, rToken, err := router.DeployUniswapV2Router02(auth, s.client, fAddr, wethAddr)
	if err != nil {
		return nil, err
	}

	// deploy multicall
	mAddr, tx, mToken, err := multicall.DeployMulticall(auth, s.client)
	if err != nil {
		return nil, err
	}
	utils.WaitPendingTx(ctx, s.client, tx.Hash())

	s.solNodes[WETHName] = wethAddr
	s.solNodes[FactoryName] = fAddr
	s.solNodes[RouterName] = rAddr
	s.solNodes[MulticallName] = mAddr
	apis = append(apis, []rpc.API{
		{
			Namespace: string(WETHName),
			Service: &weth9.WETH9Session{
				Contract:     wToken,
				TransactOpts: *auth,
				CallOpts: bind.CallOpts{
					Pending: true,
					Context: ctx,
				},
			},
			Public: true,
		},
		{
			Namespace: string(FactoryName),
			Service: &factory.UniswapV2FactorySession{
				Contract:     fToken,
				TransactOpts: *auth,
				CallOpts: bind.CallOpts{
					Pending: true,
					Context: ctx,
				},
			},
			Public: true,
		},
		{
			Namespace: string(RouterName),
			Service: &router.UniswapV2Router02Session{
				Contract:     rToken,
				TransactOpts: *auth,
				CallOpts: bind.CallOpts{
					Pending: true,
					Context: ctx,
				},
			},
			Public: true,
		},
		{
			Namespace: string(MulticallName),
			Service: &multicall.MulticallSession{
				Contract:     mToken,
				TransactOpts: *auth,
				CallOpts: bind.CallOpts{
					Pending: true,
					Context: ctx,
				},
			},
			Public: true,
		},
	}...)

	log.Info("Deploy uniswapV2 successful", "address", rAddr.String())

	return apis, nil
}
