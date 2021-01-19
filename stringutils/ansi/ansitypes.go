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

func NewAnsiSet(depth string, fg, bg, ef color) *ansiSet {

	return &ansiSet{
		depth: colorfuncs[depth],
		fg:    fmt.Sprintf(FMTansiFG, fg&ansiMask),
		bg:    fmt.Sprintf(FMTansiBG, bg&ansiMask),
		ef:    fmt.Sprintf(FMTansi, ef),
	}
}

type AnsiSet interface {
	String() string
	BG(c color) string
	FG(c color) string
	SetColors(fg, bg, ef color)
	SetType(t int)
	Build(b ...byte) string
}

type ansiSet struct {
	depth string
	fg    string
	bg    string
	ef    string
}

func (a *ansiSet) String() string             { return a.Func(3, a.color()) }
func (a *ansiSet) BG() string                 { return fmt.Sprintf(a.depth, background, a.bg) }
func (a *ansiSet) FG() string                 { return fmt.Sprintf(a.depth, foreground, c) }
func (a *ansiSet) SetColors(fg, bg, ef color) { a.fg = fg; a.bg = bg; a.ef = ef }
func (a *ansiSet) SetColorDepth(depth string) { a.depth = colorfuncs[depth] }
func (a *ansiSet) info() string               { return fmt.Sprintf("fg: %v, bg: %v, ef %v", a.fg, a.bg, a.ef) }

// todo - create a pool of stringbuilders that can go when ready?
// type sbSync struct {
// 	strings.Builder
// 	mu sync.Mutex
// }
