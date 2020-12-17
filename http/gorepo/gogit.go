package gorepo

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

// gi returns a .gitignore file from the www.gitignore.io API containing
// standard .gitignore items for the space delimited args given.
//
//      default: "macos linux windows ssh vscode go zsh node vue nuxt python django"
//
// using: https://www.toptal.com/developers/gitignore/api/
func GetGitIgnore(args string) (string, error) {

	if len(args) == 0 {
		args = defaultGitignoreItems
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

// GetPageBody - returns the body of the page at url.
func GetPageBody(url string) (*bytes.Buffer, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return bytes.NewBuffer(body), nil
}
