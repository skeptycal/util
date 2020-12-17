package gofile

import (
	"bufio"
	"bytes"
	"io"
	"io/ioutil"
	"os"
)

// BufferedFileReader implements buffering for an io.Reader object to read from a file
type BufferedFileReader struct {
	r  *bufio.Reader
	f  *os.File
	fi os.FileInfo
}

// Reader implements buffering for an io.Reader object.
// type Reader bufio.Reader

// BytesFileReader represents a buffered io.Reader optimized for file reads.
type FileReader interface {
	Close() error
	Read(p []byte) (int, error)
	ReadBytes(delim byte) ([]byte, error)
	ReadString(delim byte) (string, error)
	Reset(r io.Reader)
	Open()
}

// NewBufferedReader returns a new Reader whose buffer has at least the size of
// the specified file. If the argument io.Reader is already a Reader with large enough
// size, it returns the underlying Reader.
func NewBufferedReader(filename string) (rd *BufferedFileReader, err error) {

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

	cap := initialCapacity(fi.Size())

	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	r := bufio.NewReaderSize(f, cap)

	rd = new(BufferedFileReader)
	rd.r = r
	rd.fi = fi
	rd.f = f

	defer f.Close()

	return rd, nil
}

// ReadAll reads from r until an error or EOF and returns the data it read.
// A successful call returns err == nil, not err == EOF. Because ReadAll is
// defined to read from src until EOF, it does not treat an EOF from Read
// as an error to be reported.
func (fr *BufferedFileReader) ReadAll() ([]byte, error) {
	return ioutil.ReadAll(fr.r)
}

// Close closes the underlying file and frees any resources.
func (fr *BufferedFileReader) Close() error {
	defer fr.r.Reset(fr.f)
	return fr.f.Close()
}

// Reset discards any buffered data, resets all state, and switches
// the buffered reader to read from r.
func (fr *BufferedFileReader) Reset() {
	fr.r.Reset(fr.f)
}

func (fr *BufferedFileReader) Size() int {
	return fr.r.Size()
}

func (fr *BufferedFileReader) FileSize() int {
	return int(fr.fi.Size())
}

func (fr *BufferedFileReader) FileName() string {
	return fr.fi.Name()
}

func (fr *BufferedFileReader) FileInfo() *os.FileInfo {
	return &fr.fi
}
