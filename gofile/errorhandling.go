package gofile

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
)

var ErrHandling errorHandling // = continueOnError

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
