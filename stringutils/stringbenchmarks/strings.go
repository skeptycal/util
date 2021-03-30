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
package stringbenchmarks

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// ToString implements Stringer directly as a function call
// with a parameter instead of a method on that parameter.
func ToString(any interface{}) string {
	if v, ok := any.(fmt.Stringer); ok {
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

func JoinLines(list []string) string {
	return strings.Join(list, "\n")
}

func TabIt(s string, n int) string {
	tmp := make([]string, 0, strings.Count(s, "\n")+3)
	for _, line := range strings.Fields(s) {
		tmp = append(tmp, strings.Repeat(" ", n)+line)
	}
	return strings.Join(tmp, "\n")
}

// RuneSample prints a sample of various Unicode runes.
func RuneSample(c rune) {
	s := "日本語"
	fmt.Printf("Glyph:   %q\n", s)
	fmt.Printf("UTF-8:   [% x]\n", []byte(s))
	fmt.Printf("Unicode: %U\n", []rune(s))
}
