package main

import "fmt"
import "os"
import "bufio"
import "io"
import "strconv"
import "strings"

var rules [] int   //all rule line


func parse_line(str string) int{
    num, err := strconv.Atoi(strings.TrimSpace(str))
    if(err==nil){
        fmt.Println("num:", num )
        rules = append(rules, num)
    }
    return 0 
}

func check_rule(index int, nums int)int{
    j :=index-nums
    for j=index-nums; j<index-1; j++{
        for k:=j+1; k<index; k++{
            if(rules[j] + rules[k])==rules[index]{
                return 0
            }
        }
    }
    return 1  //found
}

func find_max_min(array[]int) (int, int){
    min := array[0]
	max := array[0]
	for i:=0; i<len(array); i++{
	   if(array[i]>max){
	       max = array[i]
	   }
	   if(array[i]<min){
	       min = array[i]
	   }
	}
	fmt.Println(array, max, min)
	return max, min
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
        if len(str)>=1 {
            parse_line(str)
        }
    }
    i := 0
    //numbers := 5
    numbers := 25
    acc := 0
    i = numbers
    for{
        if i>=len(rules) {
            break
        }
        fmt.Println(i, rules[i])
        ret := check_rule(i, numbers)
        if ret==1{
            acc=rules[i]
            break
        }
        i++
    }

    fmt.Println("index:", i)
	fmt.Println("acc:", acc)

	acc_index :=i 
	for i=0; i<acc_index; i++{
	    tmp := 0
	    for j:=i; j<acc_index; j++{
		    tmp += rules[j]
			if(tmp==acc){
			    max, min := find_max_min(rules[i:j])
			    fmt.Println("max+min:", max+min)
			}else if(tmp>acc){
			    break
			}
		}
	}
    return  acc
}

func main(){

    count := check_questions()
    fmt.Println("count:", count)
}

