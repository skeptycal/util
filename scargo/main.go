package main

import (
	"flag"
	scargo "ideas/scargo/http"
	"log"
)

func main() {

	urlPtr := flag.String("url", "https://httpbin.org/get", "url of content to fetch")
	flag.Parse()

	println("url: ", *urlPtr)

	page, err := scargo.GetPage(*urlPtr)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(page)

}
