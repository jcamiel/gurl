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

	for _, file := range flag.Args() {

		parser, err := ast.NewParserFromFile(file)
		if err != nil {
			printErr(err, parser)
		}

		hurl := parser.Parse()
		if err := parser.Err(); err != nil {
			printErr(err, parser)
			os.Exit(1)
		}

		printer := print.NewTermPrinter()
		//printer := print.NewJSONPrinter()
		//printer := print.NewHTMLPrinter()
		fmt.Print(printer.Print(hurl))

		runner := run.NewHttpRunner()
		_ = runner.Run(hurl)
	}
}

func printErr(error error, p *ast.Parser) {
	l := log.New(os.Stderr, "", 0)
	switch err := error.(type) {
	/*case *ast.SyntaxError:
		l.Println(err)
		for _, e := range p.Errs() {
			l.Println(e)
		}*/
	default:
		l.Println(err)
	}
}
