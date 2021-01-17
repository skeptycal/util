// Copyright (c) 2020 Michael Treanor
// MIT License

// Package ansi provides fast ansi escape sequence processing based on strings.Builder.
// The standard is defined by the ECMA-48 standard "Control Functions for Coded Character Sets - Fifth Edition"
package ansi

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
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
		sb.WriteString(fmt.Sprintf(FMT7bit, n))
	}
	return sb.String()
}

// Set accepts, encodes, and prints a variadic argument list of bytes
// that represent ANSI colors.
func (a *Ansi) Set(b ...byte) (int, error) {
	return fmt.Fprint(os.Stdout, a.Build(b...))
}

// func ansiFormat(n byte) string {
// 	return fmt.Sprintf(FMT7bit, n)
// }
// func aPrint(a ...byte) {
// 	fmt.Print(ansi.Build(a...))
// }

// --------------------------------------------------
type AnsiOld uint8

// String returns the string representation of an Ansi
// value as a color escape sequence.
func (a AnsiOld) String() string {
	return fmt.Sprintf("/x1b[%d;", a)
}

// Build returns a string containing multiple ANSI
// color escape sequences.
func (a AnsiOld) Build(list ...Ansi) string {
	var sb strings.Builder
	defer sb.Reset()

	for _, v := range list {
		sb.WriteString(Ansi(v).String())
	}

	return sb.String()
}

// itoa converts the integer value n into an ascii byte slice.
// Negative values produce an empty slice.
func itoa(n int) []byte {
	if n < 0 {
		return []byte{}
	}
	return []byte(strconv.Itoa(n))
}
