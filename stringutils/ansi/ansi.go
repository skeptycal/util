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
	out             ANSI      = NewANSIWriter(defaultioWriter)

	// Bold, Yellow Text, Blue Background
	DefaultAnsiFmt string = BuildAnsi(Blue, 33, 1)
	// Reset Effects; Default Foreground; Default Background
	AnsiReset string = BuildAnsi(0, 39, 49)

	defaultColorSet AnsiSet = AnsiSet{DefaultForeground, DefaultBackground, Normal}
)

type color = byte

type AnsiSet struct {
	fg color
	bg color
	ef color
}

// todo - create a pool of stringbuilders that can go when ready?
// type sbSync struct {
// 	strings.Builder
// 	mu sync.Mutex
// }

// NewANSIWriter returns a new ANSI Writer for use in terminal output.
// If w is nil, the default (os.Stdout) is used.
func NewANSIWriter(w io.Writer) ANSI {
	// if w is not a writer, use the default
	if w == nil {
		w = defaultioWriter
	}
	if wr, ok := w.(*AnsiWriter); ok {
		return wr
	}

	return &AnsiWriter{
		*bufio.NewWriter(w),
		defaultColorSet,
		DefaultForeground,
		DefaultBackground,
		Normal,
	}
}

type ANSI interface {
	io.Writer
	io.StringWriter
	Build(b ...byte)
	Flush() error
	Wrap(s string)
	String() string
}

type AnsiWriter struct {
	bufio.Writer
	ansi AnsiSet
}

func (a *AnsiWriter) String() string {
	// todo - add color from BuildAnsi(fg, bg, ef)
	return fmt.Sprintf("AnsiWriter: bytes written:%v, buffer available: %v/%v)", a.Buffered(), a.Available(), a.Size())
}

// Wrap wraps the string in the default color and effects
// set in the AnsiWriter.
func (a *AnsiWriter) Wrap(s string) {
	defer a.Writer.Flush()

	a.WriteString(a.ansiString)
	a.WriteString(s)
	a.WriteString(AnsiReset)
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
