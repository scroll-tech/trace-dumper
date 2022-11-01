# l2gethTool

Pressure test tool for L2GETH, can stress test contracts, ordinary transfer transactions. The L2GETH service and
pressure test service are integrated.Subsequent Scroll services can also be integrated.

## make and start l2geth docker

> Create a pressure test execution environment(scroll folder)

```
make docker
docker run -it -p 8545:8545 -p 8546:8546 mask-pp/l2geth:latest
```

## make and start l2gethTool

```
make l2gethTool
./bin/l2gethTool
```

## deploy contracts

```

```

## Call api

> Call api to start pressure test,you can call it repeatedly.

* erc20 call

> count: total txs
> batch: the count of parallel sender
> tps: expected tps

```
curl --location --request POST 'http://127.0.0.1:8190' \
--header 'Content-Type: application/json' \
--data-raw '{
    "jsonrpc": "2.0",
    "method": "erc20_pressure",
    "params": [60000,40,5000],
    "id": 1
}'
```

* native call

```
curl --location --request POST 'http://127.0.0.1:8190' \
--header 'Content-Type: application/json' \
--data-raw '{
    "jsonrpc": "2.0",
    "method": "native_pressure",
    "params": [60000,40,5000],
    "id": 1
}'
```

* greeter call

```
// call setValue
curl --location --request POST 'http://127.0.0.1:8190' \
--header 'Content-Type: application/json' \
--data-raw '{
    "jsonrpc": "2.0",
    "method": "greeter_setValue",
    "params": [1],
    "id": 1
}'

// call retrieve
curl --location --request POST 'http://127.0.0.1:8190' \
--header 'Content-Type: application/json' \
--data-raw '{
    "jsonrpc": "2.0",
    "method": "greeter_retrieve",
    "params": [],
    "id": 1
}'
```

## See detail

```
tail -f logs/geth.log
tail -f logs/l2gethTool.log
```

## docker compose

```
// Build service images
make docker

// Start services
cd docker
// run docker compose
docker-compose up -d

// show docker services status
docker-compose ps

// look the l2gethtool_service log
docker-compose logs l2gethtool_service -f

// look the l2geth_service log
docker-compose logs l2geth_service -f
```

## A complete guide on how to generate trace

### For greeter.sol

```
curl --location --request POST 'http://127.0.0.1:8190' \
--header 'Content-Type: application/json' \
--data-raw '{
    "jsonrpc": "2.0",
    "method": "greeter_setValue",
    "params": [1],
    "id": 1
}'
```

You can now see a new block mined, the blockHash can be found in the docker-compose console output. Let say
it's `0xf8fdf7607eab03d06ce05ba8141225df13d5ef5c808542ebc8893e88e56e8d56`. You can then run the following commands:

```
curl --location --request POST 'http://127.0.0.1:8545' --header 'Content-Type: application/json' --data-raw '{"jsonrpc": "2.0","method": "eth_getBlockResultByHash","params":["0xf8fdf7607eab03d06ce05ba8141225df13d5ef5c808542ebc8893e88e56e8d56"],"id": 1}'
```

The reponse is the trace json.

### For erc20

```
curl --location --request POST 'http://127.0.0.1:8190' \
--header 'Content-Type: application/json' \
--data-raw '{
    "jsonrpc": "2.0",
    "method": "erc20_transfer",
    "params": ["0xC5BBeAC655Df555a84029b9Da1F1C4d675456F4f",100],
    "id": 1
}'
```

You can now see a new block mined, the blockHash can be found in the docker-compose console output.

```
curl --location --request POST 'http://127.0.0.1:8545' --header 'Content-Type: application/json' --data-raw '{"jsonrpc": "2.0","method": "eth_getBlockResultByHash","params":["0xf8fdf7607eab03d06ce05ba8141225df13d5ef5c808542ebc8893e88e56e8d56"],"id": 1}'
```

The reponse is the trace json.

## Run the integration test using script

### Txs block proving test

Set `TX_COUNT` and `BLOCK_COUNT` in ENV variables. `TX_COUNT` and `BLOCK_COUNT` are 1 by default.

For example,

```
export TX_COUNT=99
export BLOCK_COUNT=2
make txs-test
```

will expect sending 99 erc20 transfer txs and 2 successful proof submissions.

If you want to test native ETH transfer, you will need to futher set `TX_TYPE` as `native`. i.e.,

```
export TX_TYPE=native
export TX_COUNT=99
export BLOCK_COUNT=2
make txs-test
```

### Greeter `set_value` proving test

We also support testing Greeter contract `set_value` by running

```
make greeter-test
```

Currently only supports 1 tx and it's in 1 block.

### Roller registration test

```
make registration-test
```

## CI & CD

Before merge PR into main branch, be sure comment run full/light test to trigger CI.  
If no CI been triggered, that PR can not be merged.  
For now, `run light test` builds the docker image and publishes it, whereas `run full test` runs `go test -v ./...`
additionally.
