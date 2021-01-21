package main

import (
	"log"
	"os"
	"os/exec"
	"path"
)

const (

    configFile = `.dotfiles/zshrc_inc/dev_mode.zsh`
)

func myfile(path string) (string, error) {

}

func getfile(path string) string {
    cmd := exec.Command("echo", path )
    b, err := cmd.Output()
    if err != nil {
        log.Fatal(err)
        return ""
    }
    return string(b)
}

func main() {
    home, _ := os.UserHomeDir()
    filename := path.Join(home, configFile)




    println(s)
}
