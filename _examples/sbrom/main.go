// Package sbrom implements funny hacky interactions between:
//  Strings, Bytes, and Runes  ... (Oh, My!)
// so ... SBROM
// anyway ... I digress
package main

import (
	"fmt"
	"strings"
)

const (
	sample           = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"
	sample2          = `bdb23dbc20e28c98`
	sampleint int32  = 0b1010101010101010101010
	ansi7fmt  string = "\033[%dm"
)

var sb = strings.Builder{}

func ansi(a ...int) string {
	defer sb.Reset()
	for i := range a {
		sb.WriteString(fmt.Sprintf(ansi7fmt, i))
	}
	return sb.String()
}

func main() {
	fmt.Println()
	fmt.Println()
	fmt.Println()

	fmt.Println(ansi(32))

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
