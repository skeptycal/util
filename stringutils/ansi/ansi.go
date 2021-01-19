// Copyright (c) 2020 Michael Treanor
// MIT License

// Package ansi provides fast ansi escape sequence processing based on strings.Builder.
// The standard is defined by the ECMA-48 standard "Control Functions for Coded Character Sets - Fifth Edition"
package ansi

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type color = byte
type ansiTypes = map[string]string

var types ansiTypes = ansiTypes{
	"basic":  "\033[%vm",
	"bright": "\033[1;%vm",
	"dim":    "\033[2;%vm",
	"256FG":  "\033[38;5;%vm",
	"256BG":  "\033[48;5;%vm",
	"24FG":   "\033[38;2;%v;%v;%vm",
	"24BG":   "\033[48;2;%v;%v;%vm",
}

var (
	DefaultioWriter io.Writer = os.Stdout
	a               ANSI      = NewANSIWriter(DefaultioWriter)

	// Reset Effects; Default Foreground; Default Background
	AnsiResetString string = BuildAnsi(DefaultForeground, DefaultBackground, Normal)
)

func NewAnsiSet(fg, bg, ef color) *AnsiSet {
	return &AnsiSet{fg, bg, ef}
}

type AnsiSet struct {
	fg color
	bg color
	ef color
}

func (a *AnsiSet) info() string {
	return fmt.Sprintf("fg: %v, bg: %v, ef %v", a.fg, a.bg, a.ef)
}

func (a *AnsiSet) String() string {
	return fmt.Sprintf(FMTansiSet, a.fg, a.bg, a.ef)
}

// todo - create a pool of stringbuilders that can go when ready?
// type sbSync struct {
// 	strings.Builder
// 	mu sync.Mutex
// }

// BuildAnsi returns a basic (3/4 bit) ANSI format code
// from a variadic argument list of bytes
func BuildAnsi(b ...byte) string {
	sb := strings.Builder{}
	defer sb.Reset()
	for _, n := range b {
		sb.WriteString(fmt.Sprintf(FMTansi, n))
	}
	return sb.String()
}
