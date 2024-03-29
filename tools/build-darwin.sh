#!/usr/bin/env bash

read -r version < version.txt

go build -ldflags="-s -w -X main.buildVersion=$version -X main.buildCommit=$(git describe --always --long --dirty)" \
    -o out/bin/darwin/amd64/gurl \
    cmd/gurl/main.go

go build -ldflags="-s -w -X main.buildVersion=$version -X main.buildCommit=$(git describe --always --long --dirty)" \
    -o out/bin/darwin/amd64/gurlfmt \
    cmd/gurlfmt/main.go