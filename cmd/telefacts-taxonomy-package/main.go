package main

import (
	"flag"
	"fmt"

	"ecksbee.com/telefacts-taxonomy-package/pkg/install"
)

var (
	zipVar    string
	volumeVar string
)

func main() {
	flag.StringVar(&zipVar, "zip", "", "taxonomy package zip file")
	flag.StringVar(&volumeVar, "volume", "", "taxonomy package zip file")
	flag.Parse()
	fmt.Printf("%s zip\n", zipVar)
	if zipVar == "" {
		fmt.Println("-zip is empty")
		return
	}
	if volumeVar == "" {
		fmt.Println("-volume is empty")
		return
	}
	_, err := install.Run(zipVar, volumeVar)
	if err != nil {
		panic(err)
	}
}
