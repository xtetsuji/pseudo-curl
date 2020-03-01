# -*- makefile -*-

usage:
	@echo "Usage:"
	@echo "  make build"
	@echo "  make clean"

build-linux:
	env GOOS=linux GOARCH=amd64 go build pseudo-curl.go

clean:
	rm -f pseudo-curl
