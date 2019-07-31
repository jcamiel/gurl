package main

import (
	"errors"
	"flag"
	"fmt"
	"gurl/ast"
	"gurl/format"
	"gurl/run"
	"log"
	"os"
)

var printFlag string
var versionFlag bool

var (
	buildVersion string = ""
	buildCommit  string = ""
)

func init() {
	const (
		printDefault   = ""
		printUsage     = "print mode (term, termws, html, json), do not run file, "
		versionDefault = false
		versionUsage   = "print current gurl version"
	)
	flag.StringVar(&printFlag, "print", printDefault, printUsage)
	flag.StringVar(&printFlag, "p", printDefault, printUsage)
	flag.BoolVar(&versionFlag, "version", versionDefault, versionUsage)
}

func main() {

	flag.Usage = func() {
		fmt.Printf("Usage of gurl\n")
		fmt.Printf("    gurl file1 file2 ...\n")
		flag.PrintDefaults()
	}
	flag.Parse()

	if versionFlag {
		fmt.Printf("%s (%s)\n", buildVersion, buildCommit)
		os.Exit(0)
	}

	for _, file := range flag.Args() {

		p, err := ast.NewParserFromFile(file)
		if err != nil {
			printErr(err, p)
		}

		hurl := p.Parse()
		if err := p.Err(); err != nil {
			printErr(err, p)
			os.Exit(1)
		}

		if len(printFlag) > 0 {
			var f format.Formatter
			switch printFlag {
			case "term":
				f = format.NewTermFormatter(false)
			case "termws":
				f = format.NewTermFormatter(true)
			case "html":
				f = format.NewHTMLFormatter()
			case "json":
				f = format.NewJSONFormatter()
			default:
				printErr(errors.New(fmt.Sprintf("unsupported print mode '%s'", printFlag)), nil)
				os.Exit(1)
			}
			fmt.Print(f.Format(hurl))
			os.Exit(0)
		}

		r := run.NewHttpRunner()
		err = r.Run(hurl)
		if err != nil {
			printErr(err, p)
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
