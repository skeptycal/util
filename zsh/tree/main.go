package main

import (
	"fmt"

	log "github.com/sirupsen/logrus"
)

func main() {
	c, err := gorepo.GetConfig()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(c.Copyright())
}
