package commands

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type jsonMap map[string]interface{}

func PrintAllCommands() {
	fmt.Println(allCommands)
}

func GetCommandList(filename string, v map[string]interface{}) (*jsonMap, error) {

	jsonFile, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	b, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	var jm jsonMap

	json.Unmarshal(b, &jm)

	return &jm, nil
}
