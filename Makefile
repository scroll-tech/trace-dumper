.PHONY: update trace_dumper geth clean docker

IMAGE_NAME=l2geth
IMAGE_VERSION=latest

update: ## Let's keep it and docker version in consistent.
	go get -u github.com/scroll-tech/go-ethereum@scroll-v3.3.1

trace_dumper: ## Builds the trace_dumper instance.
	mkdir -p ${PWD}/bin
	go build -o ${PWD}/bin/trace_dumper

clean: ## Delete generated artifacts
	rm -r ${PWD}/bin/* ${PWD}/tracedata/*.json

docker: ## Build integration-test image
	docker build -t trace-dumper/${IMAGE_NAME}:${IMAGE_VERSION} ./docker/l2geth/.
