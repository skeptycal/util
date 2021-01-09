package main

import (
	"fmt"

	"github.com/skeptycal/util/devtools/gorepo/godirs"
)

func main() {
	for k, v := range godirs.Commands {
		fmt.Printf("  %v : %v\n", k, v)
	}

}
