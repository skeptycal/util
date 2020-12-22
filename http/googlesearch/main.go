package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/prometheus/common/log"
	"github.com/skeptycal/util/http/http"
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

	s := GetURL(url)

	fmt.Println("\n", s)
}

func GetURL(url string) string {
	r, err := http.GetURL(url)
	if err != nil {
		return ""
	}
	defer r.Close()
	s, err := ioutil.ReadAll(r)
	if err != nil {
		return ""
	}
	return string(s)

}

func checks() error {
	url := `http://www.google.com/`
	s := readAllExample(url)
	log.Info(s)
	err := fetchExample(url)
	log.Info(err)
	return nil
}

func readAllExample(url string) string {
	data, err := http.ReadAllURL(url)
	if err != nil {
		return fmt.Sprintf("error reading response from server: %v", err)
	}
	return data
}

func fetchExample(url string) error {
	resp, err := http.Fetch(url)
	if err != nil {
		return fmt.Errorf("error obtaining response from server: %v", err)
	}
	defer resp.Body.Close()

	length := resp.ContentLength

	fmt.Printf("The content length was: %v\n", length)
	fmt.Printf("The URL Scheme was: %v\n", resp.Request.URL.Scheme)
	fmt.Println("")
	fmt.Printf("body: %s\n\n", resp.Body)
	return nil
}

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
