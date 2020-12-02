package file

import (
	"bytes"
	"io"
	"os"
)

// ReadFrom reads data from r until EOF and appends it to the buffer, growing
// the buffer as needed. The return value n is the number of bytes read. Any
// error except io.EOF encountered during the read is also returned. If the
// buffer becomes too large, ReadFrom will panic with ErrTooLarge.
func (b *BufferFile) ReadFrom(r io.Reader) (n int64, err error) {
	b.lastRead = opInvalid
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

// GetBuffer reads from fileName until an error or EOF and returns a
// pointer to the buffer
// allocated with a specified capacity.
// ref: modeled after Go 1.15.3 ioutil.go readAll()
func getBuffer(fileName string) (*bytes.Buffer, error) {

	var buf bytes.Buffer

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

	fi, err := os.Stat(fileName)
	if err != nil {
		return nil, err
	}
	if !fi.Mode().IsRegular() {
		return nil, os.ErrNotExist
	}

	// As initial capacity for readAll, use Size + a little extra in case Size
	// is zero, and to avoid another allocation after Read has filled the
	// buffer. The readAll call will read into its allocated internal buffer
	// cheaply. If the size was wrong, we'll either waste some space off the end
	// or reallocate as needed, but in the overwhelmingly common case we'll get
	// it just right.
	cap := initialCapacity(fi.Size())

	f, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	buff := make([]byte, 0, cap)

	buf := bytes.NewBuffer(f)

	defer buf.Reset()

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
	_, err = buf.ReadFrom()
	return &buf, err
}

func chunkMultiple(total int64, chunk int64) int64 {
	return (total/chunk + 1) * chunk
}

// initialCapacity returns the multiple of 'chunk' one more than needed
func initialCapacity(size int64) int64 {
	return (size/chunk + 1) * chunk
}

//* Notes:
// A fileStat is the implementation of FileInfo returned by Stat and Lstat.
// type fileStat struct {
// 	    name    string
// 	    size    int64
// 	    mode    os.FileMode
// 	    modTime time.Time
// 	    sys     syscall.Stat_t
// }
