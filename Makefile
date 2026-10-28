build:
	go build -o bin/hexlet-path-size ./cmd/hexlet-path-size

lint:
	golangci-lint run

format:
	golangci-lint fmt

fix:
	golangci-lint run --fix

.PHONY: build lint format fix