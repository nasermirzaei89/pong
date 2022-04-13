#@IgnoreInspection BashAddShebang
export ROOT=$(realpath $(dir $(lastword $(MAKEFILE_LIST))))
export CGO_ENABLED=1
export GO111MODULE=on

.DEFAULT_GOAL := .default

.default: format build lint test

.PHONY: help
help: ## Shows help
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.which-go:
	@which go > /dev/null || (echo "install go from https://golang.org/dl/" & exit 1)

.which-goimports:
	@which goimports > /dev/null || (echo "install goimports from https://pkg.go.dev/golang.org/x/tools/cmd/goimports" & exit 1)

.PHONY: format
format: .which-go .which-goimports ## Formats Go files
	gofmt -s -w $(ROOT)
	goimports -w .

.which-lint:
	@which golangci-lint > /dev/null || (echo "install golangci-lint from https://github.com/golangci/golangci-lint" & exit 1)

.PHONY: lint
lint: .which-lint ## Checks code with Golang CI Lint
	golangci-lint run

.PHONY: build
build: .which-go ## Builds game
	go build -v -o $(ROOT)/pong -ldflags="-s -w" $(ROOT)/*.go

.PHONY: build-windows
build-windows: .which-go ## Builds game for windows
	GOOS=windows go build -v -o $(ROOT)/pong.exe -ldflags="-s -w" $(ROOT)/*.go

.PHONY: build-wasm
build-wasm: .which-go ## Builds WASM
	GOOS=js GOARCH=wasm go build -o pong.wasm .
	cp $$(go env GOROOT)/misc/wasm/wasm_exec.js .

.PHONY: test
test: .which-go ## Tests go files
	go test -coverpkg=./... -race -coverprofile=./coverage.txt -covermode=atomic $(ROOT)/...
	go tool cover -func coverage.txt
