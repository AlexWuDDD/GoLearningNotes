package main

import (
	"fmt"
	"os"
	"strings"
)

func main()  {
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Println(os.Args[1:])
	fmt.Println(os.Args[0]) //Exercise 1.1

	//Exercise1.2
	for index, value := range os.Args{
		fmt.Println(index, value)
	}


}

