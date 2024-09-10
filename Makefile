#!make
include .env
export $(shell sed 's/=.*//' .env)
export DATE=$(shell date +%Y-%m-%d)
#SHELL:=/bin/bash
EXEC = docker-compose exec
RUN = docker-compose run --rm
START = docker-compose up -d
STOP = docker-compose stop
LOGS = docker-compose logs

LOCAL_BIN:=$(CURDIR)/bin
PATH:=$(LOCAL_BIN):$(PATH)

.PHONY: help

help: ## Display this help screen
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

status: ### Containers statuses
	@echo "\n\033[01;33m Containers statuses \033[0m"
	@docker-compose ps
.PHONY: status

up: ### Up docker-compose
	@echo "\n\033[0;33m Spinning up docker environment... \033[0m"
	@$(START)
	@$(MAKE) --no-print-directory status
.PHONY: up

stop: ### Stop docker-compose
	@echo "\n\033[0;33m Halting containers... \033[0m"
	@docker-compose stop
	@$(MAKE) --no-print-directory status
.PHONY: stop

down: ### Destroy containers
	@echo "\n\033[0;33m Destroy containers... \033[0m"
	@docker-compose down
	@$(MAKE) --no-print-directory status
.PHONY: down

down-vol: ### Destroy containers with volumes
	@echo "\n\033[0;33m Destroy containers with volumes... \033[0m"
	@docker-compose down --volumes
	@$(MAKE) --no-print-directory status
.PHONY: PHONY

restart: ### Restarting containers
	@echo "\n\033[0;33m Restarting containers... \033[0m"
	@$(STOP)
	@$(START)
	@$(MAKE) --no-print-directory status
.PHONY: restart

logs: ### Gophermart app container logs
	@$(LOGS) --tail=500 -f app
.PHONY: logs-app

bash-app: ### app container bash
	@$(EXEC) app /bin/sh
.PHONY: bash-app

tidy: ### runs app go mod tidy
	@$(RUN) app go mod tidy && go mod download
.PHONY: tidy

MIGRATION_NAME := $(or $(MIGRATION_NAME),migration_name)
migrate-create:  ### create new migration. With specific name: MIGRATION_NAME="some_name"
	GOBIN=$(LOCAL_BIN) migrate create -ext sql -dir migrations $(MIGRATION_NAME)
.PHONY: db-migrate-create

migrate-up: ### migration up
	GOBIN=$(LOCAL_BIN) migrate -path migrations -database '$(PG_URL_LOCAL)?sslmode=disable' up
.PHONY: db-migrate-up

migrate-up-force: ### migration up force to fix DB on
	GOBIN=$(LOCAL_BIN) migrate -path migrations -database '$(PG_URL_LOCAL)?sslmode=disable' force $(VERSION)
.PHONY: db-migrate-up-force

migrate-down: ### migration down
	GOBIN=$(LOCAL_BIN) migrate -path migrations -database '$(PG_URL_LOCAL)?sslmode=disable' down $(STEP)
.PHONY: db-migrate-down

linters: ### run linters
	GOBIN=$(LOCAL_BIN) golangci-lint run && go vet -vettool=$(which statictest) ./...
.PHONY: linters

bin-deps: ### install binaries
	GOBIN=$(LOCAL_BIN) go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
	GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.60.3
.PHONY: bin-deps

tests: ### run tests
	go test -v -cover ./internal/...
.PHONY: test