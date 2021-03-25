package http

import (
	"errorutils"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Fetch - get response from url and check for standard common errors
func Fetch(url string) (resp *http.Response, err error) {
	resp, err = http.Get(url)
	if errorutils.Errf(err, "error obtaining response from server: %v") {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		// Exit if rate limited ...
		if resp.StatusCode == http.StatusTooManyRequests {
			log.Fatal("You are being rate limited...")
		} else {
			return nil, fmt.Errorf("server returned error code: %v", resp.Status)
		}
	}
	return resp, nil
}

// ReadAllURL reads all available bytes from response using a buffer
func ReadAllURL(url string) (string, error) {
	resp, err := http.Get(url)
	if errorutils.Errf(err, "error obtaining response from server: %v") {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		// Exit if rate limited ...
		if resp.StatusCode == http.StatusTooManyRequests {
			log.Fatal("You are being rate limited...")
		} else {
			return "", fmt.Errorf("server returned error code: %v", resp.Status)
		}
	}

	bytes, err := ioutil.ReadAll(resp.Body)
	if errorutils.Errf(err, "error obtaining response from server: %v") {
		return "", err
	}

	return string(bytes), nil
}
