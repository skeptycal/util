package main

import (
	"os"
	"path/filepath"

	"github.com/prometheus/common/log"
)

const (
	defaultDir = ".gocom"
)

func main() {
	home := filepath.Join(os.Getenv("HOME"), defaultDir)
	log.Info("Home directory: %s", home)

}
