# Makefile for math-skills project
# Provides convenient commands for building, testing, and running the application

# Variables - Define reusable values
BINARY_NAME=math-skills
MAIN_PATH=./cmd
GO=go

# Default target - Runs when you type 'make' without arguments
.PHONY: all
all: build

# Build target - Compiles the Go program into an executable binary
.PHONY: build
build:
	@echo "Building $(BINARY_NAME)..."
	$(GO) build -o $(BINARY_NAME) $(MAIN_PATH)

# Run target - Executes the program with the sample data file
.PHONY: run
run: build
	@echo "Running $(BINARY_NAME)..."
	./$(BINARY_NAME) data.txt

# Test target - Runs all unit tests in the project
.PHONY: test
test:
	@echo "Running tests..."
	$(GO) test ./... -v

# Clean target - Removes compiled binaries and temporary files
.PHONY: clean
clean:
	@echo "Cleaning up..."
	rm -f $(BINARY_NAME)
	$(GO) clean

# Install target - Installs the binary to $GOPATH/bin
.PHONY: install
install:
	@echo "Installing $(BINARY_NAME)..."
	$(GO) install $(MAIN_PATH)

# Format target - Formats all Go source files using gofmt
.PHONY: fmt
fmt:
	@echo "Formatting code..."
	$(GO) fmt ./...

# Vet target - Examines Go source code and reports suspicious constructs
.PHONY: vet
vet:
	@echo "Vetting code..."
	$(GO) vet ./...

# Help target - Displays available make commands
.PHONY: help
help:
	@echo "Available targets:"
	@echo "  make build   - Compile the program"
	@echo "  make run     - Build and run with data.txt"
	@echo "  make test    - Run all tests"
	@echo "  make clean   - Remove built binaries"
	@echo "  make install - Install to GOPATH/bin"
	@echo "  make fmt     - Format code"
	@echo "  make vet     - Check code for issues"
