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

// ToHex returns a value (int, float, bool, nil) in hexadecimal
// notation. Any nil value is returned as zero. Rune and byte types are already uint32 and uint8, respectively, by definition.
// Other types return 'NaN'
//
// exponents are powers of two
//  e.g. ToHex() returns -0x1.23abcp+20
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
// and
//  echo "xxd uses lowercase" |xxd
//  00000000: 7878 6420 7573 6573 206c 6f77 6572 6361  xxd uses lowerca
//  00000010: 7365 0a                                  se.
func ToHex(v interface{}) string {
    if v == nil {
        return "0x0"
    }
        switch v := v.(type) {
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

// ToOctal returns a value (int, float, bool, nil) in hexadecimal
// notation. Any nil value is returned as zero. Rune and byte types are already uint32 and uint8, respectively, by definition.
// Other types return 'NaN'
func ToOctal(v interface{}) string {
    if v == nil {
        return "0o0"
    }
    switch v := v.(type) {
    case bool:
        if v {
            return "0o1"
        }
        return "0o0"
    case  int,int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
        return fmt.Sprintf("%O", v)
    case float32, float64:
        return "float"
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
