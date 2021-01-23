package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/skeptycal/util/gofile"
)

func main() {
    pwd, err := os.Getwd()
    if err != nil {
        log.Fatal(err)
    }
    err = filepath.Walk(pwd,
        func(path string, info os.FileInfo, err error) error {
            if err != nil {
                return err
            }
            fmt.Println(gofile.Base(path), info.Size())
            return nil
        })
    if err != nil {
        log.Println(err)
    }
}
