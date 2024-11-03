package main 


import (
    "os"
    "fmt"
    "bufio"

)
var ColorRed    string = "\033[31m"
var ColorGreen  string = "\033[32m"
var ColorBlue   string = "\033[34m"
var ColorReset  string = "\033[0m"

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
		scrapreCurrentSite(line,elementFile)
		lineNumber++

    }
    if err := scanner.Err(); err != nil {
        fmt.Println("Error reading file:", err)
    }
}
