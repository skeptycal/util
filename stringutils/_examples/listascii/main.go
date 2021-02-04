package main

import (
	"fmt"
	"strconv"
)

func main() {

    const maxValue = 30
    var input string

    /*
   8     8         1000
   9     9         1001
  10     a         1010
  11     b         1011
  12     c         1100
  13     d         1101
  14     e         1110
  15     f         1111
  */

    const mask = 0b00001111

    for i := 0; i < maxValue; i++ {
        fmt.Printf("%4d   %#02x  %+- 6s   %#08b\n",i,i,strconv.QuoteRuneToASCII(rune(i)),i )
        if i % 32 == 0 && i != 0 {
            fmt.Scanln(&input)
            if input == "q" || input == "Q" {
                break
            }
        }
    }
}
