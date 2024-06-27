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
	CGO_ENABLED=0 $(GOBUILD) -a -o $(GOBIN)/$(BINARY_NAME) --ldflags '-extldflags "-static"' cmd/*.go

run: clean build
	@mkdir -p $(GOWORK)
	@$(GOBIN)/$(BINARY_NAME) -dir $(GOWORK)

container: clean build
	@echo "Building container..."
	@podman build -t nginx-harness .

run-container:
	@mkdir -p $(GOWORK)
	@podman run --replace -p 9898:9898 -v $(GOWORK):/storage --name nginx-harness nginx-harness

test:
	@echo "Testing..."
	@$(GOTEST) -v ./cmd/... ./pkg/...

clean:
	@echo "Cleaning..."
	@$(GOCLEAN)
	@rm -rf $(GOBIN)
	@rm -rf $(GOWORK)
