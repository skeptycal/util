// Package format contains functions that format numeric values.
package format

import (
	"testing"
)

func TestNumSpace(t *testing.T) {
	type args struct {
		n string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"12345.54321e42", args{"12345.54321e42"}, "12345.54321e42"},
		{"1", args{"1"}, "1"},
		{"-1", args{"-1"}, "-1"},
		{"0.123", args{"0.123"}, "0.123"},
		{"-43.3234e-105", args{"-43.3234e-105"}, "-43.3234e-105"},
		{"1234567890.09876543210", args{"1234567890.09876543210"}, "1234567890.09876543210"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NumSpace(tt.args.n); got != tt.want {
				t.Errorf("NumSpace() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkReverse-8   	18962943	        60.7 ns/op	      16 B/op	       2 allocs/op
func BenchmarkReverse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Reverse("12345")
	}
}

// BenchmarkReverse2-8   	13703583	        83.6 ns/op	       8 B/op	       1 allocs/op

func BenchmarkReverse2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Reverse2("12345")
	}
}

// BenchmarkNumSpaces-8   	 2261641	       520 ns/op	      80 B/op	      12 allocs/op

func BenchmarkNumSpaces(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NumSpace("12345.54321e-42")
	}
}

func TestReverse(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"hello", args{"hello"}, "olleh"},
		{"12345", args{"12345"}, "54321"},
		{"dot.net", args{"dot.net"}, "ten.tod"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Reverse(tt.args.s); got != tt.want {
				t.Errorf("Reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
