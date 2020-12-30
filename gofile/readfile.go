package gofile

import (
	"io"
	"os"
)

// FileCloseRemover is a wrapper around FileClose that removes
// the file when it is closed, useful for temporary files.
type ReadWriteRemover interface {
	io.ReadWriteCloser
	Remove() error
}

type readWriteRemover struct {
	fileReader
}

func NewFileCloseRemover(io.ReadWriteCloser) *fileCloseRemover {
	r := FileReader()

}

func (f *ReadWriteRemover) Close() error {
	defer os.Remove(f.Name())
	f.Close()
}
