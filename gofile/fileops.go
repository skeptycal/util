package gofile

import (
	"fmt"
	"os"
	"path/filepath"
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

// GetFileInfo returns file information (after symlink evaluation)
// using os.Stat(). If the file does not exist, is not a regular file,
// or if the user lacks adequate permissions, an error is returned.
func GetRegularFileInfo(filename string) (os.FileInfo, error) {

	// EvalSymlinks also calls Abs and Clean as well as
	// checking for existance of the file.
	filename, err := filepath.EvalSymlinks(filename)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(filename)
	if err != nil {
		return nil, err
	}

	//Check 'others' permission
	m := fi.Mode()
	if m&(1<<2) == 0 {
		return nil, fmt.Errorf("insufficient permissions: %v", filename)
	}

	if fi.IsDir() {
		return nil, fmt.Errorf("the filename %s is a directory", filename)
	}

	if !fi.Mode().IsRegular() {
		return nil, fmt.Errorf("the filename %s is not a regular file", filename)
	}

	return fi, err
}

// chunkMultiple returns a multiple of chunk size closest to but greater than size.
func chunkMultiple(size int64) int64 {
	return (size/chunk + 1) * chunk
}

// InitialCapacity returns the multiple of 'chunk' one more than needed to
// accomodate the given capacity.
func InitialCapacity(capacity int64) int {
	if capacity < defaultBufSize {
		return defaultBufSize
	}
	return int(chunkMultiple(capacity))
}
