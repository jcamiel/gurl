package main

import (
	"errors"
	"flag"
	"fmt"
	"gurl/ast"
	"gurl/format"
	"log"
	"os"
)

var (
	buildVersion string = ""
	buildCommit  string = ""
)

func main() {

	flag.Usage = func() {
		fmt.Println(`usage: gurlfmt {term,termws,html,json,lint} [file]

Format Hurl file.

positional arguments:
  {term,termws,html,json,lint} output format
  file                         input Hurl file
optional arguments:
  -h, --help                   show this help message and exit
  --version                    show program's version number and exit`)
		flag.PrintDefaults()
	}
	versionFlag := flag.Bool("version", false,  "print current gurlfmt version")
	flag.Parse()

	if *versionFlag {
		fmt.Printf("%s (%s)\n", buildVersion, buildCommit)
		os.Exit(0)
	}

	if flag.NArg() != 2 {
		os.Exit(1)
	}

	ftype := flag.Arg(0)
	h := flag.Arg(1)

	p, err := ast.NewParserFromFile(h)
	if err != nil {
		printErr(err)
		os.Exit(1)
	}

	hurl := p.Parse()
	if err := p.Err(); err != nil {
		printErr(err)
		os.Exit(1)
	}

	var f format.Formatter
	switch ftype {
	case "term":
		f = format.NewTermFormatter(false)
	case "termws":
		f = format.NewTermFormatter(true)
	case "html":
		f = format.NewHTMLFormatter()
	case "json":
		f = format.NewJSONFormatter()
	default:
		printErr(errors.New(fmt.Sprintf("unsupported format mode '%s'", ftype)))
		os.Exit(1)
	}
	fmt.Print(f.Format(hurl))
	os.Exit(0)
}

func printErr(err error) {
	l := log.New(os.Stderr, "", 0)
	l.Println(err)
}
