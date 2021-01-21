// Copyright (c) 2020 Michael Treanor
// MIT License

package ansi

import (
	"fmt"
)

// AnsiSet implements the Ansi standard in strings. There are methods
// to wrap strings, encode streams, and unmarshal data structures.
//
type AnsiSet interface {
	String() string
	BG() string
	FG() string
    SetColors(fg, bg, ef color)
    Info() string
    Output() string
}

// NewAnsiSet returns a new AnsiSet with the specified features.
//
// The fg, bg, and ef values are the defaults that are used any
// time no other features are specified.
//
// The type of Ansi output is determined by the 'depth' argument.
// It can be
//  StyleNormal (the default)
// standard original 3/4 bit Ansi style with ranges
// of 8 named colors. e.g.
//  Blue, YellowBackground, Bold Green
// or 8 bit Ansi colors
//  StyleAnsi8bit
// This set includes 216 colors (numbered 0 - 231)
// plus 24 grayscale shades (numbered 232 - 255)
// They are labeled fgxxx for foreground colors and bgxxx for background
//  fg127, bg250, bg57
// Finally, 24 bit truecolor values can be used in most terminals
//  StyleAnsi24bit
// These are encoded as RGB(xxx,xxx,xxx) values. Users could also
// implement a map of colors that can be updated, adjusted to match
// a photo or camera view, or simply change as the sun sets.
func NewAnsiSet(depth AnsiStyle, fg, bg, ef color) AnsiSet {
    a := &ansiSetType{
        depth: depth,
        format: styleFormat[depth],
    }
    switch depth {
    case StyleAnsi8bit:
        b := &ansi8{*a}
        b.SetColors(fg, bg, ef)
        return b
    case StyleAnsi24bit:
        b := &ansi24{*a}
        b.SetColors(fg, bg, ef)
        return b
    default:
        b := &ansiBasic{*a}
        b.SetColors(fg, bg, ef)
        return b
    }
}

type ansiSetType struct {
    depth AnsiStyle
    format string
	fg    byte
	bg    byte
	ef    byte
	out   string
}

func (a ansiSetType) BG() string     { return fmt.Sprintf(a.format,background,a.bg) }
func (a ansiSetType) FG() string     { return fmt.Sprintf(a.format,foreground,a.fg) }
func (a ansiSetType) Info() string {  return fmt.Sprintf("fg: %q;bg: %q; ef: %q",  a.fg, a.bg,a.ef) }
func (a ansiSetType) Output() string { return fmt.Sprintf("%v;%v;%v", a.ef, a.fg, a.bg) }
func (a ansiSetType) SetColors(fg, bg, ef color) {
    a.fg = fg
    a.bg = bg
    a.ef = ef
    // not implemented
}

var styleFormat = map[AnsiStyle]string{
	StyleNormal: "\x1b%vm",
    StyleAnsi8bit:   "\x1b[%v8;5;%vm", // [38;5;${ID}m
	StyleAnsi24bit:    "\x1b[%v8;2;%vm",
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


type ansiBasic struct {ansiSetType}
func (a ansiBasic) BG() string     { return fmt.Sprintf(a.format,a.bg) }
func (a ansiBasic) FG() string     { return fmt.Sprintf(a.format,a.fg) }
func (a ansiBasic) String() string {return a.out}

// SetColors configures the ansi object according to the specified
// ANSI set
//
// struct items are set in the following way:
//      depth  AnsiStyle - set in the constructor
//      format string - copied from the styleFormat map
//      fg, bg, ef     byte - stored from args
//      out    string   // made from fg() BG(), and a
//                                  generic ansi effects string
func (a ansiBasic) SetColors(fg, bg, ef color) {
    a.format = styleFormat[a.depth]

    a.fg = fg
	a.bg = bg
    a.ef = ef

    o := fmt.Sprintf("%v;%v;%v", a.FG(), a.BG(), BuildAnsi(a.ef))

    a.out = fmt.Sprintf(a.format, o)

    fmt.Printf("fg, a.fg: %q, %q\n",fg, a.fg)
}

type ansi8 struct   {ansiSetType}
// func (a ansi8) BG() string     { return a.ansiSetType.BG()}
// func (a ansi8) FG() string     { return fmt.Sprintf(a.format,a.fg) }
func (a ansi8) String() string {return a.out}


type ansi24 struct   {ansiSetType}

func (a *ansi24) String() string { return a.out }

// todo - create a pool of stringbuilders that can go when ready?
// type sbSync struct {
// 	strings.Builder
// 	mu sync.Mutex
// }



type color = byte

type fbType = byte

const (
	foreground fbType = 3
	background fbType = 4
)

type AnsiStyle color

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
