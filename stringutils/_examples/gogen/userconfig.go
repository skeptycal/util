package gogen

import (
	"fmt"
	"strings"
	"time"
)

// YearRange returns a string representation of the year range
// that is formatted as a typical copyright notice range.
//
// examples:
//  2015-2020   1970-2019   2021
//
// It will contain the given year, a dash, and the current year.
// If the y parameter is greater than the current year or more
// than 50 years ago, it is assumed to be an error and the current
// year is used. An info level log message is generated if the
// logger is set to show INFO level, but no error is returned.
//
func YearRange(y int) string {
    now := time.Now().Year()
    // log.Infof("y input: %d", y) // todo remove dev output
    // log.Infof("now: %d", now)   // todo remove dev output
	if y > now || y < (now-50) {
        // log.Infof("year out of range: %v (older than 50 years or newer than now) ... default current year used", y)
        y = now
    }

	if y == now{
		return fmt.Sprintf("%d", time.Now().Year())
	}

	return fmt.Sprintf("%d-%d", y, time.Now().Year())
}

func NewUser(name, email, username, defaultLicense string, defaultCopyrightYear int) (*UserConfig, error) {

	if name == "" {
        if DefaultUserConfig.Name == "" {
		    return nil, fmt.Errorf("name is invalid: %v", name)
        }
		return DefaultUserConfig, nil
	}
	if email == "" {
		email = DefaultUserConfig.Email
	}
	if username == "" {
		username = DefaultUserConfig.Username
	}
	if defaultLicense == "" {
		defaultLicense = DefaultUserConfig.DefaultLicense
	}
	if defaultCopyrightYear == 0 {
		defaultCopyrightYear = DefaultUserConfig.DefaultCopyrightYear
	}

	return &UserConfig{
		Name:                 name,
		Email:                email,
		Username:             username,
		DefaultLicense:       defaultLicense,
		DefaultCopyrightYear: defaultCopyrightYear,
	}, nil
}

type UserConfig struct {
	Name                 string `json:"name"`
	Email                string `json:"email,omitempty"`
	Username             string `json:"username,omitempty"`
	DefaultLicense       string `json:"default_license,omitempty"`
	DefaultCopyrightYear int    `json:"default_copyright_year,omitempty"`
}

func (c *UserConfig) String() string {
	sb := strings.Builder{}
	defer sb.Reset()
	sb.WriteString("User Config:\n")
	sb.WriteString("  Name: " + c.Name + "\n")
	sb.WriteString("  Email: " + c.Email + "\n")
	sb.WriteString("  Username: " + c.Username + "\n")
	sb.WriteString("  DefaultLicense: " + c.DefaultLicense + "\n")
	sb.WriteString("  DefaultCopyrightYear: " + fmt.Sprintf("%d", c.DefaultCopyrightYear) + "\n")
	return sb.String()
}
