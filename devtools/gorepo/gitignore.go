package gorepo

import (
	"fmt"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
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

// gitignoreAPIList returns a list of available gitignore file languages parameters
func gitIgnoreAPIList() []string {
	// https://www.toptal.com/developers/gitignore/api/list

	body, err := http.GetPage(urlGitIgnoreAPIList)
	if err != nil {
		log.Error(err)
		return make([]string, 0)
	}

	list := strings.Split(string(body), " ")
	return list
}

// gi returns a string response from the www.gitignore.io API containing
// standard .gitignore items for the args given.
//
//      default: "macos linux windows ssh vscode go zsh node vue nuxt python django"
//
// using: https://www.toptal.com/developers/gitignore/api/macos,linux,windows,ssh,vscode,go,zsh,node,vue,nuxt,python,django
func gitIgnoreAPI(args string) string {

	if len(args) == 0 {
		args = defaultGitIgnoreItemsComma
	}

	url := urlGitIgnoreAPIPrefix + args

	body, err := http.GetPage(url)
	if err != nil {
		log.Error(err)
		return ""
	}
	return body
}

// gi returns a .gitignore file from the www.gitignore.io API containing
// standard .gitignore items for the space delimited args given.
//
//      default: "macos linux windows ssh vscode go zsh node vue nuxt python django"
//
// using: https://www.toptal.com/developers/gitignore/api/
func GetGitIgnore(args string) (string, error) {

	if len(args) == 0 {
		args = defaultGitIgnoreItems
	}

	url := "https://www.gitignore.io/api/\"${(j:,:)@}\" " + args
	// command := "curl -fLw '\n' https://www.gitignore.io/api/\"${(j:,:)@}\" " + args

	buf, err := GetPageBody(url)
	if err != nil {
		return "", err
	}
	defer buf.Reset()
	return buf.String(), nil
}

// GitIgnore writes a .gitignore file, including default items followed by the response from
// the www.gitignore.io API containing standard .gitignore items for the args given.
//
//      default: "macos linux windows ssh vscode go zsh node vue nuxt python django"
//
// using: https://www.toptal.com/developers/gitignore/api/macos,linux,windows,ssh,vscode,go,zsh,node,vue,nuxt,python,django
func GitIgnore(reponame, args string, personal string) error {
	if args == "" {
		args = defaultGitIgnoreItems
	}

	if personal == "" {
		personal = defaultGitIgnorePersonalItems
	}

	var sb strings.Builder
	defer sb.Reset()

	gifmt := fmt.Sprintf(giFormatString, reponame, gi(args))

	sb.WriteString(gi(args))

	return gofile.WriteFile(".gitignore", sb.String())
}
