package stringutils

import (
	"fmt"
	"testing"
	"unicode"
)


func BenchmarkIsASCIISpace(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, c := range ByteSamples() {
			IsASCIISpace(c)
		}
	}
}
func BenchmarkIsWhiteSpace(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, c := range ByteSamples() {
			isWhiteSpace(c)
		}
	}
}

func BenchmarkIsWhiteSpace2(b *testing.B) {
	// 	case ' ', '\t', '\n', '\f', '\r', '\v':
	for i := 0; i < b.N; i++ {
		for _, r := range ByteSamples() {
			isWhiteSpace2(r)
		}
	}
}

func BenchmarkIsWhiteSpaceLogicChain(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, r := range RuneSamples() {
			isWhiteSpaceLogicChain(r)
		}
	}
}

func BenchmarkIsUnicodeWhiteSpaceMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, r := range RuneSamples() {
			IsUnicodeWhiteSpaceMap(r)
		}
	}
}

func BenchmarkIsUnicodeWhiteSpaceMapSwitch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, r := range RuneSamples() {
			IsUnicodeWhiteSpaceMapSwitch(r)
		}
	}
}

func BenchmarkUnicode_IsSpace(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, r := range RuneSamples() {
			unicode.IsSpace(r)
		}
	}
}



func TestDedupeWhitespace(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DedupeWhitespace(tt.args.s); got != tt.want {
				t.Errorf("DedupeWhitespace() = %v, want %v", got, tt.want)
			}
		})
	}
}


func TestIsWhiteSpace(t *testing.T) {
	for _, c := range SmallByteSamples() {
		name := fmt.Sprintf("IsWhiteSpace: %v", c)
		t.Run(name, func(t *testing.T) {
			got := isWhiteSpace(c)
			want := unicode.IsSpace(rune(c))
			if got != want {
				t.Errorf("IsWhiteSpace() = %v, want %v", got, want)
			}
		})
	}
}
func TestIsWhiteSpace2(t *testing.T) {
	for _, c := range SmallByteSamples() {
		name := fmt.Sprintf("IsWhiteSpace2: %v", c)
		t.Run(name, func(t *testing.T) {
			got := isWhiteSpace2(c)
			want := unicode.IsSpace(rune(c))
			if got != want {
				t.Errorf("IsWhiteSpace2() = %v, want %v", got, want)
			}
		})
	}
}
func TestIsUnicodeWhiteSpaceMap(t *testing.T) {
	for _, c := range SmallRuneSamples() {
		t.Run( fmt.Sprintf("IsUnicodeWhiteSpaceMap: (%q)", c), func(t *testing.T) {
            if got := IsUnicodeWhiteSpaceMap(c); got != want(c) {
				t.Errorf("IsWhiteSpace3() = %v, want %v", got, want)
			}
		})
	}
}
func TestIsWhiteSpace4(t *testing.T) {
	for _, c := range SmallRuneSamples() {
		name := fmt.Sprintf("IsWhiteSpace4: %q", c)
		t.Run(name, func(t *testing.T) {
			got := isWhiteSpace4(c)
			want := unicode.IsSpace(c)
			if got != want {
				t.Errorf("IsWhiteSpace4(%q) = %v, want %v", c, got, want)
			}
		})
	}
}
func TestIsWhiteSpace5(t *testing.T) {
	for _, c := range SmallRuneSamples() {
		name := fmt.Sprintf("IsWhiteSpace5: %v", c)
		t.Run(name, func(t *testing.T) {
			got := isWhiteSpaceBoolMap(c)
			want := unicode.IsSpace(c)
			if got != want {
				t.Errorf("IsWhiteSpace5(%q) = %v, want %v", c, got, want)
			}
		})
	}
}
