package gofile

import (
	"os"
	"time"
)

type FileInfo interface {
	Name() string       // base name of the file
	Size() int64        // length in bytes for regular files; system-dependent for others
	Mode() FileMode     // file mode bits
	ModTime() time.Time // modification time
	IsDir() bool        // abbreviation for Mode().IsDir()
	Sys() interface{}   // underlying data source (can return nil)
}
type FileMethods struct {
    os.FileInfo
    ParentDir() string
    BaseName() string
    AbsPath() string
}

func (f FileMethods) PWD() string {
	root, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	path, err := filepath.Abs(root)
	if err != nil {
		return nil, err
    }
}

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
