.PHONY: trace_dumper geth clean docker start-l1geth start-l2geth

IMAGE_NAME=l2geth
IMAGE_VERSION=latest

trace_dumper: ## Builds the trace_dumper instance.
	mkdir -p ${PWD}/bin
	go build -o ${PWD}/bin/trace_dumper

clean: ## Delete generated artifacts
	rm -r ${PWD}/bin/* ${PWD}/tracedata/*.json

docker: ## Creates a new l2geth docker
	docker build --no-cache -t scroll_l2geth_creator ./docker/l2geth/

start-l1geth: ## Starts l1geth docker container
	docker run --rm -p 8560:8545 scroll_l2geth_creator

start-l2geth: ## Starts l2geth docker container
	docker run --rm -p 8561:8545 scroll_l2geth_creator
