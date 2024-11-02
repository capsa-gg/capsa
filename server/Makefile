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
build: ## Runs the Golang build for production, TODO: ./bin/build script
	go build .

.PHONY: run
run: ## Runs the application without hot reloading
	go run .

.PHONY: test
test: ## Runs the unit tests
	go test ./...

.PHONY: fmt
fmt: ## Code style formatting and linting check for the project
	goimports -w -local github.com/lucianonooijen/capsa/server .
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
	goimports -e -d -local github.com/lucianonooijen/capsa/server .
	@echo "Running tests with coverage"
	go test -coverprofile=coverage.txt -covermode count -v 2>&1 ./... | go-junit-report -set-exit-code > junit.xml
	gocover-cobertura < coverage.txt > cobertura.xml
