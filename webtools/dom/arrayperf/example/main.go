package main

import (
	"flag"
	"fmt"

	"github.com/skeptycal/util/webtools/dom/arrayperf"
	"github.com/skeptycal/util/zsh"
	"golang.org/x/mod/semver"
)

func main() {

	// hash := zsh.Hash()
	version := zsh.Version()
	fmt.Println("version: ", version)
	if version == "" || version[:5] == `/x1b[` || !semver.IsValid(version) {
		version = "unknown"
	}
	v := semver.MajorMinor(version)

	fmt.Printf("%q", v)

	fmt.Println("version: ", version)
	fmt.Printf("ArrayPerf (%s) - array performance testing in go.\n", version)

	count := flag.Int("count", 32, "number of items in array")
	size := flag.Int("size", 16, "length of strings in array")
	flag.Parse()
	a := *arrayperf.MakeArray(*count, *size)
	// fmt.Printf("%v", a)
	fmt.Print(a.Display())
}
