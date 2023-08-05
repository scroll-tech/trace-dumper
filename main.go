package main

import (
	"context"
	"flag"
	"github.com/mattn/go-colorable"
	"github.com/mattn/go-isatty"
	"github.com/scroll-tech/go-ethereum/crypto"
	"github.com/scroll-tech/go-ethereum/ethclient"
	"github.com/scroll-tech/go-ethereum/log"
	"io"
	"math/big"
	"os"
	"tool/accounts"
	"tool/api"
)

var (
	wrapJson = flag.Bool("wrap", false, "add json rpc wrapJson")
	skStr    = flag.String("sk", "", "Support a private key to do trace operation")
	// special client for bridge cmd.
	l1Endpoint = flag.String("l1.endpoint", "http://127.0.0.1:8543/l1", "The l1 chain endpoint to connect to blockchain node")
	l2Endpoint = flag.String("l2.endpoint", "ws://127.0.0.1:8546", "The l2 chain endpoint to connect to blockchain node")
	contracts  = flag.String("contracts", "./.env.local.contracts", "All the deployed scroll contracts address list")
	dump       = flag.String("dump", "erc20", "e.g: erc20, native, nft, greeter, sushi, dao, uniswapv2, multi_uniswapv2, bridge")
)

func init() {
	output := io.Writer(os.Stderr)
	usecolor := (isatty.IsTerminal(os.Stderr.Fd()) || isatty.IsCygwinTerminal(os.Stderr.Fd())) && os.Getenv("TERM") != "dumb"
	if usecolor {
		output = colorable.NewColorableStderr()
	}
	ostream := log.StreamHandler(output, log.TerminalFormat(usecolor))
	glogger := log.NewGlogHandler(ostream)
	// Set log level
	glogger.Verbosity(log.LvlInfo)
	log.Root().SetHandler(glogger)
}

func main() {
	// Parse the flags and set up the logger to print everything requested
	flag.Parse()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Enable wrap json rpc result or not.
	api.WrapJson = *wrapJson

	// create l2 client
	l2Cli, err := ethclient.DialContext(ctx, *l2Endpoint)
	if err != nil {
		log.Crit("Failed to connect l2 chain node", "err", err)
	}

	// parse private key
	sk, err := crypto.HexToECDSA(*skStr)
	if err != nil {
		log.Crit("failed to parse private key", "err", err)
	}

	// load keystore file
	var auths *accounts.Accounts
	auths, err = accounts.NewAccounts(ctx, 2, l2Cli, sk)
	if err != nil {
		log.Crit("failed to create accounts", "err", err)
	}
	root := auths.Root
	auth := auths.GetAccount()

	solName := api.SolType(*dump)
	switch solName {
	case api.BridgeName:
		// create l1 client
		var l1Cli *ethclient.Client
		l1Cli, err = ethclient.DialContext(ctx, *l1Endpoint)
		if err != nil {
			log.Crit("failed to l1 chain node", "err", err)
		}
		err = api.NewBridge(ctx, sk, *contracts, l1Cli, l2Cli)
	default:
		switch solName {
		case api.NativeName:
			err = api.Native(ctx, l2Cli, root, auth.From, big.NewInt(100))
		case api.ERC20Name:
			err = api.NewERC20(ctx, l2Cli, root, auth)
		case api.NftName:
			err = api.NewNft(ctx, l2Cli, root, auth)
		case api.GreeterName:
			err = api.NewGreeter(ctx, l2Cli, root)
		case api.SushiName:
			err = api.NewSushi(ctx, l2Cli, root)
		case api.DaoName:
			err = api.NewDao(ctx, l2Cli, root, auth)
		case api.Uniswapv2Name:
			err = api.NewUniswapv2(ctx, l2Cli, root, auth)
		case api.MultiUniswapv2Name:
			err = api.NewMultiUniswapv2(ctx, l2Cli, root, auth)
		default:
			log.Error("unexpected dump option")
			return
		}
	}
	if err != nil {
		log.Error("dump traces for contract fail", "contract name", solName, "err", err)
	} else {
		log.Info("dump traces for contract successfully", "contract name", solName)
	}
}
