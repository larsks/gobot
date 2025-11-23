GOTEST ?=  go test
GOFORMAT ?= go fmt
GOLANGCI_LINT ?= golangci-lint

all:

.PHONY: test
test:
	@echo "## TEST"
	@$(GOTEST) ./...

.PHONY: lint
lint:
	@echo "## LINT"
	@$(GOLANGCI_LINT) run

.PHONY: format
format:
	@echo "## FORMAT"
	@$(GOFORMAT) ./...
