package ascii

import (
	"testing"
)

func TestToASCII(t *testing.T) {
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
			if got := ToASCII(tt.args.s); got != tt.want {
				t.Errorf("ToASCII() = %s, want %s", got, tt.want)
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
		{"-0x1.23abcp+20", args{-0x1.23abcp+20}, "-0x1.23abcp+20"},
		{"nil", args{nil}, "0x0"},
		{"empty slice", args{[]byte{}}, "NaN"},
		{"slice", args{[]byte{50}}, "NaN"},
		{"map", args{make(map[int]int)}, "NaN"},
		{"map", args{make(map[int]int, 4)}, "NaN"},
		{"true", args{true}, "0x1"},
		{"false", args{false}, "0x0"},
		{"string", args{"string"}, "NaN"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToHex(tt.args.v); got != tt.want {
				t.Errorf("Hexed() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToOctal(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"1.234", args{1.234}, "float"},
		{"2.43e22", args{2.43e22}, "float"},
		{"0.00000008", args{0.00000008}, "float"},
		{"2", args{2}, "0o2"},
		{"17", args{17}, "0o21"},
		{"100", args{100}, "0o144"},
		{"255", args{255}, "0o377"},
		{"256", args{256}, "0o400"},
		{"1000000000", args{1000000000}, "0o7346545000"},
		{"32767", args{32767}, "0o77777"},
		{"1e12", args{1e12}, "float"},
		{"-0x1.23abcp+20", args{-0x1.23abcp+20}, "float"},
		{"nil", args{nil}, "0o0"},
		{"true", args{true}, "0o1"},
		{"false", args{false}, "0o0"},
		{"string", args{"string"}, "NaN"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToOctal(tt.args.v); got != tt.want {
				t.Errorf("ToOctal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToBinary(t *testing.T) {
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
		{"2", args{2}, "10"},
		{"17", args{17}, "10001"},
		{"100", args{100}, "1100100"},
		{"255", args{255}, "11111111"},
		{"256", args{256}, "100000000"},
		{"1000000000", args{1000000000}, "111011100110101100101000000000"},
		{"32767", args{32767}, "111111111111111"},
		{"1e12", args{1e12}, "8192000000000000p-13"},
		{"-0x1.23abcp+20", args{-0x1.23abcp+20}, "5131128709054464p-32"},
		{"nil", args{nil}, "0x0"},
		{"empty slice", args{[]byte{}}, "NaN"},
		{"slice", args{[]byte{50}}, "NaN"},
		{"map", args{make(map[int]int)}, "NaN"},
		{"map", args{make(map[int]int, 4)}, "NaN"},
		{"true", args{true}, "0x1"},
		{"false", args{false}, "0x0"},
		{"string", args{"string"}, "NaN"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToBinary(tt.args.v); got != tt.want {
				t.Errorf("ToBinary() = %v, want %v", got, tt.want)
			}
		})
	}
}
