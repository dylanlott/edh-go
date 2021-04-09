# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=edhgo
BINARY_UNIX=$(BINARY_NAME)_unix

all: test build
build:
		$(GOBUILD) -o $(BINARY_NAME) -v
short:
		$(GOTEST) -v -short ./...
test:
		$(GOTEST) -v ./...
clean:
		$(GOCLEAN)
		rm -f $(BINARY_NAME)
		rm -f $(BINARY_UNIX)
run:
		$(GOCMD) run ./
generate:
		$(GOCMD) generate ./...
migrate:
	# TODO 
dev:
		# dev target requires watchexec to be installed
		watchexec $(GOCMD) run ./
build-heroku:
	go build -o bin/edh-go -v .
build-linux:
		CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v
build-client:
	docker build -f ./frontend/Dockerfile -t dylanlott/edh-go:client ./frontend
build-server:
	docker build -f ./Dockerfile -t dylanlott/edh-go:server .