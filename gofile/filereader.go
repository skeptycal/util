package gofile

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// Reader implements buffering for an io.Reader object.
type Reader struct {
	buf          []byte
	rd           io.Reader // reader provided by the client
	r, w         int       // buf read and write positions
	err          error
	lastByte     int // last byte read for UnreadByte; -1 means invalid
	lastRuneSize int // size of last rune read for UnreadRune; -1 means invalid
}

// BufferedFileReader implements a wrapper for bufio.Reader that
// calculates the initial buffer size and contains additional information
// about the underlying file.
type BufferedFileReader struct {
	r  *bufio.Reader
	f  *os.File
	fi os.FileInfo
}

// BytesFileReader represents a buffered io.Reader optimized for file reads.
type FileReader interface {
	Close() error
	Read(p []byte) (int, error)
	ReadBytes(delim byte) ([]byte, error)
	ReadString(delim byte) (string, error)
	Reset(r io.Reader)
	Open()
}

// NewBufferedReader returns a new buffered bufio.Reader whose buffer has at least the size of
// the specified file. In addition, it stores the file Stat() information to avoid redundant calls.
/*

// It is designed to be used when accessing large files where many operations will be performed and the savings of calls to os.Stat can be substantial.
//
// It is important to use defer both file.Close() and buffer.Reset() during setup to guarantee the release of resources.
Reset discards any buffered data, resets all state, and switches the buffered reader to read from r.



The method Close() performs both of these tasks, eliminating the need to add strange habits to coding workflows. e.g.
/*
   // normal process when using buffers and file handles:
	file, err := os.Open("filename")
	if err != nil {
		return err
	}
	r := bufio.NewReader(file)
	defer r.Reset(nil)
	defer file.Close()
	// ... do stuff

*/
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

	return &BufferedFileReader{
		r:  bufio.NewReaderSize(f, cap),
		f:  f,
		fi: fi,
	}, nil

}

// ReadAll reads from r until an error or EOF and returns the data it read.
// A successful call returns err == nil, not err == EOF. Because ReadAll is
// defined to read from src until EOF, it does not treat an EOF from Read
// as an error to be reported.
func (fr *BufferedFileReader) ReadAll() ([]byte, error) {
	return ioutil.ReadAll(fr.r)
}

func (fr *BufferedFileReader) Resize() error {
	if fr.r.Size() < int(fr.fi.Size()) {
		cap := InitialCapacity(fr.fi.Size())
		fr.r.Grow(cap)
	}

}

// Close closes the underlying file and frees any resources.
func (fr *BufferedFileReader) Close() error {
	defer fr.Reset()
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
