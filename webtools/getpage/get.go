package getpage

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

type responseHeaders struct {
	AcceptEncoding string // `json: Accept-Encoding`
	Host           string
	UserAgent      string // `json: User-Agent`
	XAmznTraceID   string // `json: X-Amzn-Trace-Id`
}

type resp struct {
	args map[string]string
	responseHeaders
	origin string
	url    string
}

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
