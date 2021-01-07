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
	f *os.File
}

func (r *readWriteRemover) Close() {
	defer os.Remove(r.f.Name())
	r.f.Close()
}

func (r *readWriteRemover) Remove() error {
	// todo - this is redundant and should be accomplished in Close()
	return os.Remove(r.f.Name())
}
