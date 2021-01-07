// package gofile
// +build aix darwin dragonfly freebsd js,wasm linux netbsd openbsd solaris
package gofile

const (
	defaultPWD = "."
)

// Shorthand version of filepath.VolumeName()
func VolumeName(path string) string {
	return ""
}
