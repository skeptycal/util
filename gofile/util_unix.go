// +build aix darwin dragonfly freebsd js,wasm linux netbsd openbsd solaris

package gofile

import (
	"os"
	"strings"
)

const (
	defaultPWD = "."
)

// VolumeName is a shorthand version of filepath.VolumeName()
// todo should be removed
func VolumeName(path string) string {
	return ""
}

// Split splits path immediately following the final Separator,
// separating it into a directory and file name component.
// If there is no Separator in path, Split returns an empty dir
// and file set to path.
// The returned values have the property that path = dir+file.
func Split(path string) (dir, file string) {
	i := len(path) - 1
	for i >= 0 && !os.IsPathSeparator(path[i]) {
		i--
	}
	return path[:i+1], path[i+1:]
}

// Split2 is an alternate version for comparing benchmarks.
func Split2(path string) (dir, file string) {
	i := strings.LastIndex(path, SEP)

	return path[:i+1], path[i+1:]
}
