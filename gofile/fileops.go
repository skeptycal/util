package gofile

import (
	"fmt"
	"os"
	"path/filepath"
)

// GetRegularFileInfo returns file information (after symlink evaluation)
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
		return nil, fmt.Errorf("the filename %s refers to a directory", filename)
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
