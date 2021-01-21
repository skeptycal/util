package main

import (
	"flag"
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

func getFileUsingExec(filename string) string {
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
    modeFlag := flag.Int("mode",5,"mode for dev output (0-2)")

    flag.Parse()

    mode := *modeFlag

    if mode < 0 || mode > 2 {
        flag.Usage()
        os.Exit(1)
    }

    fmt.Println("devmode - change shell dev boot mode")
    fmt.Println("configFile: ", configFile)

    filename := myfile( configFile)

    contents := getFile(filename)



    // fmt.Printf("contents: \n%s\n",contents)

}
