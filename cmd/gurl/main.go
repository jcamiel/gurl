package main

import (
	"flag"
	"fmt"
	"gurl/ast"
	"gurl/print"
	"gurl/run"
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

	for _, file := range flag.Args() {

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
		//printer := print.NewHTMLPrinter()
		fmt.Print(printer.Print(hurl))

		runner := run.NewHttpRunner()
		runner.Run(hurl)
	}
}
