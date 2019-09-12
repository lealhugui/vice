# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GORUN=$(GOCMD) run
OUT_DIR=./bin/
MASTER_BIN_NAME=$(OUT_DIR)master
MASTER_DIR=./cmd/master
WORKER_BIN_NAME=$(OUT_DIR)worker
WORKER_DIR=./cmd/worker

ifeq ($(OS),Windows_NT)
	MASTER_BIN_NAME += .exe
	WORKER_BIN_NAME += .exe
endif

all: clean test build
build:
	$(GOBUILD) -v -o $(MASTER_BIN_NAME) $(MASTER_DIR)
	$(GOBUILD) -v -o $(WORKER_BIN_NAME) $(WORKER_DIR)
test:
	$(GOTEST) -v ./...
clean:
	$(GOCLEAN) $(MASTER_DIR)
	$(GOCLEAN) $(WORKER_DIR)
	rm -rf $(OUT_DIR)
run-master:
	$(GORUN) $(MASTER_DIR)/main.go