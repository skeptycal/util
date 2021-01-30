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
	"unicode/utf8"
)

const (
    TAB = 0x09  // '\t'
    LF = 0x0A   // '\n'
    VT = 0x0B   // '\v'
    FF = 0x0C   // '\f'
    CR = 0x0D   // '\r'
    SPACE = ' '
    RuneSelf = utf8.RuneSelf
    NBSP = 0x00A0
    NEL = 0x0085
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

// IsASCIISpace tests for the most common ASCII whitespace characters:
//  ' ', '\t', '\n', '\f', '\r', '\v', U+0085 (NEL), U+00A0 (NBSP)
func IsASCIISpace(c byte) bool {
	return c == ' ' || c == '\n' || c == '\t' || c == '\r' || c == '\f' || c == '\v' || c == 0x0085 || c == 0x00A0
}

// IsAnySpace reports whether the rune is any utf8 whitespace character
// using the broadest and most complete definition.
//
// The speed of this implementation ~25% slower than that of
// IsASCIISpace(c byte) but tests 3.75 times more possible code points.
//
// The speed is ~7% faster than that of unicode.IsSpace(r rune) from the
// standard library and covers nearly twice as many code points.
func IsUnicodeWhiteSpace(r rune) bool {
	if r < unicode.MaxLatin1 {
		return r == ' ' || (r > 8 && r < 14) || r == 0x85 || r == 0xA0
    }

	if _, ok := UnicodeWhiteSpaceMap[r]; ok {
		return true
	}
	return false
}

// IsAnySpace reports whether the rune is any utf8 whitespace character
// using the broadest and most complete definition.
//
// The speed of this implementation is 1.84 times slower than that of
// IsASCIISpace(c byte) but tests 3.75 times more possible code points.
//
// The speed is 28% slower than that of unicode.IsSpace(r rune) from the
// standard library.
//
// IsAnySpace tests nearly twice as many code points as unicode.IsSpace().
func IsAnySpace(r rune) bool {
	if uint32(r) <= unicode.MaxLatin1 {
		switch r {
		case '\t', '\n', '\v', '\f', '\r', ' ', 0x85, 0xA0:
			return true
		}
		return false
	}
	if _, ok := UnicodeWhiteSpaceMap[r]; ok {
		return true
	}
	return false
}

func isWhiteSpace(c byte) bool {
	return c == ' ' || c == '\t' || c == '\n' || c == '\f' || c == '\r' || c == '\v'
}

func isWhiteSpace2(c byte) bool {
	switch c {
	case ' ':
		return true
	case '\t', '\n', '\f', '\r', '\v':
		return true
	}
	return false
}

func isWhiteSpace3(c rune) bool {
	return unicode.IsSpace(c)
}



func isWhiteSpace4(c rune) bool {
	if _, ok := UnicodeWhiteSpaceMap[c]; ok {
		return true
	}
	return false
}

func isWhiteSpaceBoolMap(c rune) bool {
	if _, ok := unicodeWhiteSpaceMapBool[c]; ok {
		return true
	}
	return false
}

func isWhiteSpaceLogicChain(r rune) bool {
    if r < unicode.MaxLatin1 {
		return r == ' ' || r == '\t' || r == '\n' || r == '\f' || r == '\r' || r == '\v' || r == 0x85 || r == 0xA0
    }

    // 0x2000, 0x2001, 0x2002, 0x2003, 0x2004, 0x2005, 0x2006,
    // 0x2007, 0x2008, 0x2009, 0x200A,
    // 0x2028, 0x2029, 0x202F, 0x205F, 0x3000,
    return (r >= 0X2000 && r <= 0X200A) || r == 0x2028 || r == 0x2029 || r == 0x202F || r == 0x205F || r == 0x3000

    // Related Unicode characters property White_Space=no
    // 0x200B,	0x200C,	0x200D,	0x2060, 0x1680,
    // || r == 0x200B || r == 0x200C || r == 0x200D || r == 0x2060 || r == 0x1680

}

func RuneInfo(c rune) {
	s := "日本語"
	fmt.Printf("Glyph:   %q\n", s)
	fmt.Printf("UTF-8:   [% x]\n", []byte(s))
	fmt.Printf("Unicode: %U\n", []rune(s))
}

// In computer programming, whitespace is any character or series of
// characters that represent horizontal or vertical space in typography.
// When rendered, a whitespace character does not correspond to a visible
// mark, but typically does occupy an area on a page. For example, the
// common whitespace symbol U+0020 SPACE (also ASCII 32) represents a
// blank space punctuation character in text, used as a word divider in
// Western scripts.
//
// Reference: https://en.wikipedia.org/wiki/Whitespace_character
var UnicodeWhiteSpaceMap = map[rune]string{
	0x0009: `CHARACTER TABULATION <TAB>`,
	0x000A: `ASCII LF`,
	0x000B: `LINE TABULATION <VT>`,
	0x000C: `FORM FEED <FF>`,
	0x000D: `ASCII CR`,
	0x0020: `SPACE <SP>`,
	0x00A0: `NO-BREAK SPACE <NBSP>`,
	0x0085: `NEL; Next Line`,
	0x1680: `Ogham space mark, interword separation in Ogham text`,
	0x2000: `EN QUAD, 0x2002 is preferred`,
	0x2001: `EM QUAD, mutton quad, 0x2003 is preferred`,
	0x2002: `EN SPACE, "nut", &ensp, LaTeX: '\enspace'`,
	0x2003: `EM SPACE, "mutton", &emsp;, LaTeX: '\quad'`,
	0x2004: `THREE-PER-EM SPACE, "thick space", &emsp13;`,
	0x2005: `four-per-em space, "mid space", &emsp14;`,
	0x2006: `SIX-PER-EM SPACE, sometimes equated to U+2009`,
	0x2007: `FIGURE SPACE, width of monospaced char, &numsp;`,
	0x2008: `PUNCTUATION SPACE, width of period or comma, &puncsp;`,
	0x2009: `THIN SPACE, 1/5th em, thousands sep, &thinsp;; LaTeX: '\,'`,
	0x200A: `HAIR SPACE, &hairsp;`,
	0x2028: `LINE SEPARATOR`,
	0x2029: `PARAGRAPH SEPARATOR`,
	0x202F: `NARROW NO-BREAK SPACE`,
	0x205F: `MEDIUM MATHEMATICAL SPACE, MMSP, &MediumSpace, 4/18 em`,
	0x3000: `IDEOGRAPHIC SPACE, full width CJK character cell`,
	0xFFEF: `ZERO WIDTH NO-BREAK SPACE <ZWNBSP> (BOM), deprecated Unicode 3.2 (use U+2060)`,
	// Related Unicode characters property White_Space=no
	// 0x180E: `MONGOLIAN VOWEL SEPARATOR, not whitespace as of Unicode 6.3.0`,
	// 0x200B: `ZERO WIDTH SPACE, ZWSP, "soft hyphen", &ZeroWidthSpace;`,
	// 0x200C: `ZERO WIDTH NON-JOINER, ZWNJ, &zwnj;`,
	// 0x200D: `ZERO WIDTH JOINER, ZWJ, &zwj;`,
	// 0x2060: `WORD JOINER, WJ, not a line break, &NoBreak;`,
}

// In computer programming, whitespace is any character or series of
// characters that represent horizontal or vertical space in typography.
// When rendered, a whitespace character does not correspond to a visible
// mark, but typically does occupy an area on a page. For example, the
// common whitespace symbol U+0020 SPACE (also ASCII 32) represents a
// blank space punctuation character in text, used as a word divider in
// Western scripts.
var unicodeWhiteSpaceMapBool = map[rune]bool{
	0x0009: true, // CHARACTER TABULATION <TAB>
	0x000A: true, // ASCII LF
	0x000B: true, // LINE TABULATION <VT>
	0x000C: true, // FORM FEED <FF>
	0x000D: true, // ASCII CR
    0x0020: true, // SPACE <SP>
    // > utf8.RuneSelf
	0x00A0: true, // NO-BREAK SPACE <NBSP>
    0x0085: true, // NEL; Next Line
    // > unicode.MaxLatin1
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
	// 0x180E: true, // MONGOLIAN VOWEL SEPARATOR, not whitespace as of Unicode 6.3.0
	// 0x200B: true, // ZERO WIDTH SPACE, ZWSP, "soft hyphen", &ZeroWidthSpace;
	// 0x200C: true, // ZERO WIDTH NON-JOINER, ZWNJ, &zwnj;
	// 0x200D: true, // ZERO WIDTH JOINER, ZWJ, &zwj;
	// 0x2060: true, // WORD JOINER, WJ, not a line break, &NoBreak;
}
