package main

import (
	"bufio"
	"fmt"
	"os"
)

func run(urls string, elementFile string, jsonFormat bool) {
	file, err := os.Open(urls)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	lineNumber := 1

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Printf("%s[*]%s Current URL to scrape: %s\n", ColorBlue, ColorReset, line)
		scrapreCurrentSite(line, elementFile)
		lineNumber++

	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}
