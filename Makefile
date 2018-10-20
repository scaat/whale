GOCMD=go
GOTEST=$(GOCMD) test
GOBUILD=$(GOCMD) build
GORUN=$(GOCMD) run
GOIMPORT=goimports -w
LISTEN_ADDRESS=LISTEN_ADDRESS=:8910
STORAGE_ROOT=STORAGE_ROOT=~/whale

all: test build

test:
	$(GOTEST) -v ./...

build:
	$(GOBUILD) -o $(BINARY_NAME) -v

run:
	mkdir -p $(STORAGE_ROOT)
	$(LISTEN_ADDRESS) $(STORAGE_ROOT) $(GORUN) server.go

import:
	$(GOIMPORT) objects/* server.go