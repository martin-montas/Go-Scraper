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


func createJsonFile(url string) *os.File {



	jsonFile, err := os.Create(fmt.Sprintf("%s.json",url))
	if err != nil {
		fmt.Errorf("error opening file: %w", err)
	}
	defer jsonFile.Close()
	return jsonFile
}



func appendToJsonFile(filename string, value string) {
  file, err := os.Open(filename)
    if err != nil {
        fmt.Errorf("error opening file: %w", err)
    }
    defer file.Close()

    decoder := json.NewDecoder(file)
    err = decoder.Decode(&value)
    if err != nil {
         fmt.Errorf("error decoding JSON data: %w", err)
    }

}


func scrapeTheWebpage(url string, elementsFile string, jsonFormat *bool) {
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

	createJsonFile(url) 
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Printf("%s[*]%s Current URL to scrape: %s\n", ColorBlue, ColorReset, line)
		doc.Find(line).Each(func(i int, s *goquery.Selection) {

			// TODO: Create make a JSON file and add to it in j
			fmt.Println(s.Text())
		})
		lineNumber++
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}

