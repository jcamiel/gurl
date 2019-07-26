#!/usr/bin/env bash
set -eux

env GOOS=darwin GOARCH=amd64 CGO_ENABLED=1 \
    go build -o build/bin/darwin/amd64/gurl cmd/gurl/main.go

env GOOS=linux GOARCH=amd64 CGO_ENABLED=1 \
    go build -o build/bin/linux/amd64/gurl cmd/gurl/main.go

#env GOOS=windows GOARCH=amd64 go build -o build/bin/windows/amd64/gurl.exe cmd/gurl/main.go
