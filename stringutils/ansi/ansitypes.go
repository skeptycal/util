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

var styleFormat = map[ansiStyle]string{
	normal: "\x1b[%v%%vm",
	ansi8bit:   "\x1b[%v8;5;%%vm",
	ansi24bit:    "\x1b[%v8;2;%%vm",
}

func NewAnsiSet(depth ansiStyle) AnsiSet {
    switch depth {
    case ansi8bit:
        return &ansi8{ansiSetType{depth:depth}}
    case ansi24bit:
        return &ansi24{}
    default:
        return &ansiBasic{}
    }
}

type AnsiSet interface {
	String() string
	BG() string
	FG() string
	SetColors(fg, bg, ef color)
    SetStyle(style ansiStyle)
    // Build(b ...byte) string
}

type ansiBasic struct {ansiSetType}
type ansi8 struct   {ansiSetType}
type ansi24 struct   {ansiSetType}

func (a *ansiBasic) String() string {return a.out}
func (a *ansi8) String() string {return a.out}
func (a *ansi24) String() string { return a.out }

type ansiSetType struct {
	depth ansiStyle
	fg    string
	bg    string
	ef    string
	out   string
}

func (a *ansiSetType) BG() string     { return a.bg }
func (a *ansiSetType) FG() string     { return a.fg }
func (a *ansiSetType) info() string { return fmt.Sprintf("fg: %v, bg: %v, ef %v", a.fg, a.bg, a.ef) }
func (a *ansiSetType) output() string { return fmt.Sprintf("%v;%v;%v", a.ef, a.fg, a.bg) }
func (a *ansiSetType) SetColors(fg, bg, ef color) {
	o := fmt.Sprintf("%v;3%v;4%v", ef, fg&BasicMask, bg&BasicMask)

	a.fg = fmt.Sprintf(FMTansiFG, fg&BasicMask)
	a.bg = fmt.Sprintf(FMTansiBG, bg&BasicMask)
	a.ef = fmt.Sprintf(FMTansi, ef)
	a.out = fmt.Sprintf(FMTansi, o)
}
func (a *ansiSetType) SetStyle(style ansiStyle) {
    a = NewAnsiSet(style).(ansiSetType)

	a.depth = fmt.Sprintf(FMTansi, style)
}

// todo - create a pool of stringbuilders that can go when ready?
// type sbSync struct {
// 	strings.Builder
// 	mu sync.Mutex
// }
