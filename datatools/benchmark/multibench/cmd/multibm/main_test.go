package main

import "testing"

func BenchmarkMulti1(b *testing.B) {

	for i := 0; i < b.N; i++ {
		Multi1(6, 5)
	}
}

func Test_multi1(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"3 x 4", args{3, 4}, 12},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := multi1(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("multi1() = %v, want %v", got, tt.want)
			}
		})
	}
}
