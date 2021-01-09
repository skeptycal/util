package gofile

import (
	"io"
	"os"
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
func IsDir(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		if !os.IsNotExist(err) {
			_ = DoOrDie(err)
		}
		return false
	}
	return info.Mode().IsDir()
}

// FileExists checks if path exists in the current directory
// and is not a directory itself.
func FileExists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if !os.IsNotExist(err) {
			_ = DoOrDie(err)
		}
		return false
	}
	return true // !info.IsDir()
}

// Exists returns true if the file exists and is a regular file.
// Does not differentiate between path, file, or permission errors.
func Exists(file string) bool {
	info, err := os.Stat(file)
	if err != nil {
		if err == os.ErrNotExist {
			return false
		}
		return false
	}
	if m := info.Mode(); m.Perm().IsRegular() {
		return true
	}
	return false
}

// IsRegular returns true if the file exists and is regular.
func IsRegular(file string) bool {
	d, err := os.Stat(file)
	if err != nil {
		if err == os.ErrNotExist {
			return false
		}
		return false
	}
	if m := d.Mode(); !m.IsDir() && m&0111 != 0 {
		return true
	}
	return false
}

// IsExecutable returns true if the file exists and is executable.
func IsExecutable(file string) bool {
	d, err := os.Stat(file)
	if err != nil {
		if err == os.ErrNotExist {
			return false
		}
		_ = DoOrDie(err)
		return false
	}
	if m := d.Mode(); !m.IsDir() && m&0111 != 0 {
		return true
	}
	return false
}
