package main

import (
	"flag"
	"fmt"
	"gurl/ast"
	"gurl/run"
	"log"
	_ "net/http/pprof"
	"os"
	"runtime"
	"runtime/pprof"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to `file`")
var memprofile = flag.String("memprofile", "", "write memory profile to `file`")

func main() {

	flag.Usage = func() {
		fmt.Printf("Usage of gurl\n")
		fmt.Printf("    gurl file1 file2 ...\n")
		flag.PrintDefaults()
	}
	flag.Parse()

	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		defer f.Close()
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}


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

		//printer := print.NewHTMLPrinter()
		//printer := print.NewJSONPrinter()
		//printer := print.NewHTMLPrinter()
		//fmt.Print(printer.Print(hurl))

		runner := run.NewHttpRunner()
		runner.Run(hurl)
	}

	if *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			log.Fatal("could not create memory profile: ", err)
		}
		defer f.Close()
		runtime.GC() // get up-to-date statistics
		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Fatal("could not write memory profile: ", err)
		}
	}
}
