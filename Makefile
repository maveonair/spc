.PHONY: build test dev

VERSION=0.1.1

build:
	CGO_ENABLED=0  go build -o ./dist/spc -a -ldflags '-s' -installsuffix cgo ./cmd/spc

test:
	go test ./...

dev:
	go run ./cmd/spc/
