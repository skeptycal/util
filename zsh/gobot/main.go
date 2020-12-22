package main

import (
	"fmt"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"
)

type errorHandling int

const (
	continueOnError errorHandling = iota
	exitOnError
)

const ErrHandling errorHandling = continueOnError

type repo struct {
	name    string
	url     string
	author  string
	license string
}

func (r *repo) Name() string {
	if r.name == "" {
		r.name = ParentDir()
	}
	return r.name
}

func (r *repo) SetURL(url string) error {
	resp, err := http.Get(url)
	doOrDie(err)

	if resp.StatusCode != http.StatusOK {
		doOrDie(fmt.Errorf("error: server response %s", resp.Status))
	}

	r.url = url
	return nil
}

func doOrDie(err error) error {
	if err != nil {
		log.Error(err)
		if ErrHandling == exitOnError {
			os.Exit(1)
		}
	}
	return err
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
