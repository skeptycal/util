package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"

	"github.com/skeptycal/util/zsh"
)

func main() {

	// this is the demo command
	//  tr a-z A-Z
	cmd := exec.Command("tr", "a-z", "A-Z")
	err := zsh.GetStdin(cmd)
	if err != nil {
		log.Fatal(err)
	}

	var out bytes.Buffer
	cmd.Stdout = &out
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("in all caps: %q\n", out.String())
}
