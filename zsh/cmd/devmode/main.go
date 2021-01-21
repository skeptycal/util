package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
)

const (

    configFile = `.dotfiles/zshrc_inc/dev_mode.zsh`
)


func myfile(filename string) string {
    home, err := os.UserHomeDir()
    if err != nil {
        log.Fatal(err)
    }
    return path.Join(home, filename)
}

func getfile(filename string) string {
    cmd := exec.Command("cat", filename )
    b, err := cmd.Output()
    if err != nil {
        log.Fatal(err)
        return ""
    }
    return string(b)
}

func getFile(filename string)string {
    f, err := os.Open(filename)
    if err != nil {
        return ""
    }
    b, err := ioutil.ReadAll(f)
    if err != nil {
        log.Fatal(err)
    }
    return string(b)
}

func main() {
    fmt.Println("configFile: ", configFile)

    filename := myfile( configFile)
    fmt.Println("filename: ", filename)

    contents := getfile(filename)



    fmt.Printf("contents: \n%s\n",contents)

}
