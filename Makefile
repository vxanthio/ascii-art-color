# Makefile for ascii-art project
# Professional build automation for Go projects

# ==================================================================================== #
# VARIABLES
# ==================================================================================== #

# Binary name
BINARY_NAME=ascii-art

# Build output directory
BUILD_DIR=bin

# Coverage files
COVERAGE_FILE=coverage.out
COVERAGE_HTML=coverage.html

# Version information (can be overridden)
VERSION?=$(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
COMMIT=$(shell git rev-parse --short HEAD 2>/dev/null || echo "unknown")
BUILD_TIME=$(shell date -u '+%Y-%m-%d_%H:%M:%S')

# Go build flags
LDFLAGS=-ldflags "-X main.Version=$(VERSION) -X main.Commit=$(COMMIT) -X main.BuildTime=$(BUILD_TIME)"

# Color output
COLOUR_GREEN=\033[0;32m
COLOUR_RED=\033[0;31m
COLOUR_BLUE=\033[0;34m
COLOUR_END=\033[0m

# ==================================================================================== #
# HELPERS
# ==================================================================================== #

## help: Display this help message
.PHONY: help
help:
	@echo ''
	@echo 'Usage:'
	@echo '  ${COLOUR_BLUE}make${COLOUR_END} ${COLOUR_GREEN}<target>${COLOUR_END}'
	@echo ''
	@echo 'Targets:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/ /'

# ==================================================================================== #
# QUALITY CONTROL
# ==================================================================================== #

## fmt: Format all Go files
.PHONY: fmt
fmt:
	@echo "${COLOUR_BLUE}Formatting code...${COLOUR_END}"
	@gofmt -w .
	@goimports -w . 2>/dev/null || true
	@echo "${COLOUR_GREEN}✓ Code formatted${COLOUR_END}"

## vet: Run go vet on all packages
.PHONY: vet
vet:
	@echo "${COLOUR_BLUE}Running go vet...${COLOUR_END}"
	@go vet ./...
	@echo "${COLOUR_GREEN}✓ go vet passed${COLOUR_END}"

## lint: Run golangci-lint
.PHONY: lint
lint:
	@echo "${COLOUR_BLUE}Running linters...${COLOUR_END}"
	@if command -v golangci-lint > /dev/null; then \
		golangci-lint run; \
		echo "${COLOUR_GREEN}✓ Linters passed${COLOUR_END}"; \
	else \
		echo "${COLOUR_RED}✗ golangci-lint not installed${COLOUR_END}"; \
		echo "Install with: go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest"; \
		exit 1; \
	fi

## check: Run all quality checks (fmt, vet, lint)
.PHONY: check
check: fmt vet lint
	@echo "${COLOUR_GREEN}✓ All quality checks passed${COLOUR_END}"

# ==================================================================================== #
# DEVELOPMENT
# ==================================================================================== #

## run: Run the application with default arguments
.PHONY: run
run:
	@go run . "Hello World"

## run-example: Run the application with all banner styles
.PHONY: run-example
run-example:
	@echo "${COLOUR_BLUE}Standard banner:${COLOUR_END}"
	@go run . "ASCII" standard
	@echo ""
	@echo "${COLOUR_BLUE}Shadow banner:${COLOUR_END}"
	@go run . "ASCII" shadow
	@echo ""
	@echo "${COLOUR_BLUE}Thinkertoy banner:${COLOUR_END}"
	@go run . "ASCII" thinkertoy

## build: Build the binary
.PHONY: build
build:
	@echo "${COLOUR_BLUE}Building $(BINARY_NAME)...${COLOUR_END}"
	@mkdir -p $(BUILD_DIR)
	@go build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME) .
	@echo "${COLOUR_GREEN}✓ Binary built: $(BUILD_DIR)/$(BINARY_NAME)${COLOUR_END}"

## install: Install the binary to GOPATH/bin
.PHONY: install
install:
	@echo "${COLOUR_BLUE}Installing $(BINARY_NAME)...${COLOUR_END}"
	@go install $(LDFLAGS)
	@echo "${COLOUR_GREEN}✓ Installed to $(shell go env GOPATH)/bin/$(BINARY_NAME)${COLOUR_END}"

# ==================================================================================== #
# TESTING
# ==================================================================================== #

## test: Run all tests
.PHONY: test
test:
	@echo "${COLOUR_BLUE}Running tests...${COLOUR_END}"
	@go test -v -race ./...
	@echo "${COLOUR_GREEN}✓ All tests passed${COLOUR_END}"

## test-short: Run tests without integration tests
.PHONY: test-short
test-short:
	@echo "${COLOUR_BLUE}Running short tests...${COLOUR_END}"
	@go test -short -v ./...
	@echo "${COLOUR_GREEN}✓ Short tests passed${COLOUR_END}"

## coverage: Generate test coverage report
.PHONY: coverage
coverage:
	@echo "${COLOUR_BLUE}Generating coverage report...${COLOUR_END}"
	@go test -coverprofile=$(COVERAGE_FILE) ./...
	@go tool cover -html=$(COVERAGE_FILE) -o $(COVERAGE_HTML)
	@echo "${COLOUR_GREEN}✓ Coverage report generated${COLOUR_END}"
	@echo "  - Text: $(COVERAGE_FILE)"
	@echo "  - HTML: $(COVERAGE_HTML)"
	@go tool cover -func=$(COVERAGE_FILE) | grep total | awk '{print "  - Total Coverage: " $$3}'

## coverage-view: Generate and open coverage report in browser
.PHONY: coverage-view
coverage-view: coverage
	@echo "${COLOUR_BLUE}Opening coverage report...${COLOUR_END}"
	@which xdg-open > /dev/null && xdg-open $(COVERAGE_HTML) || \
	 which wslview > /dev/null && wslview $(COVERAGE_HTML) || \
	 which open > /dev/null && open $(COVERAGE_HTML) || \
	 echo "Please open $(COVERAGE_HTML) manually"

## bench: Run benchmarks
.PHONY: bench
bench:
	@echo "${COLOUR_BLUE}Running benchmarks...${COLOUR_END}"
	@go test -bench=. -benchmem ./...

## bench-cpu: Run benchmarks with CPU profiling
.PHONY: bench-cpu
bench-cpu:
	@echo "${COLOUR_BLUE}Running benchmarks with CPU profiling...${COLOUR_END}"
	@go test -bench=. -benchmem -cpuprofile=cpu.prof ./...
	@echo "${COLOUR_GREEN}✓ CPU profile saved to cpu.prof${COLOUR_END}"
	@echo "  View with: go tool pprof cpu.prof"

## bench-mem: Run benchmarks with memory profiling
.PHONY: bench-mem
bench-mem:
	@echo "${COLOUR_BLUE}Running benchmarks with memory profiling...${COLOUR_END}"
	@go test -bench=. -benchmem -memprofile=mem.prof ./...
	@echo "${COLOUR_GREEN}✓ Memory profile saved to mem.prof${COLOUR_END}"
	@echo "  View with: go tool pprof mem.prof"

# ==================================================================================== #
# BUILD & RELEASE
# ==================================================================================== #

## build-all: Build binaries for all platforms
.PHONY: build-all
build-all: build-linux build-darwin build-windows
	@echo "${COLOUR_GREEN}✓ All platform binaries built${COLOUR_END}"

## build-linux: Build Linux binaries (amd64 and arm64)
.PHONY: build-linux
build-linux:
	@echo "${COLOUR_BLUE}Building for Linux...${COLOUR_END}"
	@mkdir -p $(BUILD_DIR)
	@GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-linux-amd64 .
	@GOOS=linux GOARCH=arm64 go build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-linux-arm64 .
	@echo "${COLOUR_GREEN}✓ Linux binaries built${COLOUR_END}"

## build-darwin: Build macOS binaries (amd64 and arm64)
.PHONY: build-darwin
build-darwin:
	@echo "${COLOUR_BLUE}Building for macOS...${COLOUR_END}"
	@mkdir -p $(BUILD_DIR)
	@GOOS=darwin GOARCH=amd64 go build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-darwin-amd64 .
	@GOOS=darwin GOARCH=arm64 go build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-darwin-arm64 .
	@echo "${COLOUR_GREEN}✓ macOS binaries built${COLOUR_END}"

## build-windows: Build Windows binary (amd64)
.PHONY: build-windows
build-windows:
	@echo "${COLOUR_BLUE}Building for Windows...${COLOUR_END}"
	@mkdir -p $(BUILD_DIR)
	@GOOS=windows GOARCH=amd64 go build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-windows-amd64.exe .
	@echo "${COLOUR_GREEN}✓ Windows binary built${COLOUR_END}"

# ==================================================================================== #
# CLEANUP
# ==================================================================================== #

## clean: Remove build artifacts and coverage files
.PHONY: clean
clean:
	@echo "${COLOUR_BLUE}Cleaning...${COLOUR_END}"
	@rm -rf $(BUILD_DIR)
	@rm -f $(COVERAGE_FILE) $(COVERAGE_HTML)
	@rm -f cpu.prof mem.prof
	@go clean
	@echo "${COLOUR_GREEN}✓ Cleaned${COLOUR_END}"

## clean-all: Remove all generated files including test cache
.PHONY: clean-all
clean-all: clean
	@echo "${COLOUR_BLUE}Deep cleaning...${COLOUR_END}"
	@go clean -testcache
	@go clean -modcache
	@echo "${COLOUR_GREEN}✓ Deep cleaned${COLOUR_END}"

# ==================================================================================== #
# CI/CD
# ==================================================================================== #

## ci: Run all CI checks (quality + tests + build)
.PHONY: ci
ci: check test build
	@echo "${COLOUR_GREEN}✓ CI checks passed${COLOUR_END}"

## pre-commit: Run checks before committing
.PHONY: pre-commit
pre-commit: fmt lint test
	@echo "${COLOUR_GREEN}✓ Ready to commit${COLOUR_END}"

# ==================================================================================== #
# UTILITY
# ==================================================================================== #

## deps: Download and verify dependencies
.PHONY: deps
deps:
	@echo "${COLOUR_BLUE}Downloading dependencies...${COLOUR_END}"
	@go mod download
	@go mod verify
	@echo "${COLOUR_GREEN}✓ Dependencies verified${COLOUR_END}"

## tidy: Tidy go.mod and go.sum
.PHONY: tidy
tidy:
	@echo "${COLOUR_BLUE}Tidying modules...${COLOUR_END}"
	@go mod tidy
	@echo "${COLOUR_GREEN}✓ Modules tidied${COLOUR_END}"

## version: Display version information
.PHONY: version
version:
	@echo "Version:    $(VERSION)"
	@echo "Commit:     $(COMMIT)"
	@echo "Build Time: $(BUILD_TIME)"

## ls: List all available targets
.PHONY: ls
ls:
	@echo "${COLOUR_BLUE}Available targets:${COLOUR_END}"
	@LC_ALL=C $(MAKE) -pRrq -f $(firstword $(MAKEFILE_LIST)) : 2>/dev/null | \
		awk -v RS= -F: '/(^|\n)# Files(\n|$$)/,/(^|\n)# Finished Make data base/ {if ($$1 !~ "^[#.]") {print $$1}}' | \
		sort | grep -E -v -e '^[^[:alnum:]]' -e '^$@$$'

# ==================================================================================== #
# DEFAULT
# ==================================================================================== #

# Default target
.DEFAULT_GOAL := help
