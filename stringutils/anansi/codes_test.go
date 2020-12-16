package anansi

import (
	"testing"
)

//* --------------------------------------------------------> ansiCodes type definition
func Test_ansiCodes_Set(t *testing.T) {
	type fields struct {
		fg Attribute
		bg Attribute
		ef Attribute
	}
	type args struct {
		fg Attribute
		bg Attribute
		ef Attribute
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
			a := ansiCodes{
				fg: tt.fields.fg,
				bg: tt.fields.bg,
				ef: tt.fields.ef,
			}
			a.Set(tt.args.fg, tt.args.bg, tt.args.ef)
		})
	}
}

func Test_ansiCodes_Reset(t *testing.T) {
	type fields struct {
		fg Attribute
		bg Attribute
		ef Attribute
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := ansiCodes{
				fg: tt.fields.fg,
				bg: tt.fields.bg,
				ef: tt.fields.ef,
			}
			if got := a.Reset(); got != tt.want {
				t.Errorf("ansiCodes.Reset() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_ansiCodes_Print(t *testing.T) {
	type fields struct {
		fg Attribute
		bg Attribute
		ef Attribute
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := ansiCodes{
				fg: tt.fields.fg,
				bg: tt.fields.bg,
				ef: tt.fields.ef,
			}
			a.Print()
		})
	}
}

func Test_ansiCodes_String(t *testing.T) {
	type fields struct {
		fg Attribute
		bg Attribute
		ef Attribute
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := ansiCodes{
				fg: tt.fields.fg,
				bg: tt.fields.bg,
				ef: tt.fields.ef,
			}
			if got := a.String(); got != tt.want {
				t.Errorf("ansiCodes.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_ansiCodes_Wrap(t *testing.T) {
	type fields struct {
		fg Attribute
		bg Attribute
		ef Attribute
	}
	type args struct {
		s []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := ansiCodes{
				fg: tt.fields.fg,
				bg: tt.fields.bg,
				ef: tt.fields.ef,
			}
			if got := a.Wrap(tt.args.s...); got != tt.want {
				t.Errorf("ansiCodes.Wrap() = %v, want %v", got, tt.want)
			}
		})
	}
}
