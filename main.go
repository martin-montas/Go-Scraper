package main

import (
    "flag"
)


func main() {
	urlFile 	:= flag.String("urlfile", "./urls.txt" ,"Text file with URLs to scrape")
	elementFile	:= flag.String("elemfile", "./elements.txt" ,"The elements you wnt to scrape")
	jsonFormat 	:= flag.Bool("json", true, "True if you want to parse json. Default true")
	flag.Parse()

	// Scrapes the current file
	parseFile(urlFile,elementFile,jsonFormat)
}
