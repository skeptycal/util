package ansi

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// APrint prints a basic ansi string based on the
// variadic argument list of bytes
func APrint(a ...byte) { fmt.Print(BuildAnsi(a...)) }
func CLS()             { fmt.Fprintf(os.Stdout, "\033c") }
func HR(n int)         { fmt.Println(strings.Repeat(HrChar, n)) }
func BR()              { fmt.Println("") }

// Echo is a helper function that wraps printing to stdout
// in a default precompiled Ansi color escape sequence without
// the need to nstantiate a new ANSI color object.
//
// If the first argument is is a string that contains a %
// character, it is used as a format string for fmt.Printf,
// otherwise fmt.Println is used for all arguments.
//
// AnsiFmt is the current text color.
//
// AnsiReset is the Ansi reset code.
//
func Echo(fmtString string, a ...interface{}) {
	fmt.Printf("%s", DefaultAnsiFmt)

	if fs, ok := a[0].(string); ok {
		if strings.Contains(fs, "%") {
			fmt.Printf(fs, a[1:])
		} else {
			fmt.Println(a...)
		}
	}
	fmt.Print(AnsiReset)
}

// itoa converts the integer value n into an ascii byte slice.
// Negative values produce an empty slice.
func itoa(n int) []byte {
	if n < 0 {
		return []byte{}
	}
	return []byte(strconv.Itoa(n))
}
