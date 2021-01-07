// Package ini has resources to support configuration files.
package ini

import (
	"bufio"
	"io"
	"io/ioutil"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/skeptycal/util/gofile"
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
	fi os.FileInfo // fi is stored to avoid repeated os queries
	f  *os.File    // f is stored to avoid repeated os queries
	br *bufio.Reader
	sc *bufio.Scanner
	sb *strings.Builder // sb is reused for this ini file
}

func (i *ini) Size() int { return int(i.fi.Size()) }

func NewIni(path string) *ini {
	if !gofile.Exists(path) {
		log.Errorf("path does not exist: %v", path)
		return nil
	}
	in := ini{}
	in.fi, _ = os.Stat(path)
	f, err := os.Open(path)
	if err != nil {
		log.Error(err)
		return nil
	}
	in.f = f
	defer in.f.Close()

	in.sb = &strings.Builder{}

	in.br = bufio.NewReaderSize(in.f, int(in.Size()))
	n, err := in.br.Read()
	if err != nil {
		log.Error(err)
		return nil
	}
	if n != int(in.Size()) {
		log.Errorf("incorrect number of bytes read: %d want: %d", n, in.Size())
	}

	in.sc = bufio.NewScanner(in.br)

	// buf, err := ioutil.ReadAll(file)

}

func readfile(r io.Reader) []byte {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		log.Error(err)
		return nil
	}
	return b
}

func (i *ini) RemoveComments() (string, error) {
	file, err := os.Open(path)
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
func RemoveComments(path string) (string, error) {

}
