package main

import (
	"flag"
	"fmt"
	"time"
)

var ColorRed string = "\033[31m"
var ColorGreen string = "\033[32m"
var ColorBlue string = "\033[34m"
var ColorReset string = "\033[0m"

var currentTime = time.Now()
var formattedDate = currentTime.Format("01-02-2006")
var fileName = fmt.Sprintf("%s_elements.json", formattedDate)
var urlFile *string
var elementFile *string
var jsonFormat *bool

func main() {
	urlFile = flag.String("u", "url.txt", "Text file with URLs to scrape.")
	elementFile = flag.String("e", "elements.txt", "The element text file you want to scrape.")
	jsonFormat = flag.Bool("json", false, "If you want to parse json.")
	flag.Parse()

	runProgram()
	fmt.Printf("%s[+]%s Done!\n", ColorGreen, ColorReset)
}
