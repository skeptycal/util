package anansi

import "strings"

// Anansi - ANSI color for terminal output
// Copyright (c) 2020 Michael Treanor
// MIT License

//* --------------------------------------------------------> AString type definition

// AString represents an ANSI color formatted string builder.
type AString interface {
	Fg()
	Bg()
	Ef()
	Reset()
	Clear()
	String()
}

type aString struct {
	ansiType int
	sb       *strings.Builder
}

// Fg adds a foreground ANSI 256 color code to the string
func (a *aString) Fg(code Attribute) {
	a.sb.WriteString(code.Fg())
}

// Bg adds a background ANSI 256 color code to the string
func (a *aString) Bg(code Attribute) {
	a.sb.WriteString(code.Bg())
}

// Ef adds an ANSI code to the string
func (a *aString) Ef(code Attribute) {
	a.sb.WriteString(code.Ansi())
}

// Reset resets all ANSI escape codes
func (a *aString) Default() {
	a.sb.WriteString(Attribute(Reset).Ansi())
}

// Clear resets the string to use zero bytes so that it may be garbage-collected.
func (a *aString) Clear() {
	a.sb.Reset()
}

// String returns the complete string buffer
func (a *aString) String() string {
	return a.sb.String()
}
