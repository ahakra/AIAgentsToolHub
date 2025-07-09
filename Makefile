.PHONY: build-add build-main

build-tools:
	go build -o bin/add ./tools/add/add.go
	go build -o bin/multiply ./tools/multiply/multiply.go
build-main:
	go build  -o build/agent ./cmd/api



