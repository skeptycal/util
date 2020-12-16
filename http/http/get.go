package http

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

type rheaders struct {
	AcceptEncoding string `json: Accept-Encoding`
	Host           string
	UserAgent      string `json: User-Agent`
	XAmznTraceID   string `json: X-Amzn-Trace-Id`
}

type resp struct {
	args map[string]string
	rheaders
	origin string
	url    string
}

// sample response from test sent to http://httpbin.org
// {
// "args": {},
// "headers": {
//     "Accept-Encoding": "gzip",
//     "Host": "httpbin.org",
//     "User-Agent": "Go-http-client/2.0",
//     "X-Amzn-Trace-Id": "Root=1-5fb2dd0a-7db005a350e64f64552e1f09"
// },
// "origin": "38.75.198.14",
// "url": "https://httpbin.org/get"
// }

// GetPage - return result from url
func GetPage(url string) (*bytes.Buffer, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return bytes.NewBuffer(body), nil
}
