// Package getpage gets webpage information.
package getpage

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

// These constants are mostly modified variations or unexported parts
// of the Go standard library code (from Go 1.15.5)
const (
	defaultMaxPageCacheAge   time.Duration = 30 * time.Second
	defaultPageBufferSize    int           = 0xFFFF
	minReadBufferSize                      = 16
	smallBufferSize                        = 64
	defaultBufSize                         = 4096
	maxConsecutiveEmptyReads               = 100
	maxInt                                 = int(^uint(0) >> 1)
	chunk                                  = 512 // bytes.MinRead
	SEP                                    = string(os.PathSeparator)
)

// A PageBuilder is a strings.Builder specifically designed
// for parsing http response bodies.
// A Builder is used to efficiently build a page using Write methods.
// It minimizes memory copying. The zero value is ready to use.
// Do not copy a non-zero Builder.
type pageBuilder struct {
	url       string
	timestamp time.Time
	strings.Builder
}

func (pb *pageBuilder) SetTimestamp() {
	pb.timestamp = time.Now()
}

func (pb *pageBuilder) Timestamp() time.Time {
	return pb.timestamp
}

func (pb *pageBuilder) Age() time.Duration {
	return time.Since(pb.timestamp)
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
func GetPage(url string) (string, error) {

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	size := int(resp.ContentLength)
	if size < 0 {
		size = defaultPageBufferSize
	}

	var sb = pageSet.New(url)
	defer sb.Reset()

	// If the buffer overflows, we will get bytes.ErrTooLarge.
	// Return that as an error. Any other panic remains.
	defer func() {
		e := recover()
		if e == nil {
			return
		}
		if panicErr, ok := e.(error); ok && panicErr == bytes.ErrTooLarge {
			err = panicErr
		} else {
			panic(e)
		}
	}()

	if sb.Cap() < size {
		sb.Grow(size)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	sb.Write(body)

	return sb.String(), nil
}
