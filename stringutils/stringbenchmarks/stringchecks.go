package stringbenchmarks

import (
	"unicode"
)

// IsASCIIPrintable checks if s is ascii and printable, aka doesn't include tab, backspace, etc.
func IsASCIIPrintable(s string) bool {
	for _, r := range s {
		if !unicode.IsPrint(r) {
			return false
		}
	}
	return true
}

func IsDigit(c byte) bool {
	return '0' <= c && c <= '9'
}

// IsDigitSingleOP uses a single operation instead of
// the standard a << c && c << b form
// Another good example: very common thing is if(x >= 1 && x <= 9) which can be done as if( (unsigned)(x-1) <=(unsigned)(9-1)) Changing two conditional tests to one can be a big speed advantage; especially when it allows predicated execution instead of branches. I used this for years (where justified) until I noticed abt 10 years ago that compilers had started doing this transform in the optimizer, then I stopped. Still good to know, since there are similar situations where the compiler can't make the transform for you. Or if you're working on a compiler.
func IsDigitSingleOP(c byte) bool {
	return c >= '0' && c <= '9'
}

// IsDigitSingleOPCompare is a sample implementation used for
// benchmarking IsDigitSingleOP
func IsDigitSingleOPCompare(c byte) bool {
	c -= 48
	return c <= 9
}

func IsASCIIAlpha(c byte) bool {
	return ('a' <= c && c <= 'z') || ('A' <= c && c <= 'Z')
}

func IsHex(c byte) bool {
	return ('0' <= c && c <= '9') || ('a' <= c && c <= 'f') || ('A' <= c && c <= 'F')
}

func IsAlphaNumSwitch(c byte) bool {
	switch {
	case 'a' <= c && c <= 'z':
		return true
	case 'A' <= c && c <= 'Z':
		return true
	case '0' <= c && c <= '9':
		return true
	default:
		return false
	}
}

func IsAlphaNumUnder(c byte) bool {
	return 'a' <= c && c <= 'z' || 'A' <= c && c <= 'Z' || '0' <= c && c <= '9' || c == '_'
}

// IsAlphaNum reports whether the byte is an ASCII letter, number, or underscore
func IsAlphaNum(c byte) bool {
	return 'a' <= c && c <= 'z' || 'A' <= c && c <= 'Z' || '0' <= c && c <= '9'
}
