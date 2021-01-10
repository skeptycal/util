package gofile

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// BufferedReader implements buffering for an io.Reader object.
type BufferedReader struct {
	bufio.ReadWriter
	f *os.File
	FileInfo
}

// BufferedFileReader represents a buffered io.Reader optimized for file reads.
type BufferedFileReader interface {
	Buffered() int // also in Writer
	Discard(n int) (discarded int, err error)
	Peek(n int) ([]byte, error)
	Read(p []byte) (n int, err error)
	ReadByte() (byte, error)
	ReadBytes(delim byte) ([]byte, error)
	ReadLine() (line []byte, isPrefix bool, err error)
	ReadRune() (r rune, size int, err error)
	ReadSlice(delim byte) (line []byte, err error)
	ReadString(delim byte) (string, error)
	Reset(r io.Reader)
	Size() int // current buffer size
	UnreadByte() error
	UnreadRune() error
	WriteTo(w io.Writer) (n int64, err error)

	// from bufio.Writer (via bufio.ReadWriter)
	Flush() error
	Available() int
	Write(p []byte) (nn int, err error)
	WriteByte(c byte) error
	WriteRune(r rune) (size int, err error)
	WriteString(s string) (int, error)
	ReadFrom(r io.Reader) (n int64, err error)

	// os.FileInfo //Size() int64 replaced with FileSize()
	Name() string       // base name of the file
	Mode() os.FileMode  // file mode bits
	ModTime() time.Time // modification time
	Sys() interface{}   // underlying data source (can return nil)

	// from bufio // implement directly with bufio ReadWriter interface
	ReadAll(r io.Reader) ([]byte, error)

	// from ioutil.go // todo benchmark these 2
	ReadAll3(r io.Reader) ([]byte, error)

	// from bytes.Buffer // todo benchmark these 2
	ReadAll2(r io.Reader) ([]byte, error)

	// from os.File
	Close() error
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
func NewBufferedReader(filename string) (rd BufferedFileReader, err error) {

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

	fi, err := GetRegularFileInfo(filename)
	if err != nil {
		return nil, err
	}

	if fi.Size() == 0 {
		return nil, fmt.Errorf("file %v is empty", filename)
	}

	cap := InitialCapacity(fi.Size())

	f, err := os.Open(fi.Name())
	if err != nil {
		return nil, err
	}
	// defer f.Close() // this is the usual practice

	return &BufferedReader{*bufio.NewReaderSize(f, cap), f, fi}, nil
}

// readAll reads from r until an error or EOF and returns the data it read
// from the internal buffer allocated with a specified capacity.
//  from Go 1.15.6 (.../src/io/ioutil/ioutil.go)
func (fr *BufferedReader) readAll(r io.Reader, capacity int64) (b []byte, err error) {
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

// ReadFrom reads data from r until EOF and appends it to the buffer, growing
// the buffer as needed. The return value n is the number of bytes read. Any
// error except io.EOF encountered during the read is also returned. If the
// buffer becomes too large, ReadFrom will panic with ErrTooLarge.
func (fr *BufferedReader) ReadFrom(r io.Reader) (n int64, err error) {
	fr.ReadWriter.lastRead = opInvalid
	for {
		i := b.grow(MinRead)
		b.buf = b.buf[:i]
		m, e := r.Read(b.buf[i:cap(b.buf)])
		if m < 0 {
			panic(errNegativeRead)
		}

		b.buf = b.buf[:i+m]
		n += int64(m)
		if e == io.EOF {
			return n, nil // e is EOF, so return nil explicitly
		}
		if e != nil {
			return n, e
		}
	}
}

// ReadFrom implements io.ReaderFrom. If the underlying writer
// supports the ReadFrom method, and b has no buffered data yet,
// this calls the underlying ReadFrom without buffering.
func (fr *BufferedReader) ReadFrom(r io.Reader) (n int64, err error) {
	if b.err != nil {
		return 0, b.err
	}
	if b.Buffered() == 0 {
		if w, ok := b.wr.(io.ReaderFrom); ok {
			n, err = w.ReadFrom(r)
			b.err = err
			return n, err
		}
	}
	var m int
	for {
		if b.Available() == 0 {
			if err1 := b.Flush(); err1 != nil {
				return n, err1
			}
		}
		nr := 0
		for nr < maxConsecutiveEmptyReads {
			m, err = r.Read(b.buf[b.n:])
			if m != 0 || err != nil {
				break
			}
			nr++
		}
		if nr == maxConsecutiveEmptyReads {
			return n, io.ErrNoProgress
		}
		b.n += m
		n += int64(m)
		if err != nil {
			break
		}
	}
	if err == io.EOF {
		// If we filled the buffer exactly, flush preemptively.
		if b.Available() == 0 {
			err = b.Flush()
		} else {
			err = nil
		}
	}
	return n, err
}

// ReadAll reads from r until an error or EOF and returns the data it read.
// A successful call returns err == nil, not err == EOF. Because ReadAll is
// defined to read from src until EOF, it does not treat an EOF from Read
// as an error to be reported.
func (fr *BufferedReader) ReadAll(r io.Reader) ([]byte, error) {
	return fr.readAll(r, bytes.MinRead)
}

// Close closes the underlying file and frees any resources.
func (fr *BufferedReader) Close() error {
	defer fr.Reset(nil)
	return fr.f.Close()
}

// Sig is an experimental feature
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
