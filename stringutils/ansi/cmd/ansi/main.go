package main

import (
	"fmt"

	"github.com/skeptycal/util/stringutils/ansi"
)

func main() {
    a := ansi.NewAnsiSet(ansi.StyleBold)
    a.SetColors(35,44,ansi.Bold)
    ansi.SetupCLI(a)
    fmt.Println((a.String()))
    fmt.Println(a.Info())
    fmt.Println("a.String() before ",a.String(), "...after")


    w := ansi.NewAnsiWriter(nil)



    w.WriteString("WriteString tests\n")
    w.Write([]byte("Write Text\n"))

   s :=  w.String()
   fmt.Println(s)

   fmt.Println("Flush ... ")
   w.Flush()

   w.Wrap("Wrap test\n")

   w.Build(ansi.Italics, ansi.Red, ansi.YellowBackground)

   fmt.Println("Build test (should be: italic red text on yellow background")

   fmt.Print("still ansi color \n")
   fmt.Print(ansi.Reset)
   fmt.Print("color off\n")




}
