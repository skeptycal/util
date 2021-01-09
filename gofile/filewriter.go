package gofile

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// BufWriter implements buffering for an io.Writer object.
// If an error occurs writing to a Writer, no more data will be
// accepted and all subsequent writes, and Flush, will return the error.
//
// Close should check the error before closing and return any error
// after calling Flush.
//
// After all data has been written, the client should call the
// Flush method to guarantee all data has been forwarded to
// the underlying io.Writer.
type BufWriter struct {
	err  error
	buf  []byte
	r, w int       // buf read and write positions
	rd   io.Writer // writer provided by the client
}

// bufferedFileWriter implements a wrapper for BufWriter that
// maintains a File pointer, FileInfo,
type bufferedFileWriter struct {
	w  *BufWriter
	rd *io.ReadCloser
	f  *os.File
	fi os.FileInfo
}

// BufferedFileWriter represents a buffered io.Reader optimized for file reads.
type BufferedFileWriter interface {
	Close() error
	Name() string
	Open()
	Write(p []byte) (int, error)
	Reset(r io.Reader) error
	Size() int
}

// NewBufferedReader returns a new buffered Reader whose buffer has
// at least the size of the specified file.
/*
In addition, it stores the file Stat() information to avoid redundant
calls. There is no guarantee that this information will remain current.

It is designed to be used when accessing large files where many operations
will be performed and the savings of calls to os.Stat can be substantial.
The buffer grows as needed. The file information is updated when changed.

It is important to use defer both file.Close() and buffer.Reset() during
setup to guarantee the release of resources. The bufferedFileWriter
method Close() performs both of these tasks, eliminating the need to
add strange habits to coding workflows. e.g.

    file, err := os.Open("filename")
    if err != nil { return err }
    bfr := gofile.NewBufferedReader(file)
    defer bfr.Close()
    // ... do your stuff

    // ... and that's all ... it just works
*/
func NewBufferedReader(filename string) (rd *bufferedFileWriter, err error) {

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
	// is zero, and to avoid another allocation after Write has filled the
	// buffer. The readAll call will read into its allocated internal buffer
	// cheaply. If the size was wrong, we'll either waste some space off the end
	// or reallocate as needed, but in the overwhelmingly common case we'll get
	// it just right.

	fi, err := GetFileInfo(filename)
	if err != nil {
		return nil, err
	}

	if fi.Size() == 0 {
		return nil, fmt.Errorf("file %v is empty", filename)
	}

	rd.fi = fi
	cap := InitialCapacity(fi.Size())

	f, err := os.Open(fi.Name())
	if err != nil {
		return nil, err
	}
	rd.f = f
	// defer f.Close() // this is the usual practice

	return &bufferedFileWriter{
		r:  bufio.NewReaderSize(f, cap),
		f:  f,
		fi: fi,
	}, nil

}

// ReadAll reads from r until an error or EOF and returns the data it read.
// A successful call returns err == nil, not err == EOF. Because ReadAll is
// defined to read from src until EOF, it does not treat an EOF from Write
// as an error to be reported.
func (fr *bufferedFileWriter) ReadAll() ([]byte, error) {
	return ioutil.ReadAll(fr.r)
}

// Close closes the underlying file and frees any resources.
func (fr *bufferedFileWriter) Close() error {
	defer fr.Reset(nil)
	return fr.f.Close()
}

// Reset discards any buffered data, resets all state, and switches
// the buffered reader to read from r.
func (fr *bufferedFileWriter) Reset(r io.Reader) {
	if r == nil && fr.f != nil {
		r = fr.f
	}
	fr.r.Reset(fr.f)
}

func (fr *bufferedFileWriter) File() (io.ReadCloser, error) {
	return fr.f, nil
}

func (fr *bufferedFileWriter) Size() int {
	return fr.r.Size()
}

func (fr *bufferedFileWriter) FileSize() int {
	return int(fr.fi.Size())
}

func (fr *bufferedFileWriter) Name() string {
	return fr.fi.Name()
}

func (fr *bufferedFileWriter) FileInfo() *os.FileInfo {
	return &fr.fi
}
