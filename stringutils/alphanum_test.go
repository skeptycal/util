package stringutils

import (
	"fmt"
	"testing"
	"unicode"
)

var (
    // testBytesAlphanum results
    /********************************
    BenchmarkAlphaNumByte/Benchmark:_IsAlpha-8         	 1966162	       602 ns/op	      48 B/op	       2 allocs/op
    BenchmarkAlphaNumByte/Benchmark:_IsDigit-8         	 2037670	       542 ns/op	      48 B/op	       2 allocs/op
    BenchmarkAlphaNumByte/Benchmark:_IsAlphaNum-8      	 2000684	       606 ns/op	      48 B/op	       2 allocs/op
    BenchmarkAlphaNumByte/Benchmark:_IsAlphaSwitch-8   	 1953236	       610 ns/op	      48 B/op	       2 allocs/op
    */
    testBytesAlphaNum = []struct {
		name string
        f    func(c byte) bool // test function
        want func(r rune) bool // standard library check
	}{
		// TODO: Add test cases.
		{"IsASCIIAlpha", IsASCIIAlpha, unicodeIsASCIIAlpha},
		{"IsDigit", IsDigit, unicode.IsDigit},
		{"IsDigitSingleOP", IsDigitSingleOP, unicode.IsDigit},
		{"IsDigitSingleOPCompare", IsDigitSingleOPCompare, unicode.IsDigit},
		{"IsHex", IsHex, unicodeIsHex},
		{"IsAlphaNumUnder", IsAlphaNumUnder, unicodeIsAlphaNum},
		{"IsAlphaNum", IsAlphaNum, unicodeIsAlphaNum},
        {"IsAlphaNumSwitch", IsAlphaNumSwitch, unicodeIsAlphaNum},
        // {"IsASCIIPrintable", IsASCIIPrintable},
    }

    // testItemsAlphaNumBytes = []struct {
	// 	c byte
    // 	alpha bool
    //     digit bool
	// }{
	// 	{'\n', false, false},
	// 	{'\t', false, false},
	// 	{'\r', false, false},
	// 	{'\v', false, false},
	// 	{'\f', false, false},
	// 	{' ', false, false},
	// 	{'A', true, true},
	// 	{'c', true, true},
	// 	{'0', false, true},
	// 	{'7', false, true},
	// 	{'=', false, false},
	// 	{'?', false, false},
	// 	{'/', false, false},
	// 	{'%', false, false},
	// }
)

// unicodeIsAlphaNum uses the standard library unicode to test
// functionality of byte functions. Alphanumeric runes that are
// not valid ascii characters return false.
func unicodeIsAlphaNum(r rune) bool {
    if r > unicode.MaxASCII {
		return false
	}
	return unicode.IsLetter(r) || unicode.IsDigit(r)
}

// unicodeIsASCIIAlpha uses the standard library unicode to test
// functionality of byte functions. Alphabetic runes that are
// not valid ascii characters return false.
func unicodeIsASCIIAlpha(r rune) bool {
	if r > unicode.MaxASCII {
		return false
	}
	return unicode.IsLetter(r)
}

// unicodeIsASCIIAlpha uses the standard library unicode to test
// functionality of byte, rune, and string functions.
func unicodeIsHex(r rune) bool {
	return unicode.IsDigit(r) || ('a' <= r && r <= 'f') || ('A' <= r && r <= 'F')
}

// BenchmarkAlphaNumByte
/*
BenchmarkAlphaNumByte/Benchmark:_IsASCIIAlpha-8         	 1924887	       617 ns/op	      48 B/op	       2 allocs/op
BenchmarkAlphaNumByte/Benchmark:_IsDigit-8              	 2115994	       565 ns/op	      48 B/op	       2 allocs/op
BenchmarkAlphaNumByte/Benchmark:_IsHex-8                	 1925970	       617 ns/op	      48 B/op	       2 allocs/op
BenchmarkAlphaNumByte/Benchmark:_IsAlphaNum-8           	 2038138	       596 ns/op	      48 B/op	       2 allocs/op
BenchmarkAlphaNumByte/Benchmark:_IsAlphaNumSwitch-8     	 2014053	       595 ns/op	      48 B/op	       2 allocs/op
*/
func BenchmarkAlphaNumByte(b *testing.B) {
	benchmarks := testBytesAlphaNum
	for _, bb := range benchmarks {
		name := fmt.Sprintf("Benchmark: %s", bb.name)
		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				for _, c := range ByteSamples() {
					bb.f(c)
				}
			}
		})
	}
}

func TestAlphaNumBytes(t *testing.T) {
	tests := testBytesAlphaNum
	for _, tt := range tests {
		for _, c := range SmallByteSamples() {
			want := tt.want(rune(c)) // standard library
			name := fmt.Sprintf("Test %s (%q)", tt.name, c)
			t.Run(name, func(t *testing.T) {
				got := tt.f(c)
				if got != want {
					t.Errorf("%s = %v, want %v", name, got, want)
				}
			})
		}
	}
}

func TestIsAlphaNum2Underscore (t *testing.T) {
    		t.Run("underscore_test", func(t *testing.T) {
			if ok := IsAlphaNumUnder('_'); !ok {
				t.Errorf("IsASCIIPrintable(%v) = %v, want %v", "_",ok, true)
			}
		})
}

func TestIsASCIIPrintable(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
        // TODO: Add test cases.
        {"Hello, World!", args{"Hello, World!"}, true},
        {"Tab", args{"Testing this: \t"}, false},
        {"Newline", args{"Testing this: \n"}, false},
        {"©", args{"Testing this: ©"}, true},
        {"Block", args{"Testing this: ▓"}, true},
        {"Chr+255 (&nbsp;)", args{"Chr+255 (&nbsp;):  "}, false},
        {"Superscript 2", args{"Superscript 2: ²"}, true},
        {"Ã", args{"Testing this: Ã"}, true},
        {"Hello, World! with chr255 before !", args{"Hello, World !"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsASCIIPrintable(tt.args.s); got != tt.want {
				t.Errorf("IsASCIIPrintable(%v) = %v, want %v", tt.args.s,got, tt.want)
			}
		})
	}
}
