package main

import (
	"errors"
	"flag"
	"fmt"
	"gurl/ast"
	"gurl/print"
	"gurl/run"
	"log"
	"os"
)

var printFlag string

func init() {
	const (
		printDefault = ""
		printUsage = "print mode (term, termws, html, json), do not run file, "
	)
	flag.StringVar(&printFlag, "print", printDefault, printUsage)
	flag.StringVar(&printFlag, "p", printDefault, printUsage)
}


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

		if len(printFlag) > 0 {
			var p print.Printer
			switch printFlag {
			case "term":
				p = print.NewTermPrinter(false)
			case "termws":
				p = print.NewTermPrinter(true)
			case "html":
				p = print.NewHTMLPrinter()
			case "json":
				p = print.NewJSONPrinter()
			default:
				printErr(errors.New(fmt.Sprintf("unsupported print mode '%s'", printFlag)), nil)
				os.Exit(1)
			}
			fmt.Print(p.Print(hurl))
			os.Exit(0)
		}

		runner := run.NewHttpRunner()
		err = runner.Run(hurl)
		if err != nil {
			printErr(err, parser)
			os.Exit(1)
		}
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
