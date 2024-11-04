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

//  The new data you want to append
//  newData := Element{Element: "Example Domain"}
//  // Step 1: Open and read the existing JSON file
//  file, err := ioutil.ReadFile("data.json")
//  if err != nil {
//  	fmt.Println("Error reading file:", err)
//  	return
//  }
//
//  // Step 2: Parse the existing JSON data into a slice
//  var elements []Element
//  if len(file) > 0 {
//  	if err := json.Unmarshal(file, &elements); err != nil {
//  		fmt.Println("Error unmarshalling JSON:", err)
//  		return
//  	}
//  }
//
//  // Step 3: Append the new data to the slice
//  elements = append(elements, newData)
//
//  // Step 4: Write the updated slice back to the JSON file
//  updatedData, err := json.MarshalIndent(elements, "", "\t")
//  if err != nil {
//  	fmt.Println("Error marshalling JSON:", err)
//  	return
//  }
//
//  err = ioutil.WriteFile("data.json", updatedData, 0644)
//  if err != nil {
//  	fmt.Println("Error writing to file:", err)
//  	return
//  }
//
//  fmt.Println("Data successfully appended to file")

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

	fmt.Println("Data successfully appended to file")
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

func requestToConsole(url string, scanner *bufio.Scanner) {
	res, err := sendRequest(url)
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
			ColorRed, ColorReset, res.StatusCode, res.Status)
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
			fmt.Println("Data successfully appended to file")
		
		})
	}
}
