SHELL := /bin/bash
PATH := $(GOPATH)/bin:$(PATH)

.PHONY: install run build-linux build-darwin

install:
	go mod tidy
	go mod verify

run:
	cd ./bin && ./toggler

build-linux:
	GOOS=linux GOARCH=amd64 go build -o ./bin/toggler-linux ./cmd/toggler
	cp ./bin/toggler-linux ./bin/toggler

build-darwin:
	GOOS=darwin GOARCH=amd64 go build -o ./bin/toggler-darwin ./cmd/toggler
	cp ./bin/toggler-darwin ./bin/toggler
