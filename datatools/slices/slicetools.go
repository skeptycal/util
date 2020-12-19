package slices

import "reflect"

// FirstPointerIdx returns the index of the first pointer in the slice.
//
// Reference: https://stackoverflow.com/a/36891042
func FirstPointerIdx(s []interface{}) int {
	for i, v := range s {
		if reflect.ValueOf(v).Kind() == reflect.Ptr {
			return i
		}
	}
	return -1
}
