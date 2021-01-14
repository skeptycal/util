package zsh

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"

	log "github.com/sirupsen/logrus"
	"github.com/skeptycal/util/gofile"
)

type TrueMap map[string]bool

func (a TrueMap) String() string {
	sb := strings.Builder{}
	for k, v := range a {
		if v == true {
			sb.WriteString(k)
			sb.WriteString(",")
		}
	}
	return sb.String()
}

func (a TrueMap) List() (s []string) {
	for k, v := range a {
		if v == true {
			s = append(s, k)
		}
	}
	return
}

type Stringer interface {
	String() string
}

func Tree(pathname string) []string {

}

func Dir(pathname string) (files []os.FileInfo, err error) {
	return files, filepath.Walk(pathname,
		func(path string, info os.FileInfo, err error) error {
            if
			files = append(files, info)
			return nil
		})
}

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

func ToString(any interface{}) string {
	if v, ok := any.(Stringer); ok {
		return v.String()
	}
	switch v := any.(type) {
	case int:
		return strconv.Itoa(v)
	case float64, float32:
		return fmt.Sprintf("%.2g", v)
	}
	return "???"
}

// Contains tells whether a contains x.
func Contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

// NameSplit returns separate name and extension of file.
func NameSplit(filename string) (string, string) {
	s := strings.Split(filepath.Base(filename), ".")
	name := s[0]
	ext := ""
	if len(s) > 1 {
		ext = s[len(s)-1]
	}
	return name, ext
}

// Name returns the name of file without path information.
func Base(filename string) string {
	ns, _ := NameSplit(filename)
	return NameSplit[1]
}

// AbsPath returns the absolute path of file.
func AbsPath(file string) string {
	dir, _ := filepath.Abs(file)
	return dir
}
