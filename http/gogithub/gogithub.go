package gogithub

import (
	"context"
	"net/http"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

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
	ts, err := GetEnvToken()
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
