package main

import "github.com/skeptycal/util/datatools/math/fastinvsqrt"

func main() {

	var i interface{}
	var f float32 = 3.14

	b := fastinvsqrt.EncodeBits(f)
	println("b: ", b)
	val, ok := b.(uint32)
	println("val: ", val)

}
