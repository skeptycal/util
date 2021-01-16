// Package userinfo implements the User interface to manage
// and store user demographic information.
package userinfo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
	"time"
	"unicode/utf8"
)

type User interface {
	Copyright() string
}

type UserInfo struct {
	Name           string  `json:"Name,omitempty"`
	Email          string  `json:"Email,omitempty"`
	CopyrightStart float64 `json:"CopyrightStart,omitempty"`
	License        string  `json:"License,omitempty"`
	Github         string  `json:"Github,omitempty"`
	Website        string  `json:"Website,omitempty"`
	TwitterURL     string  `json:"TwitterURL,omitempty"`
}

func (u *UserInfo) Copyright() string {
	year := time.Now().Year()
	return fmt.Sprintf("copyright (c) %.0f-%d %s", u.CopyrightStart, year, u.Name)
}

func (u *UserInfo) Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, u)
}

func (u *UserInfo) LoadJSON(filename string) (err error) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}
	return u.Unmarshal(b, u)
}

// Marshal implements the json.Marshaler interface. v is ignored
// and should be set to nil to be clear.
func (u *UserInfo) Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(u)
}

func (u *UserInfo) SaveJSON(filename string) error {
	b, err := u.Marshal(nil)
	if err != nil {
		return err
	}

	// permissions set to owner only by default
	return ioutil.WriteFile(filename, b, 0600)
}

var defaultDevUser *UserInfo = &UserInfo{
	Name:           "Michael Treanor",
	Email:          "skeptycal@gmail.com",
	CopyrightStart: 2019,
	License:        "MIT",
	Github:         "https://github.com/skeptycal",
	Website:        "https://www.michaeltreanor.com",
	TwitterURL:     "https://twitter.com/skeptycal",
}

func FixStringUTF8(s string) string {
	sb := strings.Builder{}
	for _, r := range s {
		sb.WriteRune(FixRune(r))
		// if utf8.ValidRune(r) {
		// 	sb.WriteRune(r)
		// } else {
		// 	sb.WriteRune(utf8.RuneError)
		// }
	}
	return sb.String()
}

func FixStringUTF8Map(s string) string {
	return strings.Map(FixRune, s)
}

func FixRune(r rune) rune {
	if utf8.ValidRune(r) {
		return r
	}
	return '\uFFFD'
}
