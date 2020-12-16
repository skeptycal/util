package anansi

// Anansi - ANSI color for terminal output
// Copyright (c) 2020 Michael Treanor
// MIT License this
//
// Most of the available packages expose objects as part of the API. I find exporting and manipulating strings directly to be much simple and cleaner for the most part.
//
// Terminal output is a stream of bytes. The most efficient and flexible ways to handle them has long been using bytes.Buffer and converting to strings as needed. The strings.Builder implementation improved on this method in many ways.
//
// At any given time, only one character is 'printing' to the terminal output. Instead of thinking of the text as chunks and processing each piece of text as an object, we can think of all CLI text as a stream of characters that always have these ANSI color  attributes. They are stored somewhere and applied all the time. At any point along the stream, we can insert ANSI escape codes into the strings to change the stored attributes. This changes the current behavior and the behavior remains until it is changed.
//
// The 'anansi.go' file contains the majority of original code to focus more on inserting strings into the stream.
// The 'anansi_codes.go' file contains the new strings.Builder implementation for ANSI code sets
// The 'anansi_attribute.go' file contains the updated strings.Builder implementation of the Attribute type
// Much of the basic types, constants, and terminal checks in 'anansi_const.go' are based on the very popular and well documented color project
// https://github.com/fatih/color (MIT License)

import (
	"fmt"
	"strconv"
	"strings"
)

// stringBuilder string builder for temporary use by any function .... definitely not thread safe
//
// It's ready to use from the get-go.
// You don't need to initialize it.
var sb strings.Builder

// Ansi returns a basic (16 color) single code ansi string matching color without allocating an Attribute variable in the parent scope
//
// foreground: FgBlack, FgRed, FgGreen, FgYellow, FgBlue, FgMagenta, FgCyan, FgWhite, FgHiBlack, FgHiRed, FgHiGreen, FgHiYellow, FgHiBlue, FgHiMagenta, FgHiCyan, FgHiWhite
//
// background: BgBlack, BgRed, BgGreen, BgYellow, BgBlue, BgMagenta, BgCyan, BgWhite, BgHiBlack, BgHiRed, BgHiGreen, BgHiYellow, BgHiBlue, BgHiMagenta, BgHiCyan, BgHiWhite
//
// effects: Reset, Bold, Faint, Italic, Underline, BlinkSlow, BlinkRapid, ReverseVideo, Concealed, CrossedOut
//
func Ansi(color Attribute) string {
	return Attribute(color).Ansi()
}

// AllAnsi returns a complete ANSI color string including foreground, background
// and effect without allocating an ansiCodes variable in the parent scope
func AllAnsi(fg Attribute, bg Attribute, ef Attribute) string {
	return ansiCodes{fg, bg, ef}.String()
}

// ResetAll resets all ANSI settings to default values
func ResetAll() string {
	return AnsiDefaults.Reset()
}

// Print prints a variable number of ANSI formatted strings to Output
func Print(s ...string) {
	sb.Reset()
	defer sb.Reset()

	for _, p := range s {
		sb.WriteString(p)
	}

	fmt.Fprint(Output, sb.String())
}

// Println prints an ANSI formatted string to Output and resets the
// Output to the default values stored in AnsiDefaults
func Println(s ...string) {
	sb.Reset()
	defer sb.Reset()

	for _, p := range s {
		sb.WriteString(p)
	}
	sb.WriteString(AnsiDefaults.Reset())

	fmt.Fprintln(Output, sb.String())
}

// PadInt converts <i> to a string and adds padding to the left of <i> to make it <size> length
func PadInt(i int, size int) string {
	sb.Reset()
	defer sb.Reset()

	t := strconv.Itoa(i)

	pad := size - len(t)

	if pad > 0 {
		sb.WriteString(strings.Repeat(" ", pad))
	}
	sb.WriteString(t)

	return sb.String()
}

func contrastColor(i int) int {
	if i < 16 {
		return 0
	} else if i > 231 {
		return 232 + (255 - i)
	} else {
		return 249 - ((i - 16) % 18)
	}
}

// Sample prints a sample of all color combinations
func Sample() {
	sb.Reset()
	defer sb.Reset()

	sb.WriteString(Attribute(Reset).Ansi())

	for i := 0; i < 256; i++ {
		sb.WriteString(Attribute(contrastColor(i)).A256(bg))
		sb.WriteString(Attribute(i).A256(fg))
		sb.WriteString(" ")
		sb.WriteString(PadInt(i, 5))
		sb.WriteString(" ")
	}

	sb.WriteString(Attribute(Reset).Ansi())
	sb.WriteString("\n\n")

	for i := 0; i < 256; i++ {
		sb.WriteString(Attribute(contrastColor(i)).A256(fg))
		sb.WriteString(Attribute(i).A256(bg))
		sb.WriteString(" ")
		sb.WriteString(PadInt(i, 5))
		sb.WriteString(" ")
	}
	sb.WriteString(AnsiDefaults.String())
	sb.WriteString("\n\n")

	fmt.Fprint(Output, sb.String())
}
