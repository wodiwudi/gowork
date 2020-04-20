package main

import (
	"fmt"
	regexp2 "regexp"
)

const text = "my email is ccmouse@gmail.com"
const test2 = `1@33.com.cn  ff@4.com
               @1.com   io@c.com`

func main() {
	regexp := regexp2.MustCompile(`[a-zA-Z0-9]+@[a-zA-Z0-9.]+\.[a-zA-Z0-9]+`) //生成一个匹配器  ,反引号中的\不会被golang解析成\n等内容
	match := regexp.FindString(text)                                          //将内容放入匹配器中验证
	fmt.Println(match)                                                        //输出匹配结果 ccmouse@gmail.com

	regexp2 := regexp2.MustCompile(`([a-zA-Z0-9]+)@[a-zA-Z0-9.]+\.[a-zA-Z0-9]+`) //（）可以提取
	match2 := regexp2.FindAllStringSubmatch(test2, -1)                           //找到所有子串,-1代表所有,返回一个[]string的slice
	for _, m := range match2 {
		fmt.Println(m[1]) //1 ff io
	}
}
