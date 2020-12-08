package cli

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

// SEP - os specific path separator
const SEP = os.PathSeparator

// BaseName - return final path element by string search
func BaseName(str string) string {
	str = strings.TrimSpace(str)
	if len(str) < 2 {
		return str
	}

	if str[len(str)] == SEP {

		// str[len(str)-1] = ""   // Erase element (write zero value)
		str = str[:len(str)-1] // [A B]
	}
	n := strings.LastIndex(str, string(SEP)) + 1
	// if there is no path separator, assume the entire string is a filename
	if n < 0 {
		return str
	}
	return str[n:]
}

// basename removes trailing slashes and the leading directory name from path name.
// from path_unix.go: .../go/1.15.3/libexec/src/os/path_unix.go
func basename(name string) string {
	i := len(name) - 1
	// Remove trailing slashes
	for ; i > 0 && name[i] == '/'; i-- {
		name = name[:i]
	}
	// Remove leading directory name
	for i--; i >= 0; i-- {
		if name[i] == '/' {
			name = name[i+1:]
			break
		}
	}

	return name
}

func grabTag(tag string, data string) error {
	return nil
}

func main() {
	println("pathsep: ", SEP)
	if len(os.Args) != 2 {
		var name string = os.Args[0]
		fmt.Fprintf(os.Stderr, "Usage: %c URL\n", BaseName(name)[0])
		os.Exit(1)
	}
	response, err := http.Get(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	_, err = io.Copy(os.Stdout, response.Body)
	if err != nil {
		log.Fatal(err)
	}
}
