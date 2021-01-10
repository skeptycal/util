package gofile

import (
	"bufio"
	"os"
)

// ReadWriteRemover is a wrapper around bufio.ReadWriter that removes
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

// func NewFileCloseRemover(r *bufio.Reader, w *bufio.Writer) *readWriteRemover {

// 	rw := &readWriteRemover{
// 		bufio.NewReadWriter(r, w),
// 	}
// 	return rw
// }

// os.File notes:
/* type file struct {
	pfd         poll.FD
	name        string
	dirinfo     *dirInfo // nil unless directory being read
	nonblock    bool     // whether we set nonblocking mode
	stdoutOrErr bool     // whether this is stdout or stderr
	appendMode  bool     // whether file is opened for appending
}
*/
