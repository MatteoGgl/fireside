ifneq (,$(wildcard ./.env))
    include .env
    export
endif

# ==================================================================================== #
# HELPERS
# ==================================================================================== #

## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/ /'

.PHONY: confirm
confirm:
	@echo -n 'Are you sure? [y/N] ' && read ans && [ $${ans:-N} = y ]

# ==================================================================================== #
# DEVELOPMENT
# ==================================================================================== #

## run: run the cmd/web application
.PHONY: run
run:
	@go run ./cmd/web

## mix/watch: starts the mix watcher
.PHONY: mix/watch
mix/watch:
	@npx mix watch

## mix/build: builds the ui
.PHONY: mix/build
mix/build:
	@npx mix build

## mix/prod: builds the ui for production
.PHONY: mix/prod
mix/prod:
	@npx mix build --production

## migrations/new name=$1: create a new database migration
.PHONY: migrations/new
migrations/new:
	@echo 'Creating migration files for ${name}'
	@migrate create -seq -ext=.sql -dir=./migrations ${name}

## migrations/up: apply all up database migrations
.PHONY: migrations/up
migrations/up: confirm
	@echo 'Running up migrations...'
	@migrate -path=./migrations -database=${FIRESIDE_DB_DSN} up

## migrations/down: apply all down database migrations
.PHONY: migrations/down
migrations/down:
	@echo 'Running down migrations...'
	@migrate -path=./migrations -database=${FIRESIDE_DB_DSN} down

# ==================================================================================== #
# QUALITY
# ==================================================================================== #

## audit: tidy deps and format, vet and test all code
.PHONY: audit
audit:
	@echo 'Tidying and verifying module dependencies...'
	go mod tidy
	go mod verify
	@echo 'Formatting code...'
	go fmt ./...
	@echo 'Vetting code...'
	go vet ./...
	staticcheck ./...

# ==================================================================================== #
# BUILD
# ==================================================================================== #

## build/snapshot: build a snapshot using goreleaser
.PHONY: build/snapshot
build/snapshot: mix/prod
	goreleaser release --snapshot --rm-dist

## build/release: build a release using goreleaser
.PHONY: build/release
build/release: audit mix/prod
	goreleaser release