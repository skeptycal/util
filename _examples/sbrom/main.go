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
