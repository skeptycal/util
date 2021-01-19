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

type fbType byte

const (
	foreground fbType = 3
	background fbType = 4
)

var (
	DefaultioWriter io.Writer = os.Stdout
	AnsiResetString string    = BuildAnsi(DefaultForeground, DefaultBackground, Normal)
)

type colorDepth byte

func encodeAnsi(fb fbType, ef, c color) string {
	if fb < foreground || fb > background {
		fb = foreground
	}
	if ef < 0 || ef > 74 {
		ef = 0
	}

}

var colorfuncs = map[string]string{
	"fmtBasic":  "\x1b[%v%v%m",
	"fmtBright": "\x1b[1;%v%%vm",
	"fmtDim":    "\x1b[2;%v%%vm",
	"fmt256":    "\x1b[%v8;5;%%vm",
	"fmt24":     "\x1b[%v8;2;%%vm",
}

func NewAnsiSet(depth string, fg, bg, ef color) *AnsiSet {
	return &ansiSet{depth, fg, bg, ef}
}

type AnsiSet interface {
	BG(c color) string
	FG(c color) string
	Build(b ...byte) string
	SetType(t int)
	SetColors(fg, bg, ef color)
	String() string
}

type ansiSet struct {
	Func func(fb fbType, c interface{}) string
	fg   interface{}
	bg   interface{}
	ef   interface{}
}

func (a *ansiSet) FG(c color) string {
	if c == 0 {
		c = a.fg.(byte)
	}
	return fmt.Sprintf(a.Func(foreground, c))
	// "\x1b[%v%vm"
}

func (a *ansiSet) SetColors(fg, bg, ef color) {
	fg = fg
	bg = bg
	ef = ef
}

func (a *ansiSet) info() string {
	return fmt.Sprintf("fg: %v, bg: %v, ef %v", a.fg, a.bg, a.ef)
}

func (a *ansiSet) color() interface{} {
	// return fmt.Sprintf(FMTansiSet, a.fg, a.bg, a.ef)

}

func (a *ansiSet) String() string {
	return a.Func(3, a.color())
}

// todo - create a pool of stringbuilders that can go when ready?
// type sbSync struct {
// 	strings.Builder
// 	mu sync.Mutex
// }
