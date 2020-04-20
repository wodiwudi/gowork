package main

import (
	"encoding/base64"
	"fmt"
)

/*string转base64 */
func strToBase64() string {
	str := "test base64"
	input := []byte(str)
	base64Str := base64.StdEncoding.EncodeToString(input)
	fmt.Println("strToBase64 :", base64Str)
	return base64Str
}

/*base64转string */
func base64ToStr(base64str string) {
	str, err := base64.StdEncoding.DecodeString(base64str)
	if err != nil {
		panic(err)
	}
	fmt.Println("Base64Tostr :", string(str))
}

/*url转base64 */
func urlToBase64() string {
	url := "http://www.baidu.com"
	input := []byte(url)
	base64Url := base64.URLEncoding.EncodeToString(input)
	fmt.Println("urlToBase64 :", base64Url)
	return base64Url
}

/*base64转url */
func base64ToUrl(base64url string) {
	str, err := base64.URLEncoding.DecodeString(base64url)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(str))
}
func main() {
	strToBase64()
	base64ToStr(strToBase64())
	urlToBase64()
	base64ToUrl(urlToBase64())
}
