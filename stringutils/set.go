package stringutils

import "fmt"

type (
	SetMap = map[Any]bool

	// Set is a hashable version of a list with unique items.
	Set struct {
		name string
		data SetMap
	}
)

// NewSet returns a new Set from the given List
func NewSet(name string, data []Any) *Set {
	s := Set{}
	s.name = name
	s.data = make(SetMap, len(data))
	for _, v := range data {
		s.data[v] = true
	}
	return &s
}

// Add adds item to the Set or returns an error.
// Duplicates are not allowed.
func (s *Set) Add(item Any) error {
	if _, ok := s.data[item]; !ok {
		return fmt.Errorf("item %v could not be added to Set %v", item, s.name)
	}
	s.data[item] = true
	return nil
}

// Contains returns true if the Set contains item.
func (s *Set) Contains(item Any) bool {
	// if _, ok := s.data[item]; ok {
	return s.data[item]
}

// Len returns of elements in the Set
// If the Set is nil, Len() is zero.
func (s *Set) Len() int {
	return len(s.data)
}

// Cap returns the max number of elements in the Set
// (since cap is undefined for map types in go).
func (s *Set) Cap() int {
	return len(s.data)
}

// Name returns the name of the Set.
func (s *Set) Name() string {
	return s.name
}

// ToSlice returns the underlying data as a slice.
func (s *Set) ToSlice() []Any {
	return s.ToList().ToSlice()
}

// ToList returns the underlying data as a List.
func (s *Set) ToList() *List {
	v := make([]Any, 0, s.Len())
	for k := range s.data {
		v = append(v, k)
	}
	return &List{s.name, v}
}
