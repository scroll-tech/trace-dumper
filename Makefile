.PHONY: trace_dumper geth clean docker

IMAGE_NAME=l2geth
IMAGE_VERSION=latest

trace_dumper: ## Builds the trace_dumper instance.
	mkdir -p ${PWD}/bin
	cd ${PWD}/bin/; go build -o trace_dumper -v ../

clean: ## Delete integration test environment
	rm -r ${PWD}/scroll ${PWD}/bin/* 2>/dev/null

docker: ## Build integration-test image
	docker build -t trace-dumper/${IMAGE_NAME}:${IMAGE_VERSION} ./docker/l2geth/.
