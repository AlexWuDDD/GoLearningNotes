package main

import "fmt"

func main(){
	s := "hello world"
	fmt.Println(len(s))
	fmt.Println(s[0], s[7])
	fmt.Printf("%c %c", s[0], s[7])
}
