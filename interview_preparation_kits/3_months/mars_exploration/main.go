package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strings"
)

/*
 * Complete the 'marsExploration' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts STRING s as parameter.
 */

func marsExploration(s string) int32 {
    var tempSubIndex , sum int
    for _ , v := range s {
        if changed(string(v),tempSubIndex){
            sum += 1
        }
        if tempSubIndex == 2 {
            tempSubIndex = 0
        }else {
            tempSubIndex++
        }
    }
    return int32(sum)
}

func changed(s string   , index int ) bool {
    switch index {
        case 0:
        if s != "S" {
            return true 
        }
        case 1:
        if s != "O"{
            return true 
        }
        case 2:
        if s != "S"{
            return true 
        }
    }
    return false 
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    s := readLine(reader)

    result := marsExploration(s)

    fmt.Fprintf(writer, "%d\n", result)

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

