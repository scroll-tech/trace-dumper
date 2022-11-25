.PHONY: trace_dumper geth clean docker

IMAGE_NAME=l2geth
IMAGE_VERSION=latest

trace_dumper: ## Builds the trace_dumper instance.
	mkdir -p ${PWD}/bin
	go build -o ${PWD}/bin/trace_dumper

clean: ## Delete generated artifacts
	rm -r ${PWD}/bin/*
	rm -r ${PWD}/tracedata/*.json

docker: ## Build integration-test image
	docker build -t trace-dumper/${IMAGE_NAME}:${IMAGE_VERSION} ./docker/l2geth/.
