# Development

## Build

```bash
go build -o gurl cmd/gurl/main.go
```

## Test 

```
go test ./...
```

## Count cloc

```bash
cloc --not-match-f '_test\.go$' .
```


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
./gurl -p term

# Terminal with syntax coloring (whitespaces visible)
./gurl -p termws

# Hurl ast exported to json
./gurl -p json | jq
```

# Dev links

<https://unix.stackexchange.com/questions/105568/how-can-i-list-the-available-color-names>
