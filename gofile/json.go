package gofile

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type jsonMap map[string]interface{}

// GetJSONFile loads and returns a JSON structure representing the json file.
func GetJSONFile(filename string) (JSON, error) {
	fi, err := GetFileInfo(filename)
	if err != nil {
		return nil, err
	}

	j := &jsonStruct{fi, &jsonMap{}}
	j.Load()
	if err != nil {
		return nil, err
	}
	return j, nil
}

// func NewJSONFile(filename string) (*jsonMap, error) {
// 	f, err := os.Create(filename)

// 	j := new(jsonMap)

// }

// JSON describes a JSON file and data structure object.
type JSON interface {
	Load() error
	Name() string
	Save() error
	Size() int64
	Marshal(v interface{}) ([]byte, error)
	Unmarshal(data []byte, v interface{}) error
}

// jsonStruct implements a JSON mapping with os.FileInfo included.
type jsonStruct struct {
	os.FileInfo
	v *jsonMap
}

// Load loads the data from the JSON file
// note: variable/field names should begin with a uppercase letter
// of they will not load correctly - similar to the uppercase requirement
// for exported functions ...
func (j *jsonStruct) Load() error {
	data, err := ioutil.ReadFile(j.Name())
	if err != nil {
		return err
	}
	return j.Unmarshal(data, j.v)
}

// Save saves the data to the JSON file
// note: variable/field names should begin with a uppercase letter
// of they will not save correctly - similar to the uppercase requirement
// for exported functions ...
func (j *jsonStruct) Save() error {
	data, err := j.Marshal(j.v)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(j.Name(), data, 0644)
}

// Unmarshal is only present to satisfy the Unmarshaler interface requirement.
// If used, v is ignored and the interface{} from the internal structure is used.
func (j *jsonStruct) Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, j.v)
}

// Marshal is only present to satisfy the Marshaler interface requirement.
// If used, v is ignored and the interface{} from the internal structure is used.
func (j *jsonStruct) Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(j.v)
}
