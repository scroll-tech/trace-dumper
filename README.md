# l2gethTool

A very simple l2geth tool, it can deploy contract and call contract api and get trace store into file just in one cmd.

## Update dependence

```
go get -v github.com/scroll-tech/go-ethereum@staging
go mod tidy
```

## make and start l2geth docker

> Create environment

```
make docker
docker run -it -p 8545:8545 -p 8546:8546 trace-dumper/l2geth:latest
```

## dump traces

> the contract supports: erc20, nft, greeter, sushi, dao, uniswapv2

```
make trace_dumper
./bin/trace_dumper --help
./bin/trace_dumper -dump erc20
```

## show trace list

```
ls -l ./tracedata/erc20_*.json
```