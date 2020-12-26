package http

// Examples from https://golang.org/pkg/net/http/

// Get return text from url request
// func Get(url string) {
// 	resp, err := http.Get("http://example.com/")
// 	if err != nil {
// 		// handle error
// 	}
// 	defer resp.Body.Close()
// 	body, err := ioutil.ReadAll(resp.Body)
// }

// func client(url string) {
// 	client := &http.Client{
// 		CheckRedirect: redirectPolicyFunc,
// 	}

// 	resp, err := client.Get("http://example.com")
// 	// ...

// 	req, err := http.NewRequest("GET", "http://example.com", nil)
// 	// ...
// 	req.Header.Add("If-None-Match", `W/"wyzzy"`)
// 	resp, err := client.Do(req)
// }

// func simpleServer(myHandler http.Handler) {
// 	http.Handle("/foo", myHandler)

// 	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
// 	})

// 	log.Fatal(http.ListenAndServe(":8080", nil))
// }

// func server(myHandler http.Handler) {
// 	s := &http.Server{
// 		Addr:           ":8080",
// 		Handler:        myHandler,
// 		ReadTimeout:    10 * time.Second,
// 		WriteTimeout:   10 * time.Second,
// 		MaxHeaderBytes: 1 << 20,
// 	}
// 	log.Fatal(s.ListenAndServe())
// }
