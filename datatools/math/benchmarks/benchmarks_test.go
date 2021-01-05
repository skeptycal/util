/**

(x 1.15) ModAnd is somewhat faster for tests using a series of numbers containing a mix of odds/evens (n*3)
BenchmarkMod-8      	1000000000	         0.551 ns/op	       0 B/op	       0 allocs/op
BenchmarkModAnd-8   	1000000000	         0.478 ns/op	       0 B/op	       0 allocs/op

Only [integers] mod powers of 2

(x 1.50) Only (multiples of 8) mod (multiples of 2)
BenchmarkMod-8      	1000000000	         0.722 ns/op	       0 B/op	       0 allocs/op
BenchmarkModAnd-8   	1000000000	         0.482 ns/op	       0 B/op	       0 allocs/op

(x 1.05) Only multiples of 2
BenchmarkMod-8      	1000000000	         0.515 ns/op	       0 B/op	       0 allocs/op
BenchmarkModAnd-8   	1000000000	         0.494 ns/op	       0 B/op	       0 allocs/op

(x 1.53) Even better with only even numbers
BenchmarkMod-8      	1000000000	         0.390 ns/op	       0 B/op	       0 allocs/op
BenchmarkModAnd-8   	1000000000	         0.255 ns/op	       0 B/op	       0 allocs/op
*/
package benchmarks

import (
	"math"
	"testing"
)

const (
	loop_start = 1
	loop_step  = 1
	multipleA  = 8
	multipleB  = 2
)

func BenchmarkMod(b *testing.B) {
	for i := loop_start; i < b.N; i = i + loop_step {
		mod(i*multipleA, i)
	}
}

func BenchmarkModAnd(b *testing.B) {
	for i := loop_start; i < b.N; i = i + loop_step {
		modAnd(i*multipleA, int(math.Pow(2, float64(i))))
	}
}

func Test_mod(t *testing.T) {
	type args struct {
		n int
		r int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"3, 2", args{3, 2}, 1},
		{"8, 2", args{8, 2}, 0},
		{"27, 3", args{27, 3}, 0},
		{"48, 4", args{48, 4}, 0},
		{"32768, 256", args{32768, 256}, 0},
		{"0xFFFF, 5", args{0xFFFF, 5}, 0},
		{"0xFFFFFFFF, 0xF}", args{0xFFFFFFFF, 0xF}, 0},
		{"0xFFFFFFFF, 0xF", args{0xFFFFFFFF, 8}, 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mod(tt.args.n, tt.args.r); got != tt.want {
				t.Errorf("mod(%v, %v) = %v, want %v", tt.args.n, tt.args.r, got, tt.want)
			}
		})
	}
}

func Test_modAnd(t *testing.T) {
	type args struct {
		n int
		r int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases
		{"3, 2", args{3, 2}, 1},
		{"8, 2", args{8, 2}, 0},
		{"48, 4", args{48, 4}, 0},
		{"48, 4", args{48, 3}, 0},
		{"96, 8", args{96, 8}, 0},
		{"32768, 16", args{32768, 16}, 0},
		{"0xFFFFFFFF, 0b1010", args{0xFFFFFFFF, 0b1010}, 9},
		{"0x1234567890, 16", args{0x1234567890, 16}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := modAnd(tt.args.n, tt.args.r); got != tt.want {
				t.Errorf("modAnd(%v, %v) = %v, want %v", tt.args.n, tt.args.r, got, tt.want)
			}
		})
	}
}
