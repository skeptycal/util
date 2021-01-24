package redlogger

import (
	"bufio"
	"io"
	"os"

	"github.com/sirupsen/logrus"
	. "github.com/skeptycal/util/stringutils/ansi"
)


    var log = &logrus.Logger{
        Out: os.Stderr,
        Formatter: new(logrus.TextFormatter),
        Hooks: make(logrus.LevelHooks),
        Level: logrus.InfoLevel,
    }
func init() {
    log.SetOutput(os.Stderr)
    log.SetFormatter(new(logrus.TextFormatter))
    log.SetLevel(logrus.InfoLevel)

	log.SetOutput(NewRedLogger(os.Stderr))
	log.Info("RedLogger enabled...")

}

func NewRedLogger(w io.Writer) *redLogger {
    a  := NewAnsiWriter(os.Stdout)
    a.Build(Bold, Black, RedBackground)
	bw := bufio.NewWriter(a)
	return &redLogger{bw: bw}
}

type redLogger struct {
	bw *bufio.Writer
}

// Write wraps p with Ansi color codes and writes the result to the buffer.
func (l *redLogger) Write(p []byte) (n int, err error) {
    n, err = l.bw.WriteString("--> redlogger Write()")
    if err != nil {
        return
    }

    n, err = l.bw.Write(p)
    if err != nil {
        return
    }

	return l.bw.WriteString(Reset)
}

// WriteString wraps p with Ansi color codes and writes the result to the buffer.
func (l *redLogger) WriteString(s string) (n int, err error) {
    n, err = l.bw.WriteString("--> redlogger Write()")
    if err != nil {
        return
    }

    n, err = l.bw.WriteString(s)
    if err != nil {
        return
    }

	return l.bw.WriteString(Reset)
}
