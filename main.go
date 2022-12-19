package main

import (
	"context"
	"crypto/ecdsa"
	"flag"
	"fmt"
	"github.com/mattn/go-colorable"
	"github.com/mattn/go-isatty"
	"github.com/scroll-tech/go-ethereum/accounts/abi/bind"
	"github.com/scroll-tech/go-ethereum/common"
	"github.com/scroll-tech/go-ethereum/crypto"
	"github.com/scroll-tech/go-ethereum/ethclient"
	"github.com/scroll-tech/go-ethereum/log"
	"io"
	"math/big"
	"os"
	"time"
	"tool/contracts/uniswap/router"
)

var (
	wrapJson = flag.Bool("wrap", false, "add json rpc wrapJson")
	endpoint = flag.String("endpoint", "ws://127.0.0.1:8546", "The endpoint to connect to blockchain node")
	keystore = flag.String("keystore", "./docker/l2geth/genesis-keystore", "Keystore file path")
	password = flag.String("password", "scrolltest", "The keystore password")
	dump     = flag.String("dump", "erc20", "e.g: erc20, nft, greeter, sushi, dao, uniswapv2")
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
	/*flag.Parse()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Enable wrap json rpc result or not.
	api.WrapJson = *wrapJson

	// create client
	client, err := ethclient.DialContext(ctx, *endpoint)
	if err != nil {
		log.Crit("Failed to connect chain node", "err", err)
	}

	// load keystore file
	auths, err := accounts.NewAccounts(ctx, 2, client, *keystore, *password)
	if err != nil {
		log.Crit("failed to create accounts", "err", err)
	}
	root := auths.Root
	auth := auths.GetAccount()

	solName := api.SolType(*dump)
	switch solName {
	case api.NativeName:
		err = api.Native(ctx, client, root, auth.From, big.NewInt(100))
	case api.ERC20Name:
		err = api.NewERC20(ctx, client, root, auth)
	case api.NftName:
		err = api.NewNft(ctx, client, root, auth)
	case api.GreeterName:
		err = api.NewGreeter(ctx, client, root)
	case api.SushiName:
		err = api.NewSushi(ctx, client, root)
	case api.DaoName:
		err = api.NewDao(ctx, client, root, auth)
	case api.Uniswapv2Name:
		err = api.NewUniswapv2(ctx, client, root, auth)
	default:
		log.Error("unexpected dump option")
		return
	}

	if err != nil {
		log.Error("dump traces for contract fail", "contract name", solName, "err", err)
	} else {
		log.Info("dump traces for contract successfully", "contract name", solName)
	}*/
	test()
}

func test() {
	l2url := "https://prealpha.scroll.io/l2"
	usdcAddr := common.HexToAddress("0x5bD45286170BA4EC157d5F501231E8367C31877a")
	usdtAddr := common.HexToAddress("0xD1457fbe8F34E112A2948F3d1749AfBbdC7FBF98")

	priv := "xss"

	privateKey, _ := crypto.HexToECDSA(priv)
	publicKey := privateKey.Public()
	publicKeyECDSA, _ := publicKey.(*ecdsa.PublicKey)
	from := crypto.PubkeyToAddress(*publicKeyECDSA)

	for i := 0; ; i++ {
		l2client, _ := ethclient.Dial(l2url)
		l2chainId, _ := l2client.ChainID(context.Background())

		l2nonce, _ := l2client.PendingNonceAt(context.Background(), from)
		fmt.Printf("l2nonce:%v\n", l2nonce)

		routerAddr := common.HexToAddress("0xee0e03c1a621084ca3c542f36e4a5d0230304471") //common.HexToAddress("0x4F4Eb5aC461c115191390D1760109F1EA185e609")
		router, _ := router.NewIUniswapV2Router02(routerAddr, l2client)

		transOpt, _ := bind.NewKeyedTransactorWithChainID(privateKey, l2chainId)
		transOpt.Nonce = big.NewInt(int64(l2nonce))

		amount := big.NewInt(1e18)

		path1 := make([]common.Address, 2)
		path1[0] = usdtAddr
		path1[1] = usdcAddr

		transOpt.Nonce = big.NewInt(int64(l2nonce))
		tx1, err := router.SwapExactTokensForTokens(transOpt, amount, big.NewInt(0), path1, from, amount)
		if err == nil {
			fmt.Printf("%v swap usdt to usdc:%v\n", i, tx1.Hash())
		} else {
			fmt.Printf("%v swap usdt to usdc, err:%v\n", i, err)
		}

		path2 := make([]common.Address, 2)
		path2[0] = usdcAddr
		path2[1] = usdtAddr

		receipt, err := bind.WaitMined(context.Background(), l2client, tx1)

		time.Sleep(4 * time.Second)

		transOpt.Nonce = big.NewInt(int64(l2nonce + 1))
		tx2, err := router.SwapExactTokensForTokens(transOpt, amount, big.NewInt(0), path2, from, amount)
		if err == nil {
			fmt.Printf("%v swap usdc to usdt:%v\n", i, tx2.Hash())
		} else {
			fmt.Printf("%v swap usdc to usdt, err:%v\n", i, err)
		}

		time.Sleep(10 * time.Second)
	}
}
