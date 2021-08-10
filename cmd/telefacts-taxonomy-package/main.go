package main

import (
	"flag"
	"fmt"

	"ecksbee.com/telefacts-taxonomy-package/pkg/install"
)

func main() {
	var zipVar string
	flag.StringVar(&zipVar, "zip", "", "taxonomy package zip file")
	flag.Parse()
	fmt.Printf("%s zip\n", zipVar)
	if zipVar == "" {
		return
	}
	err := install.Run(zipVar)
	if err != nil {
		panic(err)
	}
}
