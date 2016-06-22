all: compile

get-deps:
	go get ./...

vet: get-deps
	go vet -v ./...

test: vet
	go test -v -cover ./... ./_specification

compile: test
	go build -v ./...

run: all

.PHONY: all compile test vet get-deps
