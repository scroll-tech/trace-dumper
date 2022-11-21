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

## deploy contract

> the contract supports: erc20, nft, greeter, sushi, dao, uniswapv2

```
make l2gethTool
./bin/l2gethTool --help
./bin/l2gethTool -contract erc20
```

## show trace list

```
ls -l ./tracedata/erc20_*.json
```