package stringutils

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

func IsAlphaNum2(c byte) bool {
	return 'a' <= c && c <= 'z' || 'A' <= c && c <= 'Z' || '0' <= c && c <= '9' || c == '_'
}

// IsAlphaNum reports whether the byte is an ASCII letter, number, or underscore
func IsAlphaNum(c byte) bool {
	return 'a' <= c && c <= 'z' || 'A' <= c && c <= 'Z' || '0' <= c && c <= '9'
}
