.PHONY: build
build:
	CGO_ENABLED=0 GOOS=linux godep go build -a -installsuffix cgo
