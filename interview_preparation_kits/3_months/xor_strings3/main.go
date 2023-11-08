package main
import (
"fmt" 
"os"
"bufio"
"strconv"
"strings"
)


func main() {
 //Enter your code here. Read input from STDIN. Print output to STDOUT
 
 scanner := bufio.NewScanner(os.Stdin)
 var lines = make([]string,2)
 var i int 
 for scanner.Scan() {
     lines[i] =scanner.Text()
     i++
 }
 fmt.Printf("%s" ,xorFor(lines[0],lines[1]))
}

func xorFor(s1 , s2 string ) string {
    var res = make([]string,len(s1))
    for i:=0 ; i< len(s1);i++ {
        a , _ := strconv.Atoi(string(s1[i]))
        b , _ := strconv.Atoi(string(s2[i]))
            if a != b {
                res[i] = "1"
            }else {
                res[i] = "0"
            }
    }
    return strings.Join(res,"")
}

