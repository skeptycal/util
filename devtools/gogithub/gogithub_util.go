package gogithub

import (
	"bufio"
	"fmt"
	"os"
	"syscall"

	"golang.org/x/crypto/ssh/terminal"
	"golang.org/x/oauth2"
)

// + Input - displays a prompt and waits for user input
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
	username, ok := os.LookupEnv(envGitUserNameKey)
	if !ok {
		return "", fmt.Errorf("Git username environment variable <%v> not found", envGitUserNameKey)
	}
	return username, nil
}

// GetEnvToken - get PAT from environment variable (const envGitTokenKey)
// global tokenSource is used to cache
func GetEnvToken() (oauth2.TokenSource, error) {

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
