package gogen

import (
	"fmt"
	"strings"

	"github.com/skeptycal/util/stringutils"
)

const (
	fmtRepoURL     string = "https://%s.github.io/%s"
	fmtDocURL      string = "https://github.com/%s/%s/docs"
	fmtDownloadURL string = "https://github.com/%s/%s"
)

func NewRepo(reponame, license string, year int, user *UserConfig) (r *repoConfig, err error) {
	if user == nil || user.Name == "" || !stringutils.IsASCIIPrintable(user.Name){
		user = DefaultUserConfig
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

	r = &repoConfig{
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

type repoConfig struct {
	User        *UserConfig
	name        string
	license     string `default:""`
	year        string `default:""`
	url         string `default:""`
	downloadURL string `default:""`
	docURL      string `default:""`
}

func (r *repoConfig) Name() string {
	return r.name
}

func (r *repoConfig) License() string {
	return r.license
}

func (r *repoConfig) Year() string {
	return r.year
}

func (r *repoConfig) URL() string {
	if r.url == "" {
		r.url = fmt.Sprintf(fmtRepoURL, r.User.Username , r.name)
	}
	return r.url
}

func (r *repoConfig) DownloadURL() string {
	if r.downloadURL == "" {
		r.downloadURL = fmt.Sprintf(fmtDownloadURL, r.User.Username, r.name)
	}
	return r.downloadURL
}

func (r *repoConfig) DocURL() string {
	if r.docURL == "" {
		r.docURL = fmt.Sprintf(fmtDocURL, r.User.Username, r.name)
	}
	return r.docURL
}

func (r *repoConfig) String() string {
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
