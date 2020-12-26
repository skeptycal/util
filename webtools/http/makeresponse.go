package http

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

func ResponseFromString(body string) *http.Response {
	// body := "Hello world"
	return &http.Response{
		Status:        "200 OK",
		StatusCode:    200,
		Proto:         "HTTP/1.1",
		ProtoMajor:    1,
		ProtoMinor:    1,
		Body:          ioutil.NopCloser(bytes.NewBufferString(body)),
		ContentLength: int64(len(body)),
		Request:       nil,
		Header:        make(http.Header, 0),
	}
}
