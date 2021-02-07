.PHONY: build test dev

build:
	CGO_ENABLED=0  go build -mod=vendor -o ./dist/spc -a -ldflags '-s' -installsuffix cgo ./cmd/spc

test:
	go test ./...

dev:
	go run ./cmd/spc/
