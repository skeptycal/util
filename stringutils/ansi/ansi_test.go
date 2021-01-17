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

func Test_ansi_String(t *testing.T) {
	tests := []struct {
		name string
		a    AnsiWriter
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

func TestAnsi_Build(t *testing.T) {
	var ansi AnsiWriter
	type args struct {
		list []AnsiWriter
	}
	tests := []struct {
		name string
		a    AnsiWriter
		args args
		want string
	}{
		// TODO: Add test cases.
		{"Blue", ansi, args{[]AnsiWriter{AnsiWriter(Blue)}}, "/x1b[34;"},
		{"Blue and Bold", ansi, args{[]AnsiWriter{AnsiWriter(34), AnsiWriter(Bold)}}, "/x1b[34;/x1b[1;"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.Build(tt.args.list...); got != tt.want {
				t.Errorf("Ansi.Build() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMarshalAnsi(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
		{"empty", args{""}, []byte{}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MarshalAnsi(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("MarshalAnsi() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MarshalAnsi() = %v, want %v", got, tt.want)
			}
		})
	}
}
