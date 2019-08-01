package main

import (
	"flag"
	"fmt"
	"gurl/ast"
	"gurl/run"
	"log"
	"os"
)

var versionFlag bool

var (
	buildVersion string = ""
	buildCommit  string = ""
)

func init() {
	const (
		versionDefault = false
		versionUsage   = "print current gurl version"
	)
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
