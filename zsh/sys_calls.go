package zsh

import (
	"fmt"
	"syscall"
	"unsafe"

	log "github.com/sirupsen/logrus"
)

// NewTTY returns a TTY interface used to access the
// terminal capabilities for window size.
func NewTTY() TTY {
	var err error = nil
	ws := &winsize{}
	retCode, _, errno := syscall.Syscall(
		syscall.SYS_IOCTL,
		uintptr(syscall.Stdin),
		uintptr(syscall.TIOCGWINSZ),
		uintptr(unsafe.Pointer(ws)),
	)

	if errno != 0 {
		err = errno
	}

	if err != nil {
		log.Error(err)
	}

	if int(retCode) == -1 {
		log.Errorf("syscall error ... retCode: %d errno: %d", &retCode, &errno)
	}
	return ws
}

// TTY is the interface implemented to access terminal emulation capabilities.
/*
A teleprinter or teletypewriter (TTY) is an electromechanical typewriter
paired with a communication channel or, sometimes used more generally,
any type of computer terminal.[1]

In Unix like operating systems, the `tty` command is commonly used to check
if the output medium is a terminal. The command prints the file name of the
terminal connected to standard input.

If no file is detected (in case, it's being run as part of a script or the command
is being piped) "not a tty" is printed to stdout and the command exits with an exit
status of 1. The command also can be run in silent mode (tty -s) where no output is
produced, and the command exits with an appropriate exit status.[3]

Further probing provides information about the capabilities of the terminal,
such as color support. Go provides some support for the determination of terminal
capabilities, but it is scattered and convoluted.

The TTY interface and utilities provide much more usable functionality.



1. https://en.wikipedia.org/wiki/TTY
2. https://en.wikipedia.org/wiki/Tty_(unix)
3. https://linux.die.net/man/1/tty
*/
type TTY interface {
	Col() uint
	Row() uint
	X() uint
	Y() uint
}

// winsize represents terminal constants
type winsize struct {
	row    uint16
	col    uint16
	xpixel uint16
	ypixel uint16
}

func (ws *winsize) String() string {
	return fmt.Sprintf("TTY row: %d, col: %d resolution %dx%d", ws.row, ws.col, ws.xpixel, ws.ypixel)
}

func (ws *winsize) Col() uint {
	return uint(ws.col)
}

func (ws *winsize) Row() uint {
	return uint(ws.row)
}

func (ws *winsize) X() uint {
	return uint(ws.xpixel)
}

func (ws *winsize) Y() uint {
	return uint(ws.ypixel)
}
