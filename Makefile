.PHONY: all build install uninstall clean test run help

# Binary name
BINARY_NAME=neno
INSTALL_PATH=/usr/local/bin/$(BINARY_NAME)

# Build variables
BUILD_DIR=.
GO=go
GOFLAGS=-v

# Colors
CYAN=\033[0;36m
GREEN=\033[0;32m
YELLOW=\033[1;33m
NC=\033[0m

all: build

## build: Build the binary
build:
	@echo "$(CYAN)Building $(BINARY_NAME)...$(NC)"
	@$(GO) build $(GOFLAGS) -o $(BINARY_NAME) .
	@echo "$(GREEN)✓ Build complete: ./$(BINARY_NAME)$(NC)"

## install: Build and install the binary to /usr/local/bin
install:
	@echo "$(CYAN)Installing $(BINARY_NAME)...$(NC)"
	@chmod +x install.sh
	@./install.sh

## uninstall: Remove the binary from /usr/local/bin
uninstall:
	@echo "$(CYAN)Uninstalling $(BINARY_NAME)...$(NC)"
	@chmod +x uninstall.sh
	@./uninstall.sh

## clean: Remove build artifacts
clean:
	@echo "$(CYAN)Cleaning build artifacts...$(NC)"
	@rm -f $(BINARY_NAME)
	@echo "$(GREEN)✓ Clean complete$(NC)"

## test: Run tests
test:
	@echo "$(CYAN)Running tests...$(NC)"
	@$(GO) test -v ./...

## run: Build and run the application
run: build
	@echo "$(CYAN)Running $(BINARY_NAME)...$(NC)"
	@./$(BINARY_NAME)

## fmt: Format the code
fmt:
	@echo "$(CYAN)Formatting code...$(NC)"
	@$(GO) fmt ./...
	@echo "$(GREEN)✓ Format complete$(NC)"

## vet: Run go vet
vet:
	@echo "$(CYAN)Running go vet...$(NC)"
	@$(GO) vet ./...
	@echo "$(GREEN)✓ Vet complete$(NC)"

## deps: Download dependencies
deps:
	@echo "$(CYAN)Downloading dependencies...$(NC)"
	@$(GO) mod download
	@$(GO) mod tidy
	@echo "$(GREEN)✓ Dependencies updated$(NC)"

## help: Show this help message
help:
	@echo "$(CYAN)═══════════════════════════════════════════════$(NC)"
	@echo "$(CYAN)  NENO - Makefile Commands$(NC)"
	@echo "$(CYAN)═══════════════════════════════════════════════$(NC)"
	@echo ""
	@grep -E '^## ' Makefile | sed 's/## /  $(YELLOW)/' | sed 's/:/ $(NC)-/'
	@echo ""

