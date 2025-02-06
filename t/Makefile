# Define the Go command
GO := go
MOCKERY := mockery
PKG := ./...
MOCK_OUTPUT := ./mocks

# Generate mocks using mockery
.PHONY: mocks
mocks:
	@echo "Generating mocks..."
	@$(MOCKERY) --name=PostRepoInterface --case=underscore --output=$(MOCK_OUTPUT)

# Run all tests with coverage
.PHONY: test
test:
	@echo "Running tests..."
	@$(GO) test -v $(PKG) -cover

# Run tests and display coverage in HTML
.PHONY: cover
cover:
	@echo "Generating test coverage report..."
	@$(GO) test $(PKG) -coverprofile=coverage.out
	@$(GO) tool cover -html=coverage.out -o coverage.html
	@echo "Open coverage.html in a browser to view coverage details."

# Format code
.PHONY: fmt
fmt:
	@echo "Formatting code..."
	@$(GO) fmt $(PKG)

# Tidy up Go modules
.PHONY: tidy
tidy:
	@echo "Tidying up Go modules..."
	@$(GO) mod tidy

# Install Mockery if not installed
.PHONY: install-mockery
install-mockery:
	@echo "Checking if Mockery is installed..."
	@which $(MOCKERY) || { echo "Installing mockery..."; $(GO) install github.com/vektra/mockery/v2@latest; }

# Install project dependencies
.PHONY: deps
deps:
	@echo "Installing dependencies..."
	@$(GO) mod download

# Run all necessary setup tasks
.PHONY: setup
setup: install-mockery tidy deps

# Run everything (setup, format, generate mocks, and test)
.PHONY: all
all: setup fmt mocks test