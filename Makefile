# Project directory
PROJECT_DIR = $(shell pwd)

# Path to directory with binaries
PROJECT_BIN = $(PROJECT_DIR)/bin

# Full path to binary file
BIN_PATH = $(PROJECT_BIN)/$(BIN_NAME)

# Path to bin directory in GOPATH
GO_BIN_PATH = $(shell go env GOPATH)/bin

# Full path to linter
GOLANGCI_LINT = $(GO_BIN_PATH)/golangci-lint

.SILENT:

build:
	$(shell [ -f bin ] || mkdir -p $(PROJECT_BIN))
	go build -o $(PROJECT_BIN) ./...

run: build
	$(PROJECT_BIN)/student-base

test:
	go test -v ./...

.PHONY: .install-linter
.install-linter:
	[ -f $(GOLANGCI_LINT) ] || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(GO_BIN_PATH) v1.50.0

.PHONY: lint
lint: .install-linter
	$(GOLANGCI_LINT) run ./... 
	
.PHONY: lint-fast
lint-fast: .install-linter
	$(GOLANGCI_LINT) run ./... --fast