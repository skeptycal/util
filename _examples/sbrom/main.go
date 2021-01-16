// Package sbrom implements funny hacky interactions between:
//  Strings, Bytes, and Runes  ... (Oh, My!)
// so ... SBROM
// anyway ... I digress
package main

import "fmt"

const (
	sample          = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"
	sample2         = `bdb23dbc20e28c98`
	sampleint int32 = 0x1 << 16
)

func main() {

	fmt.Println(sample)

	// fake some spacing ...
	for i := 0; i < len(sample); i++ {
		fmt.Printf("%x ", sample[i])
	}
	fmt.Println()

	// or ... just use this built-in feature
	fmt.Printf("% x\n", sample)

}
