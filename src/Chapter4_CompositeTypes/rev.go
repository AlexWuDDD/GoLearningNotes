package main

import "fmt"

func main()  {
	a := [...]int{0,1,2,3,4,5,6,7,8}
	reverse(a[:])
	fmt.Println(a)

	s := []int{0,1,2,3,4,5}
	reverse(s[:2]) //1,0,2,3,4,5
	reverse(s[2:])//105432
	reverse(s)
	fmt.Println(s)
}

func reverse(s []int)  {
	for i,j := 0, len(s)-1; i < j; i, j = i+1, j-1{
		s[i],s[j] = s[j],s[i]
	}
}

func equal(x, y []string) bool{
	if len(x) != len(y){
		return false
	}
	for i := range x{
		if x[i] != y[i]{
			return false
		}
	}
	return true
}
