// A concurrent prime sieve

package main

import "testing"

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"works on my machine (tm)"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}
