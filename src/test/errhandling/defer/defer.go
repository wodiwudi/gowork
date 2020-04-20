package main

import (
	"bufio"
	"fmt"
	"os"
	"test/functional/fib"
)

//斐波那契数列写入文件
func writeFile(num int, filename string) {
	//filename 写入文件名字
	f := fib.Fib()
	file, err := os.Create(filename) //创建一个文件
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file) //建立一个写入的缓冲槽
	defer writer.Flush()            //将缓冲数据输出
	for i := 0; i <= num; i++ {
		_, err := fmt.Fprintln(writer, f())
		if err != nil {
			panic(err)
		}
	}
}

func main() {
	writeFile(10, "fib.txt")
}
