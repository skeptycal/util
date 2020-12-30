package gofile

import (
	"os"
	"path/filepath"
)

type pathlib struct {
	name   string
	base   string
	parent string
	f      *os.File
}

func (f *pathlib) Split() (dir string, file string) {
	if f.base == "" || f.parent == "" {
		path, err := filepath.Abs(f.name)
		_ = DoOrDie(err)
		f.parent, f.base = filepath.Split(path)
	}
	return f.parent, f.base
}

// Parent returns the parent directory of the file.
func (f *pathlib) Parent() string {
	f.Split()
	return f.parent
}

// Base returns the file name of the file.
func (f *pathlib) Base() string {
	f.Split()
	return f.base
}
