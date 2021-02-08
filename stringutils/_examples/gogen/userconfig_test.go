package gogen

import (
	"reflect"
	"testing"
)

func TestNewUser(t *testing.T) {
	type args struct {
		name                 string
		email                string
		username             string
		defaultLicense       string
		defaultCopyrightYear int
	}
	tests := []struct {
		name    string
		args    args
		want    *UserConfig
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "values to match DefaultUserConfig",
			args: args{
				name:                 "Michael Treanor",
				email:                "skeptycal@gmail.com",
				username:             "skeptycal",
				defaultLicense:       "MIT",
				defaultCopyrightYear: 1975,
			},
			want:    DefaultUserConfig,
			wantErr: false,
		},
		{
			name: "blank(use DefaultUserConfig)",
			args: args{
				name:                 "",
				email:                "",
				username:             "",
				defaultLicense:       "",
				defaultCopyrightYear: 0,
			},
			want:    DefaultUserConfig,
			wantErr: false,
		},
		{
			name: "fake",
			args: args{
				name:                 "Fake Name",
				email:                "fake@home.com",
				username:             "fakely",
				defaultLicense:       "FAK",
				defaultCopyrightYear: 2001,
			},
			want:    &UserConfig{"Fake Name", "fake@home.com", "fakely", "FAK", 2001},
			wantErr: false,
		},
		{
			name: "name only (individual defaults)",
			args: args{
				name:                 "Name Only",
				email:                "",
				username:             "",
				defaultLicense:       "",
				defaultCopyrightYear: 0,
			},
			want:    &UserConfig{"Name Only", "skeptycal@gmail.com", "skeptycal", "MIT", 1975},
			wantErr: false,
		},
		{
			name: "blank(use DefaultUserConfig)",
			args: args{
				name:                 "",
				email:                "",
				username:             "",
				defaultLicense:       "",
				defaultCopyrightYear: 0,
			},
			want:    DefaultUserConfig,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewUser(tt.args.name, tt.args.email, tt.args.username, tt.args.defaultLicense, tt.args.defaultCopyrightYear)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUser() = %v, want %v", got, tt.want)
			}
		})
	}

    // save and temporarily clear DefaultUserConfig
    tmp := DefaultUserConfig
	DefaultUserConfig = &UserConfig{}
	t.Run("No DefaultUserConfig Error", func(t *testing.T) {
		_, err := NewUser("", "", "", "", 0)
		// want := DefaultUserConfig

			if err == nil {
				t.Errorf("expected NewUser() error = %v, wantErr %v", err, "invalid name")
			}
	})
    // restore DefaultUserConfig
    DefaultUserConfig = tmp
}

func TestYearRange(t *testing.T) {
	type args struct {
		y int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"2015", args{2015}, "2015-2021"},
		{"1975", args{1975}, "1975-2021"},
		{"1960", args{1960}, "2021"},
		{"3030", args{3030}, "2021"},
		{"0", args{0}, "2021"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := YearRange(tt.args.y); got[:4] != tt.want[:4] {
				t.Errorf("YearRange() = %v, want %v", got, tt.want)
			}
		})
	}
}
