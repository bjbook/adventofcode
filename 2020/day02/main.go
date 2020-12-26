package main

import "fmt"
import "os"
import "bufio"
import "io"
import "strconv"
import "strings"


func verify_pass(min int, max int, c string, str string) bool {
    i := 0
    count := 0
    for i=0; i<len(str); i++{
        if(c[0]==str[i]){
           count++
        }
    }
    if(count>=min && count<=max){
        return true
    }
    return false
}

//two positions in the password
func verify_pass2(min int, max int, c string, str string) bool {

    if c[0]==str[min-1] && c[0] != str[max-1] {
          return true
    }
    if c[0]!=str[min-1] && c[0] == str[max-1] {
        return true
    }
    return false
}

func verify_line(str string) int {
    array := strings.Fields(str)
    //fmt.Println(array[1][0:1],len(array))
    one := strings.Split(array[0], "-")
    //fmt.Println(one)
    min, err := strconv.Atoi(one[0])
    if(err!=nil){
        return 1
    }
    max, err := strconv.Atoi(one[1])
    if(err!=nil){
        return 1
    }
    ret := verify_pass2(min, max, array[1][0:1], array[2])
    if ret==true {
        return 1
    }
    return 0;
}

func main(){
    filePath := "input.txt"
    file, err := os.OpenFile(filePath, os.O_RDWR, 0666)
    if err != nil {
        fmt.Println("Open file error", err)
        return
    }
    defer file.Close()

    var count int
    reader := bufio.NewReader(file)
    for{
        str, err := reader.ReadString('\n')
        if err == io.EOF{
            break
        }
        a := verify_line(str)
        count += a
        //fmt.Print(str)
    }
    fmt.Println("count=", count)
}

