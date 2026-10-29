.PHONY: test run build
test:
	go vet ./...
	go test -race ./...
run:
	go run .
build:
	go build -o bin/demo .
