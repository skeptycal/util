package redlogger

import (
	"bufio"
	"fmt"
	"io"
	"os"

	log "github.com/sirupsen/logrus"
	. "github.com/skeptycal/util/stringutils/ansi"
)

var (
	red = Ansi(Red).String()
)

var ErrHandling errorHandling // = continueOnError

func NewRedLogger(w io.Writer) *redLogger {
	bw := bufio.NewWriter(w)
	return &redLogger{bw: bw}
}

type redLogger struct {
	bw *bufio.Writer
}

func (l *redLogger) redWrap(p []byte) (int, error) {
	l.bw.WriteString(red)
	n, err := l.bw.Write(p)
	l.bw.WriteString(Reset)
	return n, err
}

func (l *redLogger) redWrapString(s string) (int, error) {
	l.bw.WriteString(red)
	n, err := l.bw.WriteString(s)
	l.bw.WriteString(Reset)
	return n, err
}

// Write wraps p with Ansi color codes and writes the result to the buffer.
func (l redLogger) Write(p []byte) (int, error) {
	return l.redWrap(p)
}

// WriteString wraps p with Ansi color codes and writes the result to the buffer.
func (l redLogger) WriteString(s string) (int, error) {
	return l.redWrapString(s)
}

// ErrorHandling wraps error handling behavior.
type ErrorHandling interface {
	String() string
	Set()
}

// errorHandling implements the behavior for handling errors.
type errorHandling struct {
	exitOnError bool `default:"false"`
	verbose     bool
	logging     bool
	logwriter   LogWriter `default:"os.Stderr"`
}

func (e errorHandling) Check(err error) error {
	if e.exitOnError == true {
		log.Error(err)
		os.Exit(1)
	}
	return err
}

func (e errorHandling) String() string {
	s := fmt.Sprintf("Error Handling: ExitOnError(%t) , Verbose(%t), Logging(%t)", e.exitOnError, e.verbose, e.logging)

	if e.logging == true {
		s += fmt.Sprintf("\nlogwriter: %v", e.logwriter)
	}

	return s
}

// DoOrDie handles errors based on the value of ErrHandling
// by logging the error and either continuing or exiting.
func DoOrDie(err error) error {
	if err != nil {
		log.Error(err)
		if ErrHandling.exitOnError == true {
			os.Exit(1)
		}
	}
	return err
}
