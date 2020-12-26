package gofile

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
)

// CreateFileTruncate - Creates a file and returns a closer on success else error
func CreateFileTruncate(filename string) (*os.File, error) {
	file, err := os.Create(filename)
	if err != nil {
		return nil, fmt.Errorf("failed creating file (%s): %v", filename, err)
	}
	return file, nil
}

// CreateFileSafe - Creates a file and returns a closer on success else error
func CreateFileSafe(filename string) (*os.File, error) {
	file, err := os.Create(filename)
	if err != nil {
		log.Error(err)
		// err1 := file.Close()
		// log.Error(err1)
		return nil, fmt.Errorf("failed creating file (%s): %v", filename, err)
	}
	return file, nil
}
