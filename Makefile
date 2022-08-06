.PHONY: all
all: test build

.PHONY: test
test:
	go mod tidy
	go test -v -race ./...

.PHONY: build
build:
	go install github.com/crazy-max/xgo
	xgo -dest bin -out hack-chrome-data -targets darwin/amd64,darwin/arm64,windows/amd64 ./
