package main

import (
	"fmt"
	"math/rand"
)

func main()  {

	//rand.Intn [0,n)
	for i:=0;i<10;i++{
		rnd := rand.Intn(4)
		fmt.Println(rnd)
	}
}
