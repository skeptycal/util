package floatingints

import "fmt"

type float int16

func (f *float) String() string {
	return fmt.Sprintf("%e", f)
}
