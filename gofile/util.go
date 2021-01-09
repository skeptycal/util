// Package gofile provides access to the file system.
package gofile

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	log "github.com/sirupsen/logrus"
)

// Me returns the filename of the current process.
func Me() string {
	return Base(os.Args[0])
}

// Here returns the parent directory of the current process.
func Here() string {
	return Parent(os.Args[0])
}

// Which searches for an executable named file in the
// directories named by the PATH environment variable.
// If file contains a slash, it is tried directly and the PATH is not consulted.
// The result may be an absolute path or a path relative to the current directory.
// On Windows, LookPath also uses PATHEXT environment variable to match
// a suitable candidate.
func Which(file string) (string, error) {
	return exec.LookPath(file)
}

// PWD returns the current working directory. It does not return
//  any error, but instead logs the error and returns the system
// default glob pattern for current working directory.
//
// PWD runs Abs and returns an absolute representation of path.
// If the path is not absolute it will be joined with the current
// working directory to turn it into an absolute path. The absolute
// path name for a given file is not guaranteed to be unique.
// Abs calls Clean on the result.
func PWD() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Errorf("PWD could not locate current directory (using OS pwd): %v", err)

		// this is a crutch for the extremely rare case where Getwd fails
		return Abs(defaultPWD)
	}
	wd, err := filepath.Abs(dir)
	if err != nil {
		log.Errorf("PWD could not determine absolute path of pwd: %v", err)

		// this is a crutch for the extremely rare case where Abs fails
		return Abs(defaultPWD)
	}

	return Abs(wd)
}

// BaseWD returns the basename of the current directory (PWD).
func BaseWD() string {
	_, file := filepath.Split(Abs(PWD()))
	return file
}

// Abs returns an absolute representation of path.
// If the path is not absolute it will be joined with the current
// working directory to turn it into an absolute path. The absolute
// path name for a given file is not guaranteed to be unique.
// Abs calls Clean on the result.
func Abs(path string) string {
	path, _ = filepath.Abs(path)
	return path
}

// Base returns the last element of path.
// Trailing path separators are removed before extracting the last element.
// If the path is empty, Base returns ".".
// If the path consists entirely of separators, Base returns a single separator.
func Base(path string) string {
	_, file := filepath.Split(path)
	return file
}

// BaseGo returns the last element of path.
// This is a convenience version modified from Go 1.15.6
// (located at /src/path/filepath/path.go)
//
// Trailing path separators are removed before extracting the last element.
// If the path is empty, Base returns ".".
// If the path consists entirely of separators, Base returns a single separator.
func BaseGo(path string) string {
	// Strip trailing slashes.
	path = strings.TrimRight(path, SEP)

	// Throw away volume name
	path = path[len(filepath.VolumeName(path)):]

	// Find the last path separator
	i := strings.LastIndex(path, SEP)

	switch i {
	case -1: // no path separator found
		if path == "" { // if path is empty
			return emptyPath
		}
		return path
	case 0: // last path separator found at index 0
		return emptyPath
	default: // return all after the last path separator
		return path[i+1:]
	}
}

// SafeRename renames (moves) oldpath to newpath.
// If oldpath is not found or newpath already exists, SafeRename
// returns an error.
func SafeRename(oldpath string, newpath string) error {
	if Exists(newpath) {
		return os.ErrExist
	}
	return os.Rename(oldpath, newpath)
}

// Parent returns the parent directory of path.
func Parent(path string) string {
	dir, _ := filepath.Split(Abs(path))
	return dir
}

func Parents(path string) []string {
	return filepath.SplitList(Abs(path))
}
