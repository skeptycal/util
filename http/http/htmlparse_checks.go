package http

import (
	"ideas/scargo/errorutils"
	"ideas/scargo/fileutils"
)

func fileCreateChecks() {
	_, err := fileutils.CreateFileTruncate("example_rw.txt")
	_ = errorutils.Errf(err, "")
	_, err = fileutils.CreateFileSafe("example_read.txt")
	_ = errorutils.Errf(err, "")
}

func checkFileFuncs() {
	fileCreateChecks()
}
