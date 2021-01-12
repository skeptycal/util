package main

import (
	"fmt"

	"github.com/skeptycal/util/datatools/benchmark/multibench"
)

func main() {
	fmt.Println("Multi1: ", Multi1(5, 6))
}

func Multi1(x, y int) int {
	return x * y
}

func main1() {

	// funcList := make([]Func, 20)

	// funcList is a list of all function versions tested
	// funcList := []Func{
	// 	multi1,
	// 	multi2,
	// 	multi3,
	// 	multi5,
	// 	multi6,
	// 	multi7,
	// 	StandardMulti,
	// }

	funcList := make([]multibench.Func, 20)

	// argsList is a list of all sets of arguments tested
	var argsList = []struct{ x, y int }{{3, 4}, {12, 34}, {3498, 1843}}

	// fmt.Println("funcList: ", funcList)
	for _, f := range funcList {
		fmt.Println("argsList: ", argsList)
		for _, testargs := range argsList {
			fmt.Println(f(testargs.x, testargs.y))

			tt := multibench.NewFunc(testargs.x, testargs.y, f)
			fmt.Printf("function: %v", tt.Name())

			got := tt.Got()
			want := tt.Want()
			if got != want {
				fmt.Print(fmt.Errorf("multi1() = %v, want %v", got, want))
			}
		}
	}
}
