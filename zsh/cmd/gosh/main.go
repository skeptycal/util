package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {

    fmt.Println("gosh - the go shell")
    fmt.Println("\n(type 'exit' to exit ...)\n ")
    r := bufio.NewReader(os.Stdin)

    for {
        read, err := r.ReadString('\n')
        if err != nil {
            log.Fatal(err)
        }

        args := strings.Split(read, " ")

        exec.Command(args)

        n, err := fmt.Println(read)
        if err != nil {
            log.Fatal(err)
        }

        fmt.Printf(" (%d bytes written)\n>",n)

    }
}
