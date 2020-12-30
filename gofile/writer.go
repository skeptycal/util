package gofile

import (
	"fmt"
	"io"
	"os"
)

// Writer is the interface that wraps the basic Write method.
//
// Write writes len(p) bytes from p to the underlying data stream.
// It returns the number of bytes written from p (0 <= n <= len(p))
// and any error encountered that caused the write to stop early.
// Write must return a non-nil error if it returns n < len(p).
// Write must not modify the slice data, even temporarily.
//
// Implementations must not retain p.
type Writer interface {
	Write(p []byte) (n int, err error)
}

type StringWriterCloser interface {
	io.StringWriter
	io.Closer
}

// WriteString writes the contents of the string s to w, which accepts a slice of bytes.
// If w implements StringWriter, its WriteString method is invoked directly.
// Otherwise, w.Write is called exactly once.
func WriteString(w Writer, s string) (n int, err error) {
	if sw, ok := w.(io.StringWriter); ok {
		return sw.WriteString(s)
	}
	return w.Write([]byte(s))
}

// WriteFile creates the file 'fileName' and writes all 'data' to it.
// It returns any error encountered. If the file already exists, it
// will be TRUNCATED and OVERWRITTEN.
func WriteFile(fileName string, data string) (err error) {
	dataFile, err := TruncateFile(fileName)
	if DoOrDie(err) != nil {
		return
	}

	// I/O Error checking on file close
	//
	// do not defer Close() on files open for writing...
	// use a closure and named return instead ...
	// Reference: https://www.joeshaw.org/dont-defer-close-on-writable-files/
	defer func() {
		// close the file, but grab the error without
		// disturbing the err value
		cerr := dataFile.Close()
		// if there is no other error, return the value of
		// the Close() error, which is most likely, but not
		// necessarily, nil
		if err == nil {
			err = cerr
		}
	}()

	w := StringWriterCloser(dataFile)
	defer w.Close()

	n, err := dataFile.WriteString(data)
	if DoOrDie(err) != nil {
		return
	}

	if n != len(data) {
		return DoOrDie(fmt.Errorf("incorrect string length got %d want %d", n, len(data)))
	}

	// Write any buffered data to the underlying writer (standard output).
	// dataFile.Flush()

	// if err := w.Error(); err != nil {
	// 	log.Fatal(err)
	// }

	err = nil // just in case ... setup for defer closure
	return
}

// OpenTrunc creates and opens the named file for writing.
// If successful, methods on the returned file can be used for
// writing; the associated file descriptor has mode
//
//      O_WRONLY|O_CREATE|O_TRUNC
//
// If the file does not exist, it is created with mode o644.
// If the file already exists, it is TRUNCATED and overwritten.
//
// If you want other options, use:
//
//      	func os.OpenFile(name string, flag int, perm os.FileMode) (*os.File, error)
//
// If there is an error, it will be of type *PathError.
func TruncateFile(name string) (*os.File, error) {
	return os.OpenFile(name, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
}
