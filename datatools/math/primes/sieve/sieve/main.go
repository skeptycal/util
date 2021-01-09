package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/skeptycal/util/datatools/math/primes/sieve"
	"github.com/skeptycal/util/gofile"
)

func main() {

	outFlag := flag.String("output", "out", "filename to write output to")
	formatFlag := flag.String("format", "", "output format (JSON | CSV | None")
	csvFlag := flag.Bool("csv", false, "return items in CSV format")
	maxFlag := flag.Int("max", defaultSieveMax, "maximum number to search to")
	flag.Parse()

	out, err := gofile.CreateFileTruncate(*outFlag)
	if err != nil {
		log.Fatalf("error creating file: %v", err)
	}

	sieve := sieve.Sieve(*maxFlag)

	fmt.Println(sieve)
}
