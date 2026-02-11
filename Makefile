TARGET = exchange
export GOPRIVATE=github.com/itmo-english-project/*
CONFIG_PATH ?= cmd/$(TARGET)/deploy/application.yaml

all: build test

.PHONY: build
build:
	go build $(GOTAGS) -o bin/$(TARGET) cmd/$(TARGET)/main.go

.PHONY: run
run:
	CONFIG_PATH=$(CONFIG_PATH) go run cmd/$(TARGET)/main.go

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: test
test:
	go test ./... -test.v

.PHONY: coverage
coverage:
	mkdir -p .cache
	go test -cover -coverprofile .cache/cover.out ./...

.PHONY: coverage-html
coverage-html: coverage
	go tool cover -html=.cache/cover.out

.PHONY: clean
clean:
	-rm -rf bin/
	-find -type d -name '.cache' -exec rm -r {} +

.PHONY: tidy
tidy:
	go mod tidy

.PHONY: lint
lint:
	golangci-lint run -v ./...

.PHONY: generate
generate:
	go generate ./...
