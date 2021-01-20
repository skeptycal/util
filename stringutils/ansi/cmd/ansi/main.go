package main

import "github.com/skeptycal/util/stringutils/ansi"

func main() {
    w := ansi.NewANSIWriter(nil)
    w.WriteString("WriteString tests")
    w.Write("Write Text")
}
