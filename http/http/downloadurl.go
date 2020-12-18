package http

import (
	"io"
	"net/http"

	"github.com/skeptycal/util/zsh/gofile"
)

// DownloadURL - download content from a URL to <filename>
func DownloadURL(url, filename string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	f, err := gofile.CreateFileTruncate(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = io.Copy(f, resp.Body)
	if err != nil {
		return err
	}
	return nil
}
