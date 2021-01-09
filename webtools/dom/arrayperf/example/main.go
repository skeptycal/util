package main

import (
	"flag"
	"fmt"

	"github.com/skeptycal/util/webtools/dom/arrayperf"
)

func main() {

	count := flag.Int("count", 32, "number of items in array")
	size := flag.Int("size", 16, "length of strings in array")
	flag.Parse()
	a := *arrayperf.MakeParallelArray(*count, *size)
	// fmt.Printf("%v", a)
	fmt.Print(a.Display())
}
