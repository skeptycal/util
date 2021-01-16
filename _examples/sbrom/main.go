// Package sbrom implements funny hacky interactions between:
//  Strings, Bytes, and Runes  ... (Oh, My!)
// so ... SBROM
// anyway ... I digress
package main

import (
	"fmt"

	"github.com/skeptycal/util/stringutils/ansi"
)

const (
	sample          = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"
	sample2         = `bdb23dbc20e28c98`
	sampleint int32 = 0b1010101010101010101010
)

func green() string {
	return ansi.Ansi.Build(ansi.Bold, ansi.Yellow, ansi.GreenBackground)
}

func main() {
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Printf("%v", green())
	fmt.Println(sample)
	fmt.Println(sample2)
	fmt.Println(sampleint)

	// fake some spacing ...
	for i := 0; i < len(sample); i++ {
		fmt.Printf("%x ", sample[i])
	}
	fmt.Println()

	// or ... just use this built-in feature
	fmt.Printf("% x\n", sample)
	fmt.Printf("% b\n", sampleint)
	fmt.Printf("%X\n", sample)

	// escape any non-printable characters
	fmt.Printf("%q\n", sample)

	fmt.Println()
	fmt.Println()
	fmt.Println()

}
