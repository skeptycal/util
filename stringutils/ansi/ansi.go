// Copyright (c) 2020 Michael Treanor
// MIT License

// Package ansi provides fast ansi escape sequence processing based on strings.Builder.
// The standard is defined by the ECMA-48 standard "Control Functions for Coded Character Sets - Fifth Edition"
package ansi

import (
	"fmt"
	"strings"

	"github.com/skeptycal/util/stringutils"
)

const (
	// Clear screen
	ansiCLS = "\033[2J"
	// Character used for HR function
	HrChar string = "="
	// Mask to return only final nibble
	BasicMask byte = 0xF
	// Mask to return all except final nibble
	MSNibbleMask byte = 0xF0
)

type ansiCode struct {
	effect     byte
	foreground byte
	background byte
}

// set ANSI 8-bit (256 color) foreground and background color with leading effect
const (
	ansiFmtPrefix = "\033["
	ansiFmtDelim  = ";"
	ansiFmtSuffix = "m"
	ansiEncode    = "\033[%d;38;5;%d;48;5;%dm"
	ansiFgEncode  = "\033[38;5;%dm"
)

func (c *ansiCode) String() string {
	sb := strings.Builder{}
	defer sb.Reset()

	sb.WriteString(ansiFmtPrefix)

	if c.effect != 0 {
		sb.WriteByte(c.effect + 65)
	}
	return fmt.Sprintf(ansiEncode, c.effect, c.foreground, c.background)
}

func ByteToNum(c byte) byte {
	if stringutils.IsDigit(c) {

	}
	return c
}

func AnsiCode(e, f, b byte) fmt.Stringer {
	return &ansiCode{
		effect:     e,
		foreground: f,
		background: b,
	}
}

type AnsiString struct {
	color ansiCode
	string
}

func (s AnsiString) String() string {
	return fmt.Sprintf("%d%s%d")
}

// Print wraps args in an ANSI 8-bit color (256 color codes)
func Print(i ansiByte, args ...interface{}) {
	fmt.Print(ansiString(i, args...))
	fmt.Print(args...)
	fmt.Print(Reset)
}

// Println wraps args in an ANSI 8-bit color (256 color codes)
// and adds a newline character
func Println(i ansiByte, args ...interface{}) {
	fmt.Print(ansiString(i))
	fmt.Print(args...)
	fmt.Println(Reset)
}
