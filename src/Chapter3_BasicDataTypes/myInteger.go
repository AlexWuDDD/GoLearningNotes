package main

import "fmt"

func main(){
	var x uint8 = 1<<1 | 1<<5
	var y uint8 = 1<<1 | 1<<2

	fmt.Printf("%08b\n", x)
	fmt.Printf("%08b\n", y)

	fmt.Printf("%08b\n", x & y)
	fmt.Printf("%08b\n", x | y)
	fmt.Printf("%08b\n", x ^ y)
	fmt.Printf("%08b\n", x &^ y)
	//y对应位置上是1的在x上都清零

	//检测x的哪一位是1
	for i:=uint(0); i < 8; i++{
		if x&(1<<i) !=0{
			fmt.Println(i)
		}
	}

	fmt.Printf("%08b\n", x << 1)
	fmt.Printf("%08b\n", x >> 1)


	o := 0666
	fmt.Printf("%d %[1]o %#[1]o\n", o)

	m := int64(0xdeadbeef)
	fmt.Printf("%d %[1]x %#[1]x %#[1]X\n", m)


	ascii := 'a'
	unicode := '国'
	newline := '\n'
	fmt.Printf("%d %[1]c %[1]q\n", ascii)
	fmt.Printf("%d %[1]c %[1]q\n", unicode)
	fmt.Printf("%d %[1]q\n", newline)



}
