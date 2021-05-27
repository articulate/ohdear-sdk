help:
	@echo "+ $@"
	@grep -E '(^[a-zA-Z0-9\._-]+:.*?##.*$$)|(^##)' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}{printf "\033[32m%-30s\033[0m %s\n", $$1, $$2}' | sed -e 's/\[32m##      /[33m/'
.PHONY: help

mod: ## Install/sync go.mod
	@echo "+ $@"
	@go mod tidy
.PHONY: mod

test: ## Run tests
	@echo "+ $@"
	@go run github.com/onsi/ginkgo/ginkgo ./ohdear/
.PHONY: test

lint: ## Lint Go code
	@echo "+ $@"
	@golangci-lint run
.PHONY: lint

fix: ## Try to fix linting issues
	@echo "+ $@"
	@golangci-lint run --fix
.PHONY: fix
