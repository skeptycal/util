package gogithub

import (
	"context"

	"golang.org/x/oauth2"
)

// todo - rename the default environment variable keys appropriately

const (
	// environment variable for GitHub username
	envGitUserNameKey = "GIT_AUTHOR_NAME"
	// environment variable for GitHub PAT
	envGitTokenKey = "GITHUB_API_TOKEN"

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

	// MaxGitHubRateLimit - the maximum rate limit for GitHub API requests in apps
	MaxGitHubRateLimit = 12500
)

var (
	// tokenSource - used to cache the environment variable GITHUB_API_TOKEN
	//
	// GitHub Apps ask for repository contents permission and use your installation token to
	// authenticate via HTTP-based Git. The token is used as the HTTP password.
	// todo - this is likely not very useful since tokens are not refreshed very often
	tokenSource oauth2.TokenSource

	// DefaultContext represents context.Background() which is a non-nil empty context that is
	// never canceled, has no values, and has no deadline. It is typically used by the main
	// function, initialization, and tests, and as the top-level Context for incoming requests.
	DefaultContext context.Context = context.Background()
)
