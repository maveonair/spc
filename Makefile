.PHONY: build test dev

VERSION=0.1.0

build:
	CGO_ENABLED=0  go build -mod=vendor -o ./dist/spc -a -ldflags '-s' -installsuffix cgo ./cmd/spc

test:
	go test ./...

dev:
	go run ./cmd/spc/
