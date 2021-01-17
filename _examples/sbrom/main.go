// Package sbrom implements funny hacky interactions between:
//  Strings, Bytes, and Runes  ... (Oh, My!)
// so ... SBROM
// anyway ... I digress
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

const (
	sample           = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"
	sample2          = `bdb23dbc20e28c98`
	sampleint int32  = 0b1010101010101010101010
	ansi7fmt  string = "\033[%vm"
	hrChar    string = "="
)

//Ansi 7-bit color codes
const (
	Reset string = "\033[0m"

	Red    string = "\033[31m"
	Green  string = "\033[32m"
	Yellow string = "\033[33m"
	Blue   string = "\033[34m"
	Purple string = "\033[35m"
	Cyan   string = "\033[36m"
	White  string = "\033[37m"

	BgRed    string = "\033[41m"
	BgGreen  string = "\033[42m"
	BgYellow string = "\033[43m"
	BgBlue   string = "\033[44m"
	BgPurple string = "\033[45m"
	BgCyan   string = "\033[46m"
	BgWhite  string = "\033[47m"
)

var (
	ansi            ANSI      = NewANSIWriter(33, 44, 1)
	AnsiFmt         string    = ansi.Build(1, 33, 44)
	AnsiReset       string    = ansi.Build(0, 39, 49)
	defaultioWriter io.Writer = os.Stdout
)

// todo - create a pool of stringbuilders that can go when ready?
// type sbSync struct {
// 	strings.Builder
// 	mu sync.Mutex
// }

// NewANSIWriter returns a new ANSI Writer for use in terminal output.
// If w is nil, the default (os.Stdout) is used.
func NewANSIWriter(fg, bg, ef byte, w io.Writer) ANSI {
	if wr, ok := w.(io.Writer); !ok || w == nil {
		w = defaultioWriter
	}

	return &Ansi{
		fg: ansiFormat(fg),
		bg: ansiFormat(bg),
		ef: ansiFormat(ef),
		bufio.NewWriter(w),
		// sb: strings.Builder{}
	}
}

type ANSI interface {
	io.Writer
	io.StringWriter
	Build(b ...byte) string
}

type Ansi struct {
	bufio.Writer
	fg string
	bg string
	ef string
	// sb *strings.Builder
}

// Build encodes a variadic list of bytes into ANSI 7 bit escape codes.
func (a *Ansi) Build(b ...byte) string {
	sb := strings.Builder{}
	defer sb.Reset()
	for _, n := range b {
		sb.WriteString(fmt.Sprintf(ansi7fmt, n))
	}
	return sb.String()
}

// Set accepts, encodes, and prints a variadic argument list of bytes
// that represent ANSI colors.
func (a *Ansi) Set(b ...byte) (int, error) {
	return fmt.Fprint(os.Stdout, a.Build(b...))
}

func hr(n int) {
	fmt.Println(strings.Repeat(hrChar, n))
}

func br() {
	fmt.Println("")
}

// func ansiFormat(n byte) string {
// 	return fmt.Sprintf(ansi7fmt, n)
// }
// func aPrint(a ...byte) {
// 	fmt.Print(ansi.Build(a...))
// }

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

// Example code for printing Ansi color text.
// Reference:
func main() {
	br()

	// aPrint(33, 44, 1)

	hr(30)
	br()
	br()
	Echo(sample)
	Echo(sample2)
	Echo(sampleint)

	// fake some spacing ...
	for i := 0; i < len(sample); i++ {
		Echo("%x ", sample[i])
	}
	Echo()

	// or ... just use this built-in feature
	Echo("% x\n", sample)
	Echo("% b\n", sampleint)
	Echo("%X\n", sample)

	// escape any non-printable characters
	Echo("%q\n", sample)

	br()
	br()
	aPrint(0)
	hr(30)

	br()

}
