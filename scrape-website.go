package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/PuerkitoBio/goquery"
)

type Element struct {
	Tag     string `json:"tag"`
	Element string `json:"Element"`
}

var elements []Element
func saveToJSON( currentLine string, currentTag *goquery.Selection) {
	newData := Element{
		Tag:     currentLine,
		Element: currentTag.Text(),
	}

	elements = append(elements, newData)
	updatedData, err := json.MarshalIndent(elements, "", "\t")
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}
	err = os.WriteFile(fileName, updatedData, 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
}

func sendRequest(url string) (*http.Response, error) {
	// Send HTTP GET request
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	if response.StatusCode != 200 {
		log.Fatalf("failed to fetch page: %d %s", response.StatusCode, response.Status)
	}
	return response, nil
}

func scrapeToConsole(line string) {
	response, err := http.Get(line)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	if response.StatusCode != 200 {
		log.Fatalf("%s[!]%s failed to fetch page: %d %s",
			ColorRed, ColorReset, response.StatusCode, response.Status)
	}
	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	file, err := os.Open(*elementFile)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	elementMap := make(map[string]string)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Printf("%s[*]%s Current URL to scrape: %s\n", ColorBlue, ColorReset, line)
		doc.Find(line).Each(func(i int, s *goquery.Selection) {
			elementMap[line] = s.Text()
		})
	}
	for key, value := range elementMap {
		fmt.Printf("%s[+]%s %s: %s\n", ColorGreen, ColorReset, key, value)
	}
}

func executeFileRead(scanner *bufio.Scanner, doc *goquery.Document, fileName string) {
	currentLine := scanner.Text()
	fmt.Printf("%s[*]%s Current URL to scrape: %s\n", ColorBlue, ColorReset, currentLine)
	doc.Find(currentLine).Each(func(i int, currentTag *goquery.Selection) {
		saveToJSON(currentLine, currentTag)
	})
}

func scrapeToJSON(currentLine string) {
	response, err := sendRequest(currentLine)
	if err != nil {
		log.Fatal(err)
	}

	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	file, err := os.Open(*elementFile)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	executeFileRead(scanner, doc, fileName)
}
