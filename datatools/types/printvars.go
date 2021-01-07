// Package types provides utilities for dealing with types.
package types

import (
	"bufio"
	"encoding/json"
	"fmt"
	"html"
	"io"
	"math"
	"math/big"

	log "github.com/sirupsen/logrus"
)

const (
	preOpenTag  = "<pre>\n"
	preCloseTag = "</pre>\n"
)

type Any interface{}

func AddAny(things ...Any) Any {
	var x, y big.Float
	// nan := math.NaN()

	for _, v := range things {
		if v, ok := v.(float64); ok && math.IsNaN(v) {
			return math.NaN()
		}
		if _, _, err := x.Parse(fmt.Sprint(v), 10); err != nil {
			return nil
		}
		y.Add(&y, &x)
	}
	return y.String()
}

// ErrorCheckWriter implements a Writer/StringWriter that has
// specific error checking functionality build in. These
// additional methods ignore the normal flow of error
// checking and are therefore not exported.
type ErrorCheckWriter interface {
	io.Writer
	io.StringWriter
	// Write(p []byte) (n int, err error)
	write(p []byte)
	// WriteString(s string) (n int, err error)
	writeString(s string)
}

type errorCheckWriter struct {
	*bufio.Writer
}

// write handles error reporting and returns nothing.
// This method is explicitly not exported and it is only
// used for testing and debugging.
//
// Basically, write errors are logged but ignored.
func (e errorCheckWriter) write(p []byte) {
	n, err := e.Write(p)
	if err != nil {
		log.Error(err)
	}
	if n != len(p) {
		log.Errorf("incorrect number of bytes written: %d want: %d", n, len(p))
	}
}

// write handles error reporting and returns nothing.
// This method is explicitly not exported and it is only
// used for testing and debugging.
//
// Basically, write errors are logged but ignored.
func (e errorCheckWriter) writeString(s string) {
	n, err := e.WriteString(s)
	if err != nil {
		log.Error(err)
	}

	// todo - string length (utf-8) is not the number of bytes written...
	if n != len(s) {
		log.Errorf("incorrect number of bytes written: %d want: %d", n, len(s))
	}
}

func NewErrorCheckWriter(w io.Writer) ErrorCheckWriter {
	b, ok := w.(*errorCheckWriter)
	if ok {
		return b
	}

	bw := bufio.NewWriter(w)
	return &errorCheckWriter{bw}
}

func NewErrorCheckWriterSize(w io.Writer, size int) ErrorCheckWriter {
	wr := bufio.NewWriterSize(w, size)
	return NewErrorCheckWriter(wr)
}

// PrintVars prints ... vars
func PrintVars(wr io.Writer, writePre bool, vars ...interface{}) {

	w, ok := wr.(*errorCheckWriter)
	if !ok {
		w = &errorCheckWriter{bufio.NewWriter(wr)}
	}

	if writePre {
		w.writeString(preOpenTag)
	}

	for i, v := range vars {
		w.writeString(fmt.Sprintf(")» item %d type %T:\n", i, v))
		j, err := json.MarshalIndent(v, "", "    ")
		switch {
		case err != nil:
			w.writeString(fmt.Sprintf("error: %v", err))
		case len(j) < 3: // {}, empty struct maybe or empty string, usually mean unexported struct fields
			w.writeString(html.EscapeString(fmt.Sprintf("%+v", v)))
		default:
			w.write(j)
		}
		w.writeString("\n\n")
	}

	if writePre {
		w.writeString(preCloseTag)
	}
}
