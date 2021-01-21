// Copyright (c) 2020 Michael Treanor
// MIT License

// Package ansi provides fast ansi escape sequence processing based on strings.Builder.
// The standard is defined by the ECMA-48 standard "Control Functions for Coded Character Sets - Fifth Edition"
package ansi

import (
	"fmt"
)

type color = byte

type fbType = byte

const (
	foreground fbType = 3
	background fbType = 4
)

type AnsiStyle byte

const (
	StyleNormal AnsiStyle = iota
	StyleBold
	StyleAnsi8bit
	StyleItalics
	StyleUnderline
	StyleAnsi24bit
	StyleBlink
	StyleInverse
	StyleConceal
	StyleStrikeout
)

func NewAnsiSet(depth AnsiStyle) AnsiSet {
    a := ansiSetType{
        depth: depth,
        format: styleFormat[depth],
    }
    switch depth {
    case StyleAnsi8bit:
        return &ansi8{a}
    case StyleAnsi24bit:
        return &ansi24{a}
    default:
        return &ansiBasic{a}
    }
}

type AnsiSet interface {
	String() string
	BG() string
	FG() string
    SetColors(fg, bg, ef color)
    Info() string
    Output() string
}

type ansiBasic struct {ansiSetType}
type ansi8 struct   {ansiSetType}
type ansi24 struct   {ansiSetType}

func (a *ansiBasic) String() string {return a.out}
func (a *ansi8) String() string {return a.out}
func (a *ansi24) String() string { return a.out }

type ansiSetType struct {
    depth AnsiStyle
    format string
	fg    string
	bg    string
	ef    string
	out   string
}

func (a ansiSetType) BG() string     { return a.bg }
func (a ansiSetType) FG() string     { return a.fg }
func (a ansiSetType) Info() string { return fmt.Sprint(a.fg) }
func (a ansiSetType) Output() string { return fmt.Sprintf("%v;%v;%v", a.ef, a.fg, a.bg) }

var styleFormat = map[AnsiStyle]string{
	StyleNormal: "\x1b[%v%vm",
    StyleAnsi8bit:   "\x1b[%v8;5;%vm", // [38;5;${ID}m
	StyleAnsi24bit:    "\x1b[%%v8;2;%vm",
}

// SetColors creates printable strings for each of the ansi effects
/* styleFormat:

StyleNormal:
    "\x1b[%v%vm"
StyleAnsi8bit:
    "\x1b[%v8;5;%vm"
StyleAnsi24bit:
    "\x1b[%v8;2;%vm"
*/
func (a ansiSetType) SetColors(fg, bg, ef color) {
    // a.format  varieties
    /*
        StyleNormal: "\x1b[%v%vm",
        StyleAnsi8bit:   "\x1b[%v8;5;%vm", // [38;5;${ID}m
        StyleAnsi24bit:    "\x1b[%%v8;2;%vm",
    */
    a.format = styleFormat[a.depth]

    a.fg = fmt.Sprintf(a.format, foreground, fg&BasicMask)

    fmt.Printf("fg, a.fg: %q, %q\n",fg, a.fg)

	a.bg = fmt.Sprintf(a.format, background ,bg&BasicMask)
    a.ef = fmt.Sprintf(FMTansi, ef)

	o := fmt.Sprintf("%v;3%v;4%v", ef, fg&BasicMask, bg&BasicMask)

    a.out = fmt.Sprintf(a.format, o)
}


// todo - create a pool of stringbuilders that can go when ready?
// type sbSync struct {
// 	strings.Builder
// 	mu sync.Mutex
// }
