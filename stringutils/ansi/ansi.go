// Copyright (c) 2020 Michael Treanor
// MIT License

// Package ansi provides fast ansi escape sequence processing based on strings.Builder.
// The standard is defined by the ECMA-48 standard "Control Functions for Coded Character Sets - Fifth Edition"
package ansi

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// const (
// 	fmtBasic  string = "\x1b[%vm"
// 	fmtBright string = "\x1b[1;%vm"
// 	fmtDim    string = "\x1b[2;%vm"
// 	fmt256FG  string = "\x1b[38;5;%vm"
// 	fmt256BG  string = "\x1b[48;5;%vm"
// 	fmt24FG   string = "\x1b[38;2;%v;%v;%vm"
// 	fmt24BG   string = "\x1b[48;2;%v;%v;%vm"
// )

var (
	DefaultioWriter io.Writer = os.Stdout

	// Reset Effects; Default Foreground; Default Background
	AnsiResetString string = BuildAnsi(DefaultForeground, DefaultBackground, Normal)
)

type color = byte
type colorDepth byte

type colorDepths struct {
	name      string
	fmt       string
	colorfunc func()
}

var types = []colorDepths{
	{"fmtBasic", "\x1b[%vm", func() {}},
	{"fmtBright", "\x1b[1;%vm", func() {}},
	{"fmtDim", "\x1b[2;%vm", func() {}},
	{"fmt256FG", "\x1b[38;5;%vm", func() {}},
	{"fmt256BG", "\x1b[48;5;%vm", func() {}},
	{"fmt24FG", "\x1b[38;2;%v;%v;%vm", func() {}},
	{"fmt24BG", "\x1b[48;2;%v;%v;%vm", func() {}},
}

var colorfuncs = map[string]func(){}

func NewAnsiSet(depth string, fg, bg, ef color) *AnsiSet {
	return &AnsiSet{depth, fg, bg, ef}
}

type AnsiSet struct {
	depth string
	fg    color
	bg    color
	ef    color
}

func (a *AnsiSet) info() string {
	return fmt.Sprintf("fg: %v, bg: %v, ef %v", a.fg, a.bg, a.ef)
}

func (a *AnsiSet) String() string {
	return fmt.Sprintf(FMTansiSet, a.fg, a.bg, a.ef)
}

// todo - create a pool of stringbuilders that can go when ready?
// type sbSync struct {
// 	strings.Builder
// 	mu sync.Mutex
// }

// BuildAnsi returns a basic (3/4 bit) ANSI format code
// from a variadic argument list of bytes
func BuildAnsi(b ...byte) string {
	sb := strings.Builder{}
	defer sb.Reset()
	for _, n := range b {
		sb.WriteString(fmt.Sprintf(FMTansi, n))
	}
	return sb.String()
}
