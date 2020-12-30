package gofile

import (
	"fmt"
	"testing"
)

func TestUserInfo(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UserInfo(); got != tt.want {
				t.Errorf("UserInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleUserInfo(t *testing.T) {
	s := UserInfo()
	fmt.Println(s)

	// stuff
}
