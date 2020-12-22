package main

import (
	"flag"

	"github.com/skeptycal/util/coreutils"
)

func main() {

	flag.Parse()

	args := coreutils.OptionSet()
	args.Parse()

	opt := new(coreutils.Option)
}
