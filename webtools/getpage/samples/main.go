package main

import (
	"flag"
	"fmt"

	"github.com/skeptycal/util/webtools/getpage"
)

func main() {

	urlPtr := flag.String("url", "https://httpbin.org/get", "url of content to fetch")
	flag.Parse()
	// formatPtr := flag.String("format", "html", "Prettier format to use.")

	// fmt.Println("url: ", *urlPtr)

	page, err := getpage.GetPage(*urlPtr)
	if err != nil {
		fmt.Println("error retrieving url: ", err)
	}

	fmt.Println(page)
}

// sample response from test sent to http://httpbin.org
// {
// "args": {},
// "headers": {
//     "Accept-Encoding": "gzip",
//     "Host": "httpbin.org",
//     "User-Agent": "Go-http-client/2.0",
//     "X-Amzn-Trace-Id": "Root=1-5fb2dd0a-7db005a350e64f64552e1f09"
// },
// "origin": "38.75.198.14",
// "url": "https://httpbin.org/get"
// }
