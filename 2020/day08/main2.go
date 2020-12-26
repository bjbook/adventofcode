package main

import "fmt"
import "os"
import "bufio"
import "io"
import "strconv"
import "strings"


type Rule struct {
   cmd string  /* nop acc jmp */
   num int
   again int
}

var accumulator int 
var rules [] Rule   //all rule line


// light red bags contain 1 bright white bag, 2 muted yellow bags.
func parse_line(str string) int{
    var rule Rule
    array := strings.Split(str, " ")
    rule.cmd = strings.TrimSpace(array[0])
    num, err := strconv.Atoi(strings.TrimSpace(array[1]))
    if(err!=nil){
        fmt.Println(err, array[1])
    }else{
        rule.num = num
    }
    fmt.Println("Key:",rule.cmd, "value:", rule.num )
    rules = append(rules, rule)
    return 0 
}


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
    for{
        str, err := reader.ReadString('\n')
        if err == io.EOF{
            break
        }
        if len(str)>2 {
            parse_line(str)
            //fmt.Println("")
        }
    }
    i := 0
    k := 0
    for k=len(rules)-1; k!=0; k--  {
        cmd := rules[k].cmd
        for i=0; i<len(rules); i++{
            rules[i].again = 0
        }
        accumulator=0
        if(rules[k].cmd=="jmp"){
            rules[k].cmd="nop"
        }else if rules[k].cmd=="nop"{
            rules[k].cmd="jmp"
        }else{
            continue
        }

        i=0
        for {
            fmt.Println(i, rules[i])
            if rules[i].again >= 1 {
               fmt.Println("accumulator:", accumulator)
               break
            }
            if rules[i].cmd=="nop" {
                rules[i].again += 1
                //i += rules[i].num
                i += 1
            }else if rules[i].cmd=="acc"{
                rules[i].again += 1
                accumulator += rules[i].num
                i += 1
            }else if rules[i].cmd=="jmp"{
                rules[i].again += 1
                i += rules[i].num
            }
            fmt.Println("i:",i, "again", rules[i].again)
            if i==(len(rules)-1) {
                break;
            }
        }
        rules[k].cmd = cmd
        if i==len(rules)-1{
            break
        }
    }
    fmt.Println("accumulator:", accumulator)
    return  0
}

func main(){

    count := check_questions()
    fmt.Println("count:", count)
}

