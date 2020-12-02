package main

import (
	"bufio"
	"context"
	"fmt"
	"net/http"
	"os"
	"syscall"

	"github.com/google/go-github/github"
	"golang.org/x/crypto/ssh/terminal"
	"golang.org/x/oauth2"
)

// todo - rename the default environment variable keys appropriately

const (
	envGitUsernameKey = "GIT_AUTHOR_NAME"  // environment variable for GitHub username
	envGitTokenKey    = "GITHUB_API_TOKEN" // environment variable for GitHub PAT

	// MinRepositoryCountGoal - number of repositories required to receive the maximum rate limit.
	// "Installations that have more than 20 repositories receive another 50 requests per hour for each repository. The maximum rate limit for an installation is 12,500 requests per hour."
	//
	//      (20 + 7500 // 50)
	//
	// Ref: https://docs.github.com/en/free-pro-team@latest/developers/apps/differences-between-github-apps-and-oauth-apps
	MinRepositoryCountGoal = 170

	// MinGitHubRateLimit - "GitHub Apps making server-to-server requests use the installation's
	// minimum rate limit of 5,000 requests per hour."
	//
	// Ref: https://docs.github.com/en/free-pro-team@latest/developers/apps/differences-between-github-apps-and-oauth-apps
	MinGitHubRateLimit = 5000
	maxGitHubRateLimit = 12500
)

var (

	// tokenSource - used to cache the environment variable GITHUB_API_TOKEN
	//
	// GitHub Apps ask for repository contents permission and use your installation token to
	// authenticate via HTTP-based Git. The token is used as the HTTP password.
	tokenSource oauth2.TokenSource

	// DefaultContext represents context.Background() which is a non-nil empty context that is
	// never canceled, has no values, and has no deadline. It is typically used by the main
	// function, initialization, and tests, and as the top-level Context for incoming requests.
	DefaultContext context.Context = context.Background()
)

// Input - displays a prompt and waits for user input
func Input(prompt string) (string, error) {
	fmt.Print(prompt)
	return bufio.NewReader(os.Stdin).ReadString('\n')
}

// PW - displays a prompt and waits for user secure password input
func PW(prompt string) (string, error) {
	fmt.Print(prompt)
	bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
	return string(bytePassword), err
}

// getEnvUser - get Git username (const envGitUsernameKey) from environment variable
func getEnvUser() (string, error) {
	username, ok := os.LookupEnv(envGitUsernameKey)
	if !ok {
		return "", fmt.Errorf("Git username environment variable <%v> not found", envGitUsernameKey)
	}
	return username, nil
}

// getEnvToken - get PAT from environment variable (const envGitTokenKey)
// global tokenSource is used to cache
func getEnvToken() (oauth2.TokenSource, error) {

	// todo - add options for cli input if ENV variables not present?
	// username, _ := Input("GitHub Username: ")
	// password, _ := PW("GitHub Password: ")

	if tokenSource != nil { // if global is already set, pass that ...
		return tokenSource, nil
	}

	envToken, ok := os.LookupEnv(envGitTokenKey)
	if !ok {
		return nil, fmt.Errorf("github personal access token (PAT) environment variable <%v> not found", envGitTokenKey)
	}

	tokenSource = oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: envToken},
	)
	return tokenSource, nil
}

// AuthClient returns an authenticated *http.Client for use with API requests. It relies
// on getEnvToken to collect a secure Personal API Token (PAT) from an environment
// variable or other secure source.
//
// If authentication fails, AuthClient falls back and returns http.DefaultClient with
// no authentication. Any error is also returned; calling func has a choice of action.
//
// Most users will pass in ctx := DefaultContext which represents context.Background().
// It is a non-nil empty context that is never canceled, has no values, and has no
// deadline. It is typically used by the main function, initialization, and tests, and
// as the top-level Context for incoming requests.
//
func AuthClient(ctx context.Context) (*http.Client, error) {
	// todo - 2 factor authentication
	// Is this a two-factor auth error? If so, prompt for OTP and try again.
	// if _, ok := err.(*github.TwoFactorAuthError); ok {
	// 	otp, _ := Input("\nGitHub OTP: ")
	// 	gc.tp.OTP = strings.TrimSpace(otp)
	// 	user, _, err = client.Users.Get(ctx, "")
	// }
	ts, err := getEnvToken()
	if err != nil {
		return http.DefaultClient, err
	}

	return oauth2.NewClient(ctx, ts), nil
}

// NewGitHubClient returns a pointer to a new authenticated GitHub client
//
// Environment variables can be used for authentication:
//
// username = GIT_AUTHOR_NAME
//
// password = GITHUB_API_TOKEN_VSCODE
//
// Most users will pass in ctx := DefaultContext which represents context.Background().
// It is a non-nil empty context that is never canceled, has no values, and has no
// deadline. It is typically used by the main function, initialization, and tests, and
// as the top-level Context for incoming requests.
func NewGitHubClient(ctx context.Context) (*github.Client, error) {
	httpClient, err := AuthClient(ctx)
	if err != nil {
		return nil, err
	}
	return github.NewClient(httpClient), nil
}
