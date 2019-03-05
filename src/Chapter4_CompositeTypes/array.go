package main

import "fmt"

func main()  {
	//var a [3]int
	//fmt.Println(a[0])
	//fmt.Println(a[len(a)-1])
	//
	//for i, v := range a{
	//	fmt.Printf("%d %d\n", i, v)
	//}
	//
	//for _, v := range a{
	//	fmt.Printf("%d\n", v)
	//}
	//
	//var q [3]int = [3]int{1,2,3}
	//var r [3]int = [3]int{1, 2}
	//fmt.Println(r[2], q)
	//
	//s := [...]int{1,2,3}
	//fmt.Printf("%T\n", s)
	//
	//type Currency int
	//const(
	//	USD Currency = iota
	//	EUR
	//	GBP
	//	RMB
	//)
	//
	//symbol := [...]string{USD: "$",  EUR: "€", GBP: "£", RMB: "¥"}
	//fmt.Println(RMB, symbol[RMB])
	//
	//t := [...]int{99:-1}
	//fmt.Println(t)

	a := [2]int{1, 2}
	b := [...]int{1,2}
	c := [2]int{1,3}
	fmt.Println(a==b, a==c, b==c)
	d :=[3]int{1,2}
	fmt.Println(a==d)



}
