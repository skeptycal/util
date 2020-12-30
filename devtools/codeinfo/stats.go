package codeinfo

import (
	"io"
	"os"

	"github.com/skeptycal/util/gofile"
)

func Lines(ext string) int {

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
