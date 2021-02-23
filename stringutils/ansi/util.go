package ansi

import (
	"errors"
	"fmt"
	"image/color"
	"os"
	"strconv"
	"strings"
)

var errInvalidFormat = errors.New("invalid format")

func ParseHexColorFast(s string) (c color.RGBA, err error) {
    c.A = 0xff

    if s[0] != '#' {
        return c, errInvalidFormat
    }

    hexToByte := func(b byte) byte {
        switch {
        case b >= '0' && b <= '9':
            return b - '0'
        case b >= 'a' && b <= 'f':
            return b - 'a' + 10
        case b >= 'A' && b <= 'F':
            return b - 'A' + 10
        }
        err = errInvalidFormat
        return 0
    }

    switch len(s) {
    case 7:
        c.R = hexToByte(s[1])<<4 + hexToByte(s[2])
        c.G = hexToByte(s[3])<<4 + hexToByte(s[4])
        c.B = hexToByte(s[5])<<4 + hexToByte(s[6])
    case 4:
        c.R = hexToByte(s[1]) * 17
        c.G = hexToByte(s[2]) * 17
        c.B = hexToByte(s[3]) * 17
    default:
        err = errInvalidFormat
    }
    return
}

// SetupCLI clears the screen and sets the terminal
// defaults to the given AnsiSet settings.
// The returned AnsiSet may or may not be modified
// depending on the options passed.
func SetupCLI(a AnsiSet) AnsiSet{
    CLS()
    fmt.Print(a)
    fmt.Print(a.Output())
    return a
}


// APrint prints a basic ansi string based on the
// variadic argument list of bytes
func APrint(a ...ansiColor) { fmt.Print(BuildAnsi(a...)) }
func CLS()             { fmt.Fprintf(os.Stdout, "\033c") }
func HR(n int)         { fmt.Println(strings.Repeat(HrChar, n)) }
func BR()              { fmt.Println("") }

// Echo is a helper function that wraps printing to stdout
// in a default precompiled Ansi color escape sequence without
// the need to instantiate a new ANSI color object.
//
// If the first argument contains a % character, it is used as
// a format string for fmt.Printf, otherwise fmt.Print is used
// for all arguments.
//
// A newline is added in the final step when ANSI formatting
// is cleared.
//
func Echo(fmtStringMaybe interface{}, a ...interface{}) {
	fmt.Printf("%v", DefaultAnsiSet)
	defer fmt.Println(DefaultAll)

	// only first argument is given; cannot be a format string
	if len(a) == 0 {
		fmt.Print(fmtStringMaybe)
		return
	}

	// if first argument is a format string
	if fs, ok := fmtStringMaybe.(string); ok && strings.Contains(fs, "%") {
		fmt.Printf(fs, a...)
		return
	}

	// default : just print all of the arguments
	fmt.Print(append([]interface{}{fmtStringMaybe}, a...)...)

}

// itoa converts the integer value n into an ascii byte slice.
// Negative values produce an empty slice.
func itoa(n int) []byte {
	if n < 0 {
		return []byte{}
	}
	return []byte(strconv.Itoa(n))
}

// BuildAnsi returns a basic (3/4 bit) ANSI format code
// from a variadic argument list of bytes
func BuildAnsi(b ...ansiColor) string {
	sb := strings.Builder{}
	defer sb.Reset()
	for _, n := range b {
		sb.WriteString(fmt.Sprintf(FMTansi, n))
	}
	return sb.String()
}
