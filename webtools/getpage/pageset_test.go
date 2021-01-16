package getpage

import (
	"strings"
	"sync"
	"testing"
)

var (
	t = NewPageSet
)

type testPages SafePageSet

// func (t testPages) addFakePages(n int) {
// 	for i := 0; i < n; i++ {
// 		t.Add(fmt.Sprintf("FakeURL%n", i))
// 	}
// }

func TestSafePageSet_Len(t *testing.T) {
	type fields struct {
		mu    sync.Mutex
		pages map[string]*strings.Builder
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SafePageSet{
				mu:    tt.fields.mu,
				pages: tt.fields.pages,
			}
			if got := s.Len(); got != tt.want {
				t.Errorf("SafePageSet.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}
