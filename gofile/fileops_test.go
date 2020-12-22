package gofile

import "testing"

func TestChunkMultiple(t *testing.T) {
	tt := []struct {
		name     string
		n        int64
		chunk    int64
		expected int64
	}{
		{
			"1/1",
			1,
			1,
			1,
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			result := chunkMultiple(tc.n, tc.chunk)
			if result != tc.expected {
				t.Errorf("expected value <%s> does not match result: %v", tc.expected, result)
			}
		})
	}
}
