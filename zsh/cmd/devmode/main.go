package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"strings"
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
    modeFlag := flag.Int("mode",0,"mode for dev output (0-2)")

    flag.Parse()

    mode := *modeFlag

    if mode < 0 || mode > 2 {
        flag.Usage()
        os.Exit(1)
    }

    fmt.Println("devmode - change shell dev boot mode")
    fmt.Println("configFile: ", configFile)

    filename := myfile( configFile)
    bakfile :=filename + ".bak"

    contents := getFile(filename)

    find := "declare -ix SET_DEBUG=0"

    i :=     strings.Index(contents, find)

    if i < 0 {
        fmt.Printf("option not found in config file: %v\n",find)
        os.Exit(1)
    }

    fmt.Printf("Option text starts at: %v\n", i)

    optIndex := i + len(find) - 1

    fmt.Printf("Option value starts at: %v\n", optIndex)

    optString := contents[i:i+len(find)]
    fmt.Printf("optString: %v\n", optString)

    optValue := contents[optIndex] - 48
    fmt.Printf("optValue: %v\n", optValue)

    optNew := mode

    newContents := []byte(fmt.Sprintf("%v%v%v",contents[:optIndex],optNew,contents[optIndex+1:]))

    fmt.Printf("contents: \n%s\n",newContents)


    err := ioutil.WriteFile(bakfile,newContents,0644)
    if err != nil {
        log.Fatal(err)
    }

}
