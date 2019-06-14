package main

import (
	"flag"
	"fmt"
	"gurl/parser"
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
		_, err := parser.NewParserFromFile(file)
		if err != nil {
			log.Panic(err)
		}
	}
}
