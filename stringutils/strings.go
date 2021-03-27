// Copyright 2020 Michael Treanor. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package stringutils implements additional functions to
// support the go standard library strings module.
//
// The algorithms chosen are based on benchmarks from
// the stringbenchmarks module. ymmv...
//
// The current implementation at the start of this project was
// .../go/1.15.3/libexec/src/strings/strings.go
//
// For information about UTF-8 strings in Go,
// see https://blog.golang.org/strings.
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

	"golang.org/x/sys/cpu"
)

// Numbers fundamental to the encoding.
const (
	RuneError = utf8.RuneError // '\uFFFD'       // the "error" Rune or "Unicode replacement character"
	RuneSelf  = utf8.RuneSelf  // 0x80           // characters below RuneSelf are represented as themselves in a single byte.
	MaxRune   = utf8.MaxRune   // '\U0010FFFF'   // Maximum valid Unicode code point.
	UTFMax    = utf8.UTFMax    // 4              // maximum number of bytes of a UTF-8 encoded Unicode character.
)

var MaxLen int

// const alphanumerics = "0123456789abcdefghijklmnopqrstuvwxyz"

// Stringer implements the fmt.Stringer interface (for clarity)
type Stringer interface {
	String() string
}

func init() {
	_ = cpu.ARM64.HasAES
	if cpu.X86.HasAVX2 {
		MaxLen = 63
	} else {
		MaxLen = 31
	}
}

// Cutover reports the number of failures of IndexByte we should tolerate
// before switching over to Index.
// n is the number of bytes processed so far.
// See the bytes.Index implementation for details.
func Cutover(n int) int {
	// 1 error per 8 characters, plus a few slop to start.
	return (n + 16) / 8
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
	// note: 0x0B is not considered whitespace by regex so an alternative
	// solution must be considered if 0x0B detection is required.
	//
	// regex is the slowest solution (by far!) so this is likely moot.
	reWhitespace = regexp.MustCompile(`[\s\v]+`)

	// leaves out 0x85, 0xA0
	shortASCIIList = []byte{0x09, 0x0A, 0x0B, 0x0C, 0x0D, 0x20}

	// all whitespace code points that are one byte long
	shortByteList = []byte{0x09, 0x0A, 0x0B, 0x0C, 0x0D, 0x20, 0x85, 0xA0}

	// most common unicode whitespace code points
	longRuneList = []rune{0x09, 0x0A, 0x0B, 0x0C, 0x0D, 0x20, 0x85, 0xA0, 0x2000, 0x200A, 0x2028, 0x2029, 0x202F, 0x205F, 0x3000, 0xFFEF, 0x1680}

	// byte lists transformed to strings
	shortASCIIListString = string(shortASCIIList)
	shortByteListString  = string(shortByteList)
	longRuneListString   = string(longRuneList)

	// UnicodeWhiteSpaceMap provides a mapping from Unicode runes to strings
	// with descriptions of each. It is marginally slower than the bool map.
	//
	// In computer programming, whitespace is any character or series of
	// characters that represent horizontal or vertical space in typography.
	// When rendered, a whitespace character does not correspond to a visible
	// mark, but typically does occupy an area on a page. For example, the
	// common whitespace symbol SPACE (unicode: U+0020 ASCII: 32 decimal 0x20
	// hex) represents a blank space punctuation character in text, used as a
	// word divider in Western scripts.
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

func (l *line) unsafeFromStringPtr(s string) []byte {
	l = (*line)(unsafe.Pointer(&s))
	return *l
}

type list struct {
	buf *[][]byte
}

func (l list) Len() int      { return len(*l.buf) }
func (l list) Cap() int      { return cap(*l.buf) }
func (l list) Reset(cap int) { *l.buf = make([][]byte, 0, cap+1) }
func (l list) Join() []byte  { return bytes.Join(*l.buf, []byte{0x10}) }
func (l list) Make(s string) [][]byte {
	l.Reset(len(s))
	*l.buf = bytes.SplitAfter([]byte(s), []byte{0x10})
	return *l.buf
}
func (l list) Contains(b []byte) bool {
	// for some reason, bytes.Index is several times faster than
	// bytes.Contains (see Index below for improvement)
	for _, s := range *l.buf {
		if bytes.Index(s, b) < 0 {
			return false
		}
	}
	return true
}

func Sindex(s string, sep string) int {
	return strings.Index(s, sep)
}

func benStringIndex(s, substr string) int {
	return strings.Index(s, substr)
}

func benBytesIndex(a, b []byte) int {
	return bytes.Index(a, b)
}

// Index returns the index of the first instance of sep in s, or -1 if sep is not present in s.
func Index(s, sep []byte) int {
	n := len(sep)
	ns := len(s)
	switch {
	case n == 0:
		return 0
	case n == 1:
		return bytealg.IndexByte(s, sep[0])
	case n == ns:
		if string(sep) == string(s) {
			return 0
		}
		return -1
	case n > ns:
		return -1
	case n <= MaxLen:
		// Use brute force when s and sep both are small
		if ns <= bytealg.MaxBruteForce {
			return bytealg.Index(s, sep)
		}
		c0 := sep[0]
		c1 := sep[1]
		i := 0
		t := ns - n + 1
		fails := 0
		for i < t {
			if s[i] != c0 {
				// IndexByte is faster than bytealg.Index, so use it as long as
				// we're not getting lots of false positives.
				o := bytealg.IndexByte(s[i+1:t], c0)
				if o < 0 {
					return -1
				}
				i += o + 1
			}
			if s[i+1] == c1 && bytes.Equal(s[i:i+n], sep) {
				return i
			}
			fails++
			i++
			// Switch to bytealg.Index when IndexByte produces too many false positives.
			if fails > bytealg.Cutover(i) {
				r := bytealg.Index(s[i:], sep)
				if r >= 0 {
					return r + i
				}
				return -1
			}
		}
		return -1
	}
	c0 := sep[0]
	c1 := sep[1]
	i := 0
	fails := 0
	t := len(s) - n + 1
	for i < t {
		if s[i] != c0 {
			o := IndexByte(s[i+1:t], c0)
			if o < 0 {
				break
			}
			i += o + 1
		}
		if s[i+1] == c1 && Equal(s[i:i+n], sep) {
			return i
		}
		i++
		fails++
		if fails >= 4+i>>4 && i < t {
			// Give up on IndexByte, it isn't skipping ahead
			// far enough to be better than Rabin-Karp.
			// Experiments (using IndexPeriodic) suggest
			// the cutover is about 16 byte skips.
			// TODO: if large prefixes of sep are matching
			// we should cutover at even larger average skips,
			// because Equal becomes that much more expensive.
			// This code does not take that effect into account.
			j := bytealg.IndexRabinKarpBytes(s[i:], sep)
			if j < 0 {
				return -1
			}
			return i + j
		}
	}
	return -1
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
	tmp := make([]string, 0, strings.Count(s, "\n")+3)
	for _, line := range strings.Fields(s) {
		tmp = append(tmp, strings.Repeat(" ", n)+line)
	}
	return strings.Join(tmp, "\n")
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
