export GO111MODULE := on
DIST_DIR ?= ${PWD}/bin
HOME := ${HOME}
PROJECT_NAME := terraform-provider-spinnaker

.PHONY: help
help:  ## Show help messages for make targets
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[32m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: setup
setup: ## Download tools required to check and build
ifeq ($(shell uname -s), Darwin)
	@brew install golangci-lint
else
	@curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh
endif

.PHONY: build
build: ## Download dependencies and compile project
	@CGO_ENABLED=0 go build -o ${DIST_DIR}/${PROJECT_NAME} main.go

.PHONY: install
install: build ## Builds and installs the terraform provider
	@cp ${DIST_DIR}/${PROJECT_NAME} ${HOME}/.terraform.d/plugins/

.PHONY: lint
lint: ## Runs go lint checks
	@golangci-lint run ./...

.PHONY: test
test: ## Runs go tests
	@go test --cover ./...

.PHONY: test
acc: ## Runs go acceptance tests
	@TF_ACC=1 go test -v --cover ./...
