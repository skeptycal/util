package gogen

import (
	"reflect"
	"testing"
	"unicode/utf8"
)

func TestRepoMethods(t *testing.T) {
    r, err := NewRepo("testRepoMethods","",0,DefaultUserConfig)
    if err != nil {
        t.Errorf("TestRepoMethods() error: %v", err)
    }
    if !utf8.ValidString(r.name) {
        t.Errorf("invalid characters in name: %v", r.name)
    }
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
		wantR   *RepoConfig
		wantErr bool
	}{
        // TODO: Add test cases.
        // Use default user for everything ... we are not testing the user interface here
		// test default repo then use that for remaining tests
		{
			"default",
			args{"default", "MIT", 0, DefaultUserConfig},
			&RepoConfig{DefaultUserConfig, "default", "MIT", defaultYear, "", "", ""},
			false,
		},
		{
			"no repo name",
			args{"", "MIT", 0, DefaultUserConfig},
			&RepoConfig{DefaultUserConfig, "", "MIT", defaultYear, "", "", ""},
			true,
		},
		{
			"blank",
			args{"blank", "", 0, DefaultUserConfig},
			&RepoConfig{DefaultUserConfig, "blank", "MIT", defaultYear, "", "", ""},
			false,
		},
		{
			"blankNoUser",
			args{"blankNoUser", "", 0, DefaultUserConfig},
			&RepoConfig{DefaultUserConfig, "blankNoUser", "MIT", defaultYear, "", "", ""},
			false,
		},
		{
			"badYear", // year should be replaced with current year in struct
			args{"default", "MIT", 3030, DefaultUserConfig},
			&RepoConfig{DefaultUserConfig, "default", "MIT", YearRange(3030), "", "", ""},
			false,
		},
		{
			"badUser",
			args{"default", "MIT", 0, nil},
			&RepoConfig{&UserConfig{}, "default", "MIT", defaultYear, "", "", ""},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotR, err := NewRepo(tt.args.reponame, tt.args.license, tt.args.year, tt.args.user)
			if tt.wantErr { // if wantErr skip other tests ... not only if err != nil
				if err == nil {  // if no error, but wanted one ...
					t.Errorf("NewRepo() error = %v, wantErr %v", err, tt.wantErr)
                }
                // if wanted error and got one ... no output ... don't care what it was; we got it
				return
			}
			if !reflect.DeepEqual(gotR, tt.wantR) {
				t.Errorf("NewRepo() = %v, want %v", gotR, tt.wantR)
			}
		})
	}
}
