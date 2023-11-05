package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strings"
    "unicode"
)

/*
 * Complete the 'pangrams' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
 */

var alphabet = make(map[int]int,26)
var index int 

// ascii 98 - 122 lowercase alphabet 
func pangrams(s string) string {
    for _ ,v := range s {
        index = int(unicode.ToLower(v))
        alphabet[index] += 1
    }
    for  i:= 98 ; i <=122 ; i++ {
        if alphabet[i] == 0 {
            return "not pangram"
        }
    }
    return "pangram"
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    s := readLine(reader)

    result := pangrams(s)

    fmt.Fprintf(writer, "%s\n", result)

    writer.Flush()
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}

