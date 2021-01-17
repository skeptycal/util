package ansi

import (
	"fmt"
	"os"
	"strings"
)

func Cls() {
	fmt.Fprintf(os.Stdout, "\\033c")
}

func hr(n int) {
	fmt.Println(strings.Repeat(HrChar, n))
}

func br() {
	fmt.Println("")
}

// Echo is a helper function that wraps printing to stdout
// in Ansi color escape sequences.
//
// If the first argument is is a string that contains a %
// character, it is used as a format string for fmt.Printf,
// otherwise fmt.Println is used for all arguments.
//
// AnsiFmt is the current text color.
//
// AnsiReset is the Ansi reset code.
//
func Echo(a ...interface{}) {
	fmt.Print(AnsiFmt)

	if fs, ok := a[0].(string); ok {
		if strings.Contains(fs, "%") {
			fmt.Printf(fs, a[1:])
		} else {
			fmt.Println(a...)
		}
	}
	fmt.Print(AnsiReset)
}
