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

# Full path to migrate tool
MIGRATE_TOOL = $(GO_BIN_PATH)/migrate

# DB parameters
DB_HOST = localhost
DB_PORT = 5432
DB_NAME = student_information_srv
DB_USER_NAME = root
DB_USER_PASSWORD = secret

# DB connection string
CONNECTION_STRING = postgresql://$(DB_USER_NAME):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)

# Path to migration files
MIGRATION_PATH = internal/storage/db/migration

# Migration tool command string
MIGRATION_STRING = $(MIGRATE_TOOL)/migrate -path $(MIGRATION_PATH) -database "$(CONNECTION_STRING)/$(DB_NAME)?sslmode=disable" -verbose

.SILENT:

build:
	$(shell [ -f bin ] || mkdir -p $(PROJECT_BIN))
	go build -o $(PROJECT_BIN) ./...

run: build
	$(PROJECT_BIN)/student-base

test:
	go test -v ./...

createdb:
	psql $(CONNECTION_STRING) -c 'CREATE DATABASE ${DB_NAME};'

dropdb:
	psql $(CONNECTION_STRING) -c 'DROP DATABASE ${DB_NAME};'

.PHONY: migrateup
migrateup: .install-migrate-tool
	$(MIGRATION_STRING) up

.PHONY: migratedown
migratedown: .install-migrate-tool
	$(MIGRATION_STRING) down

.PHONY: .install-linter
.install-linter:
	[ -f $(GOLANGCI_LINT) ] || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(GO_BIN_PATH) v1.50.0

.PHONY: lint
lint: .install-linter
	$(GOLANGCI_LINT) run ./... 
	
.PHONY: lint-fast
lint-fast: .install-linter
	$(GOLANGCI_LINT) run ./... --fast

.PHONY: .install-migrate-tool
.install-migrate-tool:
	GOBIN=$(MIGRATE_TOOL) go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.15.2
