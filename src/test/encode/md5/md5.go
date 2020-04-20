package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
)

func strToMD5() {
	str := "testMD5"
	data := []byte(str)
	has := md5.Sum(data)
	md5str1 := fmt.Sprintf("%x", has)
	fmt.Println(md5str1)
}

func strToMD5V2() {
	str := "testMD5"
	w := md5.New()
	io.WriteString(w, str)
	bw := w.Sum(nil)
	md5str2 := hex.EncodeToString(bw) //将 bw 转成字符串
	fmt.Println(md5str2)
}

func main() {
	//strToMD5()
	strToMD5V2()
}
