package file

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"os"
)

// JSON describes a JSON object.
type JSON interface {
	new()
	load()
	size()
}

type jsonStruct struct {
	file string
	size int64
	data []byte
}

var j JSON

func (j *jsonStruct) new(fileName string) *jsonStruct {

	const chunk = bytes.MinRead
	file, err := os.Stat(fileName)
	iSize := initialSize(fSize.Size(), chunk)

	return &jsonStruct{
		file: fileName,
		size: initialSize(iSize, chunk),
		data: make([]byte, 0),
	}
	return j
}

// LoadJSON - loads the data from the JSON file
// note: variable/field names should begin with a uppercase letter
// of they will not load correctly - similar to the uppercase requirement
// for exported functions ...
func (j *jsonStruct) Load(file string) (err error) {

	data, err := ioutil.ReadFile(file)
	if err != nil {
		return
	}

	// dataBuffer := bytes.NewBuffer(data).Bytes()

	return json.Unmarshal(data, j)
}
