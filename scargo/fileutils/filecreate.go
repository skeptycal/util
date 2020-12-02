package fileutils

import (
	"fmt"
	"os"
)

// CreateFileTruncate - Creates a file and returns a closer on success else error
func CreateFileTruncate(filename string) (*os.File, error) {
	file, err := os.Create(filename)
	if err != nil {
		defer file.Close()
		return nil, fmt.Errorf("failed creating file (%s): %v", filename, err)
	}
	return file, nil
}

// CreateFileSafe - Creates a file and returns a closer on success else error
func CreateFileSafe(filename string) (*os.File, error) {
	file, err := os.Create(filename)
	if err != nil {
		defer file.Close()
		return nil, fmt.Errorf("failed creating file (%s): %v", filename, err)
	}
	return file, nil
}
