package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main()  {
	var s1, seq string

	start1 := time.Now()
	fmt.Println(start1.UnixNano())
	for i :=1; i < len(os.Args); i++{
		s1 += seq + os.Args[i]
		seq = " "
	}
	start2 := time.Now()
	fmt.Println(start2.UnixNano())
	//msecs1 := time.Since(start1)
	//fmt.Printf("for vaersion result: %s, %v\n", s1, msecs1)

	s2 := strings.Join(os.Args, " ")
	start3 := time.Now()
	fmt.Println(start3.UnixNano())
	msecs2 := time.Since(start1)
	fmt.Printf("string.Join version result: %s , %v\n", s2, msecs2)
}
