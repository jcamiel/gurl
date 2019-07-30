#!/usr/bin/env bash

go test -covermode=count -coverprofile=out/coverage.out ./...

go tool cover -html=out/coverage.out -o out/coverage.html

go tool cover -func=out/coverage.out | grep total: | tr -s '\t' | cut -f 3