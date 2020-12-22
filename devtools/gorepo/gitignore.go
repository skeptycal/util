package gorepo

import (
	"fmt"
	"os"

	"github.com/prometheus/common/log"
	"github.com/skeptycal/util/gofile"
	"github.com/skeptycal/util/http/http"
)

func NewGitIgnore(force bool, skip bool) (*os.File, error) {

	ok := gofile.Exists(".gitignore")

	if ok {
		if !force {
			return nil, fmt.Errorf(".gitignore already exists; use force option to overwrite")
		}
	}

	gitFileFlag := os.O_RDWR
	if force {
		gitFileFlag |= os.O_CREATE
	}

	file, err := os.OpenFile(".gitignore", gitFileFlag, 0644)
	if err != nil {
		return nil, err
	}

	return file, nil
}

// gi returns a string response from the www.gitignore.io API containing
// standard .gitignore items for the args given.
//
//      default: "macos linux windows ssh vscode go zsh node vue nuxt python django"
//
// using: https://www.toptal.com/developers/gitignore/api/macos,linux,windows,ssh,vscode,go,zsh,node,vue,nuxt,python,django
func gi(args string) string {

	if len(args) == 0 {
		args = defaultGitIgnoreItems
	}

	url := `https://www.toptal.com/developers/gitignore/api/` + args

	body, err := http.GetPage(url)
	if err != nil {
		log.Error(err)
	}
	body = PersonalPreferenceItems + body
	return body
}

const (
	// defaultGitignoreItems is a list of personal prefernces to download from the www.gitignore.io API
	defaultGitIgnoreItems = `macos,linux,windows,ssh,vscode,go,zsh,node,vue,nuxt,python,django`

	// personalPreferenceItems is a list of personal preferences in addition to the www.gitignore.io API
	PersonalPreferenceItems = `# Personal Preference
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
