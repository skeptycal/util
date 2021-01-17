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
	defaultioWriter io.Writer = os.Stdout
	ansi            ANSI      = NewANSIWriter(33, 44, 1, defaultioWriter)

	// Bold, Yellow Text, Blue Background
	DefaultAnsiFmt string = ansi.Build(1, 33, 44)
	// Reset Text Effects; Set Default Foreground; Default Background
	AnsiReset string = ansi.Build(0, 39, 49)
)

// todo - create a pool of stringbuilders that can go when ready?
// type sbSync struct {
// 	strings.Builder
// 	mu sync.Mutex
// }

// NewANSIWriter returns a new ANSI Writer for use in terminal output.
// If w is nil, the default (os.Stdout) is used.
func NewANSIWriter(fg, bg, ef byte, w io.Writer) ANSI {
	// if w is not a writer, use the default
	wr, ok := w.(io.Writer)
	if !ok || w == nil {
		w = defaultioWriter
	}

	return &AnsiWriter{
		*bufio.NewWriter(wr),
		DefaultAnsiFmt,
	}
}

type ANSI interface {
	io.Writer
	io.StringWriter
	Build(b ...byte) string
}

type AnsiWriter struct {
	bufio.Writer
	ansiString string // default colors and effects
}

// Build encodes a variadic list of bytes into ANSI 7 bit escape codes.
func (a *AnsiWriter) Build(b ...byte) string {
	sb := strings.Builder{}
	defer sb.Reset()
	for _, n := range b {
		sb.WriteString(fmt.Sprintf(FMT7bit, n))
	}
	return sb.String()
}

// Set accepts, encodes, and prints a variadic argument list of bytes
// that represent ANSI colors.
func (a *AnsiWriter) Set(b ...byte) (int, error) {
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
func (a AnsiOld) Build(list ...AnsiWriter) string {
	var sb strings.Builder
	defer sb.Reset()

	for _, v := range list {
		sb.WriteString(AnsiWriter(v).String())
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
