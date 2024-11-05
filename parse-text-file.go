package main

import (
	"bufio"
	"fmt"
	"os"
)

func run() {
	file, err := os.Open(*urlFile)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			fmt.Printf("%s[*]%s Current URL to scrape: %s\n",
			ColorBlue, ColorReset, line)
			if *jsonFormat {
				scrapeToJSON(line)
			}
			if !*jsonFormat {
				scrapeToConsole(line)
			}
		}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}
