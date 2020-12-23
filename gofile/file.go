package gofile

import (
	"encoding/csv"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	log "github.com/sirupsen/logrus"
)

type StringWriterCloser interface {
	io.StringWriter
	io.Closer
}

var (
	redb, _        = hex.DecodeString("1b5b33316d0a") // byte code for ANSI red
	red     string = string(redb)                     // ANSI red
)

type errorHandling uint8

const (
	SEP = string(os.PathSeparator)

	continueOnError errorHandling = iota
	exitOnError
)

var ErrHandling errorHandling = continueOnError
var emptyPath string

func init() {
	if runtime.GOOS == "windows" || runtime.GOOS == "plan9" {
		emptyPath = ".\\"
	}
	emptyPath = "."
}

// DoOrDie handles errors based on the value of ErrHandling
// by logging the error and either continuing or exiting.
func DoOrDie(err error) error {
	if err != nil {
		log.Error(err)
		if ErrHandling == exitOnError {
			os.Exit(1)
		}
	}
	return err
}

// NewCSVReader returns a new Reader that reads from file.
func NewCSVReader(file string) (*csv.Reader, error) {
	fi, err := os.Open(file)
	if DoOrDie(err) != nil {
		return nil, err
	}

	return csv.NewReader(fi), nil
}

// ReadCSV reads all records from file.
// Each record is a slice of fields.
// A successful call returns err == nil, not err == io.EOF. Because ReadAll is
// defined to read until EOF, it does not treat end of file as an error to be
// reported.
func ReadCSV(file string) ([][]string, error) {
	r, err := NewCSVReader(file)
	if DoOrDie(err) != nil {
		return nil, err
	}
	return r.ReadAll()
}

// Exists returns true if the file exists and is a regular file.
// Does not differentiate between path, file, or permission errors.
func Exists(file string) bool {
	d, err := os.Stat(file)
	if err != nil {
		return false
	}
	if m := d.Mode(); m.Perm().IsRegular() {
		return true
	}
	return false
}

// IsExecutable returns 0 if the file exists and is executable.
// Returns 1 if the file does not exist, -1 for all other errors.
//
func IsExecutable(file string) int {
	d, err := os.Stat(file)
	if err != nil {
		if err == os.ErrNotExist {
			return 1
		}
		return -1
	}
	if m := d.Mode(); !m.IsDir() && m&0111 != 0 {
		return 0
	}
	return -1
}

// FileExists checks if file exists in the current directory
// and is not a directory itself.
func FileExists(file string) bool {
	info, err := os.Stat(file)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
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

// WriteFile creates the file fileName, writes data to it, and closes the files.
//
// This will force fileName to be created or truncated.
// It returns any error encountered.
//
// Again ... If the file already exists, it will be TRUNCATED and OVERWRITTEN.
func WriteFile(fileName string, data string) error {
	dataFile, err := OpenTrunc(fileName)
	if DoOrDie(err) != nil {
		return err
	}
	_ = DoOrDie(err)

	w := StringWriterCloser(dataFile)
	defer w.Close()

	n, err := w.WriteString(data)

	if err != nil {
		return err
	}

	if n != len(data) {
		// log.Printf("incorrect string length written (wanted %d), returned %d\n", len(data), n)
		return fmt.Errorf("incorrect string length written: wanted %d, wrote %d", len(data), n)
	}
	return nil
}

// OpenTrunc creates and opens the named file for writing. If successful, methods on
// the returned file can be used for writing; the associated file descriptor has mode
//      O_WRONLY|O_CREATE|O_TRUNC
// If the file does not exist, it is created with mode o644;
//
// If the file already exists, it is TRUNCATED and overwritten
//
// If there is an error, it will be of type *PathError.
func OpenTrunc(name string) (*os.File, error) {
	return os.OpenFile(name, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
}

func RedLogger(args ...interface{}) {
	log.Infof("%s%s", red, args)
}

// type FileInfo interface {
// 	Name() string       // base name of the file
// 	Size() int64        // length in bytes for regular files; system-dependent for others
// 	Mode() FileMode     // file mode bits
// 	ModTime() time.Time // modification time
// 	IsDir() bool        // abbreviation for Mode().IsDir()
// 	Sys() interface{}   // underlying data source (can return nil)
// }

// type file struct {
// 	pfd         poll.FD
// 	name        string
// 	dirinfo     *dirInfo // nil unless directory being read
// 	nonblock    bool     // whether we set nonblocking mode
// 	stdoutOrErr bool     // whether this is stdout or stderr
// 	appendMode  bool     // whether file is opened for appending
// }

// type FileMethods struct {
// 	os.FileInfo
// 	// ParentDir() string
// 	// Base() string
// 	// AbsPath() string
// }

type filemethods struct {
	name   string
	base   string
	parent string
	f      *os.File
}

func (f *filemethods) Parent() string {
	if f.parent == "" {
		f.parent, f.base = filepath.Split(f.name)
	}
	return f.parent
}

func (f *filemethods) Base() string {
	if f.base == "" {
		f.parent, f.base = filepath.Split(f.name)
	}
	return f.base
}

// PWD returns the current working directory. It does not return any error, but instead
// logs the error and returns the system default glob pattern for current working directory
func PWD() string {
	wd, err := os.Getwd()
	if err != nil {
		log.Errorf("PWD could not locate current directory (using '.'): %v", err)
		// this is a crutch for the extremely rare case where Getwd fails
		if runtime.GOOS == "windows" || runtime.GOOS == "plan9" {
			return ".\\"
		}
		return "."
	}
	return wd
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

// Split splits path immediately following the final Separator,
// separating it into a directory and file name component.
// If there is no Separator in path, Split returns an empty dir
// and file set to path.
// The returned values have the property that path = dir+file.
func Split(path string) (dir, file string) {
	vol := filepath.VolumeName(path)
	path = path[len(vol):]
	i := len(path) - 1
	for i >= len(vol) && !os.IsPathSeparator(path[i]) {
		i--
	}
	return path[:i+1], path[i+1:]
}

func Parents(path string) []string {
	clean := filepath.Clean(path)
	return filepath.SplitList(clean)
}

func Base(path string) string {
	_, file := filepath.Split(path)
	return file
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
	dir, _ := filepath.Split(filepath.Clean(path))
	return dir
}

// BaseWD returns the basename of the current directory (PWD).
func BaseWD(path string) string {
	_, file := filepath.Split(filepath.Clean(path))
	return file
}
