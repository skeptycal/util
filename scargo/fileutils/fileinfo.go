package fileutils

import (
	"os"
)

func exists(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}
	return true
}

// SafeRename renames (moves) oldpath to newpath.
// If newpath already exists, SafeRename returns an error.
// OS-specific restrictions may apply when oldpath and newpath are in different directories.
// If there is an error, it will be of type *LinkError.

// SafeRename - moves (renames) a file or returns *LinkError
// will not overwrite existing file - returns ErrNotExist
func SafeRename(oldpath string, newpath string) error {
	if exists(newpath) {
		return os.ErrExist
	}
	return os.Rename(oldpath, newpath)
}
