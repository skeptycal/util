// Package file contains file utilities for basic scripting and shell operations
package file

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"strings"
)

// MyFile adds the current user's home directory to filename.
func MyFile(filename string) string {
    home, err := os.UserHomeDir()
    if err != nil {
        log.Fatal(err)
    }
    return path.Join(home, filename)
}

// GetFileUsingExec is an alternative to GetFile using exec.Command()
// along with cmd.Output() to gather file contents.
//
// In benchmarks, it is ~230 times slower than os.Open() with ioutil.ReadAll()
func GetFileUsingExec(filename string) []byte {
    cmd := exec.Command("cat", filename )
    b, err := cmd.Output()
    if err != nil {
        log.Fatal(err)
        return nil
    }
    return b
}

func GetFile(filename string) (string,error) {
    f, err := os.Open(filename)
    if err != nil {
        return "", err
    }
    defer f.Close()

    b, err := ioutil.ReadAll(f)
    if err != nil {
        return "", err
    }
    return string(b), nil
}

// FindOccurrence find the start and end of the first occurrence of
// sub in buf and returns the string positions.
func FindOccurrence(buf,sub string) (start, end int) {
    start = strings.Index(buf, sub)

    if start < 0 {
        return -1,-1
    }

    end = start + len(sub)
    return
}

// ChangeCharAfter finds the first occurrence of 'find' in 'content' and
// replaces one character with 'replace'
func ChangeCharAfter(content, find, replace string) (string, error){
    start, end := FindOccurrence(content, find)
    if start < 0 {
        return "", fmt.Errorf("option not found in config file: %v",string(find))
    }

    sb := strings.Builder{}
    defer sb.Reset()

    sb.WriteString(content[:end])
    sb.WriteString(replace)
    sb.WriteString(content[end+1:])
    return sb.String(), nil
}

// FileCopy copies the contents of src to dst using a buffered ReadWriter.
// Creates or truncates the dst. If dst already exists, it is truncated.
func FileCopy(src, dst string) (n int, err error) {
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


    return nil
}

// // Reference: https://stackoverflow.com/a/52684989
// func findAllOccurrences(data []byte, searches []string) map[string][]int {
//     results := make(map[string][]int)
//     for _, search := range searches {
//         searchData := data
//         term := []byte(search)
//         for x, d := bytes.Index(searchData, term), 0; x > -1; x, d = bytes.Index(searchData, term), d+x+1 {
//             results[search] = append(results[search], x+d)
//             searchData = searchData[x+1 : ]
//         }
//     }
//     return results
// }
