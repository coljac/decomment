# Makefile

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=dec
BINARY_UNIX=$(BINARY_NAME)_unix
BINARY_WINDOWS=$(BINARY_NAME)_win
BINARY_MAC=$(BINARY_NAME)_darwin
OUTPUT_DIR=bin

release: 
	cd $(OUTPUT_DIR) && xonsh ../release.xsh

all: build-all

build-all: build-linux build-linux-arm build-mac build-windows

build: 
	$(GOBUILD) -o $(BINARY_NAME) -v ./cmd/dec

test:
	$(GOTEST) -v ./...

clean:
	$(GOCLEAN)
	rm -rf $(OUTPUT_DIR)

run:
	$(GOBUILD) -o $(BINARY_NAME) -v ./cmd/dec
	./$(BINARY_NAME)

deps:
	$(GOGET) ./...

build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(OUTPUT_DIR)/linux/amd64/$(BINARY_NAME) -v ./cmd/dec

build-linux-arm:
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 $(GOBUILD) -o $(OUTPUT_DIR)/linux/arm/$(BINARY_NAME) -v ./cmd/dec

build-mac:
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 $(GOBUILD) -o $(OUTPUT_DIR)/darwin/arm64/$(BINARY_NAME) -v ./cmd/dec
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 $(GOBUILD) -o $(OUTPUT_DIR)/darwin/amd64/$(BINARY_NAME) -v ./cmd/dec

build-windows:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 $(GOBUILD) -o $(OUTPUT_DIR)/windows/amd64/$(BINARY_NAME).exe -v ./cmd/dec

.PHONY: all build test clean run deps build-linux build-mac build-windows release

