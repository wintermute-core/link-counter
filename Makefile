SHELL := /bin/bash

GOFMT_FILES?=$$(find . -name '*.go' | grep -v vendor)

.FORCE:
build: .FORCE
	go build

fmt:
	gofmt -w $(GOFMT_FILES)

