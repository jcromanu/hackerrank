package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strings"
    "strconv"
)

/*
 * Complete the 'timeConversion' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
 */

func timeConversion(s string) string {
    var pm bool 
    str := strings.Split(s,":")
     meridianAM := strings.Split(str[2], "AM")
    meridianPM := strings.Split(str[2], "PM")
    //WIP need to separate this into [12 AM] and not simply [12 ] as a result of the string AM or PM found 
     if len(meridianAM) == 2{
        str[2]=meridianAM[0]
    }else {
        str[2]=meridianPM[0]
        pm = true
    }
    if pm {
        hourVal , _:= strconv.Atoi(str[0])
        if str[0] != "12"{
        str[0] = fmt.Sprint(hourVal+12)
        }
    }else if str[0] == "12" {
        str[0] = fmt.Sprint("00")
    }
    return fmt.Sprintf("%s:%s:%s",str[0],str[1],str[2])
}


func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    s := readLine(reader)

    result := timeConversion(s)

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

