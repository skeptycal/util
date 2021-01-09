package gofile

import (
	"fmt"
	"io"
	"os"
)

// Create creates a file and returns an io.ReadCloser on success else error
func Create(filename string) (io.ReadCloser, error) {
	file, err := os.Create(filename)
	if err != nil {
		return nil, fmt.Errorf("failed creating file (%s): %v", filename, err)
	}
	return file, nil
}

// CreateSafe creates a file and returns a ReadCloser on success else error
func CreateSafe(filename string) (io.ReadCloser, error) {
	if Exists(filename) {
		return nil, fmt.Errorf("file already exists (%s): %v", filename, os.ErrExist)
	}
	return Create(filename)
}
