package main

import (
	"flag"
	"log"
	"scargo/getpage"
)

func main() {

	urlPtr := flag.String("url", "https://httpbin.org/get", "url of content to fetch")
	flag.Parse()

	println("url: ", *urlPtr)

	page, err := getpage.GetPage(*urlPtr)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(page)

}
