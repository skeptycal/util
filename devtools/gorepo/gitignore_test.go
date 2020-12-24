package gorepo

import (
	"reflect"
	"testing"
)

func Test_gitIgnoreAPIList(t *testing.T) {
	tests := []struct {
		name string
		want []string
	}{
		// TODO: Add test cases.
		{"python", []string{""}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := gitIgnoreAPIList(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("gitIgnoreAPIList() = %v, want %v", got, tt.want)
			}
		})
	}
}
