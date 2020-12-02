package http

import (
	"encoding/json"
	"io"
	"log"
)

func jsonDecode(url string) json.Decoder {
	page, err := GetPage(url)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(page)

	json.NewDecoder(page)
	io.Reader()
}
