package main

import (
	"flag"
	"fmt"
	"gurl/ast"
	"gurl/print"
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

	l := log.New(os.Stderr, "", 0)

	for _, file := range os.Args[1:] {

		parser, err := ast.NewParserFromFile(file)
		if err != nil {
			l.Println(err)
			os.Exit(1)
		}

		hurl := parser.Parse()
		if err := parser.Err(); err != nil {
			l.Println(err)
			os.Exit(1)
		}

		printer := print.NewTermPrinter()
		//printer := print.NewJSONPrinter()
		fmt.Print(printer.Print(hurl))
	}
}
