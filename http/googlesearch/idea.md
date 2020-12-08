create a cli app to search google
return top (~20??) hits (.conf file)
show them in a numbered list
choose one to show in detail
skip adds, bs, etc
show preview images (base64) for iterm2
(turn iterm2 support on/off in .conf file)

```go
// Plan:
// ----------------------------
// read args
// form search string
// use get request
// obtain result
// print top results
//  - color coded
//  - with links
//  -
```
---

and so ... work starts:



```go
// 'simple' version example:
// https://www.google.com/search?q=Tokyo+Disneyland+Rides&um=1&safe=active&hl=en&biw=1253&bih=789&tbm=isch

type query struct {
	search string //required
	Title  string //optional
	tbm    string // =isch is image search ... others??? or is this binary???
	safe   bool   // safe=active,???
	hl     string // language?
	biw    int    // 'b' image width???
	bih    int    // 'b' image height???
	// color ... &tbs=ic:specific,isc:red
	// date range ... tbs=cdr:1,cd_min:4/1/2013,cd_max:5/1/2013
	// size ... tbs=isz:l
}

// func makeURL(s string) {

// 	Url, err := url.Parse(s)
// 	if err != nil {
// 		panic("boom")
// 	}

// 	Url.Path += "/some/path/or/other_with_funny_characters?_or_not/"
// 	parameters := url.Values{}
// 	parameters.Add("hello", "42")
// 	parameters.Add("hello", "54")
// 	parameters.Add("vegetable", "potato")
// 	Url.RawQuery = parameters.Encode()

// 	fmt.Printf("Encoded URL is %q\n", Url.String())
// }
```
apparently, the rules for forming Google search queries are purposely objuscated ... there seems to be no easy way to search for this ... ironically, I'm using google to search ... so this adds to my suspicion.

```go
// Reference:
//
// This was the string I got by searching for 'chemistry lithium ion battery' while logged in.
//
// https://www.google.com/search?
// sxsrf=ALeKk01ZBzQXMzBv7acHl-P1Pm35r21b7g%3A1604000453498
// source=hp
// ei=xRqbX73WGsvwsAWKyqPAAw
// q=chemistry+lithium+ion+battery
// oq=chemistry+lithium+ion+battery
// gs_lcp=CgZwc3ktYWIQAzIFCAAQyQMyBggAEBYQHjIGCAAQFhAeMgYIABAWEB4yBggAEBYQHjI \
//      GCAAQFhAeMgYIABAWEB4yBggAEBYQHjIGCAAQFhAeMgYIABAWEB46BAgjECc6BQgAEJEC \
//      OgUIABCxAzoCCAA6CAgAELEDEIMBOgQIABBDOgsILhCxAxDHARCjAjoECC4QQzoHCAAQs \
//      QMQQzoFCC4QsQM6AgguOggIABAWEAoQHlCDFlj_SWClTmgAcAB4AIABrwKIAcwWkgEIMj \
//      IuNi4wLjGYAQCgAQGqAQdnd3Mtd2l6
// sclient=psy-ab
// ved=0ahUKEwj9zPypx9rsAhVLOKwKHQrlCDgQ4dUDCAk
// uact=5
```
