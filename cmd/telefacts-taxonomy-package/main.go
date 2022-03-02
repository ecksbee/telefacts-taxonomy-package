package main

import (
	"flag"
	"fmt"

	"ecksbee.com/telefacts-taxonomy-package/pkg/install"
	"ecksbee.com/telefacts-taxonomy-package/pkg/throttle"
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
	throttle.StartSECThrottle()
	_, err := install.Run(zipVar, volumeVar, throttle.Throttle)
	if err != nil {
		panic(err)
	}
}
