package stringbenchmarks

import (
	"fmt"
	"reflect"
	"testing"
)

var (
	tempSlice []Any = []Any{"this", 1, nil, 0, 3.14, '\t'}
	tempList  *List = NewList("tempList", tempSlice)
	tempSet   *Set  = NewSet("tempSet", tempSlice)
)

func TestNewList(t *testing.T) {
	want := &List{"testList", tempSlice}
	got := NewList("testList", tempSlice)

	t.Run(want.name, func(t *testing.T) {
		if !reflect.DeepEqual(got, want) {
			t.Errorf("NewList = %v, want %v", got, want)
		}
	})
}

func ExampleNewList() {
	fmt.Println(tempList)
	// output:
	// &{tempList [this 1 <nil> 0 3.14 9]}
}

func TestList_ToSet(t *testing.T) {
	t.Run(tempList.name, func(t *testing.T) {
		got := tempList.ToSet()
		want := NewSet("tempList", tempSlice)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Got: %v Want: %v", got, want)
		}
	})
}

func ExampleList() {
	// List.Contains()
	fmt.Println(tempList.Contains(3.14))
	fmt.Println(tempList.Contains(42))
	// List.Len()
	fmt.Println(tempList.Len())
	// List.Cap()
	fmt.Println(tempList.Cap())
	// List.Name()
	fmt.Println(tempList.Name())
	// List.Add()
	fmt.Println(tempList.Contains("fake"))
	tempList.Add("fake")
	// fmt.Println(tempList.Contains("fake"))

	// output:
	// false
	// false
	// 6
	// 6
	// tempList
	// false

}
