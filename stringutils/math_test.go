package stringutils

import (
	"testing"
)

// benchmark results
/* /// integer powers are ~6x faster than float 64 in this example.
cpu: Intel(R) Core(TM) i7-4770HQ CPU @ 2.20GHz
BenchmarkPow-8       	206724968	         5.499 ns/op	       0 B/op	       0 allocs/op
BenchmarkMathPow-8   	 41840359	         29.03 ns/op	       0 B/op	       0 allocs/op
*/

func BenchmarkPow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		pow(5, 5)
	}
}

func BenchmarkMathPow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mathPow(5, 5)
	}
}

func Test_pow(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		// {"0^0", args{0, 0}, 0}, // float conversion issues for math.Pow() tests ...
		{"0^1", args{0, 1}, 0},
		{"0^2", args{0, 2}, 0},
		{"1^0", args{1, 0}, 1},
		{"1^1", args{1, 1}, 1},
		{"-1^1", args{-1, 1}, -1},
		{"-1^3", args{-1, 3}, -1},
		{"-1^4", args{-1, 4}, 1},
		{"1^2", args{1, 2}, 1},
		{"2^2", args{2, 2}, 4},
		{"2^3", args{2, 3}, 8},
		{"3^2", args{3, 2}, 9},
		{"3^3", args{3, 3}, 27},
		{"4^4", args{4, 4}, 256},
		{"17^12", args{17, 12}, 582622237229761},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pow(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("pow(%s) = %v, want %v", tt.name, got, tt.want)
			}
		})
		t.Run(tt.name, func(t *testing.T) {
			if got := mathPow(float64(tt.args.x), float64(tt.args.y)); int(got) != tt.want {
				t.Errorf("math.pow(%s) = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}

func TestTenToThe(t *testing.T) {
	type args struct {
		y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"10^2", args{2}, 100},
		{"10^7", args{7}, 10000000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TenToThe(tt.args.y); got != tt.want {
				t.Errorf("TenToThe() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTwoToThe(t *testing.T) {
	type args struct {
		y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"2^0", args{0}, 1},
		{"2^1", args{1}, 2},
		{"2^2", args{2}, 4},
		{"2^8", args{8}, 256},
		{"2^10", args{10}, 1024},
		{"2^16", args{16}, 65536},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TwoToThe(tt.args.y); got != tt.want {
				t.Errorf("TwoToThe() = %v, want %v", got, tt.want)
			}
		})
	}
}
