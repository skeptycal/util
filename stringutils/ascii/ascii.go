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
        switch v := v.(type) {
    case nil:
        return "0o0"
    case bool:
        if v {
            return "0o1"
        }
        return "0o0"
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
