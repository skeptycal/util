package main

import (
	"fmt"

	"github.com/skeptycal/util/zsh"
)

func main() {
	ws := zsh.NewTTY()
	fmt.Println(ws)
}
