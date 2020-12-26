package main

import (
	"flag"

	log "github.com/sirupsen/logrus"
)

const (
	urlRoot = "https://www.toptal.com/developers/gitignore/api/"
)

func main() {

	// https://www.toptal.com/developers/gitignore/api/macos,linux,windows,ssh,vscode,go,zsh

	flagForced := flag.Bool("f", false, "force overwrite of previous .gitignore file (if it exists)")
	flagSkip := flag.Bool("s", false, "skip download of .gitignore items from gitignore.io")

	flag.Parse()

	giFile, err := NewGitIgnore(*flagForced, *flagSkip)
	if err != nil {
		log.Fatal(err)
	}

	// w := bufio.NewWriter(giFile)

	// w = w // todo - finish up ...

}
