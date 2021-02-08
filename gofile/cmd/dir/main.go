package main

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/skeptycal/zsh"
)

func DirShellBM(path string) string {
	command := fmt.Sprintf("ls -R %s", path)
	result := zsh.Out(command)
	return result
}

func DirShellTimerBM(path string) string {
	command := fmt.Sprintf("time ls -R . %s", path)
	result := zsh.Out(command)
	return result
}

func main() {

	testpath := "/Users/skeptycal/local_coding"
	fmt.Println("Directory Listing Benchmarks:\n ")
	log.Info("logger started")
	fmt.Println("")
	fmt.Println(DirShellBM(testpath))
}
