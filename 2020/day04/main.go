package main

import "fmt"
import "os"
import "bufio"
import "io"
import "strconv"
import "strings"

// ecl:gry pid:860033327 eyr:2020 hcl:#fffffd
// byr:1937 iyr:2017 cid:147 hgt:183cm

type Passport struct {
   ecl string
   pid string
   eyr string
   hcl string
   byr string
   iyr string
   cid string
   hgt string
}

func parse_passport(pass *Passport, arr [] string ) {
    if( "ecl" == arr[0]){
       pass.ecl=arr[1]
    }else if("pid" == arr[0]){
       pass.pid=arr[1]
    }else if("eyr" == arr[0]){
       pass.eyr=arr[1]
    }else if("hcl" == arr[0]){
       pass.hcl=arr[1]
    }else if("byr" == arr[0]){
       pass.byr=arr[1]
    }else if("iyr" == arr[0]){
       pass.iyr=arr[1]
    }else if("cid" == arr[0]){
       pass.cid=arr[1]
    }else if("hgt" == arr[0]){
       pass.hgt=arr[1]
    }
}

//byr (Birth Year) - four digits; at least 1920 and at most 2002.
func check_byr(str string) int {
    num, err := strconv.Atoi(str)
    if(err==nil){
        if num>=1920 && num<=2002 {
            return 1
        }
    }
    fmt.Println("err byr:", str)
    return 0
}

//iyr (Issue Year) - four digits; at least 2010 and at most 2020.
func check_iyr(str string) int {
    num, err := strconv.Atoi(str)
    if(err == nil){
        if num>=2010 && num<=2020 {
            return 1
        }
    }
    fmt.Println("err iyr:", str)
    return 0
}

// eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
func check_eyr(str string) int {
    num, err := strconv.Atoi(str)
    if(err == nil){
        if num>=2020 && num<=2030 {
            return 1
        }
    }
    fmt.Println("err eyr:", str)
    return 0
}

// hgt (Height) - a number followed by either cm or in:
// If cm, the number must be at least 150 and at most 193.
// If in, the number must be at least 59 and at most 76.
func check_hgt(str string) int {
    if strings.Contains(str, "cm") == true{
        num, err := strconv.Atoi(str[0:3])
        if(err == nil){
            if num>=150 && num<=193 {
                return 1
            }
        }
    }
    if strings.Contains(str, "in") == true{
        num, err := strconv.Atoi(str[0:2])
        if(err == nil){
            if num>=59 && num<=76 {
                return 1
            }
        }
    }
    fmt.Println("err hgt:", str)
    return 0
}

// hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
func check_hcl(str string) int {
    if len(str)==7 && str[0:1]== "#"{
        _, err := strconv.ParseInt(str[1:], 16, 0)
        if(err == nil){
            return 1
        }
    }
    fmt.Println("err hcl:", str)
    return 0
}

// ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
func check_ecl(str string) int {
    if str== "amb" || str =="blu" || str == "brn" || str=="gry" || str=="grn" || str =="hzl" || str == "oth" {
        return 1
    }
    fmt.Println("err ecl:", str)
    return 0;
}

// pid (Passport ID) - a nine-digit number, including leading zeroes.
func check_pid(str string) int {
    if(len(str)==9){
        num, err := strconv.Atoi(str[1:])
        if(err == nil){
            return 1
        }
        fmt.Println("err num", num)
    }
    return 0;
}

//cid (Country ID) - ignored, missing or not.

func verify_passport2(pass Passport) int {
    if check_ecl(pass.ecl)==1 && check_pid(pass.pid)==1 && check_eyr(pass.eyr)==1 && check_hcl(pass.hcl)==1 &&
        check_byr(pass.byr)==1 && check_iyr(pass.iyr)==1 && check_hgt(pass.hgt)==1 {
        return 1
    }
    return 0
}

func verify_passport(pass Passport) int {
    if len(pass.ecl)>=1 && len(pass.pid)>=1 && len(pass.eyr)>=1 && len(pass.hcl)>=1 &&
        len(pass.byr)>=1 && len(pass.iyr)>=1 && len(pass.hgt)>=1 {
        return 1
    }
    return 0
}

func check_passports() int{
    //filePath := "input0.txt"
    filePath := "input.txt"
    file, err := os.OpenFile(filePath, os.O_RDWR, 0666)
    if err != nil {
        fmt.Println("Open file error", err)
        return 0
    }
    defer file.Close()

    reader := bufio.NewReader(file)
    count := 0
    var pass Passport
    var pass2 Passport
    for{
        str, err := reader.ReadString('\n')
        if err == io.EOF{
            // calc valid
            if verify_passport2(pass)==1 {
                count ++
            }
            break
        }
        array := strings.Fields(str)
        for i:=0; i<len(array); i++{
            b := strings.Split(array[i], ":")
            parse_passport(&pass, b)
            //fmt.Println(b)
        }

        //fmt.Println("LEN:", len(str), "str:", str)
        if(len(str)<=2){
            //calc valid
            if verify_passport2(pass)==1 {
                count ++
            }
            fmt.Println(pass)
            pass = pass2
            fmt.Println("")
        }
    }
    return count
}

func main(){

    count := check_passports()
    fmt.Println("count:", count)
}

