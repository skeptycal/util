package replacer

import (
	"fmt"
	"io"
	"strings"
)

// Replace creates a new Replacer and reads from a list of old, new string pairs. Replacements are performed in the order they appear in the target string, without overlapping matches. The old string comparisons are done in argument order.
// todo - not implemented
func NewReplacer(oldnew ...string) (*strings.Replacer, error) {
	if len(oldnew)%2 != 0 {
		return nil, fmt.Errorf("NewReplacer requires an even number of arguments")
	}
	r := strings.NewReplacer(oldnew...)
	// cmd :=
	return r, nil
}

// replacer is the interface that a replacement algorithm needs to implement.
type Replacer interface {
	Replace(s string) string
	WriteString(w io.Writer, s string) (n int, err error)
}
