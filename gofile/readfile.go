package gofile

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

// Reference: https://gobyexample.com/signals
func Sig() {
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()

	fmt.Println("awaiting signal")
	<-done
	fmt.Println("exiting")

}

// FileCloseRemover is a wrapper around bufio.ReadWriter that removes
// the file when it is closed, useful for temporary files.
type ReadWriteRemover interface {
	bufio.ReadWriter
	Remove() error
}

type readWriteRemover struct {
	*bufio.Reader
	*bufio.Writer
	*os.File
}

func (f *ReadWriteRemover) Close() error {
	defer os.Remove(f.Name())
	f.Close()
}

func (f *ReadWriteRemover) Remove() error {
	// todo - this is redundant and should be accomplished in Close()
	os.Remove(f.Name())
}

func NewFileCloseRemover(r *bufio.Reader, w *bufio.Writer) *readWriteRemover {
	return &bufio.NewReadWriter(r, w)
}

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
