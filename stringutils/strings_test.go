// Copyright 2020 Michael Treanor. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package stringutils implements additional functions to
// support the go standard library strings module.
//
// The algorithms chosen are based on benchmarks from
// the stringbenchmarks module. ymmv...
//
// The current implementation at the start of this project was
// .../go/1.15.3/libexec/src/strings/strings.go
//
// For information about UTF-8 strings in Go,
// see https://blog.golang.org/strings.
package stringutils

import "testing"

func Test_benStringIndex(t *testing.T) {
	type args struct {
		s      string
		substr string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := benStringIndex(tt.args.s, tt.args.substr); got != tt.want {
				t.Errorf("benStringIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
