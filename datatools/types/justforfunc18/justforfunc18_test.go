package justforfunc18

import (
	"testing"
)

func Test_celsius_String(t *testing.T) {
	tests := []struct {
		name string
		c    celsius
		want string
	}{
		// TODO: Add test cases.
		{"10.0", celsius(10.0), "10.00 °C"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.String(); got != tt.want {
				t.Errorf("celsius.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_celsius_Value(t *testing.T) {
	tests := []struct {
		name string
		c    celsius
		want float64
	}{
		// TODO: Add test cases.
		{"10.0", celsius(10.0), 10.00},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Value(); got != tt.want {
				t.Errorf("celsius.Value() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_celsius_Farenheit(t *testing.T) {
	tests := []struct {
		name string
		c    celsius
		want string
	}{
		// TODO: Add test cases.
		{"10.0", celsius(10.0), "50.00 °F"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Farenheit(); got != tt.want {
				t.Errorf("celsius.Farenheit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewCelsius(t *testing.T) {
	type args struct {
		f float64
	}
	tests := []struct {
		name string
		args args
		want celsius
	}{
		// TODO: Add test cases.
		{"107 to 41.67 ", args{95 + 12}, NewCelsius(107)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCelsius(tt.args.f); got != tt.want {
				t.Errorf("NewCelsius() = %v, want %v", got, tt.want)
			}
		})
	}
}
