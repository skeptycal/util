// Package gofile provides support for file operations.
package gofile

import (
	"fmt"
	"io"
	"os"
)

// CheckStat returns the os.FileInfo for file if it exists. Any
// error except os.IsNotExist are logged but are not returned.
// If the file does not exist, nil is returned.
func CheckStat(file string) os.FileInfo {
	fi, err := os.Stat(file)
	if err != nil {
		if !os.IsNotExist(err) {
			Err(err)
		}
		return nil
	}
	return fi
}

// Mode returns the filemode of file.
func Mode(file string) os.FileMode { return CheckStat(file).Mode() }

// Create creates a file and returns an io.ReadCloser on success else error
func Create(filename string) (io.ReadWriteCloser, error) {
	file, err := os.Create(filename)
	if err != nil {
		return nil, fmt.Errorf("failed creating file (%s): %v", filename, err)
	}
	return file, nil
}

// CreateSafe creates a file and returns a ReadCloser on success else error
func CreateSafe(filename string) (io.ReadWriteCloser, error) {
	if Exists(filename) {
		return nil, fmt.Errorf("file already exists (%s): %v", filename, os.ErrExist)
	}
	return Create(filename)
}
