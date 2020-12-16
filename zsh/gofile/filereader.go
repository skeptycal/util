package gofile

import (
	"bufio"
	"bytes"
	"io"
	"os"
)

// This file contains the portions of code that are largely modified variations of the original standard library code (from Go 1.15.5)

const (
	minReadBufferSize        = 16
	smallBufferSize          = 64
	defaultBufSize           = 4096
	maxConsecutiveEmptyReads = 100
	maxInt                   = int(^uint(0) >> 1)
	chunk                    = bytes.MinRead
)

// BytesFileReader represents a buffered io.Reader optimized for file reads.
type BytesFileReader interface {
	Close() error
	Read(p []byte) (int, error)
	ReadBytes(delim byte) ([]byte, error)
	ReadString(delim byte) (string, error)
	Reset(r io.Reader)
	Open()
}

// FileReader defines a Reader that reads from a file
// type aFileReader struct {
// 	rd *bufio.Reader
// 	f  *os.File
// }

// FileReader implements buffering for an io.Reader object to read from a file
// buf          []byte
// rd           io.Reader // reader provided by the client
// r, w         int       // buf read and write positions
// err          error
// lastByte     int // last byte read for UnreadByte; -1 means invalid
// lastRuneSize int // size of last rune read for UnreadRune; -1       means invalid
type FileReader struct {
	*os.File
	*bufio.Reader
}

// NewFileReader returns a new Reader whose buffer has at least the size of the specified file. If the argument io.Reader is already a Reader with large enough
// size, it returns the underlying Reader.
func NewFileReader(file string) (*FileReader, error) {
	fi, err := os.Stat(file)
	if err != nil {
		return nil, err
	}

}

// Close closes the underlying file and frees any resources
func (fr *FileReader) Close() error {
	defer fr.Reset()
	return fr.Close()
}

// Reset discards any buffered data, resets all state, and switches
// the buffered reader to read from r.
func (fr *FileReader) Reset() {
	bufio.NewReaderSize(fr.Reader.Reset(), minReadBufferSize)
}

// todo - stuff

// readAll reads from r until an error or EOF and returns the data it read
// from the internal buffer allocated with a specified capacity.
func readAll(r io.Reader, capacity int64) (b []byte, err error) {
	var buf bytes.Buffer

	// If the buffer overflows, we will get bytes.ErrTooLarge.
	// Return that as an error. Any other panic remains.
	defer func() {
		e := recover()
		if e == nil {
			return
		}
		if panicErr, ok := e.(error); ok && panicErr == bytes.ErrTooLarge {
			err = panicErr
		} else {
			panic(e)
		}
	}()
	if int64(int(capacity)) == capacity {
		buf.Grow(int(capacity))
	}
	_, err = buf.ReadFrom(r)
	return buf.Bytes(), err
}

// ReadAll reads from r until an error or EOF and returns the data it read.
// A successful call returns err == nil, not err == EOF. Because ReadAll is
// defined to read from src until EOF, it does not treat an EOF from Read
// as an error to be reported.
func ReadAll(r io.Reader) ([]byte, error) {
	return readAll(r, bytes.MinRead)
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
