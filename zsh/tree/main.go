package main

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	gorepo "github.com/skeptycal/util/devtools/gorepo"
)

func main() {
	c, err := gorepo.GetConfig()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(c.Copyright())
}
