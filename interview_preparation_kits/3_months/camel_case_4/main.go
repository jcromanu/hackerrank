package main
import ("fmt"
"bufio"
"os"
"strings"
 "unicode"
)

func main() {

 scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
    line := scanner.Text()
    strs := strings.Split(line,";")
    runnes := []rune(strs[2]) 
        if strs[0] == "S"{
            fmt.Println(split(runnes))
        }else {
        fmt.Println(combine(strs[1],runnes))
        }
    }
}


func combine(varType string , runnes []rune)string {
    switch varType {
        case "V":
        splited := strings.Split(string(runnes)," ")
        for i ,v := range splited {
            if i != 0{
                splited[i] = strings.Title(v)
            }
        }
        return strings.Join(splited,"")
        case "C":
         splited := strings.Split(string(runnes)," ")
        for i ,v := range splited {
                 splited[i] = strings.Title(v)
         }
        return strings.Join(splited,"")
        case "M":
         splited := strings.Split(string(runnes)," ")
        for i ,v := range splited {
            if i != 0{
                splited[i] = strings.Title(v)
            }
        }
        splited = append(splited,"()")
        return strings.Join(splited,"")
    }
    return ""
}

func split(runnes []rune)string{
      var splited []string
 markers  := make([]int,0)
 var start int 
           for i:=0 ; i < len(runnes);i++{
 
      if unicode.IsUpper(runnes[i]){
          markers = append(markers, i)
      }
       }
          for i := 0 ; i < len(markers);i++ {
          if markers[i] != 0{
          splited = append(splited,string(runnes[start:markers[i]]))
          start = markers[i]
          if i == len(markers)-1{
              splited = append(splited,string(runnes[markers[i]:]))
          }
          }
       }
       for i , v := range splited{
           splited[i] = strings.ToLower(v)
       }
     splited[len(splited)-1] = strings.Replace( splited[len(splited)-1] ,"()","",-1)
      
      return strings.Join(splited," ")
 }
 
 

//To be worked tomorrow

