package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"os/exec"

	"github.com/skeptycal/util/gofile"
	"github.com/skeptycal/util/zsh"
)

func main() {
	var (
		cmd                *exec.Cmd
		background         = context.Background()
		app, args, command string
		promptString       string = gofile.PWD() + "\n➜ "
	)

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(promptString)
		text, _ := reader.ReadString('\n')
		fmt.Println(text)

		app, args = zsh.AppArgs(command)
		cmd = exec.CommandContext(background, app, args)
		zsh.GetStdin(cmd)
	}
}
