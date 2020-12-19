package gofile

import (
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"os/exec"

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
	if err != nil {
		return err
	}
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
