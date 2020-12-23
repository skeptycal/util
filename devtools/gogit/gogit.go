package gogit

import (
	"github.com/skeptycal/util/gofile"
	. "github.com/skeptycal/util/zsh"
)

// GitCommit creates a commit with message
func GitCommit(message string) error {
	Shell("git add --all")
	Shell("git commit -m '" + message + "'")
	return nil
}

// GitCommit creates a commit with message
func GitCommitAll(message string) error {
	Shell("git add --all")
	Shell("git commit -m '" + message + "'")
	return nil
}

// GitInit initializes the Git environment
func gitInit() error {
	if !gofile.Exists(".gitignore") {
		gorepo.GitIgnore("", "")
	}

	Shell("git init")
	GitCommitAll("initial commit")
	return nil
}
