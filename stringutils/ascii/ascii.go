package ascii

import "fmt"

// ToAscii returns the string with non-ascii characters escaped.
// This converts any non-ASCII characters to their respective
// Unicode escape sequence (in ASCII characters).
//  e.g. ToAscii("Θ") returns "\u0398"
func ToAscii(s string) string {
	return fmt.Sprintf("%+s", s)
}

// Quoted returns a a double-quoted string safely escaped with Go syntax. Non-ASCII characters are
// converted to Unicode escape sequences.
//  e.g. ToAscii("Θ") returns "\"\u0398\""
func Quoted(s string) string {
	return fmt.Sprintf("%q", s)
}

// Hexed returns a number in hexadecimal notation (with decimal power of two exponent),
//  e.g. Hexed() returns -0x1.23abcp+20
func Hexed(v interface{}) string {
	return fmt.Sprintf("%x", v)
}
