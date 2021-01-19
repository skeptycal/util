// Copyright (c) 2020 Michael Treanor
// MIT License

// Package ansi provides fast ansi escape sequence processing based on strings.Builder.
// The standard is defined by the ECMA-48 standard "Control Functions for Coded Character Sets - Fifth Edition"
package ansi

import (
	"fmt"
	"io"
	"os"
)

type color = byte

var (
	DefaultioWriter io.Writer = os.Stdout
	AnsiResetString string    = BuildAnsi(DefaultForeground, DefaultBackground, Normal)
)

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

type AnsiSet interface {
	Build(b ...byte) string
	SetType(t int)
	SetColors(fg, bg, ef color)
	String() string
}

func (s *ansiSet) SetColors(fg, bg, ef color) {
	fg = fg
	bg = bg
	ef = ef
}

type ansiSet struct {
	Func func()
	fg   color
	bg   color
	ef   color
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
