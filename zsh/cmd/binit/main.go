package main

import (
	"fmt"
	"io"
	"os"
	"path"

	log "github.com/sirupsen/logrus"
	"github.com/skeptycal/util/gofile"
)

const (
	devMode   = true
	useLogger = true
	name      = "binit"
	targetDir = "./example"
)

func init() {
	if useLogger {
		LogFormatter := new(log.TextFormatter)
		LogFormatter.TimestampFormat = "02-01-2006 15:04:05"
		LogFormatter.FullTimestamp = true
		log.SetFormatter(LogFormatter)

		log.Info("logrus initialized for ", name)

		if devMode {
			log.SetLevel(log.DebugLevel)
		} else {
			log.SetLevel(log.ErrorLevel)
		}
	}
}

// name returns the name of this file
func me() string {
	return gofile.Base(os.Args[0])
}

// here returns the location of this file
func here() string {
	return gofile.Abs(os.Args[0])
}

// OsArgs returns the slice of strings returned by os.Args[1:]
func OsArgs() []string {
	return os.Args[1:]
}

func copy(src, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}

// get source file name
// create target name
// check exist target
func main() {

	extList := []string{"py", "sh", "php"}
	args := OsArgs()

	// log.Info("me: ", me())
	// log.Info("here: ", here())
	// log.Info("target dir: ", targetDir)
	// log.Info("args: ", args)
	log.Info("======================================")

	for i := range args {
		sourceFileStat, err := os.Stat(args[i])
		if err != nil || sourceFileStat.IsDir() {
			log.Error(err)
			continue
		}

		fileName := gofile.Abs(sourceFileStat.Name())
		baseName, extension := zsh.NameSplit(fileName)
		log.Info("fileName: ", fileName)
		log.Info("baseName: ", baseName)
		log.Info("extension: ", extension)

		if !sourceFileStat.Mode().IsRegular() {
			log.Errorf("%s is not a regular file", fileName)
			continue
		}

		if extension != "" && !Contains(extList, extension) {
			baseName += "." + extension
		}
		log.Info("baseName: ", baseName)

		newName := path.Join(targetDir, baseName)
		log.Info("newName: ", newName)

		err = os.Link(baseName, newName)
		if err != nil {
			err = os.Symlink(baseName, newName)
			if err != nil {
				wf, err := os.Create(newName)
				if err != nil {
					log.Fatal(err)
				}
				rf, err := os.Open(newName)
				if err != nil {
					log.Fatal(err)
				}
				w := io.Writer(wf)
				r := io.Reader(rf)
				_, err = io.Copy(w, r)
			}
		}

		log.Info("======================================")

	}
}
