// Copyright 2020 Michael Treanor. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package strings implements additional functions to support the go library:
// .../go/1.15.3/libexec/src/strings/strings.go
//
// For information about UTF-8 strings in Go, see https://blog.golang.org/strings.
package stringutils

import (
	"fmt"
	"math/rand"
	"testing"
	"unicode"
)

/* Benchmark results
BenchmarkIsWhiteSpace-8    	957675391	         1.31 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsAlphaSwitch-8   	644505457	         1.84 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsAlphaNum-8      	773480662	         1.54 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsAlpha-8         	977785057	         1.23 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsDigit-8         	944950424	         1.25 ns/op	       0 B/op	       0 allocs/op

BenchmarkIsWhiteSpace-8    	899743548	         1.28 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsAlphaSwitch-8   	660817269	         1.81 ns/op	       0 B/op	       0 allocs/op


BenchmarkIsAlphaNum-8      	770840788	         1.51 ns/op	       0 B/op	       0 allocs/op
// removing the 'if' and simply returning the boolean result is 50% faster
BenchmarkIsAlphaNum2-8     	1000000000	         0.923 ns/op	       0 B/op	       0 allocs/op


BenchmarkIsAlpha-8         	967973194	         1.31 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsDigit-8         	971392962	         1.21 ns/op	       0 B/op	       0 allocs/op

Small sample size
================================================================
BenchmarkIsWhiteSpace-8    	227064271	         5.06 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsWhiteSpace2-8   	235199743	         5.08 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsAlphaSwitch-8   	585296522	         1.94 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsAlphaNum-8      	124885717	         9.30 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsAlphaNum2-8     	195595128	         6.16 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsAlpha-8         	990591926	         1.27 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsDigit-8         	937417184	         1.30 ns/op	       0 B/op	       0 allocs/op

================================================================
With ByteSamples() and RuneSamples()  ... consistent samples

BenchmarkIsAlpha-8         	47187997	        25.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsDigit-8         	51663846	        21.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsAlphaSwitch-8   	36453657	        32.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsWhiteSpace-8    	162498454	         7.37 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsWhiteSpace2-8   	48987033	        24.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsAlphaNum-8      	43917738	        26.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsAlphaNum2-8     	58560900	        19.4 ns/op	       0 B/op	       0 allocs/op

================================================================
sample size x8

BenchmarkIsAlpha-8         	 4349080	       273 ns/op	     144 B/op	       1 allocs/op
BenchmarkIsDigit-8         	 5197915	       231 ns/op	     144 B/op	       1 allocs/op
BenchmarkIsAlphaSwitch-8   	 3905835	       309 ns/op	     144 B/op	       1 allocs/op
BenchmarkIsWhiteSpace-8    	 6110427	       196 ns/op	     576 B/op	       1 allocs/op
BenchmarkIsWhiteSpace2-8   	 3619970	       332 ns/op	     576 B/op	       1 allocs/op
BenchmarkIsAlphaNum-8      	 4399827	       282 ns/op	     144 B/op	       1 allocs/op
BenchmarkIsAlphaNum2-8     	 2731586	       444 ns/op	     144 B/op	       1 allocs/op

================================================================
using declared variables for ByteSamples() and RuneSamples()  byteSamples / runeSamples (this is slightly slower??  )

BenchmarkIsAlpha-8         	44298718	        29.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsDigit-8         	49668928	        25.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsAlphaSwitch-8   	35638029	        33.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsWhiteSpace-8    	124931196	         9.48 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsWhiteSpace2-8   	51955320	        23.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsAlphaNum-8      	44332395	        26.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsAlphaNum2-8     	74700576	        15.9 ns/op	       0 B/op	       0 allocs/op

================================================================
using declared variables for func names ... this is MUCH slower .. it must be preventing compiler optimization since the
function is a variable and might ... vary. Thus cannot be replaced at compile time (??)
// byteSamples  = ByteSamples
// for _, c := range byteSamples()

BenchmarkIsAlpha-8         	20974160	        53.0 ns/op	      32 B/op	       1 allocs/op
BenchmarkIsDigit-8         	25314271	        49.0 ns/op	      32 B/op	       1 allocs/op
BenchmarkIsAlphaSwitch-8   	21365545	        56.5 ns/op	      32 B/op	       1 allocs/op
BenchmarkIsWhiteSpace-8    	26824587	        45.6 ns/op	      80 B/op	       1 allocs/op
BenchmarkIsWhiteSpace2-8   	19660117	        59.6 ns/op	      80 B/op	       1 allocs/op
BenchmarkIsAlphaNum-8      	22809402	        53.9 ns/op	      32 B/op	       1 allocs/op
BenchmarkIsAlphaNum2-8     	25673344	        46.1 ns/op	      32 B/op	       1 allocs/op
*/

func Samples(n int)  []byte {
    retval := make([]byte,n)
    for i := 0; i < n; i++ {
        retval := append(retval,rand.Intn(126))
    }
}

func RuneSamples() []rune {
	return []rune{
		'A', '0', 65, 't', 'n', 'f', 'r', 'v', '\t', '\n', '\f', '\r', '\v', 48, 12, ' ', 0x20, 8,
		'A', '0', 65, 't', 'n', 'f', 'r', 'v', '\t', '\n', '\f', '\r', '\v', 48, 12, ' ', 0x20, 8,
		'A', '0', 65, 't', 'n', 'f', 'r', 'v', '\t', '\n', '\f', '\r', '\v', 48, 12, ' ', 0x20, 8,
		'A', '0', 65, 't', 'n', 'f', 'r', 'v', '\t', '\n', '\f', '\r', '\v', 48, 12, ' ', 0x20, 8,
		'A', '0', 65, 't', 'n', 'f', 'r', 'v', '\t', '\n', '\f', '\r', '\v', 48, 12, ' ', 0x20, 8,
		'A', '0', 65, 't', 'n', 'f', 'r', 'v', '\t', '\n', '\f', '\r', '\v', 48, 12, ' ', 0x20, 8,
		'A', '0', 65, 't', 'n', 'f', 'r', 'v', '\t', '\n', '\f', '\r', '\v', 48, 12, ' ', 0x20, 8,
		'A', '0', 65, 't', 'n', 'f', 'r', 'v', '\t', '\n', '\f', '\r', '\v', 48, 12, ' ', 0x20, 8,
	}
}

func ByteSamples() []byte {
	return []byte{
		'A', '0', 65, 't', 'n', 'f', 'r', 'v', '\t', '\n', '\f', '\r', '\v', 48, 12, ' ', 0x20, 8,
		'A', '0', 65, 't', 'n', 'f', 'r', 'v', '\t', '\n', '\f', '\r', '\v', 48, 12, ' ', 0x20, 8,
		'A', '0', 65, 't', 'n', 'f', 'r', 'v', '\t', '\n', '\f', '\r', '\v', 48, 12, ' ', 0x20, 8,
		'A', '0', 65, 't', 'n', 'f', 'r', 'v', '\t', '\n', '\f', '\r', '\v', 48, 12, ' ', 0x20, 8,
		'A', '0', 65, 't', 'n', 'f', 'r', 'v', '\t', '\n', '\f', '\r', '\v', 48, 12, ' ', 0x20, 8,
		'A', '0', 65, 't', 'n', 'f', 'r', 'v', '\t', '\n', '\f', '\r', '\v', 48, 12, ' ', 0x20, 8,
		'A', '0', 65, 't', 'n', 'f', 'r', 'v', '\t', '\n', '\f', '\r', '\v', 48, 12, ' ', 0x20, 8,
		'A', '0', 65, 't', 'n', 'f', 'r', 'v', '\t', '\n', '\f', '\r', '\v', 48, 12, ' ', 0x20, 8,
	}
}

// This is a horrible idea ... much slower
// var (
// 	byteSamples  = ByteSamples
// 	runeSamples  = RuneSamples
// )

func BenchmarkIsAlpha(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, c := range ByteSamples() {
			IsAlpha(c)
		}
	}
}

func BenchmarkIsDigit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, c := range ByteSamples() {
			IsDigit(c)
		}
	}
}

func BenchmarkIsAlphaSwitch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, c := range ByteSamples() {
			IsAlphaSwitch(c)
		}
	}
}

func BenchmarkIsWhiteSpace(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, r := range RuneSamples() {
			IsWhiteSpace(r)
		}
	}
}

func BenchmarkIsWhiteSpace2(b *testing.B) {
	// 	case ' ', '\t', '\n', '\f', '\r', '\v':
	for i := 0; i < b.N; i++ {
		for _, r := range RuneSamples() {
			isWhiteSpace2(r)
		}
	}
}

func BenchmarkIsAlphaNum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, c := range ByteSamples() {
			IsAlphaNum(c)
		}
	}
}

func BenchmarkIsAlphaNum2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, c := range ByteSamples() {
			IsAlphaNum2(c)
		}
	}
}

func TestIsWhiteSpace(t *testing.T) {
	for _, c := range RuneSamples() {
		name := fmt.Sprintf("IsWhiteSpace: %v", c)
		t.Run(name, func(t *testing.T) {
			got := IsWhiteSpace(c)
			want := unicode.IsSpace(c)
			if got != want {
				t.Errorf("IsWhiteSpace() = %v, want %v", got, want)
			}
		})

	}
}

func TestIsAlpha(t *testing.T) {
	type args struct {
		c byte
	}
	tests := []struct {
		args args
		want bool
	}{
		// TODO: Add test cases.
		{args{'\n'}, false},
		{args{'\t'}, false},
		{args{'\r'}, false},
		{args{'\v'}, false},
		{args{'\f'}, false},
		{args{' '}, false},
		{args{'A'}, true},
		{args{'c'}, true},
		{args{'0'}, false},
		{args{'7'}, false},
		{args{'='}, false},
		{args{'?'}, false},
		{args{'/'}, false},
		{args{'%'}, false},
	}
	for _, tt := range tests {
		t.Run(string(tt.args.c), func(t *testing.T) {
			if got := IsAlpha(tt.args.c); got != tt.want {
				t.Errorf("IsAlpha() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsDigit(t *testing.T) {
	type args struct {
		c byte
	}
	tests := []struct {
		args args
		want bool
	}{
		// TODO: Add test cases.
		{args{'\n'}, false},
		{args{'\t'}, false},
		{args{'\r'}, false},
		{args{'\v'}, false},
		{args{'\f'}, false},
		{args{' '}, false},
		{args{'A'}, false},
		{args{'c'}, false},
		{args{'0'}, true},
		{args{'7'}, true},
		{args{'='}, false},
		{args{'?'}, false},
		{args{'/'}, false},
		{args{'%'}, false},
	}
	for _, tt := range tests {
		t.Run(string(tt.args.c), func(t *testing.T) {
			if got := IsDigit(tt.args.c); got != tt.want {
				t.Errorf("IsDigit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsAlphaNum(t *testing.T) {
	type args struct {
		c byte
	}
	tests := []struct {
		args args
		want bool
	}{
		// TODO: Add test cases.
		{args{'\n'}, false},
		{args{'\t'}, false},
		{args{'\r'}, false},
		{args{'\v'}, false},
		{args{'\f'}, false},
		{args{' '}, false},
		{args{'A'}, true},
		{args{'c'}, true},
		{args{'0'}, true},
		{args{'7'}, true},
		{args{'='}, false},
		{args{'?'}, false},
		{args{'/'}, false},
		{args{'%'}, false},
	}
	for _, tt := range tests {
		t.Run(string(tt.args.c), func(t *testing.T) {
			if got := IsAlphaNum(tt.args.c); got != tt.want {
				t.Errorf("IsAlphaNum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsIsAlphaSwitch(t *testing.T) {
	type args struct {
		c byte
	}
	tests := []struct {
		args args
		want bool
	}{
		// TODO: Add test cases.
		{args{'\n'}, false},
		{args{'\t'}, false},
		{args{'\r'}, false},
		{args{'\v'}, false},
		{args{'\f'}, false},
		{args{' '}, false},
		{args{'A'}, true},
		{args{'c'}, true},
		{args{'0'}, true},
		{args{'7'}, true},
		{args{'='}, false},
		{args{'?'}, false},
		{args{'/'}, false},
		{args{'%'}, false},
	}
	for _, tt := range tests {
		t.Run(string(tt.args.c), func(t *testing.T) {
			if got := IsAlphaSwitch(tt.args.c); got != tt.want {
				t.Errorf("IsAlphaSwitch() = %v, want %v", got, tt.want)
			}
		})
	}
}
