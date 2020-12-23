package main

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/skeptycal/util/gofile"
)

// New repo with all default values for testing purposes.
var r = NewRepo("", "", "", "", "")

func TestNewRepo(t *testing.T) {
	type args struct {
		name     string
		url      string
		auth     string
		username string
		license  string
	}
	tests := []struct {
		name string
		args args
		want *repo
	}{
		// all fake fields
		{"test all fields", args{"test", "localhost", "me", "the_user", "PSP"}, &repo{"test", "localhost", "me", "the_user", "PSP"}},
		// defaults tested against literals
		{"test  default literals", args{"", "", "", "", ""}, &repo{"gobot", "https://github.com/skeptycal/gobot", "Michael Treanor", "skeptycal", "MIT"}},
		// defaults tested against default values
		{"test  default constants", args{"", "", "", "", ""}, &repo{gofile.Base(gofile.PWD()), defaultRepoPrefix + "/" + defaultUserName + "/" + gofile.Base(gofile.PWD()), defaultAuthName, defaultUserName, defaultLicense}},
		// new struct tested against literals
		{"test  new struct", args{r.name, r.url, r.author, r.username, r.license}, &repo{"gobot", "https://github.com/skeptycal/gobot", "Michael Treanor", "skeptycal", "MIT"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRepo(tt.args.name, tt.args.url, tt.args.auth, tt.args.username, tt.args.license); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRepo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_doOrDie(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"nil error", args{nil}, false},
		{"generic error", args{fmt.Errorf("fake error")}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := doOrDie(tt.args.err); (err != nil) != tt.wantErr {
				t.Errorf("doOrDie() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_repo_CheckURL(t *testing.T) {
	type fields struct {
		name     string
		url      string
		author   string
		username string
		license  string
	}
	type args struct {
		url string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{"google", fields{"", "", "", "", ""}, args{"https://www.google.com"}, false},
		{"fake", fields{"", "", "", "", ""}, args{"https://fake.none"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &repo{
				name:     tt.fields.name,
				url:      tt.fields.url,
				author:   tt.fields.author,
				username: tt.fields.username,
				license:  tt.fields.license,
			}
			if err := r.CheckURL(tt.args.url); (err != nil) != tt.wantErr {
				t.Errorf("repo.CheckURL() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_repo_DefaultURL(t *testing.T) {
	type fields struct {
		name     string
		url      string
		author   string
		username string
		license  string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
		{"google", fields{"", "", "", "", ""}, "https://github.com/skeptycal/gobot"},
		{"google", fields{"FakeName", "https://www.fake.com", "FakeAuthor", "FakeUser", "FakeLicense"}, "https://github.com/FakeUser/FakeName"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &repo{
				name:     tt.fields.name,
				url:      tt.fields.url,
				author:   tt.fields.author,
				username: tt.fields.username,
				license:  tt.fields.license,
			}
			if got := r.DefaultURL(); got != tt.want {
				t.Errorf("repo.DefaultURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
