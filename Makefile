.PHONY: build-add build-main

build-add:
	go build -o bin/add ./tools/add.go
build-main:
	go build  -o build/agent ./cmd/api



