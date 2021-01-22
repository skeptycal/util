package main

import (
	"bufio"
	"bytes"
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
    modeFlag := flag.Int("mode",0,"mode for dev output (0-3)")
    helpFlag := flag.Bool("help",false,helpText)

    flag.Parse()

    mode := *modeFlag

    if mode < 0 || mode > 2 {
        flag.Usage()
        os.Exit(1)
    }

    fmt.Println(`devmode - change the shell dev boot mode

    mode = 0 is production mode

    mode is set to non-zero for dev mode
        1 - Show debug info and log to $LOGFILE
        2 - #1 plus trace and run specific tests
        3 - #2 plus display and log everything
    `)

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

    // fmt.Printf("Option text starts at: %v\n", i)

    optIndex := i + len(find) - 1

    // fmt.Printf("Option value starts at: %v\n", optIndex)

    optString := contents[i:i+len(find)]
    fmt.Printf("optString: %v\n", optString)

    optValue := contents[optIndex] - 48
    fmt.Printf("optValue: %v\n", optValue)

    optNew := mode

    newContents := []byte(fmt.Sprintf("%v%v%v",contents[:optIndex],optNew,contents[optIndex+1:]))

    // fmt.Printf("contents: \n%s\n",newContents)


    err := ioutil.WriteFile(bakfile,newContents,0644)
    if err != nil {
        log.Fatal(err)
    }

    err = filecopy(bakfile,"test.bak")
    if err != nil {
        log.Fatal(err)
    }

}

func filecopy(src, dst string) (err error) {
    fi, err := os.Stat(src)
    if err != nil {
        return err
    }

    BUFFERSIZE := int(((fi.Size() / bytes.MinRead) + 1) * bytes.MinRead)

    source, err := os.Open(src)
    if err != nil {
        return err
    }

    destination, err := os.Create(dst)
    if err != nil {
        return err
    }

    // buf := make([]byte, BUFFERSIZE)


    rw := bufio.NewReadWriter(bufio.NewReaderSize(source, BUFFERSIZE), bufio.NewWriterSize(destination, BUFFERSIZE))

    n, err :=rw.WriteTo(rw)
    if err != nil {
        return err
    }

    fmt.Printf("%v bytes written",n)


    // for {
    //     n, err := source.Read(buf)
    //     if err != nil && err != io.EOF {
    //             return err
    //     }
    //     if n == 0 {
    //             break
    //     }

    //     if _, err := destination.Write(buf[:n]); err != nil {
    //             return err
    //     }
    // }
        // buf = nil
    return nil
}


const (
    helpText string = `devmode - change the shell dev boot mode

    mode = 0 is production mode

    mode is set to non-zero for dev mode
        1 - Show debug info and log to $LOGFILE
        2 - #1 plus trace and run specific tests
        3 - #2 plus display and log everything
    `
)
