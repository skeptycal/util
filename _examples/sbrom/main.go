// Package sbrom implements funny hacky interactions between:
//  Strings, Bytes, and Runes  ... (Oh, My!)
// so ... SBROM
// anyway ... I digress
package main

import "github.com/skeptycal/util/stringutils/ansi"

const (
	sample          = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"
	sample2         = `bdb23dbc20e28c98`
	sampleint int32 = 0b1010101010101010101010
)

// Example code for printing Ansi color text.
// Reference:
func main() {

	ansi.Cls()
	// br()

	// aPrint(33, 44, 1)

	// hr(30)
	// br()
	// br()
	// Echo(sample)
	// Echo(sample2)
	// Echo(sampleint)

	// // fake some spacing ...
	// for i := 0; i < len(sample); i++ {
	// 	Echo("%x ", sample[i])
	// }
	// Echo()

	// // or ... just use this built-in feature
	// Echo("% x\n", sample)
	// Echo("% b\n", sampleint)
	// Echo("%X\n", sample)

	// // escape any non-printable characters
	// Echo("%q\n", sample)

	// br()
	// br()
	// aPrint(0)
	// hr(30)

	// br()

}
