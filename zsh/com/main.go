package main

import (
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"
)

const (
	defaultDir = ".gocom"
)

func main() {
	home := filepath.Join(os.Getenv("HOME"), defaultDir)
	log.Info("Home directory: %s", home)

}
