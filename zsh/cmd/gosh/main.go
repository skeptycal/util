package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func prompt() string {
    return `
 ➜ `
}

func main() {

    fmt.Println("gosh - the go shell")
    fmt.Println("\n(type 'exit' to exit ...)\n ")
    r := bufio.NewReader(os.Stdin)

    for {
        read, err := r.ReadString('\n')
        if err != nil {
            log.Fatal(err)
        }

        read = strings.TrimSpace(read)

        args := strings.Split(read, " ")

        cmd := exec.Command(args[0],args[1:]...)

        buf, err := cmd.CombinedOutput()
        if err != nil {
            log.Fatal(err)
        }

        out := string(buf)

        n, err := fmt.Println(out)
        if err != nil {
            log.Fatal(err)
        }

        fmt.Printf(" (%d bytes written)\n%x",n,prompt())

    }
}
