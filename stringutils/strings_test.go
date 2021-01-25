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
*/

func RuneSamples()[]rune {
    return []rune{'A','0',65,'t','n','f','r','v','\t','\n','\f','\r','\v',48,12,' ',0x20,8,}
}

func ByteSamples()[]byte {
    return []byte{'A','0',65,'t','n','f','r','v','\t','\n','\f','\r','\v',48,12,' ',0x20,8,}
}
func BenchmarkIsWhiteSpace(b *testing.B) {
    for i := 0; i < b.N; i++ {
IsWhiteSpace2(65)               // A
        IsWhiteSpace2('t')               // t
        IsWhiteSpace2('n')               // n
        IsWhiteSpace2('f')               // f
        IsWhiteSpace2('r')               // r
        IsWhiteSpace2('v')               // v
        IsWhiteSpace2('\t')               // \t
        IsWhiteSpace2('\n')               // \n
        IsWhiteSpace2('\f')               // \f
        IsWhiteSpace2('\r')               // \r
        IsWhiteSpace2('\v')               // \v
        IsWhiteSpace2(48)               // 0
        IsWhiteSpace2(12)               // \n
        IsWhiteSpace2(' ')              // space
        IsWhiteSpace2(0x20)             // space
        IsWhiteSpace2(8)                // tab
    }
}

func BenchmarkIsWhiteSpace2(b *testing.B) {
    // 	case ' ', '\t', '\n', '\f', '\r', '\v':
    for i := 0; i < b.N; i++ {
        IsWhiteSpace2(65)               // A
        IsWhiteSpace2('t')               // t
        IsWhiteSpace2('n')               // n
        IsWhiteSpace2('f')               // f
        IsWhiteSpace2('r')               // r
        IsWhiteSpace2('v')               // v
        IsWhiteSpace2('\t')               // \t
        IsWhiteSpace2('\n')               // \n
        IsWhiteSpace2('\f')               // \f
        IsWhiteSpace2('\r')               // \r
        IsWhiteSpace2('\v')               // \v
        IsWhiteSpace2(48)               // 0
        IsWhiteSpace2(12)               // \n
        IsWhiteSpace2(' ')              // space
        IsWhiteSpace2(0x20)             // space
        IsWhiteSpace2(8)                // tab
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
        IsAlphaNum(65)               // A
        IsAlphaNum('t')               // t
        IsAlphaNum('n')               // n
        IsAlphaNum('f')               // f
        IsAlphaNum('r')               // r
        IsAlphaNum('v')               // v
        IsAlphaNum('\t')               // \t
        IsAlphaNum('\n')               // \n
        IsAlphaNum('\f')               // \f
        IsAlphaNum('\r')               // \r
        IsAlphaNum('\v')               // \v
        IsAlphaNum(48)               // 0
        IsAlphaNum(12)               // \n
        IsAlphaNum(' ')              // space
        IsAlphaNum(0x20)             // space
        IsAlphaNum(8)                // tab
    }
}

func BenchmarkIsAlphaNum2(b *testing.B) {
    for i := 0; i < b.N; i++ {
        IsAlphaNum2('A')
        IsAlphaNum2('0')
        IsAlphaNum2('\n')
        IsAlphaNum2(65)               // A
        IsAlphaNum2('t')               // t
        IsAlphaNum2('n')               // n
        IsAlphaNum2('f')               // f
        IsAlphaNum2('r')               // r
        IsAlphaNum2('v')               // v
        IsAlphaNum2('\t')               // \t
        IsAlphaNum2('\n')               // \n
        IsAlphaNum2('\f')               // \f
        IsAlphaNum2('\r')               // \r
        IsAlphaNum2('\v')               // \v
        IsAlphaNum2(48)               // 0
        IsAlphaNum2(12)               // \n
        IsAlphaNum2(' ')              // space
        IsAlphaNum2(0x20)             // space
        IsAlphaNum2(8)                // tab
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
    for _, c := range RuneSamples() {
        name := fmt.Sprintf("IsWhiteSpace: %v",c)
            t.Run(name, func(t *testing.T) {
                got := IsWhiteSpace(c)
                want := strings.IsWhiteSpace(c)
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

func TestIsIsAlphaSwitch(t *testing.T) {
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
			if got := IsAlphaSwitch(tt.args.c); got != tt.want {
				t.Errorf("IsAlphaSwitch() = %v, want %v", got, tt.want)
			}
		})
	}
}
