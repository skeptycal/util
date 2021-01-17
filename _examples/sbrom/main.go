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
	"sync"
)

const (
	sample           = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"
	sample2          = `bdb23dbc20e28c98`
	sampleint int32  = 0b1010101010101010101010
	ansi7fmt  string = "\033[%vm"
	hrChar    string = "="
)

var (
	sb              strings.Builder = strings.Builder{}
	ansi            ANSI            = NewANSIWriter(33, 44, 1)
	defaultAnsiFmt  string          = ansi.Build(33, 44, 1)
	defaultioWriter io.Writer       = os.Stdout
)

type sb struct {
	strings.Builder
	mu sync.Mutex
}

// NewANSIWriter returns a new ANSI Writer for use in terminal output.
func NewANSIWriter(fg, bg, ef byte, w io.Writer) ANSI {
	if w == nil {
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
	defer sb.Reset()
	for _, n := range b {
		a.WriteString(fmt.Sprintf(ansi7fmt, n))
	}
	return sb.String()
}

// Set accepts, encodes, and prints a variadic argument list of bytes
// that represent ANSI colors.
func (a *Ansi) Set(v ...byte) (int, error) {
	return fmt.Fprint(os.Stdout, a.Build(v...))
}

// String returns the contents of the underlying strings.Builder and
// resets the buffer to nil in preparation for the next call.
func (a *Ansi) String() string {
	defer a.sb.Reset()
	fmt.Fprint(os.Stdout, a.sb.String())
	return a.sb.String()
}

// Write implements io.Writer and writes the byte slice to the underlying
// strings.Builder
func (a *Ansi) Write(b []byte) (int, error) {
	// defer fmt.Fprint(os.Stdout, a.String())
	return a.sb.Write(b)
}

// WriteString implements io.StringWriter and writes the string contents
// to the underlying strings.Builder
func (a *Ansi) WriteString(s string) (int, error) {
	// defer fmt.Fprint(os.Stdout, a.String())
	return a.sb.WriteString(s)
}

func hr(n int) {
	fmt.Println(strings.Repeat(hrChar, n))
}

func br() {
	fmt.Println("")
}

func ansiFormat(n byte) string {
	return fmt.Sprintf(ansi7fmt, n)
}
func aPrint(a ...byte) {
	fmt.Print(ansi.Build(a...))
}

func Echo(a ...interface{}) {
	fmtString := "%v\n"
	fmt.Print(defaultAnsiFmt)

	if fs, ok := a[0].(string); ok {
		if strings.Contains(fmtString, "%") {
			fmt.Printf(fs, a[1:])
		} else {
			fmt.Println(a...)
		}
	}
	aPrint(39, 49, 0)
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
