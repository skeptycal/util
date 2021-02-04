package stringutils

// Any is used to store data when the type cannot be
// determined ahead of time.
type Any interface{}

type Set struct {
    name string
    data map[Any]bool
}

type List struct {
    name string
    list []Any
    Stringer
}

// Contains tells whether a contains x.
func (v *List) Contains(x string) bool {
	for _, n := range v.list {
		if n.(string) == x {
			return true
		}
	}
	return false
}

func (v *List) Len() int {
    return len(v.list)
}

func (v *List) Cap() int {
    return cap(v.list)
}

func (v *List) Name() string {
    return v.name
}
