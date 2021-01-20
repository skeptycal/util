// Copyright (c) 2020 Michael Treanor
// MIT License

// Package ansi provides fast ansi escape sequence processing based on strings.Builder.
// The standard is defined by the ECMA-48 standard "Control Functions for Coded Character Sets - Fifth Edition"
package ansi

import (
	"bufio"
	"fmt"
	"io"
)

// NewANSIWriter returns a new ANSI Writer for use in terminal output.
// If w is nil, the default (os.Stdout) is used.
func NewANSIWriter(w io.Writer) ANSI {
	// if w is not a writer, use the default
	if w == nil {
		w = DefaultioWriter
	}
	if wr, ok := w.(*AnsiWriter); ok {
		return wr
	}

	return &AnsiWriter{*bufio.NewWriter(w), DefaultAnsiSet()}
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
	ansi *ansiSetType
}

func (a *AnsiWriter) SetColors(s *ansiSetType) {
	a.ansi = s
}

func (a *AnsiWriter) String() string {
	// todo - add color from BuildAnsi(fg, bg, ef)
	return fmt.Sprintf("AnsiWriter: bytes written:%v, buffer available: %v/%v)", a.Buffered(), a.Available(), a.Size())
}

// Wrap wraps the string in the default color and effects
// set in the AnsiWriter.
func (a *AnsiWriter) Wrap(s string) {
	defer a.Writer.Flush()

    a.WriteString(a.ansi.String())

	a.WriteString(s)
	a.WriteString(AnsiResetString)
}

// Build encodes a variadic list of bytes into ANSI codes
// and writes them to the AnsiWriter.
func (a *AnsiWriter) Build(b ...byte) {
	defer a.Writer.Flush()
	a.WriteString(BuildAnsi(b...))
}
