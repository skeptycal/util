package anansi

// Anansi - ANSI color for terminal output
// Copyright (c) 2020 Michael Treanor
// MIT License

import (
	"strings"
)

// AnsiDefaults stores the initial ANSI default values
var AnsiDefaults = ansiCodes{}

func init() {

	AnsiDefaults.Set(
		FgWhite,
		BgBlack,
		Reset,
	)
}

//* --------------------------------------------------------> ansiCodes type definition

// ansiCodes defines a custom object which is defined by a set of 3 ANSI Attributes
type ansiCodes struct {
	fg Attribute
	bg Attribute
	ef Attribute
}

// SetAnsiDefaults updates the current ANSI color default values
func (a ansiCodes) Set(fg Attribute, bg Attribute, ef Attribute) {
	a.fg = fg
	a.bg = bg
	a.ef = ef
}

// AnsiDefault returns a string that resets all ANSI escape codes to default values
func (a ansiCodes) Reset() string {
	a = AnsiDefaults
	return a.String()
}

// ResetAnsiDefault prints a string that resets all ANSI escape codes to default values
func (a ansiCodes) Print() {
	Print(a.String())
}

// String returns a string that contains a set of 3 ANSI escape codes representing
// foreground color, background color, and effect
func (a ansiCodes) String() string {
	return strings.Join([]string{a.fg.Ansi(), a.bg.Ansi(), a.ef.Ansi()}, "")
}

// Wrap returns a string that wraps s in an ansiCodes color scheme and includes a trailing reset to AnsiDefaults
func (a ansiCodes) Wrap(s ...string) string {
	sb.Reset()
	defer sb.Reset()

	sb.WriteString(a.String())

	for _, p := range s {
		sb.WriteString(p)
	}

	sb.WriteString(AnsiDefaults.String())

	return sb.String()
}
