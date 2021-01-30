// Copyright 2020 Michael Treanor. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package stringutils implements additional functions to support the go library:
// .../go/1.15.3/libexec/src/strings/strings.go
//
// For information about UTF-8 strings in Go, see https://blog.golang.org/strings.
package stringutils

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)


var reWhitespace = regexp.MustCompile(`\s+`)

func DedupeWhitespace(s string) string {
	return reWhitespace.ReplaceAllString(strings.TrimSpace(s), " ")
}

// const alphanumerics = "0123456789abcdefghijklmnopqrstuvwxyz"

// Stringer implements the fmt.Stringer interface (for clarity)
type Stringer interface {
	String() string
}

func ToString(any interface{}) string {
	if v, ok := any.(Stringer); ok {
		return v.String()
	}
	switch v := any.(type) {
	case int:
		return strconv.Itoa(v)
	case float64, float32:
		return fmt.Sprintf("%.2g", v)
	}
	return "???"
}

// Contains tells whether a contains x.
func Contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

func IsDigit(c byte) bool {
	if c >= '0' && c <= '9' {
		return true
	}
	return false
}

func IsAlpha(c byte) bool {
	if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') {
		return true
	}
	return false
}

func IsAlphaSwitch(c byte) bool {
	switch {
	case c >= 'a' && c <= 'z':
		return true
	case c >= 'A' && c <= 'Z':
		return true
	case c >= '0' && c <= '9':
		return true
	default:
		return false
	}
}

func IsAlphaNum2(c uint8) bool {
	return 'a' <= c && c <= 'z' || '0' <= c && c <= '9' || 'A' <= c && c <= 'Z' || c == '_'
}

// IsAlphaNum reports whether the byte is an ASCII letter, number, or underscore
func IsAlphaNum(c byte) bool {

	if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9') {
		return true
	}
	return false
}

func isWhiteSpace2(c rune) bool {
	switch c {
	case ' ':
		return true
	case '\t', '\n', '\f', '\r', '\v':
		return true
	}
	return false
}

func IsWhiteSpace(c rune) bool {
	return c == ' ' || c == '\t' || c == '\n' || c == '\f' || c == '\r' || c == '\v'
}



func isWhiteSpace3(c rune) bool {
    return unicode.IsSpace(c)
}

var spaceMap4 = map[rune]string{
	0x0020: "SPACE",
	0x00A0: "NO-BREAK SPACE",
	0x1680: "OGHAM SPACE MARK",
	0x2000: "EN QUAD",
	0x2001: "EM QUAD",
	0x2002: "EN SPACE",
	0x2003: "EM SPACE",
	0x2004: "THREE-PER-EM SPACE",
	0x2005: "FOUR-PER-EM SPACE",
	0x2006: "SIX-PER-EM SPACE",
	0x2007: "FIGURE SPACE",
	0x2008: "PUNCTUATION SPACE",
	0x2009: "THIN SPACE",
	0x200A: "HAIR SPACE",
	0x202F: "NARROW NO-BREAK SPACE",
	0x205F: "MEDIUM MATHEMATICAL SPACE",
	0x3000: "IDEOGRAPHIC SPACE",
}

func isWhiteSpace4(c rune) bool {

	if _, ok := spaceMap4[c]; ok {
		return true
	}
	return false
}

// In computer programming, whitespace is any character or series of
// characters that represent horizontal or vertical space in typography.
// When rendered, a whitespace character does not correspond to a visible
// mark, but typically does occupy an area on a page. For example, the
// common whitespace symbol U+0020 SPACE (also ASCII 32) represents a
// blank space punctuation character in text, used as a word divider in
// Western scripts.
var spaceMap5 = map[rune]bool{
    0x0009: true, // CHARACTER TABULATION <TAB>
    0x000A: true, // ASCII LF
    0x000B: true, // LINE TABULATION <VT>
    0x000C: true, // FORM FEED <FF>
    0x000D: true, // ASCII CR
	0x0020: true, // SPACE <SP>
    0x00A0: true, // NO-BREAK SPACE <NBSP>
    0x0085: true, // NEL; Next Line
	0x1680: true, // Ogham space mark, interword separation in Ogham text
	0x2000: true, // EN QUAD, 0x2002 is preferred
	0x2001: true, // EM QUAD, mutton quad, 0x2003 is preferred
	0x2002: true, // EN SPACE, "nut", &ensp, LaTeX: '\enspace'
	0x2003: true, // EM SPACE, "mutton", &emsp;, LaTeX: '\quad'
	0x2004: true, // THREE-PER-EM SPACE, "thick space", &emsp13;
	0x2005: true, // four-per-em space, "mid space", &emsp14;
	0x2006: true, // SIX-PER-EM SPACE, sometimes equated to U+2009
	0x2007: true, // FIGURE SPACE, width of monospaced char, &numsp;
	0x2008: true, // PUNCTUATION SPACE, width of period or comma, &puncsp;
	0x2009: true, // THIN SPACE, 1/5th em, thousands sep, &thinsp;; LaTeX: '\,'
    0x200A: true, // HAIR SPACE, &hairsp;
    0x2028: true, // LINE SEPARATOR
    0x2029: true, // PARAGRAPH SEPARATOR
	0x202F: true, // NARROW NO-BREAK SPACE
	0x205F: true, // MEDIUM MATHEMATICAL SPACE, MMSP, &MediumSpace, 4/18 em
    0x3000: true, // IDEOGRAPHIC SPACE, full width CJK character cell
    0xFFEF: true, // ZERO WIDTH NO-BREAK SPACE <ZWNBSP> (BOM), deprecated Unicode 3.2 (use U+2060)
    // Related Unicode characters property White_Space=no
    0x180E: true, // MONGOLIAN VOWEL SEPARATOR, not whitespace as of Unicode 6.3.0
    0x200B: true, // ZERO WIDTH SPACE, ZWSP, "soft hyphen", &ZeroWidthSpace;
    0x200C: true, // ZERO WIDTH NON-JOINER, ZWNJ, &zwnj;
    0x200D: true, // ZERO WIDTH JOINER, ZWJ, &zwj;
    0x2060: true, // WORD JOINER, WJ, not a line break, &NoBreak;
}

func isWhiteSpace5(c rune) bool {

	if v, ok := spaceMap5[c]; ok && v {
		return true
	}
	return false

	// return unicode.IsSpace(c)
}

func RuneInfo(c rune) {
	s := "日本語"
	fmt.Printf("Glyph:   %q\n", s)
	fmt.Printf("UTF-8:   [% x]\n", []byte(s))
	fmt.Printf("Unicode: %U\n", []rune(s))
}
