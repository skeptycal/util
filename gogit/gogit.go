package gogit

import "strings"

// gi returns a string response from the www.gitignore.io API containing
// standard .gitignore items for the args given.
//
//      default: "macos linux windows ssh vscode go zsh node vue nuxt python django"
//
// using: https://www.toptal.com/developers/gitignore/api/macos,linux,windows,ssh,vscode,go,zsh,node,vue,nuxt,python,django
func gi(args string) string {

	if len(args) == 0 {
		args = []string{"macos linux windows ssh vscode go zsh node vue nuxt python django"}
	}

	command := "curl -fLw '\n' https://www.gitignore.io/api/\"${(j:,:)@}\" "
	command += strings.Join(args, " ")

	return Shell(command)
}
