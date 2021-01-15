// Package redlogger implements a concurrent logging system.
package redlogger

import (
	"bufio"
	"io"
	"os"
)

type LogWriter struct {
	// bufio.Writer // is a StringWriter
	MultiWriter
}

type MultiWriter interface {
	Write([]byte) (int, error)
	WriteString(string) (int, error)
}

type multiWriter struct {
	writers []io.Writer
}

// NewMultiWriter returns a MultiWriter that duplicates its writes
// to all the provided writers, similar to the Unix tee(1) command.
//
//  Write(p []byte) (n int, err error) // for bytes
//
// MultiWriter implements the StringWriter interface:
//
//  WriteString(string) (int, error) // for strings
//
// Each write is written to each listed writer, one at a time.
// If a listed writer returns an error, that overall write operation
// stops and returns the error; it does not continue down the list.
//
// ======================================================
//  Reference: go version go1.15.6  src/io/multi.go
// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
func NewMultiWriter(writers ...io.Writer) MultiWriter {
	if len(writers) == 0 {
		writers = append(writers, os.Stderr)
	}
	allWriters := make([]io.Writer, 0, len(writers))
	for _, w := range writers {
		if mw, ok := w.(*multiWriter); ok {
			allWriters = append(allWriters, mw.writers...)

		} else if _, ok := w.(io.StringWriter); !ok {
			// if writer is not a StringWriter, use BufferedWriter instead.
			allWriters = append(allWriters, bufio.NewWriter(w))
		} else {
			allWriters = append(allWriters, w)
		}
	}

	return &multiWriter{allWriters}
}

func (t *multiWriter) Write(p []byte) (n int, err error) {
	for _, w := range t.writers {
		n, err = w.Write(p)
		if err != nil {
			return
		}
		if n != len(p) {
			err = io.ErrShortWrite
			return
		}
	}
	return len(p), nil
}

func (t *multiWriter) WriteString(s string) (n int, err error) {
	var p []byte // lazily initialized if/when needed
	for _, w := range t.writers {
		if sw, ok := w.(io.StringWriter); ok {
			n, err = sw.WriteString(s)
		} else {
			if p == nil {
				p = []byte(s)
			}
			n, err = w.Write(p)
		}
		if err != nil {
			return
		}
		if n != len(s) {
			err = io.ErrShortWrite
			return
		}
	}
	return len(s), nil
}
