package main

import (
	"flag"
	"fmt"
	"gurl/ast"
	"gurl/format"
	"log"
	"os"
)

func main() {

	flag.Usage = func() {
		fmt.Printf("Usage of gurl\n")
		fmt.Printf("    gurl file1 file2 ...\n")
		flag.PrintDefaults()
	}
	flag.Parse()

	for _, file := range os.Args[1:] {

		parser, err := ast.NewParserFromFile(file)
		if err != nil {
			log.Panic(err)
		}

		hurl, err := parser.Parse()
		if err != nil {
			log.Panic(err)
		}

		formatter := format.NewTermFormatter()
		fmt.Print(formatter.ToText(hurl))
	}
}
