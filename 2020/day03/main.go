package main

import "fmt"
import "os"
import "bufio"
import "io"
//import "strconv"
//import "strings"


var LEN int


func calc_line(pos int, str string) int {

    if pos>=LEN {
        pos = pos % LEN
    }
    if str[pos]=='#' {
        return 1
    }
    return 0;
}

func right_down(right int, down int) int{
    filePath := "input.txt"
    file, err := os.OpenFile(filePath, os.O_RDWR, 0666)
    if err != nil {
        fmt.Println("Open file error", err)
        return 0
    }
    defer file.Close()

    var count int
    pos := 0
    line_index := 0;
    reader := bufio.NewReader(file)
    //str, err := reader.ReadString('\n')
    for{
        str, err := reader.ReadString('\n')
        if err == io.EOF{
            break
        }

        LEN = len(str)-1
        a := calc_line(pos, str)
        count += a
        pos += right
        line_index++;
        //fmt.Println("LEN:",LEN,"line:", line_index, " pos:", pos)
        //str, err := reader.ReadString('\n')
        if(down==2){
            str, err = reader.ReadString('\n')
            if err == io.EOF{
                break
            }
        }
    }
    fmt.Println("right:", right, "down:", down, "count=", count)
    return count
}

func main(){

    a := right_down(1, 1);
    b := right_down(3, 1);
    c := right_down(5, 1);
    d := right_down(7, 1);
    e := right_down(1, 2);
    fmt.Println("all:", a*b*c*d*e)
}

