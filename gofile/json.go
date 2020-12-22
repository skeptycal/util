package gofile

import (
	"encoding/json"
	"errors"
	"io/ioutil"
)

// JSON describes a JSON object.
type JSON interface {
	New(name string) error
	Load() error
	Size() int
}

type jsonStruct struct {
	filename string
	size     int64
	data     []byte
	v        interface{}
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
