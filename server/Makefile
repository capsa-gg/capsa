help: ## Shows all commands
	@echo 'All Makefile commands:'
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: bootstrap
bootstrap: ## Installs Go packages and sets defaults
	@echo "Installing Go packages using 'go install'"
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install golang.org/x/tools/cmd/goimports@latest
	go install github.com/air-verse/air@latest
	go install go101.org/golds@latest
	go install golang.org/x/lint/golint@latest
	go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
	go install github.com/jstemmer/go-junit-report@latest
	go install github.com/t-yuki/gocover-cobertura@latest
	go install github.com/swaggo/swag/cmd/swag@latest

.PHONY: build
build: ## Runs the Golang build for production on Linux
	./bin/build

.PHONY: build-all
build-all: ## Runs the Golang build for production for all supported platforms
	./bin/build-all

.PHONY: dev
dev: ## Runs the Capsa server with hot reloading for development
	air

.PHONY: run
run: ## Runs the application without hot reloading
	go run .

.PHONY: test
test: ## Runs the unit tests
	go test ./...

.PHONY: bench
bench: ## Runs benchmarking, with N=10, printing to stdout and saving results to bench.txt
	go test -bench=. -count=10 -benchmem ./... | tee bench.txt

.PHONY: fmt
fmt: ## Code style formatting and linting check for the project
	goimports -w -local github.com/capsa-gg/capsa/server .
	go fmt ./...
	go vet ./...
	golint ./...
	@echo "Checking linting:"
	golangci-lint run --sort-results


.PHONY: docs
docs: ## Runs documentation website
	golds .

.PHONY: ci
ci: ## Runs the CI tasks
	@echo "Checking if golangci-lint is installed"
	@golangci-lint --version
	@echo "Running linting and static code analysis"
	golangci-lint run --sort-results --timeout 1h
	go vet ./...
	golint ./...
	goimports -e -d -local github.com/capsa-gg/capsa/server .
	@echo "Running tests with coverage"
	go test -coverprofile=coverage_raw.txt -covermode count -v 2>&1 ./... | go-junit-report -set-exit-code > junit.xml
	@echo "Removing packages from coverage report that don't require testing"
	./bin/clean-coverage coverage_raw.txt
	gocover-cobertura < coverage.txt > cobertura.xml

.PHONY: sql
sql: ## Checks SQL for syntax and type errors, then generates Go code for ./SQL/ files
	@sqlc compile
	@sqlc generate
	@sqlc compile

.PHONY: swagger
swagger: ## Generates Swagger (OpenAPI v2) documentation based on HTTP handlers
	swag init --output swagger --generalInfo internal/server/swagger.go

.PHONY: migrateup
migrateup: ## Runs the database migrations up
	go run . db migrate -d up

.PHONY: migratedown
migratedown: ## Runs the database migrations down
	go run . db migrate -d down

.PHONY: migration
migration: ## Asks for the new migration name and creates migration files
	./bin/create_migration

.PHONY: newdb
newdb: ## Migrates db down and up - WARNING: ALL DB DATA IS DELETED
	make migratedown
	make migrateup

.PHONY: jwkset
jwkset: ## Generates a public and private key for JWKs and writes them to disk
	go run . jwk genkeys -w
