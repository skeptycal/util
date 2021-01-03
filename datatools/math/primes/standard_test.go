package primes

import (
	"fmt"
	"testing"
)

func BenchmarkFib10(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		Fib(10)
	}
}

func TestFib(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"3", args{3}, 2},
		{"10", args{10}, 55},
		{"20", args{20}, 6765},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Fib(tt.args.n); got != tt.want {
				t.Errorf("Fib() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkIsPrime(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		IsPrime(n)
	}
}
func TestIsPrime(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"2", args{2}, 0},
		{"5", args{5}, 0},
		{"55", args{55}, 5},
		{"6", args{6}, 1},
		{"12", args{12}, 2},
		{"15", args{15}, 5},
		{"21", args{21}, 3},
		{"35", args{35}, 5},
		{"97", args{97}, -1},
		{"100", args{100}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPrime(tt.args.n); got != tt.want {
				t.Errorf("IsPrime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkSumDigitsStr(b *testing.B) {
	// BenchmarkSumDigitsStr-8   	 9055327	       139 ns/op	      16 B/op	       1 allocs/op
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		sumDigitsStr(fmt.Sprintf("%d", n))
	}
}
func Test_sumDigitsStr(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"five", args{"5"}, 5},
		{"35", args{"35"}, 8},
		{"247", args{"247"}, 13},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumDigitsStr(tt.args.s); got != tt.want {
				t.Errorf("sumDigitsStr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkSumDigits(b *testing.B) {
	// BenchmarkSumDigits-8   	94464870	        18.0 ns/op	       0 B/op	       0 allocs/op
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		sumDigits(n)
	}
}

func Test_sumDigits(t *testing.T) {
	type args struct {
		number int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"five", args{5}, 5},
		{"35", args{35}, 8},
		{"247", args{247}, 13},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumDigits(tt.args.number); got != tt.want {
				t.Errorf("sumDigits() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsOdd(t *testing.T) {
	type args struct {
		n uint8
	}
	tests := []struct {
		name string
		args args
		want uint8
	}{
		// TODO: Add test cases.
		{"five", args{5}, 1},
		{"six", args{6}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsOdd(tt.args.n); got != tt.want {
				t.Errorf("IsOdd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsEven(t *testing.T) {
	type args struct {
		n uint8
	}
	tests := []struct {
		name string
		args args
		want uint8
	}{
		// TODO: Add test cases.
		{"five", args{5}, 0},
		{"six", args{6}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsEven(tt.args.n); got != tt.want {
				t.Errorf("IsEven() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsEvenAnd(t *testing.T) {
	type args struct {
		n uint8
	}
	tests := []struct {
		name string
		args args
		want uint8
	}{
		// TODO: Add test cases.
		{"five", args{5}, 1},
		{"six", args{6}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsEvenAnd(tt.args.n); got != tt.want {
				t.Errorf("IsEvenAnd() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestAnd(t *testing.T) {
	type args struct {
		n uint8
		m uint8
	}
	tests := []struct {
		name string
		args args
		want uint8
	}{
		// TODO: Add test cases.
		{"75 & 112", args{75, 112}, 64},
		{"0 & 0", args{0, 0}, 0},
		{"1 & 0", args{1, 0}, 0},
		{"0 & 1", args{0, 1}, 0},
		{"1 & 1", args{1, 1}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := And(tt.args.n, tt.args.m); got != tt.want {
				t.Errorf("And() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestXor(t *testing.T) {
	type args struct {
		n uint8
		m uint8
	}
	tests := []struct {
		name string
		args args
		want uint8
	}{
		// TODO: Add test cases.
		{"75 ^ 112", args{75, 112}, 59},
		{"0 ^ 0", args{0, 0}, 0},
		{"1 ^ 0", args{1, 0}, 1},
		{"0 ^ 1", args{0, 1}, 1},
		{"1 ^ 1", args{1, 1}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Xor(tt.args.n, tt.args.m); got != tt.want {
				t.Errorf("Xor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOr(t *testing.T) {
	type args struct {
		n uint8
		m uint8
	}
	tests := []struct {
		name string
		args args
		want uint8
	}{
		// TODO: Add test cases.
		{"75 | 112", args{75, 112}, 123},
		{"0 | 0", args{0, 0}, 0},
		{"1 | 0", args{1, 0}, 1},
		{"0 | 1", args{0, 1}, 1},
		{"1 | 1", args{1, 1}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Or(tt.args.n, tt.args.m); got != tt.want {
				t.Errorf("Or() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestShR(t *testing.T) {
	type args struct {
		n uint8
		m uint8
	}
	tests := []struct {
		name string
		args args
		want uint8
	}{
		// TODO: Add test cases.
		{"75 >> 112", args{75, 112}, 0},
		{"0 >> 0", args{0, 0}, 0},
		{"1 >> 0", args{1, 0}, 1},
		{"0 >> 1", args{0, 1}, 0},
		{"1 >> 1", args{1, 1}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ShR(tt.args.n, tt.args.m); got != tt.want {
				t.Errorf("ShR() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestShL(t *testing.T) {
	type args struct {
		n uint8
		m uint8
	}
	tests := []struct {
		name string
		args args
		want uint8
	}{
		// TODO: Add test cases.
		{"75 << 112", args{75, 112}, 0},
		{"0 << 0", args{0, 0}, 0},
		{"1 << 0", args{1, 0}, 1},
		{"0 << 1", args{0, 1}, 0},
		{"1 << 1", args{1, 1}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ShL(tt.args.n, tt.args.m); got != tt.want {
				t.Errorf("ShL() = %v, want %v", got, tt.want)
			}
		})
	}
}
