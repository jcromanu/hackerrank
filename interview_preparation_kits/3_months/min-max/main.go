package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

/*
 * Complete the 'miniMaxSum' function below.
 *
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func miniMaxSum(arr []int32) {
    // Write your code here
    var currentMax , currentMin , sum   int64 
    //iterate from the beginning 
    for i , _ := range arr {
         for j , n := range arr {
             if j != i {
                sum += int64(n)
            }
        }
        if sum >= currentMax {
            currentMax = sum
        } 
        if i == 0{
            currentMin = currentMax
        }
        if sum <= currentMin {
            currentMin = sum
        }
         sum = 0 
    }
   //. negative numbers on the result wtf ???
    fmt.Printf("%d %d",currentMin,currentMax)
 }

 
func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

    var arr []int32

    for i := 0; i < 5; i++ {
        arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
        checkError(err)
        arrItem := int32(arrItemTemp)
        arr = append(arr, arrItem)
    }

    miniMaxSum(arr)
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

