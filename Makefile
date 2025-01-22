#!/usr/bin/env bash

.PHONY: lint test coverage help

lint: ## Lint the files
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.62.2
	@golangci-lint run

test: ## Run unittests
	@go test -race -short ./...

coverage: ## Generate global code coverage report
	@./coverage.sh;

help: ## Display this help screen
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
