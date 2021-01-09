package main

import (
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"strings"

	log "github.com/sirupsen/logrus"
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
	return FileName(os.Args[0])
}

// here returns the location of this file
func here() string {
	return AbsPath(os.Args[0])
}

// Name returns the name of file else  "."
func FileName(file string) string {
	return filepath.Base(file)
}

// AbsPath returns the absolute path of file
func AbsPath(file string) string {
	dir, _ := filepath.Abs(file)
	return dir
}

// OsArgs returns the slice of strings returned by os.Args[1:]
func OsArgs() []string {
	return os.Args[1:]
}

// Contains tells whether a contains x.
func Contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

// NameSplit returns a slice of strings containing the name and extension of file.
func NameSplit(file string) (string, string) {
	s := strings.Split(filepath.Base(file), ".")
	name := s[0]
	ext := ""
	if len(s) > 1 {
		ext = s[1]
	}
	return name, ext
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

		fileName := AbsPath(sourceFileStat.Name())
		baseName, extension := NameSplit(fileName)
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
