GET https://golang.org/

HTTP/1.1 200
[Asserts]
xpath normalize-space(//h1) contains "Go is an open source programming language"
