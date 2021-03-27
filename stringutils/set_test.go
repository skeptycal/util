package stringutils

import (
	"fmt"
	"testing"
)

func ExampleNewSet() {
	fmt.Println(tempSet)
	// output:
	// &{tempList map[<nil>:2 3.14:4 0:3 1:1 9:5 this:0]}
}

func TestSet_ToSlice(t *testing.T) {
	// checks for items in Set using Get(item) ... also tests
	// the intermediate function ToList()
	for _, got := range tempSet.ToSlice() {
		t.Run(tempSet.name, func(t *testing.T) {
			if !tempSet.Contains(got) {
				t.Errorf("slice item %v not contained in Set %v", got, tempSet.name)
			}
		})
	}
}

func ExampleSet() {
	// Set.Contains()
	fmt.Println(tempSet.Contains(3.14))
	fmt.Println(tempSet.Contains(42))
	// Set.Len()
	fmt.Println(tempSet.Len())
	// Set.Cap()
	fmt.Println(tempSet.Cap())
	// Set.Name()
	fmt.Println(tempSet.Name())
	// Set.Add()
	fmt.Println(tempSet.Contains("fake"))
	_ = tempSet.Add("fake")
	fmt.Println(tempSet.Contains("fake"))

	// output:
	// 4 <nil>
	// true
	// false
	// 6
	// 6
	// tempList
	// <nil> item fake not found in Set tempList
	// 7 <nil>

}
