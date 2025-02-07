.PHONY: build run install clean

# Binary name
BINARY_NAME=terminfo-copy

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOINSTALL=$(GOCMD) install
GOCLEAN=$(GOCMD) clean

# Build the binary
build:
	$(GOBUILD) -o $(BINARY_NAME) -v

# Run the program directly
run: build
	./$(BINARY_NAME)

# Install globally
install:
	$(GOINSTALL)

# Clean build artifacts
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_NAME).exe
