//go:generate go run makeisprint.go -output isprint.go

package stringparse

import (
	"fmt"
	"path"
	"reflect"
	"runtime"
	"strings"
)

// Configuration - program configuration settings
// environment variables have priority
type Configuration struct {
	Help    string
	Version string `env:"STRING_PARSE_VERSION"`
	Env     string `env:"STRING_PARSE_ENV"`
}

// Configure - get configuration values from JSON file / env variables
// modified version of code in the popular 'gonfig' package
func Configure(fileName string, v interface{}) (err error) {

	configType := reflect.ValueOf(v).Type()
	if configType.Kind() != reflect.Ptr || configType.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("v must be a pointer to a struct type, got: %v", configType)
	}
	return nil
	// todo
	// return LoadJSON(fileName, v)
}

// ConfigFileName - returns the name of the json config file matching the go file name
// e.g. path/to/my_program.go would return path/to/my_program.json
func ConfigFileName() string {

	_, pathName, _, _ := runtime.Caller(0)

	filePath, fileName := path.Split(pathName)
	baseName := strings.Split(fileName, ".")[0] + ".json"

	return path.Join(filePath, baseName)
}
