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

var colorfuncs = map[string]func(fb fbType, c interface{}) string{
	"fmtBasic":  func(fb fbType, c interface{}) string { return fmt.Sprintf("\x1b[%v%vm", fb, c) },
	"fmtBright": func(fb fbType, c interface{}) string { return fmt.Sprintf("\x1b[1;%v%vm", fb, c) },
	"fmtDim":    func(fb fbType, c interface{}) string { return fmt.Sprintf("\x1b[2;%v%vm", fb, c) },
	"fmt256":    func(fb fbType, c interface{}) string { return fmt.Sprintf("\x1b[%v8;5;%vm", fb, c) },
	"fmt24":     func(fb fbType, c interface{}) string { return fmt.Sprintf("\x1b[%v8;2;%v;%v;%vm", fb, r, g, b) },
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

func (a *ansiSet) FG(c color) string {
	if c == 0 {
		c = a.fg
	}
	a.Func(foreground, c)
}

func (a *ansiSet) SetColors(fg, bg, ef color) {
	fg = fg
	bg = bg
	ef = ef
}

type ansiSet struct {
	Func func(fb fbType, c interface{}) string
	fg   color
	bg   color
	ef   color
}

func (a *ansiSet) info() string {
	return fmt.Sprintf("fg: %v, bg: %v, ef %v", a.fg, a.bg, a.ef)
}

func (a *ansiSet) String() string {
	// return fmt.Sprintf(FMTansiSet, a.fg, a.bg, a.ef)
	return a.Func()
}

// todo - create a pool of stringbuilders that can go when ready?
// type sbSync struct {
// 	strings.Builder
// 	mu sync.Mutex
// }
