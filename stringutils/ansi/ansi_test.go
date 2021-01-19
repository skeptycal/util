package ansi

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"testing"
)

var testWriter = NewANSIWriter(os.Stdout)

func ExampleSetupCLI() {
	SetupCLI(&defaultAnsiSet)
	// Output:
	// c[39;49;0m
}

func TestNewANSIWriter2(t *testing.T) {
	want := &AnsiWriter{
		*bufio.NewWriter(os.Stdout),
		&defaultAnsiSet,
	}
	got := NewANSIWriter(os.Stdout)
	t.Run("NewANSIWriter test", func(t *testing.T) {
		if got.String() != want.String() {
			t.Errorf("NewANSIWriter = %v, want %v", got.String(), want.String())
		}
	})
	want = &AnsiWriter{
		*bufio.NewWriter(os.Stdout),
		&defaultAnsiSet,
	}
	got = NewANSIWriter(nil)
	t.Run("NewANSIWriter test", func(t *testing.T) {

		if got.String() != want.String() {
			t.Errorf("NewANSIWriter = %v, want %v", got.String(), want.String())
		}
	})
	want = &AnsiWriter{
		*bufio.NewWriter(os.Stdout),
		&defaultAnsiSet,
	}
	got = NewANSIWriter(want)
	t.Run("NewANSIWriter test", func(t *testing.T) {

		if got.String() != want.String() {
			t.Errorf("NewANSIWriter = %v, want %v", got.String(), want.String())
		}
	})
	want = &AnsiWriter{
		*bufio.NewWriter(os.Stdout),
		&defaultAnsiSet,
	}
	fakeFile, _ := os.Open("/tmp")
	got = NewANSIWriter(fakeFile)
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
	// [39;49;0mhello, world![39m[49m[0m
	// [39;49;0mhello, Mike[39m[49m[0m
	// [39;49;0mhi, Mike and world[39m[49m[0m
}

func ExampleAPrint() {
	APrint(1, 32)
	// Output:
	// [1m[32m
}

func ExampleAnsiSet_String() {
	testAnsiSet := NewAnsiSet(35, 44, 4)
	fmt.Print(testAnsiSet.String())
	// Output:
	// [35;44;4m
}

func ExampleAnsiSet_info() {
	testAnsiSet := NewAnsiSet(35, 44, 4)
	fmt.Print(testAnsiSet.info())
	// Output:
	// fg: 35, bg: 44, ef 4
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
		Writer bufio.Writer
		ansi   AnsiSet
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
		{"Wrap: smoke test", fields{*bufio.NewWriter(os.Stdout), defaultAnsiSet}, args{""}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &AnsiWriter{
				Writer: tt.fields.Writer,
				ansi:   &tt.fields.ansi,
			}
			a.Wrap(tt.args.s)
		})
	}
}

func TestAnsiWriter_Build(t *testing.T) {
	devNull, _ := os.Open("/dev/null")
	type fields struct {
		Writer bufio.Writer
		ansi   AnsiSet
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
		{"Wrap: smoke test", fields{*bufio.NewWriter(defaultioWriter), defaultAnsiSet}, args{[]byte{33, 44, 1}}},
		{"Wrap: smoke test", fields{*bufio.NewWriter(devNull), defaultAnsiSet}, args{[]byte{33, 44, 1}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &AnsiWriter{
				Writer: tt.fields.Writer,
				ansi:   &tt.fields.ansi,
			}
			a.Build(tt.args.b...)
		})
	}
}

func TestNewAnsiSet(t *testing.T) {
	type args struct {
		fg color
		bg color
		ef color
	}
	tests := []struct {
		name string
		args args
		want *AnsiSet
	}{
		// TODO: Add test cases.
		{"32,42,2", args{32, 42, 2}, NewAnsiSet(32, 42, 2)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAnsiSet(tt.args.fg, tt.args.bg, tt.args.ef); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAnsiSet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAnsiWriter_SetColors(t *testing.T) {
	type fields struct {
		Writer bufio.Writer
		ansi   *AnsiSet
	}
	type args struct {
		s *AnsiSet
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &AnsiWriter{
				Writer: tt.fields.Writer,
				ansi:   tt.fields.ansi,
			}
			a.SetColors(tt.args.s)
		})
	}
}
