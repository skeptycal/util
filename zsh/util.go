package zsh

import (
	"fmt"
	"os"
)

// Optional is an optional argument
type Optional interface{}

// GetEnv returns the environment variable specified by 'key'; if the value is empty, the
// default value is returned; if the value is not set, an error is also returned.
func GetEnv(key string, defValue string) (string, error) {
	value, b := os.LookupEnv(key)
	if value != "" {
		return value, nil
	}

	if b {
		return defValue, nil
	}

	return defValue, fmt.Errorf("environment variable not present: %s", key)
}

// ToString - convert any value to string
func ToString(value interface{}) string {
	return fmt.Sprintf("%v", value)
}
