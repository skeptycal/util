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
	want = &AnsiWriter{
		*bufio.NewWriter(os.Stdout),
		DefaultAnsiFmt,
	}
	got = NewANSIWriter(0, 0, 0, nil)
	t.Run("NewANSIWriter test", func(t *testing.T) {

		if got.String() != want.String() {
			t.Errorf("NewANSIWriter = %v, want %v", got.String(), want.String())
		}
	})
	want = &AnsiWriter{
		*bufio.NewWriter(os.Stdout),
		DefaultAnsiFmt,
	}
	fakeFile, _ := os.Open("/dev")
	got = NewANSIWriter(44, 33, 1, fakeFile)
	t.Run("NewANSIWriter test", func(t *testing.T) {
		if got.String() != want.String() {
			t.Errorf("NewANSIWriter = %v, want %v", got.String(), want.String())
		}
	})

}

func ExampleANSI_Wrap() {
	testWriter.Wrap("wrap this")
}

func ExampleEcho() {
	Echo("hello, world!")
	Echo("hello, %s", "Mike")
	Echo("hi, Mike", " and ", "world")
	// Output:
	// [44m[33m[1mhello, world![0m[39m[49m
	// [44m[33m[1mhello, Mike[0m[39m[49m
	// [44m[33m[1mhi, Mike and world[0m[39m[49m
}

func ExampleAPrint() {
	APrint(1, 32)
	// Output:
	// [1m[32m
}

func ExampleANSI_Build() {
	testWriter.Build(1, 32)
}

func ExampleCLS() {
	CLS()
	// Output:
	// c
}

func ExampleHR() {
	HR(10)
	// Output:
	// ==========
}

func ExampleBR() {
	BR()
	// Output:
	//
}

func TestConstants(t *testing.T) {
	tests := []struct {
		name string
		a    interface{}
		want interface{}
	}{
		// TODO: Add test cases.
		{"format string: ansi", fmt.Sprintf(FMTansi, 32), "\033[32m"},
		{"format string: bright", fmt.Sprintf(FMTbright, 32), "\033[1;32m"},
		{"format string: dim", fmt.Sprintf(FMTdim, 32), "\033[2;32m"},
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

func TestAnsiWriter_Wrap(t *testing.T) {
	type fields struct {
		Writer     bufio.Writer
		ansiString string
	}
	type args struct {
		s string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
		{"Wrap: smoke test", fields{*bufio.NewWriter(os.Stdout), DefaultAnsiFmt}, args{""}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &AnsiWriter{
				Writer:     tt.fields.Writer,
				ansiString: tt.fields.ansiString,
			}
			a.Wrap(tt.args.s)
		})
	}
}

func TestAnsiWriter_Build(t *testing.T) {
	devNull, _ := os.Open("/dev/null")
	type fields struct {
		Writer     bufio.Writer
		ansiString string
	}
	type args struct {
		b []byte
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
		{"Wrap: smoke test", fields{*bufio.NewWriter(defaultioWriter), DefaultAnsiFmt}, args{[]byte{33, 44, 1}}},
		{"Wrap: smoke test", fields{*bufio.NewWriter(devNull), DefaultAnsiFmt}, args{[]byte{33, 44, 1}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &AnsiWriter{
				Writer:     tt.fields.Writer,
				ansiString: tt.fields.ansiString,
			}
			a.Build(tt.args.b...)
		})
	}
}
