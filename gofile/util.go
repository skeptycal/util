//little-endian
// +build 386 amd64 arm arm64 ppc64le mips64le mipsle riscv64 wasm
//big-endian
// +build ppc64 s390x mips mips64

// non windows file system
// +build aix darwin dragonfly freebsd js,wasm linux netbsd openbsd solaris

// not windows
// +build !windows

// Package gofile provides access to the file system.
package gofile

import (
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	log "github.com/sirupsen/logrus"
	. "github.com/skeptycal/util/stringutils/ansi"
)

// Me returns the filename of the current process.
func Me() string {
	return Base(os.Args[0])
}

// Here returns the parent directory of the current process.
func Here() string {
	return Parent(os.Args[0])
}

// IsEmpty returns true if the directory is empty.
// If path is the empty string, the current directory is tested.
//
// Reference: https://stackoverflow.com/a/30708914
func IsEmpty(path string) (error, bool) {
	if path == "" {
		path = PWD()
	}
	f, err := os.Open(path)
	if err != nil {
		return err, false
	}
	defer f.Close()

	// [...] If n > 0, Readdirnames returns at most n names. In this case,
	// if Readdirnames returns an empty slice, it will return a non-nil
	// error explaining why. At the end of a directory, the error is io.EOF.
	_, err = f.Readdirnames(1)
	if err == io.EOF {
		return nil, true
	}
	return err, false
}

// IsDir checks to see if path is a directory in the current directory.
func IsDir(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		if !os.IsNotExist(err) {
			err = DoOrDie(err)
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
			err = DoOrDie(err)
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
		err = DoOrDie(err)
		return false
	}
	if m := d.Mode(); !m.IsDir() && m&0111 != 0 {
		return true
	}
	return false
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

func RedLogger(args ...interface{}) {
	red := Ansi(Red)
	log.Infof("%s%s", red, args)
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

// Split splits path immediately following the final Separator,
// separating it into a directory and file name component.
// If there is no Separator in path, Split returns an empty dir
// and file set to path.
// The returned values have the property that path = dir+file.
func Split(path string) (dir, file string) {
	vol := VolumeName(path)
	path = path[len(vol):]
	i := len(path) - 1
	for i >= len(vol) && !os.IsPathSeparator(path[i]) {
		i--
	}
	return path[:i+1], path[i+1:]
}

func Split2(path string) (dir, file string) {
	vol := VolumeName(path)
	path = path[len(vol):]
	i := strings.LastIndex(path, string(os.PathSeparator))

	return path[:i+1], path[i+1:]
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

// Base returns the last element of path.
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

// Split splits path immediately following the final Separator,
// separating it into a directory and file name component.
// If there is no Separator in path, Split returns an empty dir
// and file set to path.
// The returned values have the property that path = dir+file.
func Split(path string) (dir, file string) {
	vol := VolumeName(path)
	i := len(path) - 1
	for i >= len(vol) && !os.IsPathSeparator(path[i]) {
		i--
	}
	return path[:i+1], path[i+1:]
}
