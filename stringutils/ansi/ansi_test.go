package ansi

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"testing"
)

var testWriter = NewANSIWriter(44, 33, 1, os.Stdout)

func TestNewANSIWriter(t *testing.T) {
	want := &AnsiWriter{
		*bufio.NewWriter(os.Stdout),
		DefaultAnsiFmt,
	}
	got := NewANSIWriter(44, 33, 1, os.Stdout)
	t.Run("NewANSIWriter test", func(t *testing.T) {
		if got.String() != want.String() {
			t.Errorf("NewANSIWriter = %v, want %v", got.String(), want.String())
		}
	})
}

func ExampleWrap() {
	testWriter.Wrap("wrap this")
}

func ExampleBuild() {
	testWriter.Build(1, 32)
}

func TestConstants(t *testing.T) {
	tests := []struct {
		name string
		a    interface{}
		want interface{}
	}{
		// TODO: Add test cases.
		{"format string: ansi", fmt.Sprintf(FMTansi, 32), "\033[32m;"},
		{"format string: bright", fmt.Sprintf(FMTbright, 32), "\033[1;32m;"},
		{"format string: dim", fmt.Sprintf(FMTdim, 32), "\033[2;32m;"},
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
