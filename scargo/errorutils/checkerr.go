package errorutils

import (
	"fmt"
)

func die(err error) {
	if err != nil {
		err = fmt.Errorf("a fatal error occurred: %v", err)
		panic(err)
	}
}

// TryExceptPass - a mock implementation of the infamous python error
// ignoring codeblock that writes error message to stderr and then continues
// program execution as if the error never occurred.
// Not the best form, but handy during development sometimes =)
func TryExceptPass(err error, stderrOutput bool) {
	if err != nil {
		if stderrOutput {
			fmt.Printf("An error occurred: %v", err)
		}
	}
}

// Errf checks for error, adds message string, prints output, return bool (err != nil)
func Errf(err error, msg string) bool {
	if err != nil {
		if msg == "" {
			msg = "An error occurred: "
		}
		fmt.Printf(msg, err)
		return false
	}
	return true
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
