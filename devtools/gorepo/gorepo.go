package gorepo

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/skeptycal/util/devtools/gogit"
	"github.com/skeptycal/util/gofile"
	"github.com/skeptycal/zsh"
)

const (
	defaultGitignoreItems = "macos linux windows ssh vscode go zsh node vue nuxt python django"
)

type GitHubRepo struct {
	name string
	url  string
}

// gitIgnore writes a .gitignore file, including default items followed by the response from
// the www.gitignore.io API containing standard .gitignore items for the args given.
//
//      default: "macos linux windows ssh vscode go zsh node vue nuxt python django"
//
// using: https://www.toptal.com/developers/gitignore/api/macos,linux,windows,ssh,vscode,go,zsh,node,vue,nuxt,python,django
func (g *GitHubRepo) gitIgnore(args string) error {
	// notes - .gitignore header
	/*
	   # gorepo - .gitignore file

	   # generic secure items:
	   *private*
	   *secret*
	   *bak

	   # repo specific items
	   coverage.txt
	   profile.out
	*/

	if args == "" {
		args = defaultGitignoreItems
	}

	var sb strings.Builder
	defer sb.Reset()

	gifmt := fmt.Sprintf(giFormatString, g.name, gi(args))

	sb.WriteString(gi(args))

	return WriteFile(".gitignore", sb.String())
}

// Setup initializes the repo, creates files, prompts as needed, creates the
// github.com repository, and pushes the initial commit.
func Setup() error {
	err := gitRepoSetup()
	if err != nil {
		log.Printf("gitRepoSetup failed with %v", err)
		return err
	}
	err = createAutomatedFiles()
	if err != nil {
		log.Printf("createAutomatedFiles failed with %v", err)
		return err
	}
	return nil
}

// gitRepoSetup initializes the repo, prompts as needed, creates the
// github.com repository, and pushes the initial commit.
func gitRepoSetup() error {
	err := gitInit()
	if !gofile.Exists(".gitignore") {
		gitIgnore("", "")
	}
	if err != nil {
		return err
	}
	// todo - stuff
	return err
}

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

// CreateAutomatedFiles creates the automated files.
func createAutomatedFiles() error {
	zsh.Sh("touch main.go")
	return nil
}

// GoSum creates the go.sum file.
func GoSum() error {
	return nil
}

// GoTestSh creates the go.test.sh script.
func GoTestSh() error {
	return nil
}

// GoDoc creates the go.doc file.
func GoDoc() error {
	return nil
}

// BugReportMd creates the .github/ISSUE_TEMPLATE/bug_report.md file.
func BugReportMd() error {
	return nil
}

// FeatureRequestMd creates the .github/ISSUE_TEMPLATE/feature_request.md file.
func FeatureRequestMd() error {
	return nil
}

// GitWorkflows creates initial .github/workflows/... files:
// codeql-analysis.yml go.yml greetings.yml label.yml stale.yml
func GitWorkflows() error {
	return nil
}

// CodeCovYml creates the initial .codecov.yml file.
func CodeCovYml() error {
	return nil
}

// FundingYml creates the initial FUNDING.yml file.
func FundingYml() error {
	return nil
}

// PreCommitYaml creates the initial .pre-commit-config.yaml file.
func PreCommitYaml() error {
	return nil
}

// TravisYml creates the initial .travis.yml file.
func TravisYml() error {
	return nil
}

// DocGo creates the initial doc.go file.
func DocGo() error {
	return nil
}

// ReadMeMd creates the initial README.md file.
func ReadMeMd() error {
	return nil
}

// SecurityMd creates the initial SECURITY.md file.
func SecurityMd() error {
	return nil
}

// CodeOfConduct creates the initial CODE_OF_CONDUCT.md file.
func CodeOfConduct() error {
	return nil
}

// License creates the initial LICENSE file.
func License(license string) error {
	return nil
}

func NewGitHubRepo(name string, dir string) (*GitHubRepo, error) {
	wd := dir
	if wd == "" {
		wd, _ = os.Getwd()

	}
	wd, err := filepath.Abs(dir)
	if err == os.ErrNotExist {
		wd, _ = os.Getwd()

	}

	if err != nil {
		return nil, err
	}

	wd, _ = os.Getwd()
	pwd, _ := filepath.Abs(wd)
	default_path, default_name := path.Split(pwd)

	r := new(GitHubRepo)
	r.name = default_name
	r.url = default_path
	return r, nil
}

// GitHubRepoSetup initializes the repo, creates files, prompts as needed, creates the
// github.com repository, and pushes the initial commit.
func GitHubRepoSetup() (err error) {
	if GitRepoSetup() != nil {
		log.Info("GitRepoSetup failed with %v", err)
		return
	}
	if CreateAutomatedFiles() != nil {
		log.Info("CreateAutomatedFiles failed with %v", err)
		return
	}
	return nil
}

// GitRepoSetup initializes the repo, prompts as needed, creates the
// github.com repository, and pushes the initial commit.
func GitRepoSetup() error {
	return nil
}

// CreateAutomatedFiles creates the automated files.
func CreateAutomatedFiles() error {
	return nil
}

// gitIgnore writes a .gitignore file, including default items followed by the response from
// the www.gitignore.io API containing standard .gitignore items for the args given.
//
//      default: "macos linux windows ssh vscode go zsh node vue nuxt python django"
//
// using: https://www.toptal.com/developers/gitignore/api/macos,linux,windows,ssh,vscode,go,zsh,node,vue,nuxt,python,django
func MakeGitIgnore(args, repoName string) error {
	// notes - .gitignore header
	/*
	   # gorepo - .gitignore file

	   # generic secure items:
	   *private*
	   *secret*
	   *bak

	   # repo specific items
	   coverage.txt
	   profile.out
	*/

	if args == "" {
		args = defaultGitignoreItems
	}

	var sb strings.Builder
	defer sb.Reset()

	sb.WriteString(fmt.Sprintf("# %s - .gitignore file\n\n", repoName))

	sb.WriteString("# generic secure items:\n")
	sb.WriteString("*private*\n*secret*\n*bak\n\n")

	sb.WriteString("# repo specific items:\n")
	sb.WriteString("coverage.txt\nprofile.out\n\n")

	// add .gitignore contents from gitignore.io API
	sb.WriteString(gi(args))

	return WriteFile(".gitignore", sb.String())
}

// GoMod creates and initializes the repo go.mod file.
func GoMod() error {
	Shell("go mod init")
	Shell("go mod tidy")
	Shell("go mod download")
	gogit.GitCommit("go mod setup")
	return nil
}

const giFormatString = `# %s - .gitignore file

# --> This file is automatically generated <--

# personal preference items
%s/

# repo specific items:
coverage.txt
profile.out

# generic secure items
*token
*private*
*bak
*secret*

# .gitignore contents from gitignore.io API
%s

    `
