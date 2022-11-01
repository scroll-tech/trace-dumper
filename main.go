package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"os"
	"os/signal"
	"syscall"

	"github.com/mattn/go-colorable"
	"github.com/mattn/go-isatty"
	"github.com/scroll-tech/go-ethereum/core/types"
	"github.com/scroll-tech/go-ethereum/ethclient"
	"github.com/scroll-tech/go-ethereum/log"

	"tool/accounts"
	"tool/api"
	"tool/rpc"
)

var (
	endpoint = flag.String("endpoint", "ws://127.0.0.1:8546", "The endpoint to connect to blockchain node")
	keystore = flag.String("keystore", "./docker/l2geth/genesis-keystore", "Keystore file path")
	password = flag.String("password", "scrolltest", "The keystore password")
	port     = flag.Int("port", 8190, "API server port")
	limit    = flag.Int("limit", 5, "Useful internal accounts")
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

	// create client
	client, err := ethclient.DialContext(ctx, *endpoint)
	if err != nil {
		log.Crit("Failed to connect chain node", "err", err)
	}

	// load keystore file
	accounts, err := accounts.NewAccounts(ctx, *limit, client, *keystore, *password)
	if err != nil {
		log.Crit("failed to create accounts", "err", err)
	}

	// create contracts factory
	if err = api.NewFactory(accounts, client); err != nil {
		panic("failed to create contract factory, err: " + err.Error())
	}

	// start rpc api
	handler, err := rpc.RunServer(fmt.Sprintf("0.0.0.0:%d", *port))
	if err != nil {
		panic(err)
	}
	defer handler.Close()

	go loop(ctx, client)
	daemon()
}

// Monitor batch execution results
func loop(ctx context.Context, client *ethclient.Client) {
	headCh := make(chan *types.Header, 8)
	sub, err := client.SubscribeNewHead(ctx, headCh)
	if err != nil {
		log.Crit("Failed to subscribe new header", "err", err)
	}
	defer sub.Unsubscribe()

	for {
		select {
		case <-ctx.Done():
			return

		case head := <-headCh:
			block, err := client.BlockByHash(ctx, head.Hash())
			if err != nil {
				log.Warn("Failed to get block", "err", err)
			}
			log.Info("New block", "number", block.NumberU64(), "blockHash", block.Hash().String(), "tx count", block.Transactions().Len())
		}
	}
}

func daemon() {
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, syscall.SIGINT, syscall.SIGTERM)
	defer signal.Stop(sigc)
	<-sigc
}
