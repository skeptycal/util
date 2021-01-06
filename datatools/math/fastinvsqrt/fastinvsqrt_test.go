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
		{"NaN + nil", args{things: []interface{}{math.NaN(), nil}}, float64(math.NaN())},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddAny(tt.args.things...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddAny(%v, %v) = %v (%T), want %v(%T)", tt.args.things[0], tt.args.things[1], got, got, tt.want, tt.want)
			}
		})
	}
}
