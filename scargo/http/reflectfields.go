package http

import (
	"fmt"
	"net/http"
	"reflect"
)

// PrintResp - print response fields in a neater way
func PrintResp(title string, resp *http.Response) error {

	println(title)

	// x := string(resp.)

	// println(x)
	// for i, v := range *resp. {

	// }

	return nil
}

func reflectFields() {
	x := struct {
		Foo string
		Bar int
	}{"foo", 2}

	v := reflect.ValueOf(x)

	values := make([]interface{}, v.NumField())

	for i := 0; i < v.NumField(); i++ {
		values[i] = v.Field(i).Interface()
	}

	fmt.Println(values)
}
