package http

import (
	"io"
	"net/http"
)

// SimpleServer - basic server
func SimpleServer(port string) error {
	// Hello world, the web server

	if port == "" {
		port = ":8080"
	}

	Handler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, world!\n")
	}

	http.HandleFunc("/hello", Handler)
	http.ListenAndServe(port, nil)
	return nil
}
