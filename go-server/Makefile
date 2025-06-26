# Simple Makefile for Go Server
BINARY_NAME=go-server

.PHONY: build run start clean

build:
	go build -o $(BINARY_NAME) main.go

run:
	go run main.go

start: build
	./$(BINARY_NAME)

clean:
	rm -f $(BINARY_NAME)