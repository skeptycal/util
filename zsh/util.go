package zsh

import (
	"os"
	"strings"
	"syscall"

	log "github.com/sirupsen/logrus"
	"github.com/skeptycal/util/gofile"
	// "google.golang.org/appengine/log"
)

// Err calls error handling and logging routines
func Err(err error) error {
	return gofile.DoOrDie(err)
}

// isAlphaNum reports whether the byte is an ASCII letter, number, or underscore
func isAlphaNum(c uint8) bool {
	return 'a' <= c && c <= 'z' || '0' <= c && c <= '9' || 'A' <= c && c <= 'Z' || c == '_'
}

// GetEnv returns the environment variable specified by 'key'; if the value is empty or not set, the
// default value is returned.
// If the environment variable is not set, a log event is also
// triggered.
func GetEnv(key string, defValue string) string {
	// value, b := os.LookupEnv(key)
	value, b := syscall.Getenv(key)

	if !b {
		log.Infof("environment variable not set: %s", key)
		return defValue
	}

	if strings.TrimSpace(value) == "" {
		return defValue
	}
	return value
}

// AppArgs returns two strings representing the
// (app,args) arguments for os.Command
//
//  `app` is the first word (space delimited) in command
//  `args` is a string containing the remaining words
//
// Leading and trailing spaces are trimmed using
// strings.TrimSpace() so index 0 cannot contain a space.
//
// If command is a single word then, by definition, there
// are no args, but single letter commands are possible,
// e.g. aliases and short os commands like [
//
func AppArgs(command string) (string, string) {
	command = strings.TrimSpace(command)

	switch len(command) {
	case 0:
		return "", ""
	case 1:
		return command, ""
	}

	i := strings.Index(command, " ")

	if i < 1 {
		return command, ""
	}
	return command[:i], command[i+1:]
}

// Home returns the current user's home directory.
//
// On Unix, including macOS, it returns the $HOME
// environment variable. On Windows, it returns
// %USERPROFILE%. On Plan 9, it returns the $home
// environment variable.
func Home() string {
	s, err := os.UserHomeDir()
	if Err(err) != nil {
		return "~/"
	}
	return s
}
