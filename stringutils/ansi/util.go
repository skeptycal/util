package ansi

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
    DefaultAll     string = "\033[39;49;0m"
	DefaultText    string = "\033[22;39m" // Normal text color and intensity
	Reset          string = "\033[0m"     // Turn off all attributes
)

var (
    DefaultAnsiSet  = NewAnsiSet(StyleNormal, White, Black,Normal)
)



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
func APrint(a ...byte) { fmt.Print(BuildAnsi(a...)) }
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
func BuildAnsi(b ...byte) string {
	sb := strings.Builder{}
	defer sb.Reset()
	for _, n := range b {
		sb.WriteString(fmt.Sprintf(FMTansi, n))
	}
	return sb.String()
}

func encodeAnsi(fb fbType, ef, c color) string {
	if fb != foreground && fb != background {
		fb = foreground
	}
	if  ef > 74 {
		ef = 0
    }

    // todo - fixme
    return ""
}
