// Package watch implements live reloading and timed reloading of command line programs.
package watch

import (
	"time"

	"github.com/skeptycal/util/zsh"
)

func Run(command string) error {
	return zsh.Status(command)
}

func Watch(path string) error {
	return nil
}

func DoOver(path string, timeout time.Duration) error {
	for {
		err := Run(path)
		if err != nil {
			return err
		}
		time.Sleep(timeout)
	}
}
