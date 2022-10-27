SHELL := /bin/bash

GOFMT_FILES?=$$(find . -name '*.go' | grep -v vendor)

build: fmt run-lint test
	go build

run-lint:
	golangci-lint run

test:
	go test -v ./...

fmt:
	gofmt -w $(GOFMT_FILES)

