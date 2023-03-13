build:
	go build -o img-host

build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o img-host

.PHONY: build build-linux