SHELL = /bin/bash

.PHONY: all
all: pre fmt lint

.PHONY: pre
pre:
	go mod tidy

.PHONY: fmt
fmt:
	go run golang.org/x/tools/cmd/goimports@latest -w .

.PHONY: lint
lint:
	golangci-lint run
