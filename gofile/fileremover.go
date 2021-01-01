package gofile

import (
	"bufio"
	"os"
)

// FileCloseRemover is a wrapper around bufio.ReadWriter that removes
// the file when it is closed, useful for temporary files.
type ReadWriteRemover interface {
	Read(p []byte) (int, error)
	Write(p []byte) (int, error)
	Remove() error
}

type readWriteRemover struct {
	*bufio.ReadWriter
	*os.File
}

func (f *readWriteRemover) Close() error {
	defer os.Remove(f.Name())
	f.Close()
}

func (f *readWriteRemover) Remove() error {
	// todo - this is redundant and should be accomplished in Close()
	os.Remove(f.Name())
}
