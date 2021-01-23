// Package ascii provides support for ASCII strings.
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
//
// Use of lowercase is the convention based on minimal research:
// https://twitter.com/sindresorhus/status/926521127869988864
// https://meh.com/forum/topics/capital-crimes-part-1--shout-shout-let-it-all-out
// https://en.wikipedia.org/wiki/Hexadecimal
//
// In addition:
//  echo "hexdump uses lowercase" |hexdump
//  0000000 68 65 78 64 75 6d 70 20 75 73 65 73 20 6c 6f 77
//  0000010 65 72 63 61 73 65 0a
//  0000017
func Hexed(v interface{}) string {
        switch v := v.(type) {
    case nil:
        return "0x0"
    case bool:
        if v {
            return "0x1"
        }
        return "0x0"
    case  int,int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64:
    return fmt.Sprintf("%x", v)
    default:
        return "NaN"
    }
}

func ToOctal(v interface{}) string {
    switch v := v.(type) {
    case nil:
        return "0o0"
    case bool:
        if v {
            return "0o1"
        }
        return "0o0"
    case  int,int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64:
        return fmt.Sprintf("%O", v)
    default:
        return "NaN"
    }
}

// ToBinary returns a number in binary notation (with decimal power of two exponent),
//  e.g. Binned() returns -0x1.23abcp+20
func ToBinary(v interface{}) string {
    switch v := v.(type) {
    case nil:
        return "0b0"
    case bool:
        if v {
            return "0b1"
        }
        return "0b0"
    case  int,int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64:
        return fmt.Sprintf("%b", v)
    default:
        return "NaN"
    }
}
