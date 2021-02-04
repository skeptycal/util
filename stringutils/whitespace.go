package stringutils

import (
	"strings"
	"unicode"
)

//// ======================= best performing functions ... so far

// IsASCIISpace tests for the most common ASCII whitespace characters:
//  ' ', '\t', '\n', '\f', '\r', '\v', U+0085 (NEL), U+00A0 (NBSP)
//
// This includes all Unicode code points that are one byte in length
// and leaves out the unicode code points above 0xFF.
//
// The C language defines whitespace characters to be "space,
// horizontal tab, new-line, vertical tab, and form-feed."
func IsASCIISpace(c byte) bool {

    return c == 32 || (c > 8 && c < 14)

    /*

    0x09    0b00001001
    0x0a    0b00001010
    0x0b    0b00001011
    0x0c    0b00001100
    0x0d    0b00001101

    */
}

// DedupeWhitespace removes any duplicate whitespace from the string.
func DedupeWhitespace(s string) string {
	return reWhitespace.ReplaceAllString(strings.TrimSpace(s), " ")
}

//// ======================= benchmark samples
//// ======================= byte parameter

// isWhiteSpaceRegexByte is a sample implementation of a whitespace
// test function used for benchmark comparisons.
//
// regex does not recognize \v as whitespace ... so an additional
// check is required.
func isWhiteSpaceRegexByte(c byte) bool {
	return c == 0x0b || reWhitespace.Find([]byte{c}) != nil
}

// isWhiteSpace is a sample implementation of a whitespace
// test function used for benchmark comparisons.
func isWhiteSpace(c byte) bool {
	return c == ' ' || c == '\n' || c == '\t' || c == '\r' || c == '\f' || c == '\v' || c == 0x0085 || c == 0x00A0
}

// isWhiteSpace2 is a sample implementation of a whitespace
// test function used for benchmark comparisons.
func isWhiteSpace2(c byte) bool {
	switch c {
	case ' ': return true
    case '\n': return true
    case '\t': return true
    case '\f', '\r', '\v': return true
    default: return false
	}
	// return false
}

// isWhiteSpaceStringSliceBytes a sample implementation of a whitespace
// test function used for benchmark comparisons.
func isWhiteSpaceStringSliceBytes(c byte) bool {
	for _, v := range shortByteList {
		if c == v {
			return true
		}
	}
	return false
}

//// ======================= rune parameter

// IsUnicodeWhiteSpaceMap reports whether the rune is any utf8 whitespace character
// using the broadest and most complete definition.
//
// The speed of this implementation ~25% slower than that of
// IsASCIISpace(c byte) but tests 3.75 times more possible code points.
//
// The speed is ~7% faster than that of unicode.IsSpace(r rune) from the
// standard library and covers nearly twice as many code points.
//
// isWhiteSpaceLogicChain checks for any unicode whitespace rune.
//
// Included:
//  0x2000, 0x2001, 0x2002, 0x2003, 0x2004, 0x2005,
//  0x2006, 0x2007, 0x2008, 0x2009, 0x200A, 0x2028,
//  0x2029, 0x202F, 0x205F, 0x3000, 0x1680
//
// Related Unicode characters (property White_Space=no)
// Not included:
//  0x200B,	0x200C,	0x200D,	0x2060
func IsUnicodeWhiteSpaceMap(r rune) bool {
	if r < unicode.MaxLatin1 {
		return r == ' ' || (r > 8 && r < 14) || r == 0x85 || r == 0xA0
	}

	if _, ok := UnicodeWhiteSpaceMap[r]; ok {
		return true
	}
	return false
}

// isUnicodeWhiteSpaceMapSwitch is a sample implementation of a whitespace
// test function used for benchmark comparisons.
func isUnicodeWhiteSpaceMapSwitch(r rune) bool {
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

// isWhiteSpaceStringMap is a sample implementation of a whitespace
// test function used for benchmark comparisons.
func isWhiteSpaceStringMap(r rune) bool {
	if _, ok := UnicodeWhiteSpaceMap[r]; ok {
		return true
	}
	return false
}

// isWhiteSpaceBoolMap is a sample implementation of a whitespace
// test function used for benchmark comparisons.
func isWhiteSpaceBoolMap(r rune) bool {
	if _, ok := unicodeWhiteSpaceMapBool[r]; ok {
		return true
	}
	return false
}

// isWhiteSpaceLogicChain is a sample implementation of a whitespace
// test function used for benchmark comparisons.
func isWhiteSpaceLogicChain(r rune) bool {
	if r < unicode.MaxLatin1 {
		return r == ' ' || r == '\t' || r == '\n' || r == '\f' || r == '\r' || r == '\v' || r == 0x85 || r == 0xA0
	}
	return (r >= 0x2000 && r <= 0x200A) || r == 0x2028 || r == 0x2029 || r == 0x202F || r == 0x205F || r == 0x3000 || r == 0xFFEF || r == 0x1680
}

// isWhiteSpaceLogicChainNoCheck is a sample implementation of a whitespace
// test function used for benchmark comparisons.
func isWhiteSpaceLogicChainNoCheck(r rune) bool {

	return r == ' ' || r == '\t' || r == '\n' || r == '\f' || r == '\r' || r == '\v' || r == 0x85 || r == 0xA0 || (r >= 0x2000 && r <= 0x200A) || r == 0x2028 || r == 0x2029 || r == 0x202F || r == 0x205F || r == 0x3000 || r == 0xFFEF || r == 0x1680
}

// isWhiteSpaceRuneSlice a sample implementation of a whitespace
// test function used for benchmark comparisons.
func isWhiteSpaceRuneSlice(r rune) bool {
	shortlist := []rune{' ', '\t', '\n', '\f', '\r', '\v', 0x85, 0xA0}
	longlist := []rune{0x2000, 0x200A, 0x2028, 0x2029, 0x202F, 0x205F, 0x3000, 0xFFEF, 0x1680}

	for _, v := range shortlist {
		if v == r {
			return true
		}
	}

	for _, v := range longlist {
		if v == r {
			return true
		}
	}

	return false
}

// isWhiteSpaceStringSlice a sample implementation of a whitespace
// test function used for benchmark comparisons.
func isWhiteSpaceStringSlice(r rune) bool {
	if uint32(r) <= unicode.MaxLatin1 {
		for _, v := range shortByteListString {
			if r == v {
				return true
			}
		}
	}
	for _, v := range longRuneListString {
		if r == v {
			return true
		}
	}
	return false
}

//// ======================= string parameter

// isWhiteSpaceIndexByte a sample implementation of a whitespace
// test function used for benchmark comparisons.
func isWhiteSpaceIndexByte(s string) bool {
	return strings.Index(shortByteListString, s) >= 0
}

// isWhiteSpaceContainsByte a sample implementation of a whitespace
// test function used for benchmark comparisons.
func isWhiteSpaceContainsByte(s string) bool {
	return strings.Contains(shortByteListString, s)
}

// isWhiteSpaceIndexRune a sample implementation of a whitespace
// test function used for benchmark comparisons.
func isWhiteSpaceIndexRune(s string) bool {
	return strings.Index(longRuneListString, s) >= 0
}

// isWhiteSpaceContainsRune a sample implementation of a whitespace
// test function used for benchmark comparisons.
func isWhiteSpaceContainsRune(s string) bool {
	return strings.Contains(longRuneListString, s)
}

// isWhiteSpaceTrim a sample implementation of a whitespace
// test function used for benchmark comparisons.
func isWhiteSpaceTrim(s string) bool {
	return strings.TrimSpace(s) == ""
}
