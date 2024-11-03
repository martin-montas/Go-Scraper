package main

import (
	"flag"
	"fmt"
)

var ColorRed string = "\033[31m"
var ColorGreen string = "\033[32m"
var ColorBlue string = "\033[34m"
var ColorReset string = "\033[0m"

func main() {
	urlFile := flag.String("urlfile", "urls.txt", "Text file with URLs to scrape")
	elementFile := flag.String("elemfile", "elements.txt", "The elements you wnt to scrape")
	jsonFormat := flag.Bool("json", true, "True if you want to parse json. Default true")
	flag.Parse()

	run(*urlFile, *elementFile, *jsonFormat)
	fmt.Println("Done!")
}
