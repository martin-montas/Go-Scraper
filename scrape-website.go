package main
import (
	"bufio"
	"os"
	"encoding/json"
	"fmt"
	"log"
	"time"
	"net/http"
	"github.com/PuerkitoBio/goquery"
)

type elementEntry struct {
	Element  string
}


func saveJSON(fileName string, key interface{}) {
	file , err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	err = encoder.Encode(key)
	file.Close()
}

func sendRequest(url string) (*http.Response, error){
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

func requestToConsole(url string, scanner *bufio.Scanner) {
	res , err := sendRequest(url)
	if err != nil {
		log.Fatal(err)
	}
	lineNumber := 1
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Printf("%s[*]%s Current URL to scrape: %s\n", ColorBlue, ColorReset, line)
		doc, err := goquery.NewDocumentFromReader(res.Body)
		if err != nil {
			log.Fatal(err)
		}
		doc.Find(line).Each(func(i int, s *goquery.Selection) {
			fmt.Printf("fuck off!!", ColorBlue, ColorReset, line)
			fmt.Println(s.Text())
		})
		lineNumber++
	}
}

func scanFile(filename string) *os.File {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	return file
}

func scrapeWebPage(url string, elementsFile string, jsonFormat bool) {
	elementFileObject := scanFile(elementsFile)
	scanner := bufio.NewScanner(elementFileObject)
	
	// Send request and handle errors
	res, err := sendRequest(url)
	if err != nil {
		log.Fatalf("Error sending request to %s: %v", url, err)
	}
	defer res.Body.Close()  // Ensure response body is closed

	if jsonFormat {
		// Create the goquery document once
		doc, err := goquery.NewDocumentFromReader(res.Body)
		if err != nil {
			log.Fatalf("Error creating goquery document: %v", err)
		}
		lineNumber := 1
		for scanner.Scan() {
			line := scanner.Text()
			fmt.Printf("%s[*]%s Current URL to scrape: %s\n", ColorBlue, ColorReset, line)

			doc.Find(line).Each(func(i int, s *goquery.Selection) {
				// Prepare element entry
				currentElem := elementEntry{
					Element: s.Text(),
				}
				
				currentDate := time.Now()
				formattedDate := currentDate.Format("2006-01-02")
				stringJSONFile := fmt.Sprintf("%s.json", formattedDate)

				// Save to JSON
				saveJSON(stringJSONFile, currentElem)
			})
			lineNumber++
		}
		if err := scanner.Err(); err != nil {
			log.Fatalf("Error scanning elements file: %v", err)
		}
	} else { 
		requestToConsole(url, scanner)
	}
}
func previousCommit(url string, elementsFile string) {
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
