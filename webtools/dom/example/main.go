package main

import (
	"fmt"
	"syscall/js"
)

func main() {
	fmt.Println("Hello, WebAssembly!") // original example

	doc := js.Global().Get("document")
	body := doc.Call("getElementById", "thebody")
	body.Set("innerHTML", "Dynamic Content")
}
