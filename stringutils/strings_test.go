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

BenchmarkIsAlpha-8         	45196527	        25.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsDigit-8         	51888256	        22.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsAlphaSwitch-8   	34407850	        33.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsWhiteSpace-8    	164546520	         7.41 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsWhiteSpace2-8   	52679346	        22.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsAlphaNum-8      	44141539	        26.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsAlphaNum2-8     	71749123	        15.7 ns/op	       0 B/op	       0 allocs/op
*/

func RuneSamples() []rune {
	return []rune{'A', '0', 65, 't', 'n', 'f', 'r', 'v', '\t', '\n', '\f', '\r', '\v', 48, 12, ' ', 0x20, 8}
}

func ByteSamples() []byte {
	return []byte{'A', '0', 65, 't', 'n', 'f', 'r', 'v', '\t', '\n', '\f', '\r', '\v', 48, 12, ' ', 0x20, 8}
}

var (
	byteSamples []byte = ByteSamples()
	runeSamples []rune = RuneSamples()
)

func BenchmarkIsAlpha(b *testing.B) {
	samples := ByteSamples()
	for i := 0; i < b.N; i++ {
		for _, c := range samples {
			IsAlpha(c)
		}
	}
}

func BenchmarkIsDigit(b *testing.B) {
	samples := ByteSamples()
	for i := 0; i < b.N; i++ {
		for _, c := range samples {
			IsDigit(c)
		}
	}
}

func BenchmarkIsAlphaSwitch(b *testing.B) {
	samples := ByteSamples()
	for i := 0; i < b.N; i++ {
		for _, c := range samples {
			IsAlphaSwitch(c)
		}
	}
}

func BenchmarkIsWhiteSpace(b *testing.B) {
	samples := RuneSamples()
	for i := 0; i < b.N; i++ {
		for _, r := range samples {
			IsWhiteSpace(r)
		}
	}
}

func BenchmarkIsWhiteSpace2(b *testing.B) {
	// 	case ' ', '\t', '\n', '\f', '\r', '\v':
	samples := RuneSamples()
	for i := 0; i < b.N; i++ {
		for _, r := range samples {
			isWhiteSpace2(r)
		}
	}
}

func BenchmarkIsAlphaNum(b *testing.B) {
	samples := ByteSamples()
	for i := 0; i < b.N; i++ {
		for _, c := range samples {
			IsAlphaNum(c)
		}
	}
}

func BenchmarkIsAlphaNum2(b *testing.B) {
	samples := ByteSamples()
	for i := 0; i < b.N; i++ {
		for _, c := range samples {
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
