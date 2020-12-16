package main

import "fmt"

type Shape interface {
	Area() float64
	Perimeter() float64
}

func both() (int, int) {
	return 1, 2
}

func thing() string {
	x := 0
	x, y := both()
	return fmt.Sprintf("%d, %d", x, y)
}

func main() {
	// println(thing())
	var s Shape
	fmt.Println(s)
	fmt.Printf("type: %T\n", s)
}
