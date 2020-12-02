package http

import (
	"testing"
)

func TestMakeQueryURL(t *testing.T) {
	tt := []struct {
		name     string
		params   []string
		prefix   string
		sep      string
		expected string
	}{
		{
			"chemistry lithium ion battery",
			[]string{"chemistry", "lithium", "ion", "battery"},
			"http://www.google.com/search?q=",
			"+",
			"http://www.google.com/search?q=chemistry+lithium+ion+battery",
		},
		{
			"my name but not the actor",
			[]string{"Michael", "Treanor", "developer", "-actor"},
			"http://www.google.com/search?q=",
			"+",
			"http://www.google.com/search?q=Michael+Treanor+developer+-actor",
		},
		{
			"golang theory filetype:pdf",
			[]string{"golang", "theory", "filetype:pdf"},
			"http://www.google.com/search?q=",
			"+",
			"http://www.google.com/search?q=golang+theory+filetype:pdf",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			result := MakeQueryURL(tc.params, tc.prefix, tc.sep)
			if result != tc.expected {
				t.Errorf("expected value <%s> does not match result: %v", tc.expected, result)
			}
		})
	}
}
