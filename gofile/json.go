package gofile

import (
	"bufio"
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
)

type jsonMap map[string]interface{}

// buf is a reusable buffer for reading JSON files.
// a global buffer does not seem to help much
// and definitely isn't concurrency friendly
// var buf bufio.Reader

// JSON describes a JSON object.
type JSON interface {
	Load() error
	Size() int
}




func New(filename string) (*JSON, error) {

    r, err := NewBufferedReader(filename)


	jsonFile, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	size := jsonFile.

	buf := bufio.NewReader

	b, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	var jm jsonMap

	json.Unmarshal(b, &jm)

	return &jm, nil
}

type jsonStruct struct {
	filename string
	size     int64
	data     []byte
	v        *jsonMap
}

// Load - loads the data from the JSON file
// note: variable/field names should begin with a uppercase letter
// of they will not load correctly - similar to the uppercase requirement
// for exported functions ...
func (j *jsonStruct) Load(file string) (err error) {
	j.data, err = ioutil.ReadFile(j.filename)
	if err != nil {
		return
	}

	// dataBuffer := bytes.NewBuffer(data).Bytes()

	return json.Unmarshal(j.data, j.v)
}

// New returns a new JSON object.
func (j *jsonStruct) New(name string) error {
	return errors.New("Not Implemented")
}

// Size returns the number of bytes occupied by the JSON data.
// Returns -1 on error.
func (j *jsonStruct) Size() int {
	// todo - not implemented
	return int(j.size)
}
