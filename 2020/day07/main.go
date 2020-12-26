package main

import "fmt"
import "os"
import "bufio"
import "io"
//import "strconv"
import "strings"


type Rule struct {
   key string
   value [12]string
   count int
   ok int
}

var rules [] Rule   //all rule line


// light red bags contain 1 bright white bag, 2 muted yellow bags.
func parse_line(str string) int{
    var rule Rule
    array := strings.Split(str, "bags contain")
    //fmt.Println("len:", len(array))
    bags :=strings.Split(array[1], ",")
    //fmt.Println("bags:", bags)
    for i:=0; i<len(bags); i++{
        tmp := strings.Split(strings.TrimSpace(bags[i]), "bag")
        //fmt.Println("tmp:", tmp[0][2:])
        rule.value[i] = strings.TrimSpace(tmp[0][2:])
        rule.count=i+1
    }
    rule.key = strings.TrimSpace(array[0])
    rules = append(rules, rule)
    return 0 
}

var new_colors map[string]int

func check_rule(key string) int{
    ret := 0 
    for i:=0; i<len(rules); i++{
        for j:=0; j<rules[i].count; j++{
            if(rules[i].value[j]==key){
                rules[i].ok += 1
                new_colors[rules[i].key]=1
                ret = 1
                break
            }
        }
    }
    return ret
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
    new_colors = make(map[string]int)
    for i:=0; i<len(rules); i++{
        //fmt.Println(rules[i])
        for j:=0; j<rules[i].count; j++{
            if rules[i].value[j]=="shiny gold"{
                rules[i].ok += 1
                new_colors[rules[i].key]=1
                break
            }
        }
        //fmt.Println(rules[i])
    }

    for k:=0; k<20; k++{
        for key := range new_colors {
            check_rule(key)
        }
    }
    //fmt.Println("len:", len(new_colors), new_colors)
    count := 0
    for i:=0; i<len(rules); i++{
        if(rules[i].ok>=1){
            count += 1
        }
    }
    return  count
}

func main(){

    count := check_questions()
    fmt.Println("count:", count)
}

