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
 * Complete the 'hourglassSum' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts 2D_INTEGER_ARRAY arr as parameter.
 */

func hourglassSum(arr [][]int32) int32 {
    // calculate number of hourglass depending on size of arrays
    var height , width , hourGlassIndex int
    //height := len(arr)
    //width := len(arr[0])
    //Calculate horizontal hourglasses available
    if len(arr[0])-2 > 0 {
        width = len(arr[0])-2
    }
    if len(arr)-2 > 0{
        height = len(arr)-2
    }
    numHourglass := width * height
        if numHourglass < 1 {
        return 0
    }
    hourGlasses := make([]Hourglass,numHourglass)
    
    //Travel the arrays
    for rowIndex:=0 ; rowIndex < len(arr)-2;rowIndex++ {
        for colIndex:=0 ; colIndex < len(arr[0])-2;colIndex++  {
            //get top three digits of hourglass
            hourGlasses[hourGlassIndex].values[0] = arr[rowIndex][colIndex]
            hourGlasses[hourGlassIndex].values[1] = arr[rowIndex][colIndex+1]
            hourGlasses[hourGlassIndex].values[2] = arr[rowIndex][colIndex+2]
            //get middle digit of hourglas
            hourGlasses[hourGlassIndex].values[3] = arr[rowIndex+1][colIndex+1]
             //get last digits of hourglass
            hourGlasses[hourGlassIndex].values[4] = arr[rowIndex+2][colIndex]
            hourGlasses[hourGlassIndex].values[5] = arr[rowIndex+2][colIndex+1]
            hourGlasses[hourGlassIndex].values[6] = arr[rowIndex+2][colIndex+2]
            hourGlassIndex++
        }
    }
    fmt.Printf("%v\n",hourGlasses)
    sums := calculateHourGlassSums(hourGlasses)
    
return  int32(calculateGreatest(sums))
}


//calculate sum 
func calculateHourGlassSums(hourGlasses []Hourglass) []int32{ 
        hourglassSum := make([]int32,len(hourGlasses))
        //Iterate through hourglasses
        for i:=0 ; i < len(hourGlasses) ; i++{
        var sum int32 = 0 
            for j:= 0; j < len(hourGlasses[i].values);j++{
                sum += hourGlasses[i].values[j]
            }
            fmt.Println(sum)
            hourglassSum[i] = sum
         }
         return hourglassSum
}

func calculateGreatest(hourGlassSums []int32) int32 {
    var top int32 
     for i , v := range hourGlassSums {
         if i == 0{
             top = v 
         } 
        if top < v {
            top = v 
        }
    }
    return top
}

type Hourglass struct {
    values [7]int32 
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    var arr [][]int32
    for i := 0; i < 6; i++ {
        arrRowTemp := strings.Split(strings.TrimRight(readLine(reader)," \t\r\n"), " ")

        var arrRow []int32
        for _, arrRowItem := range arrRowTemp {
            arrItemTemp, err := strconv.ParseInt(arrRowItem, 10, 64)
            checkError(err)
            arrItem := int32(arrItemTemp)
            arrRow = append(arrRow, arrItem)
        }

        if len(arrRow) != 6 {
            panic("Bad input")
        }

        arr = append(arr, arrRow)
    }

    result := hourglassSum(arr)

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

