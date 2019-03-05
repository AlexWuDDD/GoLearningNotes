package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

func main(){
	prefix := "http://"
	for _,url := range os.Args[1:]{
		//Exercise1.8
		if !strings.HasPrefix(url, prefix){
			url = prefix+ url
		}
		resp, err := http.Get(url)
		if err != nil{
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		//Exercise 1.9
		fmt.Println("HTTP Status code: ", resp.Status)

		//b, err := ioutil.ReadAll(resp.Body)
		//Exercise1.7
		//b, err = io.Copy(os.Stdout,resp.Body)
		//resp.Body.Close()
		//if err!= nil{
		//	fmt.Fprintf(os.Stderr, "fecth: reading %s: %v\n",
		//		url, err)
		//}
		//fmt.Printf("%s", b)

		//Exercise1.7
		//written, err := io.Copy(os.Stdout,resp.Body)
		resp.Body.Close()
		//if(err != nil){
		//	fmt.Fprintf(os.Stdout, "fetch:Copy: %v",err)
		//}
		//fmt.Println(reflect.TypeOf(written), written)


	}



}
