package ansi

import (
	"testing"
)

func TestToAscii(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"to ascii", args{"to ascii"}, "to ascii"},
		{"Theta (Θ)", args{"Theta (Θ)"}, "Theta (\u0398)"},
		{"Theta (Θ)", args{"Theta (Θ)"}, "Theta (\u0398)"},
		{"Theta (Θ)", args{"Theta (\u0398)"}, "Theta (Θ)"},
		{"newline", args{"new\\nline"}, "new\\nline"},
		{"newline", args{"new\nline"}, "new\nline"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToAscii(tt.args.s); got != tt.want {
				t.Errorf("ToAscii() = %s, want %s", got, tt.want)
			}
		})
	}
}

func TestQuoted(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"to ascii", args{"to ascii"}, "\"to ascii\""},
		{"Theta (Θ)", args{"Theta (Θ)"}, "\"Theta (Θ)\""},
		{"Theta (Θ)", args{"Theta (Θ)"}, "\"Theta (\u0398)\""},
		{"Theta (Θ)", args{"Theta (\u0398)"}, "\"Theta (Θ)\""},
		{"newline", args{"new\nline"}, "\"new\\nline\""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Quoted(tt.args.s); got != tt.want {
				t.Errorf("Quoted() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHexed(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"1.234", args{1.234}, "0x1.3be76c8b43958p+00"},
		{"2.43e22", args{2.43e22}, "0x1.49538f9945421p+74"},
		{"0.00000008", args{0.00000008}, "0x1.5798ee2308c3ap-24"},
		{"2", args{2}, "2"},
		{"17", args{17}, "11"},
		{"100", args{100}, "64"},
		{"255", args{255}, "ff"},
		{"256", args{256}, "100"},
		{"1000000000", args{1000000000}, "3b9aca00"},
		{"32767", args{32767}, "7fff"},
		{"1e12", args{1e12}, "0x1.d1a94a2p+39"},
		{"-0x1.23abcp+20", args{-0x1.23abcp+20}, "3"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Hexed(tt.args.v); got != tt.want {
				t.Errorf("Hexed() = %v, want %v", got, tt.want)
			}
		})
	}
}
