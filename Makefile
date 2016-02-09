.PHONY: build
build:
	CGO_ENABLED=0 GOOS=linux godep go build -a -installsuffix cgo

run:
	source "./.env" && ../../../../bin/fresh godep go run main.go
