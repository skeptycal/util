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

// Getwd returns a rooted path name corresponding to the
// current directory. If the current directory can be
// reached via multiple paths (due to symbolic links),
// Getwd may return any one of them.
func Getwd() (dir string, err error) {

	// Clumsy but widespread kludge:
	// if $PWD is set and matches ".", use it.
	dot, err := statNolog(".")
	if err != nil {
		return "", err
	}
	dir = Getenv("PWD")
	if len(dir) > 0 && dir[0] == '/' {
		d, err := statNolog(dir)
		if err == nil && SameFile(dot, d) {
			return dir, nil
		}
	}
}
