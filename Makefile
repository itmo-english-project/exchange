TARGET = cmd/exchange/main.go
OUT = bin/exchange

all: build test

.PHONY: build
build:
	go build -o $(OUT) $(TARGET)

.PHONY: run
run:
	go run $(TARGET)

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