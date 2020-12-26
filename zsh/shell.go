package zsh

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"

	log "github.com/sirupsen/logrus"
	. "github.com/skeptycal/util/stringutils/ansi"
)

var (
	DefaultContext      = context.Background()
	red            Ansi = Ansi(Red)
)

// Status executes a shell command and returns the exit status as an error.
func Status(command string) error {
	cmd := CmdPrep(command)
	return cmd.Run()
}

// Sh executes a shell command line string and returns the result.
// Any error encountered is returned as an ANSI red string.
func Sh(command string) string {
	cmd := CmdPrep(command)
	stdout, err := cmd.Output()

	if err != nil {
		return fmt.Sprintf("%verror: %v", red, err)
	}

	return string(stdout)
}

// Repl executes the command and returns the the result.
// Unlike Shell(), the Repl() process has access to the parent's
// stdin, stdout, and stderr streams.
func Repl(command string) (string, error) {
	return shell(CmdPrep(command), os.Stdin, os.Stdout, os.Stderr)
}

// ShellIn executes the command and returns the the result.
// Unlike Shell(), the ShellIn() process has access to the parent's stdin.
// This can be used to query stdin for parameters like 'size'
//      ShellIn("stty size")
//      out: "36 118\n"
//      err: <nil>
//
// Or to preload the stdin buffer.
func ShellIn(command string) (string, error) {
	return shell(CmdPrep(command), os.Stdin, nil, nil)
}

// ShellOut executes the command and returns the the result.
// Unlike Shell(), the ShellOut() process has access to the parent's
// stdout and stderr streams.
func ShellOut(command string) (string, error) {
	return shell(CmdPrep(command), nil, os.Stdout, os.Stderr)
}

// Shell executes a command line string and returns the result.
func Shell(command string) (string, error) {
	return shell(CmdPrep(command), nil, nil, nil)
}

// shell executes a prepared command structure and returns the result.
// []byte values are converted to string
// sin, sout, and serr are used to redirect input and output of the command.
func shell(cmd *exec.Cmd, sin io.Reader, sout, serr io.Writer) (string, error) {
	if sin != nil {
		cmd.Stdin = sin
	}
	if sout != nil {
		cmd.Stdout = sout
	}
	if serr != nil {
		cmd.Stderr = serr
	}

	out, err := cmd.Output()

	if err != nil {
		log.Error(err)
		return string(out), err
	}
	return string(out), err
}

// cmdPrep prepares a Cmd struct from a command line string.
func cmdPrep(command string, ctx context.Context) *exec.Cmd {
	if ctx == nil {
		ctx = DefaultContext
	}
	s := strings.Split(command, " ")
	return exec.CommandContext(ctx, s[0], s[1:]...)
}

// CmdPrep returns a Cmd struct from CommandContext.
// It is like Command but includes a context. Since ctx is a private
// field, this is the only way to add a context.
//
// The provided context is used to kill the process (by calling os.Process.Kill)
// if the context becomes done before the command completes on its own.
// If nil is passed as the context, the default context is used.
func CmdPrep(command string) *exec.Cmd {
	return cmdPrep(command, DefaultContext)
}

// WriteFile creates the file 'fileName' and writes all 'data' to it.
// It returns any error encountered. If the file already exists, it
// will be TRUNCATED and OVERWRITTEN.
func WriteFile(fileName string, data string) error {
	dataFile, err := OpenTrunc(fileName)
	if err != nil {
		log.Error(err)
		return err
	}
	defer dataFile.Close()

	n, err := dataFile.WriteString(data)
	if err != nil {
		log.Println(err)
		return err
	}
	if n != len(data) {
		log.Printf("incorrect string length written (wanted %d): %d\n", len(data), n)
		return fmt.Errorf("incorrect string length written (wanted %d): %d", len(data), n)
	}
	return nil
}

// type result struct {
// 	stdout string
// 	stderr string
// 	retval int
// 	err    error
// }

// OpenTrunc creates and opens the named file for writing. If successful, methods on
// the returned file can be used for writing; the associated file descriptor has mode
//      O_WRONLY|O_CREATE|O_TRUNC
// If the file does not exist, it is created with mode o644;
//
// If the file already exists, it is TRUNCATED and overwritten
//
// If there is an error, it will be of type *PathError.
func OpenTrunc(name string) (*os.File, error) {
	return os.OpenFile(name, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 644)
}

// appArgs preps (app,args) for os.Command
func appArgs(command string) (string, []string) {
	a := strings.Split(command, " ")
	return a[0], a[1:]
}

// fileExists checks if a file exists and is not a directory
// func fileExists(fileName string) bool {
// 	info, err := os.Stat(fileName)
// 	if os.IsNotExist(err) {
// 		return false
// 	}
// 	return !info.IsDir()
// }

// Notes: Cmd struct summary:
/*
type Cmd struct {
	Path            string
	Args            []string
	Env             []string
	Dir             string
	Stdin           io.Reader
	Stdout          io.Writer
	Stderr          io.Writer
	ExtraFiles      []*os.File
	SysProcAttr     *syscall.SysProcAttr
	Process         *os.Process
	ProcessState    *os.ProcessState
	ctx             context.Context // nil means none
	lookPathErr     error           // LookPath error, if any.
	finished        bool            // when Wait was called
	childFiles      []*os.File
	closeAfterStart []io.Closer
	closeAfterWait  []io.Closer
	goroutine       []func() error
	errch           chan error // one send per goroutine
	waitDone        chan struct{}
}
*/

// Notes: Cmd struct summary:
/*
type Cmd struct {
	Path            string
	Args            []string
	Env             []string
	Dir             string
	Stdin           io.Reader
	Stdout          io.Writer
	Stderr          io.Writer
	ExtraFiles      []*os.File
	SysProcAttr     *syscall.SysProcAttr
	Process         *os.Process
	ProcessState    *os.ProcessState
	ctx             context.Context // nil means none
	lookPathErr     error           // LookPath error, if any.
	finished        bool            // when Wait was called
	childFiles      []*os.File
	closeAfterStart []io.Closer
	closeAfterWait  []io.Closer
	goroutine       []func() error
	errch           chan error // one send per goroutine
	waitDone        chan struct{}
}
*/
