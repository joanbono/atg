package main

import (
	"flag"

	"github.com/gocaio/atg/modules/parser"
)

var jsonfile string

func init() {
	flag.StringVar(&jsonfile, "json", "", "JSON file to use")
	flag.Parse()
}

func main() {

	if jsonfile == "" {
		onStdin := parser.CheckStdin()
		if !onStdin {
			flag.PrintDefaults()
		} else {
			parser.ParseStdin("AA")
		}
	} else {
		parser.ParseJSON(jsonfile)
	}
}
