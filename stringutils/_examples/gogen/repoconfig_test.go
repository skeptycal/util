package gogen

import (
	"fmt"
	"testing"

	"github.com/prometheus/common/log"
)

func TestRepoConfigName(t *testing.T) {
	t.Parallel()
}

func TestRepoMethods(t *testing.T) {
	t.Parallel()

	// func ErrTest(name string, test interface{}, wantErr bool, t *testing.T) {
	tests := []struct {
		name    string
		subname string
		f       func()
		wantErr bool
	}{}

	log.Info(tests)

	name := "RepoConfig_Methods"
	r, err := NewRepo("default", "", 0, DefaultUserConfig)
	if err != nil {
		t.Errorf("TestRepoMethods() error: %v", err)
	}

	GotWant(name, "Name()", r.Name(), "default", false, t)
	GotWant(name, "License()", r.License(), DefaultUserConfig.DefaultLicense, false, t)
	GotWant(name, "Year()", r.Year(), YearRange(DefaultUserConfig.DefaultCopyrightYear), false, t)
	GotWant(name, "URL()", r.URL(), fmt.Sprintf(fmtRepoURL, r.User.Username, r.name), false, t)
	GotWant(name, "DownloadURL()", r.DownloadURL(), fmt.Sprintf(fmtDownloadURL, r.User.Username, r.name), false, t)
	GotWant(name, "DocURL()", r.DocURL(), fmt.Sprintf(fmtDocURL, r.User.Username, r.name), false, t)
	GotWant(name, "String()", r.String(), fmt.Sprintf("%v", r), false, t)

	// ErrTest(name, func() error {}, true, t)

	// r, err = NewRepo("testRepoMethods", "", 0, DefaultUserConfig)
	// if err != nil {
	// 	t.Errorf("TestRepoMethods() error: %v", err)
	// }
	// if !utf8.ValidString(r.Name()) {
	// 	t.Errorf("invalid characters in name: %v", r.Name())
	// }
	// name := "newline in name\n"
	// r, err = NewRepo(name,"",0,DefaultUserConfig)
	// if err == nil {
	//     t.Errorf("should invalid repo name error (name: %v): %v",name, err)
	// }

	// || unicode.IsControl(r rune)
	// html.EscapeString(s string)
}

func TestNewRepo(t *testing.T) {
	// var wantRepo *RepoConfig = &RepoConfig{
	// 	DefaultUserConfig,
	// 	"default",
	// 	"MIT",
	// 	defaultYear,
	// 	"", "", "",
	// }
	defaultYear := YearRange(DefaultUserConfig.DefaultCopyrightYear)
	type args struct {
		reponame string
		license  string
		year     int
		user     *UserConfig
	}
	tests := []struct {
		name    string
		args    args
		wantR   *repoConfig
		wantErr bool
	}{
		// TODO: Add test cases.
		// Use default user for everything ... we are not testing the 'user' interface here
		// test default repo then use that for remaining tests
		{
			"default",
			args{"default", "MIT", 0, DefaultUserConfig},
			&repoConfig{DefaultUserConfig, "default", "MIT", defaultYear, "", "", ""},
			false,
		},
		{
			"no repo name",
			args{"", "MIT", 0, DefaultUserConfig},
			&repoConfig{DefaultUserConfig, "", "MIT", defaultYear, "", "", ""},
			true,
		},
		{
			"blank",
			args{"blank", "", 0, DefaultUserConfig},
			&repoConfig{DefaultUserConfig, "blank", "MIT", defaultYear, "", "", ""},
			false,
		},
		{
			"blankNoUser",
			args{"blankNoUser", "", 0, DefaultUserConfig},
			&repoConfig{DefaultUserConfig, "blankNoUser", "MIT", defaultYear, "", "", ""},
			false,
		},
		{
			"badYear", // year should be replaced with current year in struct
			args{"default", "MIT", 3030, DefaultUserConfig},
			&repoConfig{DefaultUserConfig, "default", "MIT", YearRange(3030), "", "", ""},
			false,
		},
		{
			"badUser",
			args{"badUser", "MIT", 0, nil},
			&repoConfig{DefaultUserConfig, "badUser", "MIT", defaultYear, "", "", ""},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotR, err := NewRepo(tt.args.reponame, tt.args.license, tt.args.year, tt.args.user)
            if err != nil != tt.wantErr {
                t.Errorf("not expected error: %v", err)
            }

			err = GotWant(tt.name, tt.args.reponame, gotR, tt.wantR, tt.wantErr, t)
            if err != nil {
                t.Error(err)
            }
		})
	}
}
