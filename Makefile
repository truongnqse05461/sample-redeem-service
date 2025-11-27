.PHONY: build run

build:
	go build -tags netgo -ldflags '-s -w' -o app cmd/server/main.go

run:
	go run cmd/server/main.go
