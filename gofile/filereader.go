package gofile

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/signal"
	"syscall"
)

// Reader implements buffering for an io.Reader object.
type BufReader struct {
	buf          []byte
	rd           io.Reader // reader provided by the client
	r, w         int       // buf read and write positions
	err          error
	lastByte     int // last byte read for UnreadByte; -1 means invalid
	lastRuneSize int // size of last rune read for UnreadRune; -1 means invalid
}

// BufferedFileReader represents a buffered io.Reader optimized for file reads.
type BufferedFileReader interface {
	Close() error
	Name() string
	Open()
	Read(p []byte) (int, error)
	Reset(r io.Reader) error
	Size() int
}

// bufferedFileReader implements a wrapper for bufio.Reader that
// calculates the initial buffer size and contains additional information
// about the underlying file.
type bufferedFileReader struct {
	r  *bufio.Reader
	rd *io.ReadCloser
	f  *os.File
	fi os.FileInfo
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
setup to guarantee the release of resources. The bufferedFileReader
method Close() performs both of these tasks, eliminating the need to
add strange habits to coding workflows. e.g.

    file, err := os.Open("filename")
    if err != nil { return err }
    bfr := gofile.NewBufferedReader(file)
    defer bfr.Close()
    // ... do your stuff

    // ... and that's all ... it just works
*/
func NewBufferedReader(filename string) (rd *bufferedFileReader, err error) {

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

	return &bufferedFileReader{
		r:  bufio.NewReaderSize(f, cap),
		f:  f,
		fi: fi,
	}, nil

}

// ReadAll reads from r until an error or EOF and returns the data it read.
// A successful call returns err == nil, not err == EOF. Because ReadAll is
// defined to read from src until EOF, it does not treat an EOF from Read
// as an error to be reported.
func (fr *bufferedFileReader) ReadAll() ([]byte, error) {
	return ioutil.ReadAll(fr.r)
	b := bytes.Buffer{}
	buf := bufio.NewReader(&b)
	buf.Reset(nil)

	print(buf)
	print(b)
}

// Close closes the underlying file and frees any resources.
func (fr *bufferedFileReader) Close() error {
	defer fr.Reset(nil)
	return fr.f.Close()
}

// Reset discards any buffered data, resets all state, and switches
// the buffered reader to read from r.
func (fr *bufferedFileReader) Reset(r io.Reader) {
	if r == nil && fr.f != nil {
		r = fr.f
	}
	fr.r.Reset(fr.f)
}

func (fr *bufferedFileReader) File() (io.ReadCloser, error) {
	return fr.r, nil
}

func (fr *bufferedFileReader) Size() int {
	return fr.r.Size()
}

func (fr *bufferedFileReader) FileSize() int {
	return int(fr.fi.Size())
}

func (fr *bufferedFileReader) Name() string {
	return fr.fi.Name()
}

func (fr *bufferedFileReader) FileInfo() *os.FileInfo {
	return &fr.fi
}

// Reference: https://gobyexample.com/signals
func Sig() {
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()

	fmt.Println("awaiting signal")
	<-done
	fmt.Println("exiting")

}

// func NewFileCloseRemover(r *bufio.Reader, w *bufio.Writer) *readWriteRemover {

// 	rw := &readWriteRemover{
// 		bufio.NewReadWriter(r, w),
// 	}
// 	return rw
// }

// os.File notes:
/* type file struct {
	pfd         poll.FD
	name        string
	dirinfo     *dirInfo // nil unless directory being read
	nonblock    bool     // whether we set nonblocking mode
	stdoutOrErr bool     // whether this is stdout or stderr
	appendMode  bool     // whether file is opened for appending
}
*/
