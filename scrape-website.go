package main
import (
	"bufio"
	"os"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/PuerkitoBio/goquery"
)

func scrapeTheWebpage(url string, elementsFile string) {
	// Send HTTP GET request
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("failed to fetch page: %d %s", res.StatusCode, res.Status)
	}

	// Load HTML document with goquery
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Open(elementsFile)
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
		doc.Find(line).Each(func(i int, s *goquery.Selection) {
			// TODO: Create make a json file and add to it in j
			fmt.Println(s.Text())
		})
		lineNumber++
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}

func createAddJsonFile(data string, filename string) error {
	jsonFile, err := os.Create(filename)
	jsonData, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return fmt.Errorf("error opening file: %w", err)
	}

	defer jsonFile.Close()
	os.WriteFile(filename, jsonData, 0644)
}
