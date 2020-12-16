package anansi

// Anansi - ANSI color for terminal output
// Copyright (c) 2020 Michael Treanor
// MIT License

import (
	"fmt"
)

//* --------------------------------------------------------> Attribute type definition

// Attribute defines a single SGR Code
type Attribute int

// Ansi returns a basic (16 color) single code ANSI string matching Attribute:
//
// foreground: FgBlack, FgRed, FgGreen, FgYellow, FgBlue, FgMagenta, FgCyan, FgWhite, FgHiBlack, FgHiRed, FgHiGreen, FgHiYellow, FgHiBlue, FgHiMagenta, FgHiCyan, FgHiWhite
//
// background: BgBlack, BgRed, BgGreen, BgYellow, BgBlue, BgMagenta, BgCyan, BgWhite, BgHiBlack, BgHiRed, BgHiGreen, BgHiYellow, BgHiBlue, BgHiMagenta, BgHiCyan, BgHiWhite
//
// effects: Reset, Bold, Faint, Italic, Underline, BlinkSlow, BlinkRapid, ReverseVideo, Concealed, CrossedOut
//
func (a Attribute) Ansi() string {
	return fmt.Sprintf(fmtAnsi, a)
}

// Bright returns  a basic (16 color) single code 'bright' ANSI string matching Attribute:
func (a Attribute) Bright() string {
	return fmt.Sprintf(fmtAnsiBright, a)
}

// \u001b[37;1m

// A256 returns a string that represents an ANSI 256 color value
// fb is a foreground / background choice - use constants fg and bg
//
// codes 0 - 231 are colors; 232 - 255 are greyscale
func (a Attribute) A256(fb string) string {
	return fmt.Sprintf(fmtAnsi256, fb, a)
}

// Fg returns a string that represents an ANSI 256 foreground color value
//
// codes 0 - 231 are colors; 232 - 255 are greyscale
func (a Attribute) Fg() string {
	return fmt.Sprintf(fmtAnsi256Fg, a)
}

// Bg returns a string that represents an ANSI 256 background color value
//
// codes 0 - 231 are colors; 232 - 255 are greyscale
func (a Attribute) Bg() string {
	return fmt.Sprintf(fmtAnsi256Bg, a)
}
