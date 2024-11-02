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

func requestToJson(url string, scanner *bufio.Scanner, elementsFile string) {
	// Send HTTP GET request
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("failed to fetch page: %d %s", res.StatusCode, res.Status)
	}
	stringJsonFile := createJsonFile()
	lineNumber := 1
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		fmt.Println("Creating JSON file: ")
		log.Fatal(err)
	}
	file, err := os.Open(elementsFile)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	defer file.Close()
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Printf("%s[*]%s Current URL to scrape: %s\n", ColorBlue, ColorReset, line)
		doc.Find(line).Each(func(i int, s *goquery.Selection) {
			formatedValue := fmt.Sprintf("{test,%s}",line)

			// TODO: Try to write to json the created Json file 
			saveStringAsJSON(stringJsonFile, formatedValue) 
		})
		lineNumber++
	}
}

func createJsonFile()  string {
	currentDate := time.Now()
	formattedDate := currentDate.Format("2006-01-02")
	stringjsonFile := fmt.Sprintf("%s.json", formattedDate) 
	jsonFile, err := os.Create(stringjsonFile)
	if err != nil {
		fmt.Errorf("error opening file: %w", err)
	}
	defer jsonFile.Close()
	return stringjsonFile
}

func saveStringAsJSON(filename string, jsonString string) error {
	jsonData := []byte(jsonString)
	var temp map[string]interface{}
	if err := json.Unmarshal(jsonData, &temp); err != nil {
		return fmt.Errorf("invalid JSON string: %w", err)
	}
	return os.WriteFile(filename, jsonData, 0644)
}

func sendRequest(url string) *http.Response {
	// Send HTTP GET request
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("failed to fetch page: %d %s", res.StatusCode, res.Status)
	}
	return res
}

func requestToConsole(url string,  scanner *bufio.Scanner) {

	res := sendRequest(url)
	lineNumber := 1
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Printf("%s[*]%s Current URL to scrape: %s\n", ColorBlue, ColorReset, line)

		doc, err := goquery.NewDocumentFromReader(res.Body)
		if err != nil {
			log.Fatal(err)
		}
		doc.Find(line).Each(func(i int, s *goquery.Selection) {
			fmt.Println(s.Text())
			// TODO: create a way to minimize the amount of requests
			// on the servers by limiting the number of requests per second
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
	elemetFileOject := scanFile(elementsFile)
	scanner := bufio.NewScanner(elemetFileOject)
	if jsonFormat {
		requestToJson(url, scanner,elementsFile)
	} else { 
		requestToConsole(url, scanner)
	}
}
