package zsh

import (
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"os/exec"
)

var (
	redb, _        = hex.DecodeString("1b5b33316d0a") // byte code for ANSI red
	red     string = string(redb)                     // ANSI red
)

// Exists returns true if the file exists and is executable.
// Does not differentiate between path, file, or permission errors.
func Exists(file string) bool {
	d, err := os.Stat(file)
	if err != nil {
		return false
	}
	if m := d.Mode(); !m.IsDir() && m&0111 != 0 {
		return true
	}
	return false
}

// IsExecutable returns nil if the file exists and is executable
// If the permissions are incorrect, os.ErrPermission is returned;
// otherwise, error will be of type *PathError.
func IsExecutable(file string) error {
	d, err := os.Stat(file)
	if err != nil {
		return err
	}
	if m := d.Mode(); !m.IsDir() && m&0111 != 0 {
		return nil
	}
	return os.ErrPermission
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

// WriteFile creates the file 'fileName' and writes 'data' to it.
// It returns any error encountered. If the file already exists, it
// will be TRUNCATED and OVERWRITTEN.
func WriteFile(fileName string, data string) error {
	dataFile, err := OpenTrunc(fileName)
	if err != nil {
		log.Println(err)
		return err
	}
	defer dataFile.Close()

	n, err := dataFile.WriteString(data)
	if err != nil {
		log.Println(err)
		return err
	}
	if n != len(data) {
		log.Printf("incorrect string length written (wanted %d): %d\n", len(data), n)
		return fmt.Errorf("incorrect string length written (wanted %d): %d", len(data), n)
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
	return os.OpenFile(name, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 644)
}
