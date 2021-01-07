// package fastinvsqrt

package fastinvsqrt

import (
	"math"
	"reflect"
	"testing"
)

func BenchmarkInvSqrtBasic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// Pi is to slow down the benchmark somewhat
		invSqrtBasic(float64(i) * math.Pi)
	}
}

func Test_invSqrtBasic(t *testing.T) {
	type args struct {
		x float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
		{"64", args{64.0}, 1.0 / 8.0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := invSqrtBasic(tt.args.x); got != tt.want {
				t.Errorf("invSqrtBasic() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_AddAny(t *testing.T) {

	// special procedure for testing NaN results
	/*
	   according to IEEE 754, NaN is not equivalent to itself, so equality tests require
	   the use of the math.IsNaN() function
	   --- FAIL: Test_AddAny (0.00s)
	   --- FAIL: Test_AddAny/NaN_+_nil (0.00s)
	       github.com/skeptycal/util/datatools/math/fastinvsqrt/fastinvsqrt_test.go:62:
	       AddAny(NaN, <nil>) = NaN(float64), want NaN(float64)

	               {"NaN + nil", args{things: []interface{}{math.NaN(), nil}}, float64(math.NaN())},
	*/
	t.Run("AddAny NaN result", func(t *testing.T) {
		if got := AddAny(math.NaN(), int(3)); !math.IsNaN(got.(float64)) {
			t.Errorf("AddAny() = %v, want %v", got, math.NaN())
		}
	})

	type args struct {
		things []interface{}
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		// TODO: Add test cases.
		{"byte + int", args{things: []interface{}{byte(32), int(32)}}, "64"},
		{"float32 + uint16", args{things: []interface{}{float32(32), uint16(32)}}, "64"},
		{"int16 + uint16", args{things: []interface{}{int16(32), uint16(32)}}, "64"},
		{"string + int", args{things: []interface{}{"32", 32}}, "64"},
		{"func + int", args{things: []interface{}{AddAny(), 32}}, "32"},
		{"func + nil", args{things: []interface{}{AddAny(), nil}}, nil},
		// NaN cannot be tested with an equality
		// {"NaN + nil", args{things: []interface{}{math.NaN(), nil}}, float64(math.NaN())},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddAny(tt.args.things...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddAny(%v, %v) = %v(%T), want %v(%T)", tt.args.things[0], tt.args.things[1], got, got, tt.want, tt.want)
			}
		})
	}
}

func Test_parabola(t *testing.T) {
	type args struct {
		x float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
		{"2", args{2}, 4},
		{"-2", args{-2}, 4},
		{"3", args{3}, 9},
		{"1.5854332985", args{1.5854332985}, 2.5135987439925898},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parabola(tt.args.x); got != tt.want {
				t.Errorf("parabola() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_derivativeOfParabola(t *testing.T) {
	type args struct {
		x float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
		{"2", args{2}, 4},
		{"-2", args{-2}, -4},
		{"3", args{3}, 6},
		{"1.5854332985", args{1.5854332985}, 3.170866597},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := derivativeOfParabola(tt.args.x); got != tt.want {
				t.Errorf("derivativeOfParabola() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEncodeBits(t *testing.T) {
	type args struct {
		f float32
	}
	tests := []struct {
		name string
		args args
		want uint32
	}{
		// TODO: Add test cases.
		{"3.14", args{3.14}, 1078523331},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EncodeBits(tt.args.f); got != tt.want {
				t.Errorf("EncodeBits() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecodeBits(t *testing.T) {
	type args struct {
		b uint32
	}
	tests := []struct {
		name string
		args args
		want float32
	}{
		// TODO: Add test cases.
		{"back to 3.14 (answer from TestEncodeBits)", args{EncodeBits(3.14)}, 3.14},
		{"mantissaBitMask", args{uint32(mantissaBitMask)}, 1.1754942e-38},
		{"3.14 encoded", args{uint32(1078523331)}, 3.14},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DecodeBits(tt.args.b); got != tt.want {
				t.Errorf("DecodeBits() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBits_Shift(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		b    Bits
		args args
		want uint32
	}{
		// TODO: Add test cases.
		{"Bits(zero)", Bits(0), args{1}, 0},
		{"Bits(zero)", 0, args{1}, 0},
		{"Bits(signBitMask)", Bits(signBitMask), args{0}, signBitMask},
		{"Bits(signBitMask)", Bits(signBitMask), args{1}, 0},
		{"Bits(all32BitMask)", Bits(all32BitMask), args{0}, all32BitMask},
		{"Bits(all32BitMask)", Bits(all32BitMask), args{31}, signBitMask},
		{"Bits(all32BitMask)", Bits(all32BitMask), args{0}, all32BitMask},
		{"Bits(mantissaBitMask)", Bits(mantissaBitMask), args{0}, 0x7FFFFF},
		{"Bits(1)", Bits(1), args{1}, 2},
		{"Bits(1)", Bits(1), args{2}, 4},
		{"Bits(1)", Bits(1), args{3}, 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.Shift(tt.args.n); got != tt.want {
				t.Errorf("Bits.Shift() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBits_Bytes(t *testing.T) {
	tests := []struct {
		name string
		b    Bits
		want []byte
	}{
		// TODO: Add test cases.
		{"", Bits(0), []byte{0, 0, 0, 0, 0}},
		{"Bits(signBitMask)", Bits(signBitMask), []byte{0, 0, 0, 0, 0}},
		{"Bits(signBitMask)", Bits(signBitMask), []byte{0, 0, 0, 0, 0}},
		{"Bits(all32BitMask)", Bits(all32BitMask), []byte{0, 0, 0, 0, 255}},
		{"Bits(all32BitMask)", Bits(all32BitMask), []byte{0, 0, 0, 0, 255}},
		{"Bits(all32BitMask)", Bits(all32BitMask), []byte{0, 0, 0, 0, 255}},
		{"Bits(mantissaBitMask)", Bits(mantissaBitMask), []byte{0, 0, 0, 0, 255}},
		{"Bits(1)", Bits(1), []byte{0, 0, 0, 0, 1}},
		{"Bits(1)", Bits(2), []byte{0, 0, 0, 0, 2}},
		{"Bits(1)", Bits(0xFFFFFF), []byte{0, 0, 0, 0, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.Bytes(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Bits.Bytes() = %v, want %v", got, tt.want)
			}
		})
	}
}
