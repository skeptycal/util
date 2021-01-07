package main

import (
	"fmt"
	"go/parser"
	"go/token"
	"math"

	"github.com/skeptycal/util/datatools/math/tree"
)

func RunFuncSeries(min, max, step int, f func(args ...interface{}) ([]interface{}, error)) {

	var retval []interface{}

	for _, v := range args {

	}

	for i := min; i < max; i += step {
		a, err := f(args)
		if err != nil {
			break
		}
		retval = append(retval, a)
	}

	return retval, nil
}

// showTable creates a table of example data using
// a range of [min .. max] as x values and inputs for f
func showTable(min, max, cols int, f func(args ...interface{}) ([]interface{}, error)) {

	for i := min; i < max; i++ {
		c, err := f(i) // uint8(math.Pow(2, float64(i)))
		if err != nil {
			println(err)
		}
		fmt.Printf("%4d : %08b  ", c, c)
		if (i+1)%cols == 0 {
			fmt.Print("\n")
		}
	}
	fmt.Print("\n")
}

func main() {
	// err := floatingints.BpPush()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// ------------------------------------------------------------

	const maxNum = 8
	const colNum = 8

	var (
		u8  uint8  = 0
		u16 uint16 = 0
		u32 uint32 = 0
		u64 uint64 = 0
		i8  int8   = 0
		i16 int16  = 0
		i32 int32  = 0
		i64 int64  = 0

		s = tree.Things{u8, u16, u32, u64, i8, i16, i32, i64}
	)

	for i, v := range s {
		fmt.Printf("  thing%d ( %-7T) : %v\n", i, v, v)
	}

	showTable(1, maxNum, 2, math.Pow10(v))

	// ------------------------------------------------------------

	fset := token.NewFileSet() // positions are relative to fset

	src := `package foo

import (
    "fmt"
    "time"
)

func bar() {
    fmt.Println(time.Now())
}`

	// Parse src but stop after processing the imports.
	f, err := parser.ParseFile(fset, "", src, parser.ImportsOnly)
	if err != nil {
		fmt.Println(err)
		return
	}
	// Print the imports from the file's AST.
	for _, s := range f.Imports {
		fmt.Println(s.Path.Value)
	}
}
