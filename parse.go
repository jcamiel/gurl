package gurl

import (
	"fmt"
	"io/ioutil"
)

func ParseFile(file string) error {

	fmt.Printf("Parsing %s\n", file)

	dat, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	return ParseString(string(dat))
}

func ParseString(text string) error {
	return nil
}
