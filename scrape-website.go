package main
import (
	"bufio"
	"os"
	"encoding/json"
	"fmt"
	"log"
	// "time"
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

func scrapreCurrentSite(url string, elementsFile string) {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("%s[!]%s failed to fetch page: %d %s",
		ColorRed,ColorReset ,res.StatusCode, res.Status)
	}
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
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Printf("%s[*]%s Current URL to scrape: %s\n", ColorBlue, ColorReset, line)
		doc.Find(line).Each(func(i int, s *goquery.Selection) {
			// TODO: Create make a json file and add to it in j
			fmt.Println(s.Text())
		})
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}
