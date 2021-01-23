package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/skeptycal/util/zsh/file"
)

const (
    configFile = `.dotfiles/zshrc_inc/dev_mode.zsh`
)

func main() {

    filename := file.MyFile( configFile)

    modeFlag := flag.Int("mode",-1,"mode for dev output (0-3)")
    helpFlag := flag.Bool("help",false,helpText)

    flag.Parse()

    mode := *modeFlag
    help := *helpFlag

    if help {
        fmt.Printf(helpText,filename)
        os.Exit(0)
    }


    // parse mode 0 .. 3
    if mode > -1 && mode < 4 {
        err := changeDevMode(filename, mode)
        if err != nil {
            log.Fatal(err)
        }
    }

}

// changeDevMode changes the debug 'DEV' mode in the .zshrc utility file
// named in the constant configFile to 'mode'
//
// values of mode may be 0 - 3
func changeDevMode(filename string, mode int) error {

    // get file
    contents, err:= file.GetFile(filename)
    if err != nil {
        return fmt.Errorf("error reading file %v: %v", filename, err)
    }

    // change contents
    find := "declare -ix SET_DEBUG="
    contents, err = file.ChangeCharAfter(contents,find,fmt.Sprintf("%d",mode))
    if err != nil {
        return fmt.Errorf("error changing contents of file %v: %v", filename, err)
    }

    // make a backup copy
    n, err := file.FileCopy(filename, filename+".bak")
    if err != nil {
        return fmt.Errorf("error copying backup file %v: %v", filename + ".bak", err)
    }

    // write new file
    err = ioutil.WriteFile(filename,[]byte(contents),0644)
    if err != nil {
        return fmt.Errorf("error writing file %v: %v", filename, err)
    }

    fmt.Printf("%v bytes written\n",n)

    return nil
}




const (
    helpText string = `
DEVMODE(1)                       User Commands                      DEVMODE(1)



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
