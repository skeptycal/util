package gofile

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"runtime"
)

var (
	ErrHandling errorHandling = continueOnError
	emptyPath   string        = getEmptyPath()
)

// These constants are mostly modified variations or unexported parts
// of the Go standard library code (from Go 1.15.5)
const (
	minReadBufferSize        = 16
	smallBufferSize          = 64
	defaultBufSize           = 4096
	maxConsecutiveEmptyReads = 100
	maxInt                   = int(^uint(0) >> 1)
	chunk                    = bytes.MinRead
	SEP                      = string(os.PathSeparator)
)

// getEmptyPath returns a valid empty path for the current OS
/*
macOS results:
  getEmptyPath()                 .
  filepath.Clean("")            .

Windows results:
  GOOS=windows go build
  getEmptyPath()                .\
  filepath.Clean("")            .
*/
func getEmptyPath() string {
	// return filepath.Clean("") // alternative
	if runtime.GOOS == "windows" || runtime.GOOS == "plan9" {
		return ".\\"
	}
	return "."
}

// ErrorHandling wraps error handling behavior.
type ErrorHandling interface {
	String() string
	Set()
}

func (e errorHandling) String() string {
	return fmt.Sprintf("error handling: ")
}

type MultiWriter interface {
	Write([]byte) (int, error)
	WriteString(string) (int, error)
}

// NewMultiWriter returns a MultiWriter that duplicates its writes to all the
// provided writers, similar to the Unix tee(1) command.
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
func NewMultiWriter(writers ...Writer) Writer {
	if len(writers) == 0 {
		writers = append(writers, os.Stderr)
	}
	allWriters := make([]Writer, 0, len(writers))
	for _, w := range writers {
		if mw, ok := w.(*multiWriter); ok {
			allWriters = append(allWriters, mw.writers...)
		} else {
			allWriters = append(allWriters, w)
		}
	}

	// if writer is not a StringWriter, use BufferedWriter instead.
	for _, w := range writers {
		if _, ok := w.(io.StringWriter); ok {
			w = bufio.NewWriter(w)
			w = FileWriter(w)
		}
	}
	return &multiWriter{allWriters}
}

type multiWriter struct {
	writers []Writer
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

type logWriter struct {
	bufio.Writer // is a StringWriter
}

// errorHandling implements the behavior for handling errors.
type errorHandling struct {
	exitOnError bool // "default=false"
	verbose     bool
	logging     bool
	logfile     MultiWriter `default:"os.Stderr"`
}

// The readOp constants describe the last action performed on
// the buffer, so that UnreadRune and UnreadByte can check for
// invalid usage. opReadRuneX constants are chosen such that
// converted to int they correspond to the rune size that was read.
// (from Go 1.15.5 bytes/buffer.go)
type readOp int8

// Don't use iota for these, as the values need to correspond with the
// names and comments, which is easier to see when being explicit.
// (from Go 1.15.5 bytes/buffer.go)
const (
	opRead      readOp = -1 // Any other read operation.
	opInvalid   readOp = 0  // Non-read operation.
	opReadRune1 readOp = 1  // Read rune of size 1.
	opReadRune2 readOp = 2  // Read rune of size 2.
	opReadRune3 readOp = 3  // Read rune of size 3.
	opReadRune4 readOp = 4  // Read rune of size 4.
)
