package zsh

import (
	"os"
	"strings"
)

// GetEnv returns the environment variable specified by 'key'; if the value is empty, the
// default value is returned; if the value is not set, an error is also returned.
func GetEnv(key string, defValue string) string {
	value, b := os.LookupEnv(key)

	if !b {
		return defValue
	}

	if strings.TrimSpace(value) == "" {
		return defValue
	}
	return value
}
