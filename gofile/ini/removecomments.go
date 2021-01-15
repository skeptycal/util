// Package ini has resources to support configuration files.
package ini

import (
	"bufio"
	"io"
	"io/ioutil"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/skeptycal/util/gofile/filereader"
)

// RemoveComments returns the contents of a text file with
// comments removed.
//
// Comments are defined as lines of text that begin with a
// comment prefix. Leading whitespace characters are  ignored,
// and optionally removed. The default implementation uses
// the following list of comment strings:
//
//  "#", "//", ";"
//
// Other options include:
//  removeWhiteSpace - strip all leading and trailing whitespace
//  remove trailing comments - strip comments at end of lines
//
var (
	commentPrefixes = []string{"#", ";", "//"}
)

type ini struct {
	filereader.BufferedReadWriter                  // file readwriter
	sc                            *bufio.Scanner   // data parser
	sb                            *strings.Builder // string buffer
}

// type ini struct {
// 	w  Writer
// 	r  Reader
// 	fi os.FileInfo // fi is stored to avoid repeated os queries
// 	f  *os.File    // f is stored to avoid repeated os queries
// 	br *bufio.Reader
// 	sc *bufio.Scanner
// 	sb *strings.Builder // sb is reused for this ini file
// }

// func (i *ini) Size() int { return int(i.fi.Size()) }

// func NewIni(filename string, data []byte) (*ini, error) {
// if !gofile.Exists(filename) {
// 	log.Errorf("path does not exist: %v", filename)
// 	return nil, nil
// }
// r, err := filereader.NewBufferedReader(filename)
// w, err := filereader.NewBufferedWriter(filename, data)
// bufio.NewReadWriter(r *bufio.Reader, w *bufio.Writer)

// in := ini{r, w, bufio.NewScanner(*r), &strings.Builder{}}
// in.BufferedFileReader, _ = os.Stat(filename)
// f, err := os.Open(filename)
// if err != nil {
// 	log.Error(err)
// 	return nil, nil
// }
// in.BufferedFileWriter = f
// defer in.BufferedFileWriter.Close()

// in.sb = &strings.Builder{}

// in.BufferedFileReader = bufio.NewReaderSize(in.BufferedFileReader, int(in.BufferedFileReader.Size()))
// if err != nil {
// 	log.Error(err)
// 	return nil, nil
// }

// in.sc = bufio.NewScanner(in.BufferedFileReader)

// buf, err := ioutil.ReadAll(file)

// }

func readfile(r io.Reader) []byte {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		log.Error(err)
		return nil
	}
	return b
}

func (i *ini) RemoveComments() (string, error) {
	file, err := os.Open(i.BufferedReadWriter.)
	if err != nil {
		log.Error(err)
		return "", err
	}
	defer file.Close()

	sb := strings.Builder{}
	defer sb.Reset()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := strings.TrimSpace(scanner.Text())
		if IsComment(s, commentPrefixes) {
			continue
		}
		sb.WriteString(s)
		sb.WriteString("\n")
	}

	if err := scanner.Err(); err != nil {
		log.Error(err)
		return sb.String(), err
	}
	return sb.String(), nil

}

func IsComment(comment string, commentPrefixes []string) bool {
	for _, c := range commentPrefixes {
		if strings.HasPrefix(comment, c) {
			return true
		}
	}
	return false
}

// func RemoveComments(path string) (string, error) {

// }
