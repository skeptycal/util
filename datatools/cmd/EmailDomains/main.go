package main

import (
	"os"
	"strings"

	"github.com/skeptycal/util/datatools/format/email"
)

func main() {
    list := strings.Join(os.Args[1:]," ")
    out := email.GetDomainNames(list)

}
