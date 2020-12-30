package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func DirWalk(path string) (err error) {
	var files []string

	root, err := os.Getwd()
	if err != nil {
		root = "."
	}

	err = filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	if err != nil {
		return (err)
	}
	for _, file := range files {
		fmt.Println(filepath.Base(file))
	}
	return nil
}

func main() {

	var p percent = 0.0

	for i := 0; i < 15; i++ {
		p += percent(float64(i) * 1.1)
		fmt.Println(string(p))
		fmt.Println(p.Decimal())
	}

}
