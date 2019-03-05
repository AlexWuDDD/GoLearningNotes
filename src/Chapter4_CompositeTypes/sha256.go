package main

import (
	"crypto/sha256"
	"fmt"
)

func main()  {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1==c2, c1)
	fmt.Println(countDiffBitsBetweenSHA256(&c1, &c2))
}

//Exercise4.1
func countDiffBitsBetweenSHA256(hash1 *[32]uint8, hash2 *[32]uint8) int{
	var sum int;
	for i:=0; i<32; i++{
		ret := (*hash1)[i] ^ (*hash2)[i]
		for j:=uint(0);j<8;j++{
			fmt.Println(((ret>>j) & 1))
			sum += int((ret>>j) & 1)
		}
	}
	return sum
}


