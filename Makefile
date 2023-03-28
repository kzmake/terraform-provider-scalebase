SHELL = /bin/bash

.PHONY: all
all: pre fmt lint build install

.PHONY: pre
pre:
	go mod tidy

.PHONY: fmt
fmt:
	go run golang.org/x/tools/cmd/goimports@latest -w .

.PHONY: lint
lint:
	golangci-lint run

.PHONY: gen
gen:
	go generate ./...

.PHONY: build
build:
	go build -ldflags="-s -w" -trimpath -o terraform-provider-scalebase

.PHONY: install
install:
	go install .
