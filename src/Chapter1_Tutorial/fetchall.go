package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main(){
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:]{
		go fetch(url, ch) //start a goroutine
	}
	for range os.Args[1:]{
		fmt.Println(<-ch) // receive from channel ch
	}
	fmt.Printf("%.2fs elapased\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan <- string)  {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil{
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	//Exercise 1.10
	//data, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil{
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	//err = ioutil.WriteFile("alex.txt", data, os.ModeAppend)
	//if err != nil{
	//	ch <- fmt.Sprintf("while wrting file: %v", err)
	//}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
	//ch <- fmt.Sprintf("%.2fs %s", secs, url)

}