// Copyright (c) 2020 Michael Treanor
// MIT License

// Package ansi provides fast ansi escape sequence processing based on strings.Builder.
// The standard is defined by the ECMA-48 standard "Control Functions for Coded Character Sets - Fifth Edition"
package ansi

import (
	"fmt"
)

type color = byte

type fbType byte

const (
	foreground fbType = 3
	background fbType = 4
)

type ansiStyle byte

const (
	normal ansiStyle = iota
	bold
	ansi8bit
	italics
	underline
	ansi24bit
	blink
	inverse
	conceal
	strikeout
)



var colordepth = map[string]string{
	"fmtBasic": "\x1b[%v%v%m",
	"fmt256":   "\x1b[%v8;5;%%vm",
	"fmt24":    "\x1b[%v8;2;%%vm",
}

func NewAnsiSet(depth ansiStyle, fg, bg, ef color) *ansiSet {

    switch depth {
    case ansi8bit:
        a := ansi8bit{}
    case: ansi24bit:
        a := ansi24bit{}
    default:
        a := ansiSet{}
    }

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


func (a *ansiSet) BG() string     { return a.bg }
func (a *ansiSet) FG() string     { return a.fg }


func (a *ansiSet) SetColors(fg, bg, ef color) {
	o := fmt.Sprintf("%v;3%v;4%v", ef, fg&BasicMask, bg&BasicMask)

	a.fg = fmt.Sprintf(FMTansiFG, fg&BasicMask)
	a.bg = fmt.Sprintf(FMTansiBG, bg&BasicMask)
	a.ef = fmt.Sprintf(FMTansi, ef)
	a.out = fmt.Sprintf(FMTansi, o)
}
func (a *ansiSet) output() string { return fmt.Sprintf("%v;%v;%v", a.ef, a.fg, a.bg) }
func (a *ansiSet) SetStyle(style ansiStyle) {
	if style == 2 {
        // 8 bit colors
        a = NewAnsiSet(ansi8bits)
		a.depth = FMT8bit
	}
	if style == 5 {
		// 24 bit colors
		a.depth = FMT24bit
	}

	a.depth = fmt.Sprintf(FMTansi, style)
}
func (a *ansiSet) info() string { return fmt.Sprintf("fg: %v, bg: %v, ef %v", a.fg, a.bg, a.ef) }


type ansi8     ansiSet

func (a *ansi8) String() string {return a.out}
func (a *ansi8) String() string { return a.out }


type ansi24     ansiSet

func (a *ansi24) String() string {
    return a.out
}



// todo - create a pool of stringbuilders that can go when ready?
// type sbSync struct {
// 	strings.Builder
// 	mu sync.Mutex
// }
