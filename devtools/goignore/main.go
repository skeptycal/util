package main

import (
	"bufio"
	"flag"
	"log"
	"os"
	"strings"
)

type giFile struct {
	bufio.ReadWriter
	filename string
	flag     int
	file     *os.File
	force    bool
	skip     bool
}

func NewGitIgnore(force bool, skip bool) (*os.File, error) {

	gitFileFlag := os.O_RDWR
	if force {
		gitFileFlag |= os.O_CREATE
	}

	file, err := os.OpenFile(".gitignore", gitFileFlag, 0644)
	if err != nil {
		return nil, err
	}

	return &giFile{
		filename: ".gitignore",
		force:    force,
		skip:     skip,
	}, nil
}

func main() {

	// https://www.toptal.com/developers/gitignore/api/macos,linux,windows,ssh,vscode,go,zsh

	flagForced := flag.Bool("force", false, "force overwrite of previous .gitignore file (if it exists)")
	flagSkip := flag.Bool("skip", false, "skip download of .gitignore items from toptal")

	flag.Parse()

	giFile, err := NewGitIgnore(*flagForced, *flagSkip)
	if err != nil {
		log.Fatal(err)
	}
	rw := bufio.NewReadWriter(giFile)

}

// gi returns a string response from the www.gitignore.io API containing
// standard .gitignore items for the args given.
//
//      default: "macos linux windows ssh vscode go zsh node vue nuxt python django"
//
// using: https://www.toptal.com/developers/gitignore/api/macos,linux,windows,ssh,vscode,go,zsh,node,vue,nuxt,python,django
func gi(args string) string {

	if len(args) == 0 {
		args = []string{"macos linux windows ssh vscode go zsh node vue nuxt python django"}
	}

	command := "curl -fLw '\n' https://www.gitignore.io/api/\"${(j:,:)@}\" "
	command += strings.Join(args, " ")

	return Shell(command)
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
