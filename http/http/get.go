package http

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/skeptycal/util/gofile"
)

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
func GetPage(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("server returned error code: %v", resp.Status)
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(bufio.NewReaderSize(resp.Body, gofile.InitialCapacity(resp.ContentLength)))
	if err != nil {
		return "", err
	}

	return string(b), nil
}

// GetPage - return result from url
func BuffPage(sb *strings.Builder, url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("server returned error code: %v", resp.Status)
	}

	defer resp.Body.Close()

	_, err = io.Copy(sb, resp.Body)

	if err != nil {
		return err
	}

	return nil
}
