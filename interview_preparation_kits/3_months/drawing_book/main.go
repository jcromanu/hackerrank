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
 * Complete the 'pageCount' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER n
 *  2. INTEGER p
 */

func pageCount(n int32, p int32) int32 {
    // Write your code here
    fw := countForward(n,p)
    bw := countBackward(n,p)
    if fw >= bw {
        return bw 
    }
    return fw
}

func countBackward(n,p int32 ) int32 {
      var  pageCounter , currentPage int32 
      var emptyPage int
      // caso edge , la ultima pagina tiene la pagina izquierda vacia 
      // 01 , 23 , 45 , 60 
      newP := int(n)
    // si el numero total de paginas es par siempre se agrega una pagina adicional
      if n%2 == 0{
          emptyPage++
      }
      // vamos desde la ultima pagina hasta el inicio del libro 
      for i:=newP+emptyPage ; i >= 0 ;i-- {
          // si estamos en el final del libro donde estan las ultimas dos paginas no incrementamos el contador 
          if n%2 == 0 {
            if i!=newP && i!= newP+1 {
                currentPage ++ 
            }
          }else {
              if i != newP && i != newP-1  {
                  currentPage++
              }
          }
        if currentPage == 1 {
            pageCounter++
        }
        if i == int(p){
            return pageCounter
        }
        if currentPage == 2 {
            currentPage = 0 
        }
     }
     return pageCounter
}

func countForward(n,p int32) int32 {
      var  pageCounter , currentPage int32 = 0 , 0  
      for i:=0 ; i <= int(p) ;i++ {
         if i!=0 && i != 1{
            currentPage ++ 
         }
        if currentPage == 1 {
            pageCounter++
        }
        if i == int(p){
            return pageCounter
        }
        if currentPage == 2 {
            currentPage = 0 
        }
     }
     return pageCounter
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    n := int32(nTemp)

    pTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    p := int32(pTemp)

    result := pageCount(n, p)

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

