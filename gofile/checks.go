package gofile

import (
	"io"
	"os"

	"github.com/skeptycal/util/gofile/redlogger"
)

// IsEmpty returns true if the directory is empty.
// If path is the empty string, the current directory is tested.
//
// Reference: https://stackoverflow.com/a/30708914
func IsEmpty(path string) (bool, error) {
	if path == "" {
		path = PWD()
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

// IsDir checks to see if path is a directory in the current directory.
// This is not recursive and does not walk the tree. Use Find or Tree
// for recursive searches.
func IsDir(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		if !os.IsNotExist(err) {
			Err(err)
		}
		return false
	}
	return info.Mode().IsDir()
}

// Exists returns true if the file or directory exists.
// Does not differentiate between path, file, or permission errors.
func Exists(file string) bool {
	_, err := os.Stat(file)
	if err != nil {
		if !os.IsNotExist(err) {
			Err(err)
		}
		return false
	}
	return true
}

// NotDir returns true if file exists and is not a directory.
func NotDir(file string) bool { return !Mode(file).IsDir() }

func IsReg(file string) bool       { return Mode(file).IsRegular() }
func IsExec(file string) bool      { return Mode(file)&0111 == 0111 }
func IsExecOwner(file string) bool { return Mode(file)&0100 != 0 }
func IsExecGroup(file string) bool { return Mode(file)&0010 != 0 }
func IsExecOther(file string) bool { return Mode(file)&0001 != 0 }

// Err calls error handling and logging routines
// without returning the error.
func Err(err error) { redlogger.DoOrDie(err) }
