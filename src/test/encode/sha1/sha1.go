package main

import (
	"crypto/sha1"
	"fmt"
)

func strToSha1() {
	str := "test sha1"
	data := []byte(str)
	has := sha1.Sum(data)
	shastr1 := fmt.Sprintf("%x", has)
	fmt.Println(shastr1)
}

func main() {
	strToSha1()
}
