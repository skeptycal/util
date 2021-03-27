package stringutils

// Any is used to store data when the type cannot be
// determined ahead of time.
type (
	Any interface{}
)

// List is a wrapper around a slice of items. It offers
// formatting options and convenience functions.
type List struct {
	name string
	list []Any
}

// NewList returns a new List from the given data.
func NewList(name string, data []Any) *List {
	return &List{name, data}
}

// Contains tells whether a contains x.
func (v *List) Contains(item Any) bool {
	for i := range v.list {
		if i == item {
			return true
		}
	}
	return false
}

// Add adds item to the List
// Duplicates are allowed.
func (v *List) Add(item Any) {
	v.list = append(v.list, item)
}

// Len returns of count of elements in the Set.
// If the Set is nil, Len() is zero.
func (v *List) Len() int {
	return len(v.list)
}

// Cap returns the max number of elements in the List.
func (v *List) Cap() int {
	return cap(v.list)
}

// Name returns the name of the List.
func (v *List) Name() string {
	return v.name
}

// ToSlice returns the underlying data as a slice.
func (v *List) ToSlice() []Any {
	return v.list
}

// ToSet returns the underlying data as a Set.
func (v *List) ToSet() *Set {
	return NewSet(v.name, v.list)
}
