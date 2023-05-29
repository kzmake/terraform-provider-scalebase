SHELL = /bin/bash

.PHONY: all
all: pre fmt lint build

.PHONY: pre
pre:
	go mod tidy

.PHONY: fmt
fmt:
	terraform fmt -recursive
	go run golang.org/x/tools/cmd/goimports@latest -w .

.PHONY: lint
lint:
	golangci-lint run

.PHONY: gen
gen:
	go generate ./...
	patch -f < api/patch/protobuf_any.go.patch

.PHONY: build
build:
	go build -ldflags="-s -w" -trimpath -o terraform-provider-scalebase

.PHONY: install
install:
	go install -ldflags="-s -w" -trimpath .

.PHONY: test
test: test/unittest test-acctest

.PHONY: test/unittest
test/unittest:
	SCALEBASE_HOST=localhost TF_ACC=false go test github.com/kzmake/terraform-provider-scalebase/scalebase/...

.PHONY: test-acctest
test-acctest:
	SCALEBASE_HOST=${SCALEBASE_HOST} TF_ACC=true go test github.com/kzmake/terraform-provider-scalebase/scalebase/...
