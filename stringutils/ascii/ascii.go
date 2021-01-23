package ascii

import "fmt"

// ToASCII returns the string with non-ascii characters escaped.
// This converts any non-ASCII characters to their respective
// Unicode escape sequence (in ASCII characters).
//  e.g. ToAscii("Θ") returns "\u0398"
func ToASCII(s string) string {
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

// Binned returns a number in binary notation (with decimal power of two exponent),
//  e.g. Binned() returns -0x1.23abcp+20
func Binned(v interface{}) string {
    switch v := v.(type) {
    case bool:
        return fmt.Sprintf("%t", v)
    case  int,int8, int16, int32, int64, uint8, uint16, uint32, uint:
        return fmt.Sprintf("%b", v)
    case float32, float64:
        return fmt.Sprintf("%b", v)
    default:
        return fmt.Sprintf("%b", v)
    }
}
