package gorepo

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"

	log "github.com/sirupsen/logrus"
)

const (
	defaultGitignoreItems = "macos linux windows ssh vscode go zsh node vue nuxt python django"
	sep                   = string(os.PathSeparator)
)

type GitHubRepo struct {
	name string
	url  string
}

// PWD returns the current working directory. It does not return any error, but instead
// logs the error and returns the system default glob pattern for current working directory
func PWD() string {
	wd, err := os.Getwd()
	if err != nil {
		log.Println(err)
		// this is a crutch for the extremely rare case where Getwd fails
		if runtime.GOOS == "windows" || runtime.GOOS == "plan9" {
			return ".\\"
		}
		return "."
	}
	return wd
}

// DirName returns the last element of a path.
// similar to filepath.Base()
//
// todo  - btw filepath.Base() has an error (rather redundant check)
//
// (see below) if path is "" ... "." is returned ... so ...
// the next line `len(path) > 0` is redundant ... it must be > 0
// or it would have been returned in the first check ...
// however, if one slash is removed ... and it is the only one ...
// then this check would be valid ... but ... if the only thing in the path is
// one slash ... it should be returned directly
//	    if path == "" {
//      	return "."
//      }
//      // Strip trailing slashes.
//      for len(path) > 0 && os.IsPathSeparator(path[len(path)-1]) {
//      	path = path[0 : len(path)-1]
//      }
func DirName(path string) string {
	i := strings.LastIndex(path, sep)
	return path[i+1:]
}

// Base returns the last element of path.
// This is a convenience version modified from Go 1.15.6
// (located at /src/path/filepath/path.go)
//
// Trailing path separators are removed before extracting the last element.
// If the path is empty, Base returns ".".
// If the path consists entirely of separators, Base returns a single separator.
func Base(path string) string {
	if path == "" {
		return "."
	}
	// Strip trailing slashes.
	path = strings.TrimRight(path, string(os.PathSeparator))

	// Throw away volume name
	path = path[len(filepath.VolumeName(path)):]
	// Find the last element
	i := len(path) - 1
	for i >= 0 && !os.IsPathSeparator(path[i]) {
		i--
	}
	if i >= 0 {
		path = path[i+1:]
	} else {
		// If empty now, it had only slashes.
		return string(os.PathSeparator)
	}
	return path
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
func MakeGitIgnore(args string) error {
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

// GitInit initializes the Git environment
func gitInit() error {
	if !fileExists(".gitignore") {
		gitIgnore("")
	}

	Shell("git init")
	GitCommit("initial commit")
	return nil
}

// GoMod creates and initializes the repo go.mod file.
func GoMod() error {
	Shell("go mod init")
	Shell("go mod tidy")
	Shell("go mod download")
	GitCommit("go mod setup")
	return nil
}

// GitCommit creates a commit with message
func GitCommit(message string) error {
	Shell("git add --all")
	Shell("git commit -m '" + message + "'")
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
