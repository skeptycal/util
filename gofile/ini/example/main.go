package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/skeptycal/util/gofile/ini"
)

func main() {
	filename := "sample.ini"

	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	b, _ := ioutil.ReadAll(f)
	println(string(b))
	println("")
	println(ini.RemoveComments(filename))
}
