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

var colordepth = map[string]string{
	"fmtBasic": "\x1b[%v%v%m",
	"fmt256":   "\x1b[%v8;5;%%vm",
	"fmt24":    "\x1b[%v8;2;%%vm",
}

func NewAnsiSet(depth string, fg, bg, ef color) *ansiSet {

	a := &ansiSet{}
	a.SetColorDepth(depth)
	a.SetColors(fg, bg, ef)

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
	out   string
}

func (a *ansiSet) String() string { return a.out }
func (a *ansiSet) BG() string     { return a.bg }
func (a *ansiSet) FG() string     { return a.fg }
func (a *ansiSet) SetColors(fg, bg, ef color) {
	o := fmt.Sprintf("%v;3%v;4%v", ef, fg&ansiMask, bg&ansiMask)

	a.fg = fmt.Sprintf(FMTansiFG, fg&ansiMask)
	a.bg = fmt.Sprintf(FMTansiBG, bg&ansiMask)
	a.ef = fmt.Sprintf(FMTansi, ef)
	a.out = fmt.Sprintf(FMTansi, o)
}
func (a *ansiSet) out() string                { return fmt.Sprintf("%v;%v;%v", a.ef, a.fg, a.bg) }
func (a *ansiSet) SetColorDepth(depth string) { a.depth = colorfuncs[depth] }
func (a *ansiSet) info() string               { return fmt.Sprintf("fg: %v, bg: %v, ef %v", a.fg, a.bg, a.ef) }

// todo - create a pool of stringbuilders that can go when ready?
// type sbSync struct {
// 	strings.Builder
// 	mu sync.Mutex
// }
