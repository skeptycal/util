package getpage

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	defaultPageBufferSize = 32768
)

// A PageBuilder is a strings.Builder specifically designed
// for parsing http response bodies.
// A Builder is used to efficiently build a string using Write methods.
// It minimizes memory copying. The zero value is ready to use.
// Do not copy a non-zero Builder.
type PageBuilder struct {
	addr *PageBuilder // of receiver, to detect copies by value
	buf  []byte
}

type Builder interface {
	String() string
	Len() int
	Cap() int
	Reset()
	Grow(n int)
	Write(p []byte) (n int, err error)
	WriteByte(c byte) error
	WriteRune(r rune) (int, error)
	WriteString(s string) (n int, err error)
}

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

	size := int(resp.ContentLength)
	if size < 0 {
		size = defaultPageBufferSize
	}

	// buf := make([]byte, size, size)

	var sb strings.Builder
	defer sb.Reset()

	buf := Builder()

	sb.Grow(size)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func GetBuilder(ctx context.Context, b PageBuilder) (Builder, error) {
	return nil
}
