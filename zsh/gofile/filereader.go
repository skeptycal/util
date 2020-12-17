package gofile

import (
	"bufio"
	"bytes"
	"io"
	"os"
)

// FileReader implements buffering for an io.Reader object to read from a file
type FileReader struct {
	r  *bufio.Reader
	f  *os.File
	fi os.FileInfo
}

// Reader implements buffering for an io.Reader object.
type Reader bufio.Reader

// BytesFileReader represents a buffered io.Reader optimized for file reads.
// type FileReader interface {
// 	Close() error
// 	Read(p []byte) (int, error)
// 	ReadBytes(delim byte) ([]byte, error)
// 	ReadString(delim byte) (string, error)
// 	Reset(r io.Reader)
// 	Open()
// }

// NewBufferedReader returns a new Reader whose buffer has at least the size of
// the specified file. If the argument io.Reader is already a Reader with large enough
// size, it returns the underlying Reader.
func NewBufferedReader(filename string) (r *FileReader, err error) {

	// panic recover:
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

	// It's a good but not certain bet that FileInfo will tell us exactly how much to
	// read, so let's try it but be prepared for the answer to be wrong.

	// As initial capacity for readAll, use Size + a little extra in case Size
	// is zero, and to avoid another allocation after Read has filled the
	// buffer. The readAll call will read into its allocated internal buffer
	// cheaply. If the size was wrong, we'll either waste some space off the end
	// or reallocate as needed, but in the overwhelmingly common case we'll get
	// it just right.

	fi, err := GetFileInfo(filename)
	if err != nil {
		return nil, err
	}

	r.fi = *fi

	cap := initialCapacity(r.fi.Size())

	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	r = NewBufferedReaderSize(f, int(cap))

	defer f.Close()

	return r, nil
}

// NewReaderSize returns a new Reader whose buffer has at least the specified
// size. If the argument io.Reader is already a Reader with large enough
// size, it returns the underlying Reader.
func NewBufferedReaderSize(rd io.Reader, size int) *FileReader {

	r = bufio.NewReaderSize(rd, size)
	fr := &FileReader{
		r: r,
	}
}

// Close closes the underlying file and frees any resources
func (fr *FileReader) Close() error {
	defer fr.r.Reset(fr.f)
	return fr.f.Close()
}

// Reset discards any buffered data, resets all state, and switches
// the buffered reader to read from r.
func (fr *FileReader) Reset() {
	fr.r.Reset(fr.f)
}

// todo - stuff

func (fr *FileReader) Size() {
	return fr.r.Size()
}

// readAll reads from r until an error or EOF and returns the data it read
// from the internal buffer allocated with a specified capacity.
func (fr *FileReader) readAll(b []byte, err error) {

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
	if int64(int(f.r.Size())) == capacity {
		buf.Grow(int(capacity))
	}
	_, err = buf.ReadFrom(r)
	return buf.Bytes(), err

}

// ReadAll reads from r until an error or EOF and returns the data it read.
// A successful call returns err == nil, not err == EOF. Because ReadAll is
// defined to read from src until EOF, it does not treat an EOF from Read
// as an error to be reported.
func ReadAll(fr *FileReader) ([]byte, error) {
	n, err := fr.r.Read(p)

}
