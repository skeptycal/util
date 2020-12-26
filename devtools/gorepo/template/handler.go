package template

import (
	"net/http"
	"os"
	"text/template"

	log "github.com/sirupsen/logrus"
)

// handler is a template handler
/*
   As you can see, it's very easy to use, load and render data in templates in Go, just as in other programming languages.

   For the sake of convenience, we will use the following rules in our examples:

   Use Parse to replace ParseFiles.
   Parse can test content directly from strings, so we don't need any extra files.

   Use main for every example and do not use handler.

   Use os.Stdout to replace http.ResponseWriter.
   (os.Stdout also implements the io.Writer interface.)

   GetUser() is a generic username and password authentication handler.
   It should not be used in production.
*/
func handler(w http.ResponseWriter, r *http.Request) {
	t := template.New("some template")           // Create a template.
	t, _ = t.ParseFiles("tmpl/welcome.html", "") // Parse template file.
	// user := GetUser()                             // Get current user infomration.
	err := t.Execute(w, os.Stdout) // merge.
	if err != nil {
		log.Error(err)
	}
}
