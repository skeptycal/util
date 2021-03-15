package main

import (
	"fmt"
	"os"
	"strconv"

	goflag "flag"

	flag "github.com/spf13/pflag"
)

func main() {
	flag.CommandLine.AddGoFlagSet(goflag.CommandLine)

	binary := flag.Bool("binary", true, "list binary equivalents")
	hex := flag.Bool("hex", true, "list hex equivalents")
	dec := flag.Bool("decimal", true, "list decimal equivalents")

	flags.BoolP("verbose", "v", false, "verbose output")
	flags.String("coolflag", "yeaah", "it's really cool flag")
	flags.Int("usefulflag", 777, "sometimes it's very useful")
	flags.SortFlags = false
	flags.PrintDefaults()

	// deprecate a flag by specifying its name and a usage message
	// the following is here as an example of flag usage only.
	flag.MarkDeprecated("decimal", "deprecated - int values are shown by default.")

	// range := flag.String() // ("r", "32-126", "ASCII range to list (default printable characters: 32-126)")

	flag.Parse()

	fmt.Println("args: ", os.Args[1:])
	fmt.Println("binary: ", *binary)
	fmt.Println("hex: ", *hex)

	const maxValue = 30
	var input string

	// whitespace examples
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
		fmt.Printf("%4d   %#02x  %+- 6s   %#08b\n", i, i, strconv.QuoteRuneToASCII(rune(i)), i)
		if i%32 == 0 && i != 0 {
			fmt.Scanln(&input)
			if input == "q" || input == "Q" {
				break
			}
		}
	}
}
