.PHONY: build test clean

# Project name
PROJECT_NAME := nginx-harness

# Go related variables.
GOBASE := $(shell pwd)
GOBIN := $(GOBASE)/.bin
GOWORK := $(GOBASE)/.workdir
GOCMD := go
GOBUILD := $(GOCMD) build
GOTEST := $(GOCMD) test
GOCLEAN := $(GOCMD) clean

# Binary output
BINARY_NAME := $(PROJECT_NAME)

all: test build

bin:
	@mkdir -p $(GOBIN)

build: bin
	@echo "Building..."
	@export CGO_ENABLED=0
	@$(GOBUILD) -o $(GOBIN)/$(BINARY_NAME) -ldflags -s cmd/*.go

run: clean build
	@mkdir $(GOWORK)
	@$(GOBIN)/$(BINARY_NAME) -dir $(GOWORK)

test:
	@echo "Testing..."
	@$(GOTEST) -v ./cmd/... ./pkg/...

clean:
	@echo "Cleaning..."
	@$(GOCLEAN)
	@rm -rf $(GOBIN)
	@rm -rf $(GOWORK)
