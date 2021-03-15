// Copyright (c) 2020 Michael Treanor
// MIT License

// Package ansi provides fast ansi escape sequence processing based on strings.Builder.
// The standard is defined by the ECMA-48 standard "Control Functions for Coded Character Sets - Fifth Edition"
package ansi

import (
	"fmt"
)

const (
	// Character used for HR function
	HrChar string = "="
	// Mask to return only final nibble
	BasicMask byte = 0xF
	// Mask to return all except final nibble
	MSNibbleMask byte = 0xF0
)

const (
	ansiEncode = "\033[38;5;%dm" // set ANSI foreground color to code %d using printf
	ansiReset  = "\033[39;49;0m" // reset ANSI terminal output to default foreground and background colors
	// PrintColor = "\033[38;5;%dm%s\033[39;49m\n" // wraps text %s in a color %d
)

func ansiString(i int) string {
	return fmt.Sprintf(ansiEncode, i)
}

func Cprint(i int, args ...interface{}) {
	fmt.Print(ansiString(i))
	fmt.Print(args...)
	fmt.Print(ansiReset)
}

func Cprintln(i int, args ...interface{}) {
	fmt.Print(ansiString(i))
	fmt.Print(args...)
	fmt.Println(ansiReset)
}
