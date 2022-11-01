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
# contract_name: erc20, nft, greeter, sushi, vote, uniswapv2
curl --location --request POST 'http://127.0.0.1:8190' \
--header 'Content-Type: application/json' \
--data-raw '{
    "jsonrpc": "2.0",
    "method": "rpc_deployContract",
    "params": ["contract_name"],
    "id": 1
}'
```

## show deployed contracts' addresses.

```
curl --location --request GET 'http://localhost:8190' \
--header 'Content-Type: application/json' \
--data-raw '{
    "jsonrpc": "2.0",
    "method": "rpc_addresses",
    "params": [],
    "id": 1
}'
```

## Call contracts' apis

> Call api to start pressure test,you can call it repeatedly.

* call erc20 transfer

```
curl --location --request POST 'http://127.0.0.1:8190' \
--header 'Content-Type: application/json' \
--data-raw '{
    "jsonrpc": "1.0",
    "method": "erc20_transfer",
    "params": ["0xc1e376BFA31B7DDC1f6B6b4020064326f986d189", 1],
    "id": 1
}'
```

* call greeter setValue

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

## get trace

```
curl --location --request POST 'http://127.0.0.1:8545' \
--header 'Content-Type: application/json' \
--data-raw '{
    "jsonrpc": "1.0",
    "method": "debug_getBlockResultByNumberOrHash",
    "params": ["0x6ca58af908dc955dc10ebed6653be95394489808d9e48480e66e98975ca06714"],
    "id": 1
}'
```