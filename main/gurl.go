package main

import (
	"flag"
	"fmt"
	"gurl"
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
		_, err := gurl.NewParserFromFile(file)
		if err != nil {
			log.Panic(err)
		}
	}
}
