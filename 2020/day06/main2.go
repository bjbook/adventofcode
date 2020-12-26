package main

import "fmt"
import "os"
import "bufio"
import "io"
//import "strconv"
//import "strings"


func check_questions() int{
    //filePath := "input0.txt"
    filePath := "input.txt"
    file, err := os.OpenFile(filePath, os.O_RDWR, 0666)
    if err != nil {
        fmt.Println("Open file error", err)
        return 0
    }
    defer file.Close()

    reader := bufio.NewReader(file)

    people := 0 
    var arr [26] int 
    for i:=0; i<len(arr); i++{
        arr[i]=0
    }

    count := 0
    for{
        str, err := reader.ReadString('\n')
        if err == io.EOF{
            if people<1 {
                break
            }
            c := 0
            for i:=0; i<len(arr); i++{
                if arr[i]==people {
                    c++
                }
            }
            //fmt.Println("c:", c, "count:", count)
            count += c
            break
        }
        for i:=0; i<len(str)-1; i++{
            arr[str[i]-'a'] += 1
        }

        //fmt.Println("LEN:", len(str), "str:", str)
        if(str=="\n"){
            c := 0 
            for i:=0; i<len(arr); i++{
                if arr[i]==people {
                    c++
                }
            }
            //fmt.Println("c:", c, "count:", count, "people:", people)
            count += c
            for i:=0; i<len(arr); i++{
                arr[i]=0
            }
            people = 0
        }else
        {
            people += 1
        }
    }
    return count
}

func main(){

    count := check_questions()
    fmt.Println("count:", count)
}

