# l2gethTool

Pressure test tool for L2GETH, can stress test contracts, ordinary transfer transactions. The L2GETH service and
pressure test service are integrated.Subsequent Scroll services can also be integrated.

## Update dependence

```
go get -v github.com/scroll-tech/go-ethereum@staging
go mod tidy
```

## make and start l2geth docker

> Create a pressure test execution environment(scroll folder)

```
make docker
docker run -it -p 8545:8545 -p 8546:8546 mask-pp/l2geth:latest
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