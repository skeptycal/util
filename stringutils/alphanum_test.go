package stringutils

import "testing"

func TestIsAlpha(t *testing.T) {
	type args struct {
		c byte
	}
	tests := []struct {
		args args
		want bool
	}{
		// TODO: Add test cases.
		{args{'\n'}, false},
		{args{'\t'}, false},
		{args{'\r'}, false},
		{args{'\v'}, false},
		{args{'\f'}, false},
		{args{' '}, false},
		{args{'A'}, true},
		{args{'c'}, true},
		{args{'0'}, false},
		{args{'7'}, false},
		{args{'='}, false},
		{args{'?'}, false},
		{args{'/'}, false},
		{args{'%'}, false},
	}
	for _, tt := range tests {
		t.Run(string(tt.args.c), func(t *testing.T) {
			if got := IsAlpha(tt.args.c); got != tt.want {
				t.Errorf("IsAlpha() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsDigit(t *testing.T) {
	type args struct {
		c byte
	}
	tests := []struct {
		args args
		want bool
	}{
		// TODO: Add test cases.
		{args{'\n'}, false},
		{args{'\t'}, false},
		{args{'\r'}, false},
		{args{'\v'}, false},
		{args{'\f'}, false},
		{args{' '}, false},
		{args{'A'}, false},
		{args{'c'}, false},
		{args{'0'}, true},
		{args{'7'}, true},
		{args{'='}, false},
		{args{'?'}, false},
		{args{'/'}, false},
		{args{'%'}, false},
	}
	for _, tt := range tests {
		t.Run(string(tt.args.c), func(t *testing.T) {
			if got := IsDigit(tt.args.c); got != tt.want {
				t.Errorf("IsDigit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsAlphaNum(t *testing.T) {
	type args struct {
		c byte
	}
	tests := []struct {
		args args
		want bool
	}{
		// TODO: Add test cases.
		{args{'\n'}, false},
		{args{'\t'}, false},
		{args{'\r'}, false},
		{args{'\v'}, false},
		{args{'\f'}, false},
		{args{' '}, false},
		{args{'A'}, true},
		{args{'c'}, true},
		{args{'0'}, true},
		{args{'7'}, true},
		{args{'='}, false},
		{args{'?'}, false},
		{args{'/'}, false},
		{args{'%'}, false},
	}
	for _, tt := range tests {
		t.Run(string(tt.args.c), func(t *testing.T) {
			if got := IsAlphaNum(tt.args.c); got != tt.want {
				t.Errorf("IsAlphaNum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsIsAlphaSwitch(t *testing.T) {
	type args struct {
		c byte
	}
	tests := []struct {
		args args
		want bool
	}{
		// TODO: Add test cases.
		{args{'\n'}, false},
		{args{'\t'}, false},
		{args{'\r'}, false},
		{args{'\v'}, false},
		{args{'\f'}, false},
		{args{' '}, false},
		{args{'A'}, true},
		{args{'c'}, true},
		{args{'0'}, true},
		{args{'7'}, true},
		{args{'='}, false},
		{args{'?'}, false},
		{args{'/'}, false},
		{args{'%'}, false},
	}
	for _, tt := range tests {
		t.Run(string(tt.args.c), func(t *testing.T) {
			if got := IsAlphaSwitch(tt.args.c); got != tt.want {
				t.Errorf("IsAlphaSwitch() = %v, want %v", got, tt.want)
			}
		})
	}
}
