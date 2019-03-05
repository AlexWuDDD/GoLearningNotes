package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"os"
	"strconv"
)

func main()  {
	if len(os.Args) != 3{
		fmt.Println("Please following below format:")
		fmt.Println("... + 384/256 + 'your text'")
		return;
	}

	if os.Args[1] == strconv.Itoa(256){
		message := []byte(os.Args[2])
		hash := sha256.New()
		hash.Write(message)
		bytes := hash.Sum(nil)
		hashcode := hex.EncodeToString(bytes)
		fmt.Println("sha256: ", hashcode)
	} else if os.Args[1] == strconv.Itoa(384){
		message := []byte(os.Args[2])
		hash := sha512.New384()
		hash.Write(message)
		bytes := hash.Sum(nil)
		hashcode := hex.EncodeToString(bytes)
		fmt.Println("sha256: ", hashcode)
	} else{
		fmt.Println("Error", os.Args[0], os.Args[1], os.Args[2])
	}
}
