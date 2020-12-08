package http

import (
	"ideas/scargo/fileutils"
	"io"
	"net/http"
)

// DownloadURL - download content from a URL to <filename>
func DownloadURL(url, filename string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	f, err := fileutils.CreateFileTruncate(filename)

	_, err = io.Copy(f, resp.Body)
	return nil
}
