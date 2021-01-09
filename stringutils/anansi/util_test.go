package anansi

import "testing"

func TestBR(t *testing.T) {
	tests := []struct {
		name    string
		wantN   int
		wantErr bool
	}{
		// TODO: Add test cases.
		{"", 1, false},
		{"", 1, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotN, err := BR()
			if (err != nil) != tt.wantErr {
				t.Errorf("BR() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotN != tt.wantN {
				t.Errorf("BR() = %v, want %v", gotN, tt.wantN)
			}
		})
	}
}
