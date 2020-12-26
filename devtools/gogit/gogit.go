package gogit

import (
	"fmt"
	"strings"

	"github.com/skeptycal/util/zsh"
)

const (
	gitCommitFormatString = `git commit -m '%s'`
)

// GitCommit creates a commit with message
func GitCommit(message string) error {
	s := fmt.Sprintf(gitCommitFormatString, message)
	return zsh.Status(s)
}

// GitCommit creates a commit with message
func GitCommitAll(message string) error {
	err := zsh.Status("git add --all")
	if err != nil {
		return err
	}
	return GitCommit(message)
}

// GitInit initializes the Git environment in the current directory with:
//  git init
func GitInit() error {
	err := zsh.Status("git init")
	if err != nil {
		return err
	}
	return GitCommitAll("initial commit")
}

// GitTag Create, list, delete or verify a tag object signed with GPG
func GitTag(s string) error {
	// todo check tag with regex
	args := strings.Split(s, " ")
	command := strings.TrimSpace(args[0])
	fmt.Printf("command: %s", command)
	tag := s[1:]
	return zsh.Status(fmt.Sprintf("git tag %s", tag))
}
