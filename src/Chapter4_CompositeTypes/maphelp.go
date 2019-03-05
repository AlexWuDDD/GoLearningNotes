package main

import "fmt"

var m = make(map[string]int)

func k(list []string) string {return fmt.Sprintf("%q", list)}
func Add(list []string)  {
	m[k(list)]++
}
func Count(list []string) int { return m[k(list)]}

func main()  {
	mylist := []string{"1","2","3","4","5","6"}
	fmt.Println(k(mylist))
}