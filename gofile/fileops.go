package gofile

import (
	"os"
)

// ReadFrom reads data from r until EOF and appends it to the buffer, growing
// the buffer as needed. The return value n is the number of bytes read. Any
// error except io.EOF encountered during the read is also returned. If the
// buffer becomes too large, ReadFrom will panic with ErrTooLarge.
// func (b *BufferFile) ReadFrom(r io.Reader) (n int64, err error) {
// 	b.lastRead = opInvalid
// 	for {
// 		i := b.grow(MinRead)
// 		b.buf = b.buf[:i]
// 		m, e := r.Read(b.buf[i:cap(b.buf)])
// 		if m < 0 {
// 			panic(errNegativeRead)
// 		}

// 		b.buf = b.buf[:i+m]
// 		n += int64(m)
// 		if e == io.EOF {
// 			return n, nil // e is EOF, so return nil explicitly
// 		}
// 		if e != nil {
// 			return n, e
// 		}
// 	}
// }

// GetFileSize returns the size of the file using os.Stat. If an error occurred
// while reading the file, the function will return -1.
func GetFileInfo(filename string) (os.FileInfo, error) {
	fi, err := os.Stat(filename)
	if err != nil {
		return nil, err
	}
	if !fi.Mode().IsRegular() {
		return nil, os.ErrNotExist
	}
	return fi, err
}

// chunkMultiple returns a multiple of chunk size closest to but greater than size.
func chunkMultiple(size int64, chunk int64) int64 {
	return (size/chunk + 1) * chunk
}

// InitialCapacity returns the multiple of 'chunk' one more than needed to
// accomodate the given capacity.
func InitialCapacity(capacity int64) int {
	if capacity < defaultBufSize {
		return defaultBufSize
	}
	return int((capacity/chunk + 1) * chunk)
}
