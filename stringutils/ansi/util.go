package ansi

import "fmt"

// ToAscii returns the string with non-ascii characters escaped.
// This converts any non-ASCII characters to their respective
// Unicode escape sequence (in ASCII characters).
//  e.g. ToAscii("Θ") returns "\u0398"
func ToAscii(s string) string {
	return fmt.Sprintf("%+s", s)
}

// Quoted returns a string wrapped in double quotes and with non-ASCII
// characters to their respective Unicode escape sequence (in ASCII characters).
// Any double quotes or characters with special meaning in the string are also
// escaped with slashes.
//  e.g. ToAscii("Θ") returns "\"\u0398\""
func Quoted(s string) string {
	return fmt.Sprintf("%q", s)
}
