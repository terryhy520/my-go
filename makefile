# Makefile
.PHONY: build run test clean

build:
	go build -o myapp cmd/main.go

run:
	go run cmd/main.go

test:
	go test ./...

clean:
	rm -f myapp