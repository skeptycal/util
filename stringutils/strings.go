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
	"strconv"
)

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
    switch  {
    case c>='a' && c<='z':
        return true
    case c>='A' && c<='Z':
        return true
    case c>='0' && c<='9':
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

var spaceMap3 = map[string]string {
        "U+0020":	"SPACE",
        "U+00A0":	"NO-BREAK SPACE",
        "U+1680":	"OGHAM SPACE MARK",
        "U+2000":	"EN QUAD",
        "U+2001":	"EM QUAD",
        "U+2002":	"EN SPACE",
        "U+2003":	"EM SPACE",
        "U+2004":	"THREE-PER-EM SPACE",
        "U+2005":	"FOUR-PER-EM SPACE",
        "U+2006":	"SIX-PER-EM SPACE",
        "U+2007":	"FIGURE SPACE",
        "U+2008":	"PUNCTUATION SPACE",
        "U+2009":	"THIN SPACE",
        "U+200A":	"HAIR SPACE",
        "U+202F":	"NARROW NO-BREAK SPACE",
        "U+205F":	"MEDIUM MATHEMATICAL SPACE",
        "U+3000":	"IDEOGRAPHIC SPACE",
    }
func isWhiteSpace3(c rune) bool {

    if _, ok := spaceMap3[string(c)]; ok {
        return true
    }
    return false
    // return unicode.IsSpace(c)
}

var spaceMap4 = map[rune]string{
        0x0020:	"SPACE",
        0x00A0:	"NO-BREAK SPACE",
        0x1680:	"OGHAM SPACE MARK",
        0x2000:	"EN QUAD",
        0x2001:	"EM QUAD",
        0x2002:	"EN SPACE",
        0x2003:	"EM SPACE",
        0x2004:	"THREE-PER-EM SPACE",
        0x2005:	"FOUR-PER-EM SPACE",
        0x2006:	"SIX-PER-EM SPACE",
        0x2007:	"FIGURE SPACE",
        0x2008:	"PUNCTUATION SPACE",
        0x2009:	"THIN SPACE",
        0x200A:	"HAIR SPACE",
        0x202F:	"NARROW NO-BREAK SPACE",
        0x205F:	"MEDIUM MATHEMATICAL SPACE",
        0x3000:	"IDEOGRAPHIC SPACE",
    }

func isWhiteSpace4(c rune) bool {

    if _, ok := spaceMap4[c]; ok {
        return true
    }
    return false
}

var spaceMap5 = map[rune]bool{
        0x0020:	true,
        0x00A0:	true,
        0x1680:	true,
        0x2000:	true,
        0x2001:	true,
        0x2002:	true,
        0x2003:	true,
        0x2004:	true,
        0x2005:	true,
        0x2006:	true,
        0x2007:	true,
        0x2008:	true,
        0x2009:	true,
        0x200A:	true,
        0x202F:	true,
        0x205F:	true,
        0x3000:	true,
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
