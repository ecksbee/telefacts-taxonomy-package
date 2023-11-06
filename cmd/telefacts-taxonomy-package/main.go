package main

import (
	"flag"
	"fmt"
	"time"

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
	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()
	if zipVar == "" && volumeVar != "" {
		fmt.Println("-zip is empty")
		return
	}
	if volumeVar == "" && zipVar != "" {
		fmt.Println("-volume is empty")
		return
	}
	if zipVar != "" && volumeVar != "" {
		throttle.StartSECThrottle()
		_, err := install.Run(zipVar, volumeVar, throttle.Throttle)
		if err != nil {
			panic(err)
		}
		return
	}
	fmt.Println("no valid command parameters")
}
