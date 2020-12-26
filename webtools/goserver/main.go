package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

const ()

var (
	port int = 8080
	url      = fmt.Sprintf("http://localhost:%d", port)
)

// Route stores information to match a request and build URLs.
type Route interface {
	Handler() *Route
	HandlerFunc() *Route
}

// BasicServer is a test server that serves static text only.
type BasicServer interface {
	HandleFunc() *mux.Route
	ListenAndServe() error
}

type basicServer struct {
	name     string
	response string
	route    Route
}

// A SimpleHandler responds to an HTTP request.
//
// ServeHTTP should write reply headers and data to the ResponseWriter
// and then return. Returning signals that the request is finished; it
// is not valid to use the ResponseWriter or read from the
// Request.Body after or concurrently with the completion of the
// ServeHTTP call.
//
// Depending on the HTTP client software, HTTP protocol version, and
// any intermediaries between the client and the Go server, it may not
// be possible to read from the Request.Body after writing to the
// ResponseWriter. Cautious handlers should read the Request.Body
// first, and then reply.
//
// simpleHandler sets a handler function for the route.
//
// Except for reading the body, handlers should not modify the
// provided Request.
func (b *basicServer) SimpleHandler() *mux.Route {
	func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello Medium, you've requested: %s\n", r.URL.Path)
	}

}

func startBasicServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello Medium, you've requested: %s\n", r.URL.Path)
	},
	)

	http.ListenAndServe(":8080", nil)
}

func routeServer() error {

	fs := http.FileServer(http.Dir("static/"))

	http.Handle("/static/", http.StripPrefix("/static/", fs))

	r := mux.NewRouter()

	r.HandleFunc("/user/{username}", func(w http.ResponseWriter, r *http.Request) {
		// TODO - route code ...

	}).Methods("GET")

	err := http.ListenAndServe(":80", r)
	if err != nil {
		return err
	}
	return nil
}

func htmlServer() error {
	tmpl := template.Must(template.ParseFiles("layout.html"))
	// var t = template.Must(template.New("name").Parse("html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := tmpl.Execute(w, "")
		if err != nil {
			log.Error(err)
		}
	})

	err := http.ListenAndServe(":80", nil)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	fmt.Printf("Server is active at: %s\n", url)
	fmt.Println("   (Press <ctrl>-C to stop server.)")
	err := routeServer()
	if err != nil {
		log.Panicln(err)
	}
}
