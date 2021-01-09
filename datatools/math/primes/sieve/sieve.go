package sieve

import (
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"strings"

	"github.com/skeptycal/util/gofile"
	"github.com/skeptycal/zsh"
	"github.com/yourbasic/bit"
)

const defaultSieveMax = 1000


// GetFlags gets the values of the command line flags.
func GetFlags(args CLIoptions) *flag.FlagSet {

	flags := flag.NewFlagSet("sieve", flag.ExitOnError)

	flags.String("output", "out", "filename to write output to")
	flags.String("format", "", "output format (JSON | CSV | None")

	flags.Int("max", defaultSieveMax, "maximum number to search to")

	flags.Parse(os.Args[1:])

	for k, _ := range args {
		args[k] = strings.ToLower(flags.Lookup(k).Value.String())
	}

	filename := flags.Lookup("output").Value.String()
	if !zsh.Exists(filename) {
		log.Fatalf("error creating file: %v", filename)
    }

    out OutputFormat("json")=

	switch args["format"] {
	case "json", "csv", "md", "text":
		fallthrough
	default:
		log.Fatalf("format is not supported: %v", args["format"])
	}

	return flags
}

// Sieve of Eratosthenes
//
// returns a bit.Set of sieve values between 2 and maxSieve
//
// Reference: https://yourbasic.org/golang/bitmask-flag-set-clear/
func Sieve(maxSieve int) *bit.Set {

	if maxSieve <= 3 {
		maxSieve = defaultSieveMax
	}

	sieve := bit.New().AddRange(2, maxSieve)
	sqrtN := int(math.Sqrt(float64(maxSieve)))
	for p := 2; p <= sqrtN; p = sieve.Next(p) {
		for k := p * p; k < maxSieve; k += p {
			sieve.Delete(k)
		}
	}
	return sieve
}
func main() {

	args := CLIoptions{
		"filename": "",
		"format":   "",
	}

	flags := GetFlags(args)

	f, err := gofile.CreateFileTruncate(flags.Lookup("out").Value.String())
	if err != nil {
		flags.Set("output", "")
		log.Fatalf("error creating file: %v", err)
	}

	sieve := Sieve(*maxFlag)

	fmt.Println(sieve)
}
