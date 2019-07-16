#!/usr/bin/env bash
set -eux

GOOS=darwin
GOARCH=amd64
export $GOOS
export $GOARCH

go build -o build/bin/darwin/amd64/gurl cmd/gurl/main.go

GOOS=linux
GOARCH=amd64
export $GOOS
export $GOARCH

go build -o build/bin/linux/amd64/gurl cmd/gurl/main.go

GOOS=windows
GOARCH=amd64
export $GOOS
export $GOARCH

go build -o build/bin/windows/amd64/gurl.exe cmd/gurl/main.go
