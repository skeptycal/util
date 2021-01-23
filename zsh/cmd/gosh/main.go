package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

    r := bufio.NewReader(os.Stdin)
    w := bufio.NewWriter(os.Stdout)

    rw := bufio.NewReadWriter(r,w)

    for {
        read, err := rw.ReadString('\n')
        if err != nil {
            log.Fatal(err)
        }

        n, err := rw.WriteString(read)
        if err != nil {
            log.Fatal(err)
        }

        rw.WriteString(fmt.Sprintf(" (%d bytes written) "))
        rw.WriteString("\n")

    }
}
