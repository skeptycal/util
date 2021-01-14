// Package fileset contains a custom slice type for manipulating files and directories.
package fileset

import (
	"os"
	"path/filepath"
	"regexp"
)

type Any interface{}

type Set interface {
	Len() int
	Contains(v interface{}) bool
	Get(v interface{})
}

type FileSet map[string]os.FileInfo

func (s FileSet) Len() int { return len(s) }

func (s FileSet) Contains(a string) bool {
	for i := range s {
		if s[i].Name() == a {
			return true
		}
	}
	return false
}

func (s FileSet) Search(pattern, searchpath string) FileSet {
	s = ReDir(pattern, searchpath)
}

func (s FileSet) Dirs() []string {
	retval := []string{}
	for _, dir := range s {
		if dir.IsDir() {
			retval = append(retval, dir.Name())
		}
	}
	return retval
}

func (s FileSet) Extensions() []string {
	ext := ""
	seen := make(map[string]bool)
	for _, file := range s {
		ext = filepath.Ext(file.Name())
		if !seen[ext] {
			seen[ext] = true
		}
	}
	return nil
}

func (s FileSet) Get(a string) interface{} {
	if k, ok := s[a]; ok {
		return k
	}
	return nil
}

func (s FileSet) Add(v interface{}) { s = append(s, v) }

func (s FileSet) Remove(v interface{}) {}

// ReDir Walks a directory tree looking for files that match the
// pattern: re.
//
// A 'WalkFunc' is used to traverse the directory tree and match files:
//
//  type WalkFunc func(path string, info os.FileInfo, err error) error
//
// Uses standard library regexp package and populates the files
// variable with the os.FileInfo of files that match the pattern.
//
//	fmt.Printf("Found %[1]d files.\n", len(files))
func ReDir(pattern, dir string) (fs FileSet) {

	re := regexp.MustCompile(pattern)

	walk := func(fn string, fi os.FileInfo, err error) error {

		// skip files that are not a match
		if !re.MatchString(fn) {
			return nil
		}

		fs[fi.Name()] = fi

		return nil
	}

	err := filepath.Walk(dir, walk)
	if err != nil {
		return nil
	}
	return
}
