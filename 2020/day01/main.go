package main

import "fmt"
import "os"
import "bufio"
import "io"
import "strconv"
import "strings"

//进行遍历即可
func twosum(nums [] int, target int) int {
    i := 0
    j := 1
    for i = 0; i<len(nums)-1; i++{
        for j=i+1; j<len(nums); j++{
            if((nums[i]+nums[j])==target){
                return nums[i]*nums[j] 
            }
        }
    }
    return target
}

func threesum(nums [] int, target int) int {
    i := 0
    j := 1
    k := 2
    for i = 0; i<len(nums)-2; i++{
        for j=i+1; j<len(nums)-1; j++{
            for k=j+1; k<len(nums); k++{
                if((nums[i]+nums[j]+nums[k])==target){
                    return nums[i]*nums[j]*nums[k]
                }
            }
        }
    }
    return target
}

func main(){
    filePath := "input.txt"
    file, err := os.OpenFile(filePath, os.O_RDWR, 0666)
    if err != nil {
        fmt.Println("Open file error", err)
        return
    }
    defer file.Close()

    var array []int
    reader := bufio.NewReader(file)
    for{
        str, err := reader.ReadString('\n')
        if err == io.EOF{
            break
        }
        str2 := strings.Trim(str, "\n")
        a, err := strconv.Atoi(str2)
        array = append(array, a)
        //fmt.Print(str)
    }
    //fmt.Print(array)
    r2 := twosum(array, 2020)
    fmt.Println("r2=(twosum)", r2)
    r3 := threesum(array, 2020)
    fmt.Println("r3=(threesum)", r3)

}

