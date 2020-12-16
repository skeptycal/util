package anansi

// Anansi - ANSI color for terminal output
// Copyright (c) 2020 Michael Treanor
// MIT License

import (
	"os"
	"strconv"
	"sync"

	"github.com/mattn/go-colorable"
	"github.com/mattn/go-isatty"
)

var ( // globally available variables
	// NoColor defines if the output is colorized or not. It's dynamically set to
	// false or true based on the stdout's file descriptor referring to a terminal
	// or not. This is a global option and affects all colors. For more control
	// over each color block use the methods DisableColor() individually.
	NoColor = os.Getenv("TERM") == "dumb" ||
		(!isatty.IsTerminal(os.Stdout.Fd()) && !isatty.IsCygwinTerminal(os.Stdout.Fd()))

	// Output defines the standard output of the print functions. By default
	// os.Stdout is used.
	Output = colorable.NewColorableStdout()

	// Error defines a color supporting writer for os.Stderr.
	Error = colorable.NewColorableStderr()

	// colorsCache is used to reduce the count of created Color objects and
	// allows to reuse already created objects with required Attribute.
	colorsCache   = make(map[Attribute]*Color)
	colorsCacheMu sync.Mutex // protects colorsCache
)

// Color defines a custom color object which is defined by SGR parameters.
type Color struct {
	params  []Attribute
	noColor *bool
}

// ScreenWidth returns the cli screen width in 'characters' or 'columns'
func ScreenWidth() int {
	i, _ := strconv.Atoi(os.Getenv("COLUMNS"))
	return i
}

// String formatting constants
const (
	escape = "\x1b"

	// ANSI escape sequence for ANSI 16 color set. variables are: constants for 16 color sets (below)
	fmtAnsi = "\x1b[%dm"

	// ANSI escape sequence for ANSI 16 color 'bright' colors. Use the regular (non 'Hi') constants.
	fmtAnsiBright = "\x1b[%d;1m"

	// ANSI escape sequence for 256 colors. variables are: 3 for fg / 4 for bg, color 0-255
	// equivalent to BASH "\u001b[38;5;${ID}m"
	fmtAnsi256 = "\x1b[%s8;5;%dm"

	// ANSI escape sequence for 256 foreground colors; color 0-255
	fmtAnsi256Fg = "\x1b[38;5;%dm"

	// ANSI escape sequence for 256 foreground colors; color 0-255
	fmtAnsi256Bg = "\x1b[48;5;%dm"

	// foreground code for ANSI 256 color strings
	fg = "3"

	// background code for ANSI 256 color strings
	bg = "4"
)

// Color Styles
const (
	None int = iota
	Ansi8
	Ansi16
	Ansi256
	Hex
	RGB
	HSV
	HSL
	CMYK
)

// Base attributes
const (
	Reset Attribute = iota
	Bold
	Faint
	Italic
	Underline
	BlinkSlow
	BlinkRapid
	ReverseVideo
	Concealed
	CrossedOut
)

// Foreground text colors
const (
	FgBlack Attribute = iota + 30
	FgRed
	FgGreen
	FgYellow
	FgBlue
	FgMagenta
	FgCyan
	FgWhite
)

// Foreground Hi-Intensity text colors
const (
	FgHiBlack Attribute = iota + 90
	FgHiRed
	FgHiGreen
	FgHiYellow
	FgHiBlue
	FgHiMagenta
	FgHiCyan
	FgHiWhite
)

// Background text colors
const (
	BgBlack Attribute = iota + 40
	BgRed
	BgGreen
	BgYellow
	BgBlue
	BgMagenta
	BgCyan
	BgWhite
)

// Background Hi-Intensity text colors
const (
	BgHiBlack Attribute = iota + 100
	BgHiRed
	BgHiGreen
	BgHiYellow
	BgHiBlue
	BgHiMagenta
	BgHiCyan
	BgHiWhite
)
