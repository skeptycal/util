package http

import (
	"io"
	"net/http"
)

// Get is a wrapper around DefaultClient.Get.
// To make a request with custom headers, use NewRequest and DefaultClient.Do.

// GetURL - return content from a URL
func GetURL(url string) (io.ReadCloser, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return resp.Body, nil
}

func request(url string) (io.ReadCloser, error) {
	return nil, nil
}
