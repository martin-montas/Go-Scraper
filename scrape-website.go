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
	Element string `json:"Element"`
}

func saveJSON(element Element) {
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	err = encoder.Encode(element)
	var elements []Element
	elements = append(elements, element)
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
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("failed to fetch page: %d %s", res.StatusCode, res.Status)
	}
	return res, nil
}

func scrapeToConsole(line string) {
	res, err := http.Get(line)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("%s[!]%s failed to fetch page: %d %s",
			ColorRed, ColorReset, res.StatusCode, res.Status)
	}
	doc, err := goquery.NewDocumentFromReader(res.Body)
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

func scrapeToJSON(line string) {
	res, err := http.Get(line)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("%s[!]%s failed to fetch page: %d %s",
			ColorRed, ColorReset, res.StatusCode, res.Status)
	}
	doc, err := goquery.NewDocumentFromReader(res.Body)
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
	var elements []Element
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Printf("%s[*]%s Current URL to scrape: %s\n", ColorBlue, ColorReset, line)
		doc.Find(line).Each(func(i int, s *goquery.Selection) {
			newData := Element{Element: s.Text()}
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
		})
	}
}
