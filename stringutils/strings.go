// Copyright 2020 Michael Treanor. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package strings implements additional functions to support the go library:
// .../go/1.15.3/libexec/src/strings/strings.go
//
// For information about UTF-8 strings in Go, see https://blog.golang.org/strings.
package strings

// const alphanumerics = "0123456789abcdefghijklmnopqrstuvwxyz"

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
    case c>'a' && c<'z':
        return true

    case c>'A' && c<'Z':
        return true
    case c>'0' && c<'9':
        return true
    default:
        return false
    }
}

func IsAlphaNum(c byte) bool {

	if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9') {
		return true
	}
	return false
}

func IsWhiteSpace(c rune) bool {
	switch c {
	case ' ', '\t', '\n', '\f', '\r', '\v':
		return true
	}
	return false
}
