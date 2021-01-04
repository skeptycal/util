package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/skeptycal/util/gofile"
	"github.com/skeptycal/util/stringutils/cli"
)

func main() {
	println("pathsep: ", cli.SEP)

	if len(os.Args) != 2 {
		var name string = os.Args[0]
		fmt.Fprintf(os.Stderr, "Usage: %c URL\n", gofile.Base(name)[0])
		os.Exit(1)
	}

	response, err := http.Get(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	_, err = io.Copy(os.Stdout, response.Body)
	if err != nil {
		log.Fatal(err)
	}
}
