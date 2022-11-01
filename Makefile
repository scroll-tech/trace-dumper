.PHONY: l2gethTool geth clean docker

IMAGE_NAME=l2geth
IMAGE_VERSION=latest

l2gethTool: ## Builds the l2gethTool instance.
	mkdir -p ${PWD}/bin
	cd ${PWD}/bin/; go build -o l2gethTool -v ../ 

geth: ## Builds the L2Geth instance.
	./scripts/build_geth.sh

clean: ## Delete integration test environment
	rm -r ${PWD}/scroll ${PWD}/bin/* 2>/dev/null

docker: ## Build integration-test image
	docker build -t mask-pp/${IMAGE_NAME}:${IMAGE_VERSION} ./docker/l2geth/.
