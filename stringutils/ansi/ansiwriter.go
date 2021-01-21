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
)

// NewAnsiWriter returns a new ANSI Writer for use in terminal output.
// If w is nil, the default (os.Stdout) is used.
//
// // var     DefaultioWriter = os.Stdout
func NewAnsiWriter(w io.Writer) ANSI {
	// if w is nil, use the default
	if w == nil {
		w = os.Stdout
    }

    // if w is an AnsiWriter writer, reuse it
	if wr, ok := w.(*AnsiWriter); ok {
		return wr
	}


	return &AnsiWriter{
        *bufio.NewWriter(w),
        NewAnsiSet(StyleNormal),
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

func (a *AnsiWriter) SetColors(s AnsiSet) {
	a.ansi = s
}

func (a *AnsiWriter) String() string {
	// todo - add color from BuildAnsi(fg, bg, ef)
	return fmt.Sprintf("%sAnsiWriter: bytes written:%v, buffer available: %v/%v)", a.ansi.String(), a.Buffered(), a.Available(), a.Size())
}

// Wrap wraps the string in the default color and effects
// set in the AnsiWriter.
func (a *AnsiWriter) Wrap(s string) {
    defer	a.WriteString(DefaultAll)

	defer a.Writer.Flush()

    a.WriteString(a.ansi.String())

	a.WriteString(s)
}

// Build encodes a variadic list of bytes into ANSI codes
// and writes them to the AnsiWriter.
func (a *AnsiWriter) Build(b ...byte) {
	defer a.Writer.Flush()
	a.WriteString(BuildAnsi(b...))
}
