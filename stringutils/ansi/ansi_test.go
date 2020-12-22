package ansi

import (
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
		{"test1", Underline, 4},
		{"test1", Blue, 34},
		{"test1", Bold, 1},
		{"test1", BlueText, "\x1b[34m"},
		{"test1", BoldYellowText, "\x1b[1;33m"},
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := itoa(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("itoa() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_ansi_String(t *testing.T) {
	tests := []struct {
		name string
		a    Ansi
		want string
	}{
		// TODO: Add test cases.
		{"test1", 34, "/x1b[34;"},
		{"test1", 44, "/x1b[44;"},
		{"test1", 1, "/x1b[1;"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.String(); got != tt.want {
				t.Errorf("ansi.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
