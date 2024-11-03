package main

import (
	"flag"
	"fmt"
)

func main() {
	urlFile 	:= flag.String("urlfile", "urls.txt" ,"Text file with URLs to scrape")
	elementFile	:= flag.String("elemfile", "elements.txt" ,"The elements you wnt to scrape")
	jsonFormat 	:= flag.Bool("json", false, "True if you want to parse json. Default true")
	flag.Parse()

	run(*urlFile,*elementFile, *jsonFormat)
	fmt.Println("Done!")
}
