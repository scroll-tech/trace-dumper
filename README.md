# trace-dumper

A very simple tool, it can deploy contract and call contract api and get trace store into file just in one cmd.

## upgrade dependence

1. upgrade go mod
```
go get -v github.com/scroll-tech/go-ethereum@staging
go mod tidy
```

2. manaully edit `scrolltech/l2geth`'s tag in `docker/l2geth/Dockerfile` to the corresponding version

For example, `scrolltech/l2geth:prealpha-v3.1`

## make and start l2geth docker

create environment (**need to keep it running**)

```
make docker
docker run -it -p 8545:8545 -p 8546:8546 trace-dumper/l2geth:latest
```

## dump traces

```
make trace_dumper
./bin/trace_dumper --help
./bin/trace_dumper -dump erc20 # options: erc20, nft, greeter, sushi, dao, uniswapv2
```

## show trace list

```
ls -l ./tracedata/erc20_*.json
```
