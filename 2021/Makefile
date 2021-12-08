export GOFLAGS=-mod=vendor

BIN_DIR=./bin

# allows to specify which tests to be run (ex: TEST_PATTERN=FooTest make test)
TEST_PATTERN ?= .

OUTPUT_DIR ?= .

# allow passing -ldflags, etc for release builds
BUILD_ARGS ?=

.PHONY: build
build: ## build the binary
	@printf ">> building binary..."
	@go build $(BUILD_ARGS) -o ./bin/day1 ./cmd/day1/...
	@go build $(BUILD_ARGS) -o ./bin/day2 ./cmd/day2/...
	@go build $(BUILD_ARGS) -o ./bin/day3 ./cmd/day3/...
	@go build $(BUILD_ARGS) -o ./bin/day4 ./cmd/day4/...
	@go build $(BUILD_ARGS) -o ./bin/day5 ./cmd/day5/...
	$(call display_check)

.PHONY: test
test: ## run the tests
	@printf ">> running tests..."
	@go test -count=1 -v -race -coverprofile=$(OUTPUT_DIR)/coverage.out ./... -run $(TEST_PATTERN) -timeout=30s

.PHONY: cover
cover: test ## Run all the tests and opens the coverage report
	@printf ">> generating test coverage..."
	@go tool cover -html=$(OUTPUT_DIR)/coverage.out

.PHONY: tidy
tidy: ## go mod tidy
	@printf ">> tidying up..."
	@go mod tidy
	$(call display_check)

.PHONY: vendor
vendor: tidy ## vendor dependencies
	@printf ">> vendoring..."
	@go mod vendor
	$(call display_check)

.PHONY: fmt
fmt: ## go fmt everything
	@printf ">> formatting..."
	@find . -name '*.go' -not -wholename './vendor/*' | while read -r file; do gofmt -w -s "$$file"; done
	$(call display_check)

.PHONY: lint
lint: ## run the linters
	@printf ">> running golangci-lint..."
	@./_tools/bin/golangci-lint run
	$(call display_check)

.DEFAULT_GOAL := build
default: build

define display_check
	@printf " \033[32m✔︎\033[0m\n"
endef
