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

func getFileUsingExec(filename string) []byte {
    cmd := exec.Command("cat", filename )
    b, err := cmd.Output()
    if err != nil {
        log.Fatal(err)
        return nil
    }
    return b
}

func getFile(filename string)[]byte {
    f, err := os.Open(filename)
    if err != nil {
        return nil
    }
    b, err := ioutil.ReadAll(f)
    if err != nil {
        log.Fatal(err)
    }
    return b
}

func main() {

    filename := myfile( configFile)
    bakfile :=filename + ".bak"

    modeFlag := flag.Int("mode",-1,"mode for dev output (0-3)")
    helpFlag := flag.Bool("help",false,helpText)

    flag.Parse()

    mode := *modeFlag
    help := *helpFlag

    if help {
        fmt.Printf(helpText,filename)
        os.Exit(0)
    }

    contents := getFile(filename)

    // parse mode 0 .. 3
    if mode > -1 && mode < 4 {
        err := changeDevMode(mode)
        if err != nil {
            log.Fatal(err)
        }
    }

}

// Reference: https://stackoverflow.com/a/52684989
func findAllOccurrences(data []byte, searches []string) map[string][]int {
    results := make(map[string][]int)
    for _, search := range searches {
        searchData := data
        term := []byte(search)
        for x, d := bytes.Index(searchData, term), 0; x > -1; x, d = bytes.Index(searchData, term), d+x+1 {
            results[search] = append(results[search], x+d)
            searchData = searchData[x+1 : len(searchData)]
        }
    }
    return results
}

func findOccurrence(buf []byte, sub []byte) (start, end int) {
    start := bytes.Index(buf, sub)
    if start < 0 {
        return -1,-1
    }

    end = start + len(sub) - 1
    return
}

func changeDevMode(contents []byte, mode int) error {

    find := []byte("declare -ix SET_DEBUG=0")

    start, end := findOccurrence(contents, find)

    // i :=     strings.Index(contents, find)

    if start < 0 {
        fmt.Printf("option not found in config file: %v\n",string(find))
        os.Exit(1)
    }

    contents[end+1] = []byte(mode)

    newContents := []byte(fmt.Sprintf("%v%v%v",contents[:optIndex],mode,contents[optIndex+1:]))

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
    helpText string = `DEVMODE(1)                       User Commands                      DEVMODE(1)



NAME
       devmode - set login script debug options

SYNOPSIS
       devmode [OPTION]...

DESCRIPTION
       Apply changes to login script options. The main option that enables the
       dev mode is the 'mode' option. This option is stored in the devmode
       config script:

           %s

       mode = 0 is production mode

       mode is set to non-zero for dev mode
          1 - Show debug info and log to $LOGFILE
          2 - #1 plus trace and run specific tests
          3 - #2 plus display and log everything

       Mandatory  arguments  to  long  options are mandatory for short options
       too.

       -m <mode>
              set debug mode in login script

       -h, -help
              show this help message

       -l, -list
              list all available options

`
)
