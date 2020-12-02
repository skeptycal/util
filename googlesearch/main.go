package main

import (
	"fmt"
	"ideas/googlesearch/http"
	"os"
	"strings"
)

const googleSearchStringPrefix = "http://www.google.com/search?q="
const searchStringSep = "+"

func main() {
	var args []string = os.Args[1:]
	urlPrefix := os.Args[1]
	if urlPrefix == "google" {
		urlPrefix = googleSearchStringPrefix
	}
	url := http.MakeQueryURL(args, urlPrefix, searchStringSep)
	fmt.Println("The google search string is: ", url)

	// fetchExample(url)
	// readAllExample(url)
	urls := []string{url}

	tokenizerExample(urls)
	//
	// -----------------------------------------------------------
	//
}

// MakeQueryURL - join args into a google-style query string url e.g. <site><search terms>
func MakeQueryURL(args []string, SearchStringPrefix string, SearchStringSep string) string {
	// https://www.google.com/search?q=trump+biden
	return SearchStringPrefix + strings.Join(args, SearchStringSep)
}

func tokenizerExample(urls []string) {
	seedUrls := urls

	// Kick off the crawl process (concurrently)
	for _, url := range seedUrls {
		go http.Crawl(url, "a", "href")
	}
}

func readAllExample(url string) {
	data, err := http.ReadAllURL(url)
	if err != nil {
		fmt.Printf("error reading response from server: %v", err)
		return
	}
	fmt.Println("HTML:\n\n", data)
}

func fetchExample(url string) {
	resp, err := http.Fetch(url)
	defer resp.Body.Close()
	if err != nil {
		fmt.Printf("error obtaining response from server: %v", err)
		return
	}

	length := resp.ContentLength

	fmt.Printf("The content length was: %v\n", length)
	fmt.Printf("The URL Scheme was: %v\n", resp.Request.URL.Scheme)
	fmt.Println("")
	fmt.Printf("body: %s\n\n", resp.Body)
}

const htm = `<!DOCTYPE html>
<html>
<head>
    <title></title>
</head>
<body>
    body content
    <p>more content</p>
</body>
</html>`

// Plan:
// ----------------------------
// read args
// form search string
// use get request
// obtain result
// print top results
//  - color coded
//  - with links
//  -
