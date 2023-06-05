.PHONY: update trace_dumper geth clean docker start_docker

update: ## Let's keep it and docker version in consistent.
	go get -u github.com/scroll-tech/go-ethereum@scroll-v3.3.1

trace_dumper: ## Builds the trace_dumper instance.
	mkdir -p ${PWD}/bin
	go build -o ${PWD}/bin/trace_dumper

clean: ## Delete generated artifacts
	rm -rf ${PWD}/bin/*  ${PWD}/tracedata/*

start_docker:
	docker run --rm -p 8545:8545 -p 8546:8546 trace-dumper/l2geth

docker: ## Build integration-test image
	docker build --no-cache -t trace-dumper/l2geth:latest ./docker/l2geth/.
