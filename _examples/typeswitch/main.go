package main

import (
	"fmt"
)

// TypeIt returns the type of the value.
//
// https://golang.org/doc/effective_go.html#type_switch
func TypeIt(t interface{}) string {

	switch t := t.(type) {
	default:
		fmt.Printf("unexpected type %T\n", t) // %T prints whatever type t has
	case bool:
		fmt.Printf("boolean %t\n", t) // t has type bool
	case int:
		fmt.Printf("integer %d\n", t) // t has type int
	case *bool:
		fmt.Printf("pointer to boolean %t\n", *t) // t has type *bool
	case *int:
		fmt.Printf("pointer to integer %d\n", *t) // t has type *int
	}
	return fmt.Sprintf("%T\n", t)
}

func main() {

	i := 42
	b := false

	fmt.Printf("%s\n", TypeIt(10))
	fmt.Printf("%s\n", TypeIt(true))
	fmt.Printf("%s\n", TypeIt(&i))
	fmt.Printf("%s\n", TypeIt(&b))

	fmt.Printf("%d\n", FirstPointerIdx([]interface{}{i, b, &i, &i, &b}))

	dict := map[string]string{"foo": "1", "bar": "2"}
	fmt.Println("Value of Key in Dict: ", maptools.InMap("foo", dict))
	fmt.Println("Value of Key in Dict: ", maptools.InMap("fee", dict))

}
