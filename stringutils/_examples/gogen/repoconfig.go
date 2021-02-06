package gogen

import (
	"fmt"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/skeptycal/util/stringutils"
)

const (
	fmtRepoURL     string = "https://%s.github.io/%s"
	fmtDocURL      string = "https://github.com/%s/%s/docs"
	fmtDownloadURL string = "https://github.com/%s/%s"
)

func NewRepo(reponame, license string, year int, user *UserConfig) (r *RepoConfig, err error) {
	if user == nil {
		user = DefaultUserConfig
	}
	if user.Name == "" || !stringutils.IsASCIIPrintable(user.Name) {
		return nil, fmt.Errorf("user is invalid: %v", user)
	}
	if reponame == "" {
		return nil, fmt.Errorf("repo name is invalid: %v", reponame)
	}
	if license == "" {
		license = user.DefaultLicense
    }
    if year == 0 {
        year = user.DefaultCopyrightYear
    }

	repoyear := YearRange(year)

	r = &RepoConfig{
		User:    user,
		name:    reponame,
		license: license,
		year:    repoyear,
	}

	// these are lazily initialized, could be done here specifically
	// _ = r.License()
	// _ = r.Year()
	// _ = r.URL()
	// _ = r.DocURL()
	// _ = r.DownloadURL()

	return r, nil
}

// type Repo interface {
// 	String() string
// 	URL() string
// 	DownloadURL() string
// 	DocURL() string
// }

type RepoConfig struct {
	User        *UserConfig
	name        string
	license     string `default:""`
	year        string `default:""`
	url         string `default:""`
	downloadURL string `default:""`
	docURL      string `default:""`
}

func (r *RepoConfig) Name() string {
	if r.name == "" {
        log.Errorf("repo name is invalid: %v", r.name)
        return ""
	}
	return r.name
}

func (r *RepoConfig) License() string {
	if r.license == "" {
		r.license = r.User.DefaultLicense
	}
	return r.license
}

func (r *RepoConfig) Year() string {
	if r.year == "" {
		r.year = YearRange(r.User.DefaultCopyrightYear)
	}
	return r.year
}

func (r *RepoConfig) URL() string {
	if r.User.Username == "" {
		log.Fatalf("a valid username is required: %v", r.User.Username)
	}
	if r.url == "" {
		r.url = fmt.Sprintf(fmtRepoURL, r.User.Username , r.name)
	}
	return r.url
}

func (r *RepoConfig) DownloadURL() string {
	if r.name == "" {
		log.Fatalf("a valid repo name is required")
	}
	if r.downloadURL == "" {
		r.downloadURL = fmt.Sprintf(fmtDownloadURL, r.User.Username, r.name)
	}
	return r.downloadURL
}

func (r *RepoConfig) DocURL() string {
	if r.name == "" {
		log.Fatalf("a valid repo name is required")
	}
	if r.docURL == "" {
		r.docURL = fmt.Sprintf(fmtDocURL, r.User.Username, r.name)
	}
	return r.docURL
}

func (r *RepoConfig) String() string {
	sb := strings.Builder{}
	defer sb.Reset()
	sb.WriteString("Repo Config:\n")
	sb.WriteString("  name: " + r.name + "\n")
	sb.WriteString("  license: " + r.license + "\n")
	sb.WriteString("  year: " + r.year + "\n")
	sb.WriteString("  url: " + r.url + "\n")
	sb.WriteString("  downloadURL: " + r.downloadURL + "\n")
	sb.WriteString("  docURL: " + r.docURL + "\n\n")
	sb.WriteString(r.User.String())
	return sb.String()
}
