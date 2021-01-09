// Package godirs sets up the directory structure
// for a new go module.
package godirs

import "path"

var reponame = "gotemp"

var Commands = map[string]interface{}{
	"dirs": []string{
		reponame,
		path.Join(reponame, "cmd", reponame),
	},
	"files": map[string]string{
		"cmdMainSample": `package main


func main() {

}
`,
	},
}
