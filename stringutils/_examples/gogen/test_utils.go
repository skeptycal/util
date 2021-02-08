package gogen

import (
	"fmt"
	"reflect"
	"testing"
)

func ErrTest(name string, test interface{}, wantErr bool, t *testing.T) (typ string, ok bool) {
    // var ok bool = false

    switch typ := i.(type) {
        case nil:
            ok = false
		case bool:
			ok = test == true
		case int:
			ok = test != 0
        case float64, float32:
            ok = test != 0.0
        case string:
            ok = test != ""
        case interface{}:
            ok = ErrTest(name, test, wantErr, t)
		default:
			ok = true
        return ok
	}
    return ok
}

func GotWant(name, subname string, got, want interface{}, wantErr bool, t *testing.T) (err error) {

	if !reflect.DeepEqual(got, want) != wantErr {
		// t.Errorf("NewRepo() = %v, want %v", got, want)
		err = fmt.Errorf("%s.%s got: %v want: %v", name, subname, got, want)
	}
	return
}
