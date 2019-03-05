package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

//Exercise4.3
func reverseArr(arr *[5]int) {
	for i, j := 0, len(*arr)-1; i < j; i, j = i+1, j-1{
		arr[i], arr[j] = arr[j], arr[i]
	}
}

//Exercise4.4
func rotata(slice []int)  {
	tmp := slice[len(slice)-1]
	copy(slice[1:], slice[:len(slice)-1])
	slice[0] = tmp
}

//Exercise4.5
func eliminateadjdup(strings []string) []string {
	tmp := strings[0]
	indexNum := []int{}
	for i:=1; i<len(strings); i++{
		if tmp == strings[i]{
			indexNum = append(indexNum, i)
		}else {
			tmp = strings[i]
		}
	}
	for i:= len(indexNum)-1; i >= 0; i--{
		v:= indexNum[i]
		fmt.Println(v)
		if v != len(strings)-1{
			copy(strings[v:], strings[v+1:])
		}
		strings = strings[:len(strings)-1]
	}
	return strings
}

func isExist(str string, strings []string) bool{
	fmt.Println(str, strings)
	for _, item := range strings{
		if item == str{
			return true;
		}
	}
	return false
}

//Exercise 4.6
func SquashUtf8Space2AsciiSpace(str string) {
	for _, r := range str {
		fmt.Printf("%c = %v\n", r, unicode.IsSpace(r))
		if unicode.IsSpace(r) {
			r = ' '
		}
	}

}

//Exercise 4.7
func Exercise4_7(str string){
	fmt.Println(strings.Split(strconv.QuoteToASCII(str), "\\u"))
	fmt.Println([]byte(str))
}



func main(){
	arr1 := [5]int{1,2,3,4,5}
	reverseArr(&arr1)
	fmt.Println(arr1)

	s1 := []int{1,2,3,4,5}
	rotata(s1)
	fmt.Println(s1)

	strings := []string{"alex", "is", "is", "cool", "is", "cool", "cool"}
	strings = eliminateadjdup(strings)
	fmt.Println(strings)

	message := "吴康俊 是 一个好人\t\n"
	SquashUtf8Space2AsciiSpace(message)

	message2 := "你好世界！"
	Exercise4_7(message2)

	translate();
}

//golang实现unicode码和中文之间的转换
func translate(){
	sText := "中文"
	textQuoted := strconv.QuoteToASCII(sText)
	textUnquoted := textQuoted[1 : len(textQuoted)-1]
	fmt.Println(textUnquoted)
	sUnicodev := strings.Split(textUnquoted, "\\u")
	var context string
	for _, v := range sUnicodev {
	    if len(v) < 1 {
	        continue
	    }
	    temp, err := strconv.ParseInt(v, 16, 32)
	    if err != nil {
	        panic(err)
	    }
	    context += fmt.Sprintf("%c", temp)
	}
	fmt.Println(context)
}




