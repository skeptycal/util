package userinfo

import (
	"testing"
)

type fields = UserInfo

var testFields fields = *defaultDevUser
var sampleUserJSONFilename = "sample_user.json"

func TestUserInfo_Copyright(t *testing.T) {
	type fields struct {
		Name           string
		Email          string
		CopyrightStart float64
		License        string
		Github         string
		Website        string
		TwitterURL     string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
		{"year only: 2019", fields{CopyrightStart: 2019}, "copyright (c) 2019-2021 "},
		{"2019 and 'fakename'", fields{CopyrightStart: 2019, Name: "fakename"}, "copyright (c) 2019-2021 fakename"},
		{"use defaultDevUser", fields{CopyrightStart: defaultDevUser.CopyrightStart, Name: defaultDevUser.Name}, "copyright (c) 2019-2021 Michael Treanor"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserInfo{
				Name:           tt.fields.Name,
				Email:          tt.fields.Email,
				CopyrightStart: tt.fields.CopyrightStart,
				License:        tt.fields.License,
				Github:         tt.fields.Github,
				Website:        tt.fields.Website,
				TwitterURL:     tt.fields.TwitterURL,
			}
			if got := u.Copyright(); got != tt.want {
				t.Errorf("UserInfo.Copyright() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserInfo_SaveJSON(t *testing.T) {

	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{"json save to sample file", testFields, args{sampleUserJSONFilename}, false},
		{"empty json", UserInfo{}, args{"/dev"}, true},
		// invalid runes are (apparently) not tested by json.Marshaller
		// {"invalid utf8 rune", UserInfo{Name: string(utf8.RuneError)}, args{"/dev/null"}, true},
		// invalid strings are (apparently) not tested by json.Marshaller
		// {"invalid utf8 string", UserInfo{Name: "a\xc5z"}, args{"/dev/null"}, true},
		{"error", testFields, args{"/dev"}, true},
		{"write to /dev/null", testFields, args{"/dev/null"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserInfo{
				Name:           tt.fields.Name,
				Email:          tt.fields.Email,
				CopyrightStart: tt.fields.CopyrightStart,
				License:        tt.fields.License,
				Github:         tt.fields.Github,
				Website:        tt.fields.Website,
				TwitterURL:     tt.fields.TwitterURL,
			}
			if err := u.SaveJSON(tt.args.filename); (err != nil) != tt.wantErr {
				t.Errorf("UserInfo.SaveJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUserInfo_LoadJSON(t *testing.T) {

	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{"json load from sample file", testFields, args{sampleUserJSONFilename}, false},
		{"error = open : no such file or directory", UserInfo{}, args{"fakefile"}, true},
		{"error = unexpected end of JSON input", testFields, args{"/dev/null"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserInfo{
				Name:           tt.fields.Name,
				Email:          tt.fields.Email,
				CopyrightStart: tt.fields.CopyrightStart,
				License:        tt.fields.License,
				Github:         tt.fields.Github,
				Website:        tt.fields.Website,
				TwitterURL:     tt.fields.TwitterURL,
			}
			if err := u.LoadJSON(tt.args.filename); (err != nil) != tt.wantErr {
				t.Errorf("UserInfo.LoadJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
			if *u != testFields && !tt.wantErr {
				t.Errorf("UserInfo.LoadJSON() struct %v, want %v", *u, testFields)
			}
		})
	}
}
func Test_FixRune(t *testing.T) {
	type args struct {
		r rune
	}
	tests := []struct {
		name string
		args args
		want rune
	}{
		// TODO: Add test cases.
		{"rune('a')", args{'a'}, 'a'},
		{"int(0)", args{0}, 0},
		{"error (-1)", args{-1}, '\uFFFD'},
		{"rune('\uF000')", args{'\uF000'}, 61440},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FixRune(tt.args.r); got != tt.want {
				t.Errorf("FixRune() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFixStringUTF8Map(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"a\xc5z", args{"a\xc5z"}, "a�z"},
		{"hello", args{"hello"}, "hello"},
		{"rune(-1)", args{string(rune(-1))}, "\uFFFD"},
		{"posic�o", args{"posic�o"}, "posic�o"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FixStringUTF8Map(tt.args.s); got != tt.want {
				t.Errorf("FixStringUTF8Map() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFixStringUTF8(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"a\xc5z", args{"a\xc5z"}, "a�z"},
		{"hello", args{"hello"}, "hello"},
		{"rune(-1)", args{string(rune(-1))}, "�"},
		{"posic�o", args{"posic�o"}, "posic�o"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FixStringUTF8(tt.args.s); got != tt.want {
				t.Errorf("FixStringUTF8() = %v, want %v", got, tt.want)
			}
		})
	}
}
