.PHONY: test run build
test:
	go vet ./...
	go test -race ./...
	node --test docs/*.test.mjs
run:
	go run .
build:
	go build -o bin/demo .
