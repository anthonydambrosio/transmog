package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/anthonydambrosio/transmog"
)

func main() {
	var file = flag.String("f", "", "file")
	var asYAML = flag.Bool("y", false, "output as YAML")
	var asXML = flag.Bool("x", false, "output as XML")
	flag.Parse()

	var transmogrifier transmog.Transmog
	if len(*file) == 0 {
		os.Exit(1)
	}

	var err error
	var data []byte
	transmogrifier.Load(*file)
	if *asYAML {
		data, err = transmogrifier.ToYaml()
	} else {
		if *asXML {
			data, err = transmogrifier.ToXML("  ")
		} else {
			data, err = transmogrifier.ToJSON()
		}
	}

	if err != nil {
		fmt.Print(fmt.Errorf("%v", err))
		os.Exit(1)
	}
	fmt.Println(string(data))
}
