// Copyright 2020 Michael Treanor. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package stringutils implements additional functions to support the go library:
// .../go/1.15.3/libexec/src/strings/strings.go
//
// For information about UTF-8 strings in Go, see https://blog.golang.org/strings.
package stringutils

import (
	"bytes"
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
	"unsafe"
)

// Numbers fundamental to the encoding.
const (
	RuneError = '\uFFFD'     // the "error" Rune or "Unicode replacement character"
	RuneSelf  = 0x80         // characters below RuneSelf are represented as themselves in a single byte.
	MaxRune   = '\U0010FFFF' // Maximum valid Unicode code point.
	UTFMax    = 4            // maximum number of bytes of a UTF-8 encoded Unicode character.
)

// const alphanumerics = "0123456789abcdefghijklmnopqrstuvwxyz"

// Stringer implements the fmt.Stringer interface (for clarity)
type Stringer interface {
	String() string
}

// ToString implements Stringer directly as a function call
// with a parameter instead of a method on that parameter.
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

var (
	reWhitespace = regexp.MustCompile(`[\s\v]+`)

	shortByteList       = []byte{0x09, 0x0A, 0x0B, 0x0C, 0x0D, 0x20, 0x85, 0xA0}
	longRuneList        = []rune{0x09, 0x0A, 0x0B, 0x0C, 0x0D, 0x20, 0x85, 0xA0, 0x2000, 0x200A, 0x2028, 0x2029, 0x202F, 0x205F, 0x3000, 0xFFEF, 0x1680}
	shortByteListString = string(shortByteList)
	longRuneListString  = string(longRuneList)

	// UnicodeWhiteSpaceMap provides a mapping from Unicode runes to strings.
	// In computer programming, whitespace is any character or series of
	// characters that represent horizontal or vertical space in typography.
	// When rendered, a whitespace character does not correspond to a visible
	// mark, but typically does occupy an area on a page. For example, the
	// common whitespace symbol U+0020 SPACE (also ASCII 32) represents a
	// blank space punctuation character in text, used as a word divider in
	// Western scripts.
	//
	// Reference: https://en.wikipedia.org/wiki/Whitespace_character
	UnicodeWhiteSpaceMap = map[rune]string{
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

	// UnicodeWhiteSpaceMap provides a mapping from Unicode runes to bool{true}.
	// In computer programming, whitespace is any character or series of
	// characters that represent horizontal or vertical space in typography.
	// When rendered, a whitespace character does not correspond to a visible
	// mark, but typically does occupy an area on a page. For example, the
	// common whitespace symbol U+0020 SPACE (also ASCII 32) represents a
	// blank space punctuation character in text, used as a word divider in
	// Western scripts.
	unicodeWhiteSpaceMapBool = map[rune]bool{
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
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type line []byte // todo ?? use [255]byte or similar?? make it user choosable?

func (l *line) unsafeToStringPtr() string {
    return *(*string)(unsafe.Pointer(&l))
}

func (l *line) unsafeMakeFromPtr(s string) {
    l = (*line)(unsafe.Pointer(&s))
}

type list struct {
    buf *[][]byte
}

func (l list) Len() int { return len(*l.buf) }
func (l list) Cap() int { return cap(*l.buf) }
func (l list) Reset(cap int) { *l.buf = make([][]byte,0,cap+1) }
func (l list) Join() []byte { return bytes.Join(*l.buf, []byte{0x10}) }
func (l list) Make(s string) [][]byte {
    l.Reset(len(s))
    *l.buf = bytes.SplitAfter([]byte(s), []byte{0x10})
    return *l.buf
}
func (l list) Contains(b []byte) bool {
    for _, s := range *l.buf {
        if bytes.Index(s, b) == -1 {
            return false
        }
    }
    return true
}

// SplitNoSave - modified from standard library generic split:
// splits after each instance of sep.
func SplitNoSave(s, sep []byte) [][]byte {

    n := bytes.Count(s, sep) + 1
    sepSave := 0
	// if n == 0 {
	// 	return nil
	// }
	// if len(sep) == 0 {
	// 	return Explode(s, n)
	// }
	// if n < 0 {
	// 	n = bytes.Count(s, sep) + 1
	// }

	a := make([][]byte, n)
	n--
	i := 0
	for i < n {
		m := bytes.Index(s, sep)
		if m < 0 {
			break
		}
		a[i] = s[: m+sepSave : m+sepSave]
		s = s[m+len(sep):]
		i++
	}
	a[i] = s
	return a[:i+1]
}

func bEqual(s, sub []byte) int {
    if bytes.Equal(s, sub) {
        return 1
    }
    return 0
}

func bContains(s, sub []byte) int {
    if bytes.Contains(s, sub) {
        return 1
    }
    return 0
}

// Count counts the number of non-overlapping instances
// of sep in s. If sep is an empty slice, Count returns
// 1 + the number of UTF-8-encoded code points in s.
// if len(s) == 1, a fast processor specific implementation
// is used.
func Count(s, sep []byte) int {
	// special case
	if len(sep) == 0 {
		return utf8.RuneCount(s) + 1
	}
	if len(sep) == 1 {
		return bytes.Count(s, sep)
	}
	n := 0
	for {
		i := bytes.Index(s, sep)
		if i == -1 {
			return n
		}
		n++
		s = s[i+len(sep):]
	}
}

// Explode splits s into a slice of UTF-8 sequences, one per Unicode code point (still slices of bytes),
// up to a maximum of n byte slices. Invalid UTF-8 sequences are chopped into individual bytes.
func Explode(s []byte, n int) [][]byte {
	if n <= 0 {
		n = len(s)
	}
	a := make([][]byte, n)
	var size int
	na := 0
	for len(s) > 0 {
		if na+1 >= n {
			a[na] = s
			na++
			break
		}
		_, size = utf8.DecodeRune(s)
		a[na] = s[0:size:size]
		s = s[size:]
		na++
	}
	return a[0:na]
}


func Join(list []string) string {
    return strings.Join(list, "\n")
}

func TabIt(s string, n int) string {
    tmp := make([]string,0,strings.Count(s,"\n") + 3)
    for  _, line := range strings.Fields(s) {
        tmp = append(tmp, strings.Repeat(" ", n)+line)
    }
    return strings.Join(tmp,"\n")
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

// RuneInfo prints a sample of various Unicode runes.
func RuneInfo(c rune) {
	s := "日本語"
	fmt.Printf("Glyph:   %q\n", s)
	fmt.Printf("UTF-8:   [% x]\n", []byte(s))
	fmt.Printf("Unicode: %U\n", []rune(s))
}
