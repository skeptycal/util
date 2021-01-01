package codeinfo

import (
	"io"
	"os"

	"github.com/skeptycal/util/gofile"
)

// DirNames returns the contents of the directory path and returns
// a slice containing names of files in the directory, in directory
// order. Subsequent calls on the same file will yield further names.
//
// [...] If n <= 0, Readdirnames returns all the names from the
// directory in a single slice. In this case, if Readdirnames succeeds (reads all the way to the end of the directory), it returns the slice and a nil error. If it encounters an error before the end of the directory, Readdirnames returns the names read until that point and a non-nil error.
func DirNames(path, ext string) ([]string, error) {
	if path == "" {
		path = gofile.PWD()
	}
	f, err := os.Open(path)
	if err != nil {
		return []string{}, err
	}
	defer f.Close()
	return f.Readdirnames(-1)
}

func IsEmpty(path string) (bool, error) {
	if path == "" {
		path = gofile.PWD()
	}
	f, err := os.Open(path)
	if err != nil {
		return false, err
	}
	defer f.Close()

	// [...] If n > 0, Readdirnames returns at most n names. In this case,
	// if Readdirnames returns an empty slice, it will return a non-nil
	// error explaining why. At the end of a directory, the error is io.EOF.
	_, err = f.Readdirnames(1)
	if err == io.EOF {
		return true, nil
	}
	return false, err
}
