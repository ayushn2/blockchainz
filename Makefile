# Minimal Makefile for Go project

build:
	go build -o ./bin/blockchainz

run: build
	./bin/blockchainz

test:
	go test ./...

clean:
	rm -f ./bin/blockchainz