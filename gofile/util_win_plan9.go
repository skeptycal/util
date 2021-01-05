// package gofile
// +build windows plan9
package gofile

import (
	"syscall"
)

const (
	defaultPWD = ".\\"
)

// Getwd returns a rooted path name corresponding to the
// current directory. If the current directory can be
// reached via multiple paths (due to symbolic links),
// Getwd may return any one of them.
func Getwd() (dir string, err error) {
	return syscall.Getwd()
}

// VolumeName returns leading volume name.
// Given "C:\foo\bar" it returns "C:" on Windows.
// Given "\\host\share\foo" it returns "\\host\share".
// On other platforms it returns "".
func VolumeName(path string) string {
	return path[:volumeNameLen(path)]
}
