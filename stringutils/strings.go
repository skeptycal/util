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

func IsWhiteSpace(c rune) bool {
	switch c {
	case ' ', '\t', '\n', '\f', '\r', '\v':
		return true
	}
	return false
}

func IsWhiteSpace2(c rune) bool {
    return c == ' ' || c == '\t' || c == '\n' || c == '\f' || c == '\r' || c == '\v'
}
