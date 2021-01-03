package float

import "fmt"

type Float int16

func (f *Float) String() string {
	return fmt.Sprintf("%e", f)
}
