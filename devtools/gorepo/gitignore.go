package gorepo

import (
	"fmt"
	"os"
	"reflect"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/skeptycal/util/gofile"
	"github.com/skeptycal/util/webtools/http"
)

func ReturnCheck(f func()) ([]interface{}, error) {
	foo := reflect.TypeOf(f)

	fmt.Printf("f is type: %v", foo)
	return nil, nil
}

func CreateGitIgnore(force bool) error {

	if gofile.Exists(".gitignore") {
		if !force {
			return fmt.Errorf(".gitignore already exists; use force option to overwrite")
		}
	}

	f, err := os.Create(".gitignore")
	if err != nil {
		return err
	}

	log.Infof("Creating gitignore file %v", f.Name())

	return nil
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
