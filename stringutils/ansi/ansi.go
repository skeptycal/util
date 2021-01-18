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
	"strings"
)

var (
	defaultioWriter io.Writer = os.Stdout
	ansi            ANSI      = NewANSIWriter(33, 44, 1, defaultioWriter)

	// Bold, Yellow Text, Blue Background
	DefaultAnsiFmt string = BuildAnsi(44, 33, 1)
	// Reset Effects; Default Foreground; Default Background
	AnsiReset string = BuildAnsi(0, 39, 49)
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
	if w == nil {
		w = defaultioWriter
	} else {
		wr, ok := w.(io.Writer)
		if !ok {
			wr = defaultioWriter
		}
	}
	return &AnsiWriter{
		*bufio.NewWriter(wr),
		DefaultAnsiFmt,
	}
}

type ANSI interface {
	io.Writer
	io.StringWriter
	Build(b ...byte)
	Wrap(s string)
	String() string
}

type AnsiWriter struct {
	bufio.Writer
	ansiString string // default colors and effects
}

func (a *AnsiWriter) String() string {
	return fmt.Sprintf("AnsiWriter: bytes written:%v, buffer available: %v/%v)", a.Buffered(), a.Available(), a.Size())
}

// Wrap wraps the string in the default color and effects
// set in the AnsiWriter.
func (a *AnsiWriter) Wrap(s string) {
	a.WriteString(a.ansiString)
	a.WriteString(s)
	a.WriteString(AnsiReset)
	a.Writer.Flush()
}

// Build encodes a variadic list of bytes into ANSI codes
// and writes them to the AnsiWriter.
func (a *AnsiWriter) Build(b ...byte) {
	a.WriteString(BuildAnsi(b...))
	a.Flush()
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
