// Copyright 2020 Michael Treanor. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package strings implements additional functions to support the go library:
// .../go/1.15.3/libexec/src/strings/strings.go
//
// For information about UTF-8 strings in Go, see https://blog.golang.org/strings.
package strings

import "testing"

/* Benchmark results
BenchmarkIsWhiteSpace-8    	923692542	         1.29 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsAlphaSwitch-8   	639313898	         1.88 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsAlphaNum-8      	780350467	         1.54 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsAlpha-8         	972538683	         1.28 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsDigit-8         	935870430	         1.37 ns/op	       0 B/op	       0 allocs/op
*/
func BenchmarkIsWhiteSpace(b *testing.B) {
    for i := 0; i < b.N; i++ {
        IsWhiteSpace(65) // A
        IsWhiteSpace(48) // 0
        IsWhiteSpace(12) // \n
    }
}

func BenchmarkIsAlphaSwitch(b *testing.B) {
    for i := 0; i < b.N; i++ {
        IsAlphaSwitch(65) // A
        IsAlphaSwitch(48) // 0
        IsAlphaSwitch(12) // \n
    }
}

func BenchmarkIsAlphaNum(b *testing.B) {
    for i := 0; i < b.N; i++ {
        IsAlphaNum('A')
        IsAlphaNum('0')
        IsAlphaNum('\n')
    }
}

func BenchmarkIsAlpha(b *testing.B) {
    for i := 0; i < b.N; i++ {
        IsAlpha('A')
        IsAlpha('0')
        IsAlpha('\n')
    }
}

func BenchmarkIsDigit(b *testing.B) {
    for i := 0; i < b.N; i++ {
        IsDigit('A')
        IsDigit('0')
        IsDigit('\n')
    }
}

func TestIsWhiteSpace(t *testing.T) {
	type args struct {
		c rune
	}
	tests := []struct {
		args args
		want bool
	}{
        // TODO: Add test cases.
        { args{'\n'}, true},
        { args{'\t'}, true},
        { args{'\r'}, true},
        { args{'\v'}, true},
        { args{'\f'}, true},
        { args{' '}, true},
        { args{'A'}, false},
        { args{'c'}, false},
        { args{'0'}, false},
        { args{'7'}, false},
        { args{'='}, false},
        { args{'?'}, false},
        { args{'/'}, false},
        { args{'%'}, false},
	}
	for _, tt := range tests {
		t.Run(string(tt.args.c), func(t *testing.T) {
			if got := IsWhiteSpace(tt.args.c); got != tt.want {
				t.Errorf("IsWhiteSpace() = %v, want %v", got, tt.want)
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
        { args{'\n'}, false},
        { args{'\t'}, false},
        { args{'\r'}, false},
        { args{'\v'}, false},
        { args{'\f'}, false},
        { args{' '}, false},
        { args{'A'}, true},
        { args{'c'}, true},
        { args{'0'}, false},
        { args{'7'}, false},
        { args{'='}, false},
        { args{'?'}, false},
        { args{'/'}, false},
        { args{'%'}, false},
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
        { args{'\n'}, false},
        { args{'\t'}, false},
        { args{'\r'}, false},
        { args{'\v'}, false},
        { args{'\f'}, false},
        { args{' '}, false},
        { args{'A'}, false},
        { args{'c'}, false},
        { args{'0'}, true},
        { args{'7'}, true},
        { args{'='}, false},
        { args{'?'}, false},
        { args{'/'}, false},
        { args{'%'}, false},
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
        { args{'\n'}, false},
        { args{'\t'}, false},
        { args{'\r'}, false},
        { args{'\v'}, false},
        { args{'\f'}, false},
        { args{' '}, false},
        { args{'A'}, true},
        { args{'c'}, true},
        { args{'0'}, true},
        { args{'7'}, true},
        { args{'='}, false},
        { args{'?'}, false},
        { args{'/'}, false},
        { args{'%'}, false},
	}
	for _, tt := range tests {
		t.Run(string(tt.args.c), func(t *testing.T) {
			if got := IsAlphaNum(tt.args.c); got != tt.want {
				t.Errorf("IsAlphaNum() = %v, want %v", got, tt.want)
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
        { args{'\n'}, false},
        { args{'\t'}, false},
        { args{'\r'}, false},
        { args{'\v'}, false},
        { args{'\f'}, false},
        { args{' '}, false},
        { args{'A'}, true},
        { args{'c'}, true},
        { args{'0'}, true},
        { args{'7'}, true},
        { args{'='}, false},
        { args{'?'}, false},
        { args{'/'}, false},
        { args{'%'}, false},
	}
	for _, tt := range tests {
		t.Run(string(tt.args.c), func(t *testing.T) {
			if got := IsAlphaNum(tt.args.c); got != tt.want {
				t.Errorf("IsAlphaNum() = %v, want %v", got, tt.want)
			}
		})
	}
}
