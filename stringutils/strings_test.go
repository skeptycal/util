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

// Benchmark results
/*
IsWhiteSpace switch/case is significantly slower than multiple || with 6 checks

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
using n = 1024
BenchmarkIsAlpha-8         	   39996	     30380 ns/op	    1024 B/op	       1 allocs/op
BenchmarkIsDigit-8         	   42748	     28097 ns/op	    1024 B/op	       1 allocs/op
BenchmarkIsAlphaSwitch-8   	   39072	     31276 ns/op	    1024 B/op	       1 allocs/op
BenchmarkIsWhiteSpace-8    	   44443	     27086 ns/op	    5120 B/op	       2 allocs/op
BenchmarkIsWhiteSpace2-8   	   43936	     27800 ns/op	    5120 B/op	       2 allocs/op
BenchmarkIsAlphaNum-8      	   40113	     29856 ns/op	    1024 B/op	       1 allocs/op
BenchmarkIsAlphaNum2-8     	   43183	     28851 ns/op	    1024 B/op	       1 allocs/op

(without preallocation of rune make buffer; e.g. the 'n' in retval := make([]byte, 0, n)
BenchmarkIsWhiteSpace-8    	   43213	     28103 ns/op	    9208 B/op	      11 allocs/op
BenchmarkIsWhiteSpace2-8   	   41742	     29355 ns/op	    9208 B/op	      11 allocs/op

(without preallocation of either buffer)
BenchmarkIsAlpha-8         	   39459	     30819 ns/op	    2040 B/op	       8 allocs/op
BenchmarkIsDigit-8         	   44440	     26921 ns/op	    2040 B/op	       8 allocs/op
BenchmarkIsAlphaSwitch-8   	   39086	     31031 ns/op	    2040 B/op	       8 allocs/op
BenchmarkIsWhiteSpace-8    	   42594	     28711 ns/op	   10224 B/op	      18 allocs/op
BenchmarkIsWhiteSpace2-8   	   40940	     30592 ns/op	   10224 B/op	      18 allocs/op
BenchmarkIsAlphaNum-8      	   39778	     30075 ns/op	    2040 B/op	       8 allocs/op
BenchmarkIsAlphaNum2-8     	   43225	     28086 ns/op	    2040 B/op	       8 allocs/op

================================================================
n = 65535 (with preallocation) ( ... preallocation is generally good)

BenchmarkIsAlpha-8         	     619	   1899081 ns/op	   65539 B/op	       1 allocs/op
BenchmarkIsDigit-8         	     721	   1660150 ns/op	   65536 B/op	       1 allocs/op
BenchmarkIsAlphaSwitch-8   	     636	   1888789 ns/op	   65536 B/op	       1 allocs/op
BenchmarkIsWhiteSpace-8    	     708	   1717852 ns/op	  327681 B/op	       2 allocs/op
BenchmarkIsWhiteSpace2-8   	     688	   1718280 ns/op	  327681 B/op	       2 allocs/op
BenchmarkIsAlphaNum-8      	     643	   1875367 ns/op	   65536 B/op	       1 allocs/op
BenchmarkIsAlphaNum2-8     	     690	   1746706 ns/op	   65536 B/op	       1 allocs/op

(with special 'common' case of space)
BenchmarkIsWhiteSpace-8    	     657	   1740744 ns/op	  327681 B/op	       2 allocs/op
BenchmarkIsWhiteSpace2-8   	     676	   1800894 ns/op	  327681 B/op	       2 allocs/op

(no preallocation)
BenchmarkIsAlpha-8         	     613	   1914118 ns/op	  284666 B/op	      23 allocs/op
BenchmarkIsDigit-8         	     702	   1726401 ns/op	  284666 B/op	      23 allocs/op
BenchmarkIsAlphaSwitch-8   	     594	   1990682 ns/op	  284669 B/op	      23 allocs/op
BenchmarkIsWhiteSpace-8    	     608	   1941723 ns/op	 1693433 B/op	      49 allocs/op
BenchmarkIsWhiteSpace2-8   	     594	   2052508 ns/op	 1693435 B/op	      49 allocs/op
BenchmarkIsAlphaNum-8      	     614	   1972096 ns/op	  284665 B/op	      23 allocs/op
BenchmarkIsAlphaNum2-8     	     666	   1828022 ns/op	  284666 B/op	      23 allocs/op

*/

// Alternate methods that generated worse results...
/*
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

================================================================
using maps for checks (3 [string]string/ 4 [rune]string ) (this is much worse than I expected!!)
BenchmarkIsWhiteSpace-8    	     708	   1681603 ns/op	  327680 B/op	       2 allocs/op
BenchmarkIsWhiteSpace2-8   	     666	   1775882 ns/op	  327681 B/op	       2 allocs/op

 using map [string]string
BenchmarkIsWhiteSpace3-8   	      18	  63820745 ns/op	76760858 B/op	   68792 allocs/op
 using [rune]string
BenchmarkIsWhiteSpace4-8   	      22	  51020315 ns/op	47036008 B/op	   68787 allocs/op
 using [rune]bool (returned on the fly from a function call)
BenchmarkIsWhiteSpace5-8   	      30	  38028585 ns/op	15449897 B/op	   78547 allocs/op
 using [rune]bool from var(map literal)
BenchmarkIsWhiteSpace5-8   	     417	   2779469 ns/op	  327681 B/op	       2 allocs/op


using pre-made maps (var = map literal) for 4 and 5 .. 3 is still generated within function
BenchmarkIsWhiteSpace-8    	10117532	       119 ns/op	      16 B/op	       2 allocs/op
BenchmarkIsWhiteSpace2-8   	 9799669	       117 ns/op	      16 B/op	       2 allocs/op
BenchmarkIsWhiteSpace3-8   	  402494	      2942 ns/op	    3515 B/op	       5 allocs/op
BenchmarkIsWhiteSpace4-8   	 7176753	       164 ns/op	      16 B/op	       2 allocs/op
BenchmarkIsWhiteSpace5-8   	 7355563	       164 ns/op	      16 B/op	       2 allocs/op
*/

const (
	defaultSamples = 1<<8 - 1  // 1<<8 - 1
	maxSamples     = 1<<32 - 1
)

func Samples(n int) []byte {
	if n < 2 || n > maxSamples {
		n = defaultSamples
	}
	retval := make([]byte, 0, n)
	for i := 0; i < n; i++ {
		retval = append(retval, byte(rand.Intn(126)))
	}
	return retval
}

func SmallRuneSamples() []rune {
	return []rune{
		'A', '0', 65, 't', 'n', 'f', 'r', 'v', '\t', '\n', '\f', '\r', '\v', 48, 12, ' ', 0x20, 8,
	}
}

func SmallByteSamples() []byte {
	return []byte{
		'A', '0', 65, 't', 'n', 'f', 'r', 'v', '\t', '\n', '\f', '\r', '\v', 48, 12, ' ', 0x20, 8,
	}
}

func ByteSamples() []byte {
	return Samples(defaultSamples)
}

func RuneSamples() []rune {
	const n = defaultSamples

	retval := make([]rune, 0, n)

	for _, c := range Samples(n) {
		retval = append(retval, rune(c))
	}
	return retval
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

func BenchmarkIsWhiteSpace3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, r := range RuneSamples() {
			isWhiteSpace3(r)
		}
	}
}

func BenchmarkIsWhiteSpace4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, r := range RuneSamples() {
			isWhiteSpace4(r)
		}
	}
}

func BenchmarkIsWhiteSpace5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, r := range RuneSamples() {
			isWhiteSpace5(r)
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
func TestIsWhiteSpace2(t *testing.T) {
	for _, c := range RuneSamples() {
		name := fmt.Sprintf("IsWhiteSpace2: %v", c)
		t.Run(name, func(t *testing.T) {
			got := isWhiteSpace2(c)
			want := unicode.IsSpace(c)
			if got != want {
				t.Errorf("IsWhiteSpace2() = %v, want %v", got, want)
			}
		})
	}
}
func TestIsWhiteSpace3(t *testing.T) {
	for _, c := range RuneSamples() {
		name := fmt.Sprintf("IsWhiteSpace3: (%q)", c)
		t.Run(name, func(t *testing.T) {
            got := isWhiteSpace3(c)
            want := unicode.IsSpace(c)
			want := unicode.IsSpace(c)
			if got != want {
				t.Errorf("IsWhiteSpace3() = %v, want %v", got, want)
			}
		})
	}
}
func TestIsWhiteSpace4(t *testing.T) {
	for _, c := range RuneSamples() {
		name := fmt.Sprintf("IsWhiteSpace4: %v", c)
		t.Run(name, func(t *testing.T) {
			got := isWhiteSpace4(c)
			want := unicode.IsSpace(c)
			if got != want {
				t.Errorf("IsWhiteSpace4() = %v, want %v", got, want)
			}
		})
	}
}
func TestIsWhiteSpace5(t *testing.T) {
	for _, c := range RuneSamples() {
		name := fmt.Sprintf("IsWhiteSpace5: %v", c)
		t.Run(name, func(t *testing.T) {
			got := isWhiteSpace5(c)
			want := unicode.IsSpace(c)
			if got != want {
				t.Errorf("IsWhiteSpace5() = %v, want %v", got, want)
			}
		})
	}
}
