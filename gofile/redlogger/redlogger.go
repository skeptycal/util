package redlogger

import (
	"bufio"
	"io"
	"os"

	"github.com/sirupsen/logrus"
	. "github.com/skeptycal/util/stringutils/ansi"
)

var (
	logColor = AnsiWriter.Build(Bold, RedBackground)
)

func init() {
	var log = logrus.New()

	log.SetOutput(NewRedLogger(os.Stderr))
	log.Info("RedLogger enabled...")

}

func NewRedLogger(w io.Writer) *redLogger {
	bw := bufio.NewWriter(w)
	return &redLogger{bw: bw}
}

type redLogger struct {
	bw *bufio.Writer
}

// Write wraps p with Ansi color codes and writes the result to the buffer.
func (l *redLogger) Write(p []byte) (n int, err error) {
	_, _ = l.bw.WriteString(logColor)
	n, err = l.bw.Write(p)
	_, _ = l.bw.WriteString(Reset)
	return
}

// WriteString wraps p with Ansi color codes and writes the result to the buffer.
func (l *redLogger) WriteString(s string) (n int, err error) {
	_, _ = l.bw.WriteString(logColor)
	n, err = l.bw.WriteString(s)
	_, _ = l.bw.WriteString(Reset)
	return
}
