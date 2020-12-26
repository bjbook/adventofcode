package main

import "fmt"
import "os"
import "bufio"
import "io"
import "math"
//import "strconv"
//import "strings"


func calc_line(str string) int {
    row := 0
    column := 0
    if(len(str)>=10){
        for i:=0; i<7; i++ {
            if(str[i]=='B'){
                row += int(math.Pow(2, float64(6-i)))
                //fmt.Println("i:",i, "row:", row)
            }
        }
        for j:=7; j<10; j++{
            if(str[j]=='R'){
                column += int(math.Pow(2, float64(9-j)))
            }

        }
        //fmt.Println("str:",str, "row:", row, "column:", column)
    } 
    return row*8+column;
}

func main(){
    filePath := "input.txt"
    file, err := os.OpenFile(filePath, os.O_RDWR, 0666)
    if err != nil {
        fmt.Println("Open file error", err)
        return
    }
    defer file.Close()

    max := 0
    var array [1024] int
    for i := 0; i < 1024; i++ {
      array[i] = -1
    }
    reader := bufio.NewReader(file)
    for{
        str, err := reader.ReadString('\n')
        if err == io.EOF{
            break
        }
        a := calc_line(str)
        //fmt.Println("a:", a)
        if(a<1024){
            array[a]=a
        }
        if(a> max) {
            max = a
        }
        //fmt.Println(a)
    }
    fmt.Println("max=", max)
    your := 0
    for i := 0; i < 1024-2; i++ {
         if array[i] == i && array[i+1]==-1 && array[i+2]== i+2 {
             your = i+1
         }
    }
    fmt.Println("your=", your)
}

