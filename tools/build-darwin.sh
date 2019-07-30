#!/usr/bin/env bash

go build -ldflags="-s -w" -o out/bin/darwin/amd64/gurl cmd/gurl/main.go