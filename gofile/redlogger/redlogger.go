package redlogger

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/skeptycal/util/stringutils/ansi"
)



func init() {
    var log = &logrus.Logger{
        Out: New(os.Stderr),
        Formatter: new(logrus.TextFormatter),
        Hooks: make(logrus.LevelHooks),
        Level: logrus.InfoLevel,
    }
    // log.SetFormatter(new(logrus.TextFormatter))
    log.SetLevel(logrus.InfoLevel)

	log.Info("RedLogger enabled...")
}

func New(w io.Writer) *ansi.AnsiWriter {
    if w == nil {
        w = os.Stdout
    }
    a  := ansi.NewWriter(w)
    a.Build(ansi.Bold, ansi.Black, ansi.RedBackground)
	return a
}

type redLogger ansi.AnsiWriter

// Write wraps p with Ansi color codes and writes the result to the buffer.
func (l *redLogger) Write(p []byte) (n int, err error) {
    n, err = l.WriteString("--> redlogger Write()")
    if err != nil {
        return
    }

    n, err = l.Write(p)
    if err != nil {
        return
    }

	return l.WriteString(Reset)
}

// WriteString wraps p with Ansi color codes and writes the result to the buffer.
func (l *redLogger) WriteString(s string) (n int, err error) {
    n, err = l.WriteString("--> redlogger WriteString()")
    if err != nil {
        return
    }

    n, err = l.WriteString(s)
    if err != nil {
        return
    }

	return l.WriteString(Reset)
}
