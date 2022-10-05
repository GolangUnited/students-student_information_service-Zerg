PROJECT_DIR = $(shell pwd)

SOURCE_PATH = $(PROJECT_DIR)/cmd/main.go

BIN_NAME = stud
PROJECT_BIN = $(PROJECT_DIR)/bin
BIN_PATH = $(PROJECT_BIN)/$(BIN_NAME)

$(shell [ -f bin ] || mkdir -p $(PROJECT_BIN))
PATH := $(PROJECT_BIN):$(PATH)
GOLANGCI_LINT = $(PROJECT_BIN)/golangci-lint

.SILENT:

build:
	CGO_ENABLED=0 GOOS="linux" go build -o $(BIN_PATH) $(SOURCE_PATH)

test:
	go test -v ./...

.PHONY: lint
lint:
	golangci-lint run ./...

.PHONY: lint-fast
lint-fast:
	golangci-lint run ./... --fast

