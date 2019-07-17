![Gurl](docs/logo-readme.png)


# Development

## Build

```bash
go build -o build/gurl cmd/gurl/main.go
```

## Test 

```
go test ./...
```

## Misc

```bash
# Build:
tools/build.sh

# Unit tests:
tools/test.sh

# Unit test and coverage report:
tools/coverage.sh

# Line count:
tools/cloc.sh
# Only Go:
tools/cloc.sh | grep Go | tr -s ' ' | cut -d ' ' -f 5
```

<https://jeanchristophamiel.pages.gitlab.si.francetelecom.fr/coverage.html>

# Usage

## Basic

```bash
./gurl ~/Documents/Dev/reunion/integration/hurl/generated/jdd-26-rue-des-bancouliers.hurl
```

## HTML or terminal formatter

```bash
# HTML with syntax coloring
./gurl -p html ~/Documents/Dev/reunion/integration/hurl/generated/jdd-26-rue-des-bancouliers.hurl > /tmp/x.html && open -a "Safari" /tmp/x.html

# Terminal with syntax coloring (whitespaces not visible)
./gurl -p term snippet.hurl

# Terminal with syntax coloring (whitespaces visible)
./gurl -p termws snippet.hurl

# Hurl ast exported to json
./gurl -p json snippet.hurl | jq
```

# Dev links

<https://unix.stackexchange.com/questions/105568/how-can-i-list-the-available-color-names>
