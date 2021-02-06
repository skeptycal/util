package main

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/skeptycal/util/stringutils/_examples/gogen"
)

func main() {
	u, err := gogen.NewUser("", "", "", "", 0)
	if err != nil {
		log.Fatal(err)
	}
	r, err := gogen.NewRepo("fake", "", 0, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println()
	fmt.Printf("print UserConfig: %v\n", stringutils.TabIt(u.String(), 2))
	fmt.Println()
	fmt.Println(r)
	fmt.Println()
	fmt.Printf("repo name: %s\n", r.Name())
	fmt.Printf("repo license: %s\n", r.License())
	fmt.Printf("repo copyright year: %s\n", r.Year())
	fmt.Printf("repo URL: %s\n", r.URL())
	fmt.Printf("repo Download URL: %s\n", r.DownloadURL())
	fmt.Printf("repo Documentation: %s\n", r.DocURL())
}
