// Package sbrom implements funny hacky interactions between:
//  Strings, Bytes, and Runes  ... (Oh, My!)
// so ... SBROM
// anyway ... I digress
package main

import (
	"bufio"
	"fmt"
	"strings"
)

const (
	sample           = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"
	sample2          = `bdb23dbc20e28c98`
	sampleint int32  = 0b1010101010101010101010
	ansi7fmt  string = "\033[%dm"
	hrChar    string = "="
)

var (
	defaultAnsiFmt string = a.Build(33, 44, 1)
)

func NewANSIWriter(fg, bg, ef []byte) ANSI {
	w := bufio.NewWriter(w)
	return &Ansi{
		fg: fg,
		bg: bg,
		ef: ef,
		w:  bufio.NewWriter(w),
		sb: &strings.Builder{},
	}
}

type ANSI interface {
	Build([]byte) string
	Write([]byte) (int, error)
	WriteString(string) (int, error)
	String() string
}

type Ansi struct {
	fg []byte
	bg []byte
	ef []byte
	w  *bufio.Writer
	sb *strings.Builder
}

func (a *Ansi) Build(b []byte) string {
	defer a.sb.Reset()
	for i, n := range b {
		a.sb.WriteString(fmt.Sprintf(ansi7fmt, n))
	}
	return a.sb.String()
}

func (a *Ansi) String() string {
	defer a.sb.Reset()
	return a.sb.String()
}

func (a *Ansi) Write(b []byte) (int, error) {
	return a.w.Write(b)
}

func (a *Ansi) WriteString(s string) (int, error) {
	return a.w.WriteString(s)
}

func hr(n int) {
	fmt.Println(strings.Repeat(hrChar, n))
}

func br() {
	fmt.Println("")
}

func aPrint(a ...int) {
	fmt.Print(a.Build(a...))
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
