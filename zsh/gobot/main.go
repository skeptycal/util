package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/skeptycal/util/gofile"
)

const (
	defaultAuthName   = "Michael Treanor"
	defaultLicense    = "MIT"
	defaultUserName   = "skeptycal"
	defaultRepoPrefix = "https://github.com"
)

// NewRepo returns a new repo instance with default values initialized.
func NewRepo(name, url, auth, username, license string) *repo {
	if auth == "" {
		auth = defaultAuthName
	}
	if username == "" {
		username = defaultUserName
	}
	if license == "" {
		license = defaultLicense
	}
	if name == "" {
		name = gofile.Base(gofile.PWD())
	}
	if url == "" {
		url = fmt.Sprintf("%s/%s/%s", defaultRepoPrefix, username, name)
	}
	return &repo{
		name:     name,
		url:      url,
		author:   auth,
		username: username,
		license:  license,
	}
}

type repo struct {
	name     string
	url      string
	author   string
	username string
	license  string
}

func (r *repo) Name() string {
	if r.name == "" {
		r.name = gofile.Base(gofile.PWD())
	}
	return r.name
}
func (r *repo) URL() string {
	if r.url == "" {
		url := r.DefaultURL()
		err := r.CheckURL(r.DefaultURL())
		if err != nil {
			return ""
		}
		r.url = url
	}
	return r.url
}
func (r *repo) Author() string {
	if r.author == "" {
		r.author = defaultAuthName
	}
	return r.author
}
func (r *repo) User() string {
	if r.username == "" {
		r.username = defaultUserName
	}
	return r.username
}
func (r *repo) License() string {
	if r.license == "" {
		r.license = defaultLicense
	}
	return r.license
}

func (r *repo) CheckURL(url string) error {
	resp, err := http.Get(url)
	if doOrDie(err) != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		_ = doOrDie(fmt.Errorf("error: server response %s", resp.Status))
	}

	r.url = url
	return nil
}
func (r *repo) DefaultURL() string {
	return fmt.Sprintf("%s/%s/%s", defaultRepoPrefix, r.User(), r.Name())
}

func (r *repo) ReadFile(path string) ([]byte, error) {
	if gofile.Exists(path) {
		return ioutil.ReadFile(path)
	}
	return nil, os.ErrNotExist
}

// WriteFile writes data to a file named by filename. If the file does not exist, WriteFile creates it; otherwise WriteFile truncates it before writing, without changing permissions.
func (r *repo) WriteFile(path string, data []byte) error {
	return ioutil.WriteFile(path, data, 0644)
}

func doOrDie(err error) error {
	return gofile.DoOrDie(err)
}

const (
	zsh_shebang = "#!/usr/bin/env zsh\n"
	zsh_section = `#? -----------------------------> `
)
const _ = `#!/usr/bin/env zsh

#? -----------------------------> parameter expansion tips
 #? ${PATH//:/\\n}    - replace all colons with newlines
 #? ${PATH// /}       	- strip all spaces
 #? ${VAR##*/}        - return only final element in path (program name)
 #? ${VAR%/*}         - return only path (without program name)

. $(which ansi_colors)

REPO_NAME=${PWD##*/}
REPO_URL=$(git remote get-url origin)
BORDER_CHAR='='
BORDER_COLOR=$LIME
SIDE_INDENT=">>----------->    "

function hr() {
	printf -v BORDER_TEMPLATE '%*s' $COLUMNS '';
	printf '%b%s%b\n' $BORDER_COLOR ${BORDER_TEMPLATE// /${1:-$BORDER_CHAR}} $RESET
}

function side() {
	printf '%b%s%s%b\n'  $BORDER_COLOR $SIDE_INDENT ${@:-} $RESET
}

function basic_readme() {
	if ! [ -f README.md ]; then
	(
		echo "Repo: ${REPO_NAME}"
		echo ""
		echo "go version: $(go version)"
		echo ""

	) >> README.md
	fi
}

function refresh() {
	hr
	side "REF --- REPOSITORY REFRESH"
	br

	side "Repo: $REPO_NAME"
	side "URL: $REPO_URL"
	hr
	br

	side  "go build and mod tidy"
	go mod tidy && go mod verify
	br

	side "go doc update"
	go doc | tail -n 5
	go doc >| go.doc
	br

	side "git add all"
	git add --all
	git status | tail -n 5
	br

	side "git commit -m 'GoBot: dev (pre v1.0) progress and formatting'"
	git commit -m "tidy and formatting documentation"
	br

	side "git push origin $BRANCH"
	git push origin $BRANCH
	br
	hr
}

refresh
`
