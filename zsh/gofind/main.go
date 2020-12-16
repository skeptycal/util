package main

import (
	"log"
	"os"
	"path/filepath"
	"regexp"
	"runtime/trace"
)

// walkFn type WalkFunc func(path string, info os.FileInfo, err error) error
// WalkFunc is the type of the function called for each file or directory
// visited by Walk. The path argument contains the argument to Walk as a
// prefix; that is, if Walk is called with "dir", which is a directory
// containing the file "a", the walk function will be called with argument
// "dir/a". The info argument is the os.FileInfo for the named path.
//

func main() {
	err := trace.Start(os.Stdout)
	if err != nil {
		log.Fatal(err)
	}
	defer trace.Stop()

	libRegEx, err := regexp.Compile(`^.+\.(go)$`)
	if err != nil {
		log.Fatal(err)
	}

	err = filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err == nil && libRegEx.MatchString(info.Name()) {
			println(info.Name())
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}

/*
cat git_repos.txt | cut -d '/' -f 2 | uniq | find \! -name 'keep' -empty type -f
find / \! -name "*.c" -print
*/
