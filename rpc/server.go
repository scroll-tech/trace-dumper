package rpc

import (
	"net"
	"net/http"

	"github.com/scroll-tech/go-ethereum/log"
	"github.com/scroll-tech/go-ethereum/rpc"
)

var (
	srv = rpc.NewServer()
)

func RunServer(url string) (*http.Server, error) {
	handler, _, err := startHTTPEndpoint(url, rpc.DefaultHTTPTimeouts, srv)
	if err != nil {
		return nil, err
	}
	log.Info("HTTP endpoint opened", "url", url)
	return handler, nil
}

func Register(apis []rpc.API) error {
	for _, api := range apis {
		if err := srv.RegisterName(api.Namespace, api.Service); err != nil {
			return err
		}
	}
	return nil
}

// startHTTPEndpoint starts the HTTP RPC endpoint.
func startHTTPEndpoint(endpoint string, timeouts rpc.HTTPTimeouts, handler http.Handler) (*http.Server, net.Addr, error) {
	// start the HTTP listener
	var (
		listener net.Listener
		err      error
	)
	if listener, err = net.Listen("tcp", endpoint); err != nil {
		return nil, nil, err
	}
	// Bundle and start the HTTP server
	httpSrv := &http.Server{
		Handler:      handler,
		ReadTimeout:  timeouts.ReadTimeout,
		WriteTimeout: timeouts.WriteTimeout,
		IdleTimeout:  timeouts.IdleTimeout,
	}
	go func() {
		_ = httpSrv.Serve(listener)
	}()
	return httpSrv, listener.Addr(), err
}
