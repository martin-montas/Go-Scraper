# Go-Scraper

This is a small program for scraping data from websites. It uses the 
[goquery](https://github.com/PuerkitoBio/goquery) library for parsing HTML.
Made just to learn Go. it is actually a CLI tool so you should pass it arguments
on the command line to be able to use it.

## Usage
  -e string
        The element text file you want to scrape. (default "elements.txt")
  -json
        If you want to parse json. (default true)
  -u string
        Text file with URLs to scrape. (default "url.txt")

## Example
 Here is a simple example:
 ```bash
    ./ -e elements.txt -u url.txt -json true
 ```
 In this case you are passing `elements.txt` and `url.txt` as arguments.
 `elements.txt` is the elements to scrape. `url.txt` is the URLs to scrape.
 `true` means that you want to JSON as the output.

