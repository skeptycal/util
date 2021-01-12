// A concurrent prime sieve

package main

import (
	"fmt"
	"testing"
)

func BenchmarkMain(b *testing.B) {
	for j := 0; j < 20; j++ {
		for i := 0; i < b.N; i++ {
			main()
		}
		fmt.Printf("%v", b.Name())
	}
}

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
