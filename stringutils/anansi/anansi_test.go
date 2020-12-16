package anansi

import (
	"testing"
)

func TestAnsi(t *testing.T) {
	type args struct {
		color Attribute
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"FgBlack color test", args{FgBlack}, "\x1b[30m"},
		{"FgRed color test", args{FgRed}, "\x1b[31m"},
		{"FgGreen color test", args{FgGreen}, "\x1b[32m"},
		{"FgYellow color test", args{FgYellow}, "\x1b[33m"},
		{"FgBlue color test", args{FgBlue}, "\x1b[34m"},
		{"FgMagenta color", args{FgMagenta}, "\x1b[35m"},
		{"FgCyan color test", args{FgCyan}, "\x1b[36m"},
		{"FgWhite color test", args{FgWhite}, "\x1b[37m"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Ansi(tt.args.color); got != tt.want {
				t.Errorf("Ansi() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAllAnsi(t *testing.T) {
	type args struct {
		fg Attribute
		bg Attribute
		ef Attribute
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"Normal white on black", args{FgWhite, BgBlack, Reset}, "\x1b[37m\x1b[40m\x1b[0m"},
		{"Bold Red on Yellow", args{FgRed, BgYellow, Bold}, "\x1b[31m\x1b[43m\x1b[1m"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AllAnsi(tt.args.fg, tt.args.bg, tt.args.ef); got != tt.want {
				t.Errorf("AllAnsi() = %v, want %v", got, tt.want)
			}
		})
	}
}

// todo - fix this test
// func TestResetAll(t *testing.T) {
// 	tests := []struct {
// 		name string
// 		want string
// 	}{
// 		// TODO: Add test cases.
// 		{name: "Reset all to default", want: "\x1b[37m\x1b[40m\x1b[0m"},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := ResetAll(); got != tt.want {
// 				t.Errorf("ResetAll() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

func TestPrint(t *testing.T) {
	type args struct {
		s []string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{"Test Print", args{[]string{"Test", "Print"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Print(tt.args.s...)
		})
	}
}

func TestPrintln(t *testing.T) {
	type args struct {
		s []string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{"Test Println", args{[]string{"Test", "Println"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Println(tt.args.s...)
		})
	}
}

func TestSample(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{"Test Sample"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Sample()
		})
	}
}

func Test_PadInt(t *testing.T) {
	type args struct {
		i    int
		size int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"Pad integer 1 to 4 spaces", args{1, 4}, "   1"},
		{"Pad integer 348321 to 10 spaces", args{348321, 10}, "    348321"},
		{"Pad integer 234 to 7 spaces", args{-234, 7}, "   -234"},
		{"Pad integer 100 to 2 spaces", args{100, 2}, "100"},
		{"Pad integer -42 to -1 spaces", args{-42, -1}, "-42"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PadInt(tt.args.i, tt.args.size); got != tt.want {
				t.Errorf("padInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
