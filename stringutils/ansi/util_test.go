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
