package gofile

import (
	"bufio"
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"syscall"
	"time"
)

// bufferedWriter implements buffering for an io.Writer object.
/*
If an error occurs writing to a Writer, no more data will be
accepted and all subsequent writes, and Flush, will return the error.

Close should check the error before closing and return any error
after calling Flush. After all data has been written, the client
should call the Flush method to guarantee all data has been forwarded
to the underlying io.Writer.

It maintains a File pointer and FileInfo interface. The file pointer
is the io.Writer supplied by the client in bufio.Writer. It is stored
to avoid repetitive system calls.

The FileInfo interface provides access to the following methods:

    Name() string       // base name of the file
    Size() int64        // length in bytes for regular files; system-dependent for others
    Mode() FileMode     // file mode bits
    ModTime() time.Time // modification time
    IsDir() bool        // abbreviation for Mode().IsDir() // not used
    Sys() interface{}   // underlying data source (can return nil) // not used
*/
type bufferedWriter struct {
	bufio.Writer
	f *os.File
	FileInfo
}

// BufferedFileWriter represents a buffered io.Reader optimized for file reads.
/*
type FileInfo interface {
    Name() string       // base name of the file
    Size() int64        // length in bytes for regular files; system-dependent for others
    Mode() FileMode     // file mode bits
    ModTime() time.Time // modification time
    IsDir() bool        // abbreviation for Mode().IsDir()
    Sys() interface{}   // underlying data source (can return nil)
}
*/
type BufferedFileWriter interface {
	Available() int
	Buffered() int
	Flush() error
	ReadFrom(r io.Reader) (n int64, err error)
	Reset(w io.Writer)
	Size() int
	Write(p []byte) (n int, err error)
	WriteByte(c byte) error
	WriteRune(r rune) (size int, err error)
	WriteString(s string) (int, error)

	// from FileInfo
	Name() string       // base name of the file
	Mode() os.FileMode  // file mode bits
	ModTime() time.Time // modification time

	// from ioutil.go
	WriteFile(filename string, data []byte, perm os.FileMode) error

	// from os.File
	Close() error
}

// NewBufferedWriter returns a new buffered Reader whose buffer has
// at least the size of the specified file.
/*
In addition, it stores the file Stat() information to avoid redundant
calls. There is no guarantee that this information will remain current.

It is designed to be used when accessing large files where many operations
will be performed and the savings of calls to os.Stat can be substantial.
The buffer grows as needed. The file information is updated when changed.

It is important to use defer both file.Close() and buffer.Reset() during
setup to guarantee the release of resources. The BufferedFileWriter
method Close() performs both of these tasks, eliminating the need to
add strange habits to coding workflows. e.g.

    file, err := os.Open("filename")
    if err != nil { return err }
    bfr := gofile.NewBufferedReader(file)
    defer bfr.Close()
    // ... do your stuff

    // ... and that's all ... it just works
*/
func NewBufferedWriter(filename string, data []byte) (w BufferedFileWriter, err error) {

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

	f, err := TruncateFile(filename)
	if err != nil {
		return nil, err
	}

	fi, err := GetRegularFileInfo(filename)
	if err != nil {
		return nil, err
	}

	cap := InitialCapacity(int64(len(data)))

	// defer f.Close() // this is the usual practice

	return &bufferedWriter{*bufio.NewWriterSize(f, cap), f, fi}, nil
}

// WriteFile writes a file.
func (fr *bufferedWriter) WriteFile(filename string, data []byte, perm os.FileMode) error {
	return ioutil.WriteFile(fr.Name(), data, 0644)
}

// Close closes the underlying File, rendering it unusable for I/O. The
// buffer is also reset to nil and its resources freed up to be garbage
// collected. This has the effect of rendering the bufferedWriter unusable.
//
// On files that support SetDeadline, any pending I/O operations will be
// canceled and return immediately with an error. Close will return an
// error if it has already been called.
func (fr bufferedWriter) Close() error {
	defer fr.Reset(nil)
	fr.Flush()
	return fr.Close()
}

// FileInfo implements os.FileInfo except for Size() and IsDir().
//
// Size() conflicts with bufio.Writer Size().
//
// IsDir() is nonsensical because the File used as the io.Writer as
// the target for the buffer in bufferedWriter is never a directory.
//
// The Mode() and Name() are implemented and necessary.
//
// The remaining methods, ModTime() and Sys(), are of dubious
// importance but there is no good reason to exclude them.
type FileInfo interface {
	Name() string       // base name of the file
	Mode() os.FileMode  // file mode bits
	ModTime() time.Time // modification time
	Sys() interface{}   // underlying data source (can return nil)
}

// FileMethods implements a selection of os.File methods that are useful
// and do not conflict with other interfaces.
//
// Methods not implemented because of conflicts are;
//  Name, Read, Write, WriteString, Truncate
// Methods not implemented because of other issues are
//  Chdir (could cause awkward problems ...)
//
// Methods implemented, but not used, are:
//  Chmod, Chown, Fd, Readdir, Readdirnames, Seek, SetDeadline,
//  SetReadDeadline, SetWriteDeadline, Sync, SyscallConn, WriteAt
type FileMethods interface {
	Chmod(mode os.FileMode) error
	Chown(uid, gid int) error
	Close() error
	Fd() uintptr
	Readdir(n int) ([]FileInfo, error)
	Readdirnames(n int) (names []string, err error)
	Seek(offset int64, whence int) (ret int64, err error)
	SetDeadline(t time.Time) error
	SetReadDeadline(t time.Time) error
	SetWriteDeadline(t time.Time) error
	Stat() (os.FileInfo, error)
	Sync() error
	SyscallConn() (syscall.RawConn, error)
	Truncate(size int64) error
	WriteAt(b []byte, off int64) (n int, err error)
}
