.PHONY: update trace_dumper geth clean docker start_docker

VERSION=scroll-v4.2.2

update: ## Let's keep it and docker version in consistent.
	go get -u github.com/scroll-tech/go-ethereum@${VERSION}

trace_dumper: ## Builds the trace_dumper instance.
	mkdir -p ${PWD}/bin
	go build -o ${PWD}/bin/trace_dumper

clean: ## Delete generated artifacts
	rm -rf ${PWD}/bin/*  ${PWD}/tracedata/*

start_docker:
	docker run --rm -p 8545:8545 -p 8546:8546 trace-dumper/l2geth:${VERSION}

docker: ## Build integration-test image
	docker build --no-cache -t trace-dumper/l2geth:${VERSION} ./docker/l2geth/.
