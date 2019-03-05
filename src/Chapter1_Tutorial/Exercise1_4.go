package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main(){
	for _, filename := range os.Args[1:]{
		data, err := ioutil.ReadFile(filename)
		if err != nil{
			fmt.Fprintf(os.Stderr, "Exercise1_4: %v\n", err)
		}
		counts := make(map[string]int)
		for _, line := range strings.Split(string(data), "\n"){
			counts[line]++
		}
		for _, value := range counts {
			if value > 2 {
				fmt.Printf("file %s has duplicated line\n", filename)
				break
			}
		}
	}
}