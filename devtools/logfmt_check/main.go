package main

import (
	"bytes"
	"fmt"
	"io"
	"sync"
	"unicode/utf8"
)

var hex = "0123456789abcdef"

var bufferPool = sync.Pool{
	New: func() interface{} {
		return &bytes.Buffer{}
	},
}

func getBuffer() *bytes.Buffer {
	return bufferPool.Get().(*bytes.Buffer)
}

func poolBuffer(buf *bytes.Buffer) {
	buf.Reset()
	bufferPool.Put(buf)
}

type QuoteWriter interface {
	io.Writer
	Quote() string
}

type quoteBufferType struct {
	buf *bytes.Buffer
	io.Writer
}

func NewQuoteWriter(s string) QuoteWriter {
	q := new(quoteBufferType)
	return q

	// return &quoteBufferType{
	// 	s: s,
	// 	buf: getBuffer(),
	// }
}

func (q *quoteBufferType) Quote() string {
	q.buf = getBuffer() // sync.Pool entry
	q.buf.WriteRune(34)
	// the meat of it all ...

	q.buf.WriteRune(34)
	q.writeQuotedString()
}

func examplewriteQuotedString(w io.Writer, s string) (int, error) {
	// part of quoteBufferType
	// buf := getBuffer()

	// ---------> to func (q *quoteBufferType) writeQuoteChar()
	// buf.WriteByte('"')
	start := 0

	for i := 0; i < len(s); {
		if b := s[i]; b < utf8.RuneSelf {
			if 0x20 <= b && b != '\\' && b != '"' {
				i++
				continue
			}
			if start < i {
				buf.WriteString(s[start:i])
			}
			switch b {
			case '\\', '"':
				buf.WriteByte('\\')
				buf.WriteByte(b)
			case '\n':
				buf.WriteByte('\\')
				buf.WriteByte('n')
			case '\r':
				buf.WriteByte('\\')
				buf.WriteByte('r')
			case '\t':
				buf.WriteByte('\\')
				buf.WriteByte('t')
			default:
				// This encodes bytes < 0x20 except for \n, \r, and \t.
				buf.WriteString(`\u00`)
				buf.WriteByte(hex[b>>4])
				buf.WriteByte(hex[b&0xF])
			}
			i++
			start = i
			continue
		}
		c, size := utf8.DecodeRuneInString(s[i:])
		if c == utf8.RuneError {
			if start < i {
				buf.WriteString(s[start:i])
			}
			buf.WriteString(`\ufffd`)
			i += size
			start = i
			continue
		}
		i += size
	}

	if start < len(s) {
		buf.WriteString(s[start:])
	}

	// ---------> to func (q *quoteBufferType) writeQuoteChar()
	// buf.WriteByte('"')

	// ---------> to func (q *quoteBufferType) writeQuotedString() (int, error)
	// n, err := w.Write(buf.Bytes())
	// poolBuffer(buf)
	// return n, err
}

func (q *quoteBufferType) writeQuoteChar() {
	q.buf.WriteByte('"')
}

func (q *quoteBufferType) writeQuotedString() (int, error) {
	n, err := q.Write(q.buf.Bytes())
	poolBuffer(q.buf)
	return n, err
}

type byteStringBuffer struct {
	*bytes.Buffer
}

func (bs *byteStringBuffer) WriteString(s interface{}) (n int, err error) {
	b, ok := s.(*[]byte)
	if !ok {
		return 0, fmt.Errorf("string to []byte convertion failed for %v (type %T) ", s, s)
	}
	return bs.Write(*b)
}
