/*
Package zsh contains command line utilities designed for use with
zsh on macOS. Other environments may be supported, but are not exhaustively
tested.

List of utilities available:
    `binit` creates links to command line tools in ~/bin to provide path access
    `tree` creates a visual tree of directories and files
    `gofind` searches for files and directories matching a pattern
    `ls` lists files and directories in a directory

*/
package zsh

import (
	"context"
	"fmt"
	"io"
	"os/exec"
	"strings"
)

var (
	DefaultContext = context.TODO()
)

// Status executes a shell command and returns the exit status.
func Status(command string) error {
	cmd := CmdPrep(command, DefaultContext)
	return cmd.Run()
}

// Sh executes a shell command line string and returns the result.
// Any error encountered is returned as an ANSI red string.
func Sh(command string) string {
	cmd := CmdPrep(command, DefaultContext)
	stdout, err := cmd.Output()

	if err != nil {
		return fmt.Errorf("%Terror: %v", red, err).Error()
	}

	return string(stdout)
}

// Source is a placeholder // TODO
func Source() string {
	return ""
}

// Replace creates a new Replacer and reads from a list of old, new string pairs. Replacements are performed in the order they appear in the target string, without overlapping matches. The old string comparisons are done in argument order.
// todo - not implemented
func Replace(oldnew []string) *strings.Replacer {
	if len(oldnew)%2 != 0 {
		return nil
	}
	r := strings.NewReplacer(oldnew)
	// cmd :=
	return r
}

// replacer is the interface that a replacement algorithm needs to implement.
type Replacer interface {
	Replace(s string) string
	WriteString(w io.Writer, s string) (n int, err error)
}

// CmdPrep returns a Cmd struct from CommandContext.
// It is like Command but includes a context. Since ctx is a private
// field, this is the only way to add a context.
//
// The provided context is used to kill the process (by calling os.Process.Kill)
// if the context becomes done before the command completes on its own.
// If nil is passed as the context, the default context is used.
func CmdPrep(command string, ctx context.Context) *exec.Cmd {
	if ctx == nil {
		ctx = context.Background()
	}
	app, args := appArgs(command)
	return exec.CommandContext(ctx, app, args...)
}

// CmdPrepNoCtx returns a Cmd struct from Command. It is used to execute
// the named program with the given arguments.
//
// It sets only the Path and Args in the returned structure.
//
// If name contains no path separators, Command uses LookPath to
// resolve name to a complete path if possible. Otherwise it uses name
// directly as Path.
//
// The returned Cmd's Args field is constructed from the command name
// followed by the elements of arg, so arg should not include the
// command name itself. For example, Command("echo", "hello").
// Args[0] is always name, not the possibly resolved Path.
//
// On Windows, processes receive the whole command line as a single string
// and do their own parsing. Command combines and quotes Args into a command
// line string with an algorithm compatible with applications using
// CommandLineToArgvW (which is the most common way). Notable exceptions are
// msiexec.exe and cmd.exe (and thus, all batch files), which have a different
// unquoting algorithm. In these or other similar cases, you can do the
// quoting yourself and provide the full command line in SysProcAttr.CmdLine,
// leaving Args empty.
func CmdPrepNoCtx(command string) *exec.Cmd {
	app, args := appArgs(command)
	return exec.Command(app, args...)
}

// appArgs preps (app,args) for os.Command
func appArgs(command string) (string, []string) {
	a := strings.Split(command, " ")
	return a[0], a[1:]
}

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
