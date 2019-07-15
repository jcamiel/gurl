#!/usr/bin/env bash

go test -covermode=count -coverprofile=build/coverage.out ./...

go tool cover -html=build/coverage.out -o build/coverage.html

go tool cover -func=build/coverage.out | grep total: | tr -s '\t' | cut -f 3