package main

import (
	"bufio"
	"flag"
	"os"

	log "github.com/sirupsen/logrus"
)

type giFile struct {
	bufio.ReadWriter
	filename string
}

// NewGitIgnore creates a new empty .gitignore file and returns a file
// descriptor. If 'force' is true, any previous .gitignore file will be
// overwritten.
func NewGitIgnore(force bool) (*os.File, error) {

	gitFileFlag := os.O_RDWR
	if force {
		gitFileFlag |= os.O_CREATE
	}

	return os.OpenFile(".gitignore", gitFileFlag, 0644)
}

func main() {

	// https://www.toptal.com/developers/gitignore/api/macos,linux,windows,ssh,vscode,go,zsh

	flagForced := flag.Bool("force", false, "force overwrite of previous .gitignore file (if it exists)")
	flagSkip := flag.Bool("skip", false, "skip download of .gitignore items from toptal")

	flag.Parse()

	// get file handle
	giFile, err := NewGitIgnore(*flagForced)
	if err != nil {
		log.Fatal(err)
	}

	// create .gitignore file contents

	// write .gitignore contents to file

}

// gi returns a string response from the www.gitignore.io API containing
// standard .gitignore items for the args given.
//
//      default: "macos linux windows ssh vscode go zsh node vue nuxt python django"
//
// using: https://www.toptal.com/developers/gitignore/api/macos,linux,windows,ssh,vscode,go,zsh,node,vue,nuxt,python,django
func gi(args string) string {

	if len(args) == 0 {
		args = "macos linux windows ssh vscode go zsh node vue nuxt python django"
	}

	command := `curl -fLw '\n' https://www.gitignore.io/api/\"${(j:,:)@}\" `
	command += args

	return gofile.Shell(command)
}

var (
	// defaultGitignoreItems is a list of personal prefernces to download from the www.gitignore.io API
	DefaultGitignoreItems string = "macos linux windows ssh vscode go zsh node vue nuxt python django"

	// personalPreferenceItems is a list of personal preferences in addition to the www.gitignore.io API
	PersonalPreferenceItems string = `# Personal Preference
ideas
notes.md
*[Pp]rivate*
*[Ss]ecret*

# used by go.test.sh
coverage.txt
profile.out

# generic items
**/*/node_modules/
*[Bb]ak
*temp
bak/
temp/
*ssh*
*history*
*hst*

`
)
