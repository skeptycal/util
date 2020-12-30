package gofile

import (
	"encoding/csv"
	"os"
)

// NewCSVReader returns a new Reader that reads from file.
func NewCSVReader(file string) (*csv.Reader, error) {
	fi, err := os.Open(file)
	if DoOrDie(err) != nil {
		return nil, err
	}

	return csv.NewReader(fi), nil
}

// ReadCSV reads all records from file.
// Each record is a slice of fields.
// A successful call returns err == nil, not err == io.EOF. Because ReadAll is
// defined to read until EOF, it does not treat end of file as an error to be
// reported.
func ReadCSV(file string) ([][]string, error) {
	r, err := NewCSVReader(file)
	if DoOrDie(err) != nil {
		return nil, err
	}
	return r.ReadAll()
}
