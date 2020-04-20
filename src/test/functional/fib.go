package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

//高级斐波那契数列 -将数列存入byte数组文件，再读出来
type intGen func() int //定义type函数类型

func (g intGen) Read(p []byte) (n int, err error) { //intGen函数实现Read接口，将内容加进byte数组
	next := g()
	if next > 1000 { //当斐波那契数到1000停止
		return 0, io.EOF
	}
	s := fmt.Sprintf("%d\n", next)      //将斐波那契数列返回内容变为string,因为string实现了Read接口
	return strings.NewReader(s).Read(p) //返回string里面实现的Read的接口
}

func fib() intGen {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

//读取io流内容
func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	f := fib()
	printFileContents(f)
}
