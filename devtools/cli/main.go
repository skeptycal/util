package main

import (
	"context"
	"fmt"
	"os/exec"

	"github.com/skeptycal/util/zsh"
)

func main() {
	var (
		cmd                *exec.Cmd
		background         = context.Background()
		app, args, command string
	)

	for {
		fmt.Scan

		app, args = zsh.AppArgs(command)
		cmd = exec.CommandContext(background, app, args)
		zsh.GetStdin(cmd)
	}
}
