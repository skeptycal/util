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

// NewWriter returns a new ANSI Writer for use in terminal output.
// If w is nil, the default (os.Stdout) is used.
//
func NewWriter(w io.Writer) *AnsiWriter {
	// if w is nil, use the default
	if w == nil {
		w = os.Stdout
    }

    // if w is an AnsiWriter writer, reuse it
	if wr, ok := w.(*AnsiWriter); ok {
		return wr
	}


    // otherwise, create a new AnsiWriter from w
	return &AnsiWriter{
        *bufio.NewWriter(w),
        Config{},
    }
}

type ANSI interface {
    Write([]byte) (int, error)
    WriteString(string) (int, error)
    Build(b ...byte)
    SetColors(s AnsiSet)
	Flush() error
	Wrap(s string)
	String() string
}

type AnsiWriter struct {
    bufio.Writer
    config Config
}

func (a *AnsiWriter) SetColors(s AnsiSet) {
	a.config.ansi = s
}

func (a *AnsiWriter) String() string {
	// todo - add color from BuildAnsi(fg, bg, ef)
	return fmt.Sprintf("%sAnsiWriter: bytes written:%v, buffer available: %v/%v)", a.config.ansi.String(), a.Buffered(), a.Available(), a.Size())
}

// Wrap wraps the string in the default color and effects
// set in the AnsiWriter.
func (a *AnsiWriter) Wrap(s string) {
	defer a.Writer.Flush()

    a.WriteString(a.config.ansi.String())
    a.WriteString(s)
	a.WriteString(DefaultAll)
}

// Build encodes a variadic list of bytes into ANSI codes
// and writes them to the AnsiWriter.
func (a *AnsiWriter) Build(b ...byte) {
	defer a.Writer.Flush()
	a.WriteString(BuildAnsi(b...))
}
