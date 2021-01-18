package ansi

import (
	"fmt"
	"reflect"
	"testing"
)

func TestConstants(t *testing.T) {
	tests := []struct {
		name string
		a    interface{}
		want interface{}
	}{
		// TODO: Add test cases.
		{"format string: ansi", fmt.Sprintf(FMTansi, 32), ""},
		{"Underline", Underline, byte(4)},
		{"Blue", Blue, byte(34)},
		{"Bold", Bold, byte(1)},
		{"BlueText", BlueText, "\x1b[34m"},
		{"BoldYellowText", BoldYellowText, "\x1b[1;33m"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a; got != tt.want {
				t.Errorf("constant value = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_itoa(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		// TODO: Add test cases.
		{"42", args{42}, []byte{52, 50}},
		{"testing", args{128}, []byte{49, 50, 56}},
		{"42", args{255}, []byte{50, 53, 53}},
		{"-1", args{-1}, []byte{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := itoa(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("itoa() = %v, want %v", got, tt.want)
			}
		})
	}
}
