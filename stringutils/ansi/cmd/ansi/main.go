package main

import (
	"fmt"

	"github.com/skeptycal/util/stringutils/ansi"
)

func main() {
    w := ansi.NewANSIWriter(nil)
    w.WriteString("WriteString tests\n")
    w.Write([]byte("Write Text\n"))

   s :=  w.String()
   fmt.Println(s)

   fmt.Println("Flush ... ")
   w.Flush()

   w.Wrap("Wrap test\n")

   w.Build(ansi.Italics, ansi.Red,ansi.YellowBackground)

   fmt.Println("Build test (should be: italic red text on yellow background")
}
