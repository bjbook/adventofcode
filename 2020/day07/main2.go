package main

import "fmt"
import "os"
import "bufio"
import "io"
import "strconv"
import "strings"

type Num struct {
   key string
   num int
}

type Rule struct {
   key string
   value [12] Num
   count int
   ok int
}

var rules [] Rule   //all rule line


// light red bags contain 1 bright white bag, 2 muted yellow bags.
func parse_line(str string) int{
    var rule Rule
    array := strings.Split(str, "bags contain")
    //fmt.Println("len:", len(array))
    fmt.Println("------")
    rule.key = strings.TrimSpace(array[0])
    index := strings.Index(array[1], "no other")
    if(index != -1){
        rule.count = 0
        rules = append(rules, rule)
        fmt.Println("rule count:", rule.count)
        return 0
    }

    bags :=strings.Split(array[1], ",")
    //fmt.Println("bags:", bags)
    for i:=0; i<len(bags); i++{
        tmp := strings.Split(strings.TrimSpace(bags[i]), "bag")
        //fmt.Println("tmp:", tmp[0][2:])
        rule.value[i].key = strings.TrimSpace(tmp[0][2:])
        rule.value[i].num,_ = strconv.Atoi(strings.TrimSpace(tmp[0][0:2]))
        fmt.Println("rule count:", rule.value[i].key, rule.value[i].num)
        rule.count=i+1
    }
    rules = append(rules, rule)
    return 0 
}

var new_colors map[string]int

func check_rule(key string) int{
    ret := 0 
    for i:=0; i<len(rules); i++{
        for j:=0; j<rules[i].count; j++{
            if(rules[i].value[j].key==key){
                rules[i].ok += 1
                new_colors[rules[i].key]=1
                ret = 1
                break
            }
        }
    }
    return ret
}

func calc_count(key string) int{
    count :=0
    for i:=0; i<len(rules); i++{
        if rules[i].key==key{
            for j:=0; j<rules[i].count; j++{
                sub := rules[i].value[j].num
                count += sub
                count += sub*calc_count(rules[i].value[j].key)
                //fmt.Println("count:", count)
            }
            break
        }
    }
    fmt.Println("key:", key, "count:", count)
    return count
}

func check_questions() int{
    //filePath := "input00.txt"
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
    count := calc_count("shiny gold")
    return  count
}

func main(){

    count := check_questions()
    fmt.Println("count:", count)
}

