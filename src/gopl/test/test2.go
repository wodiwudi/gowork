package main

import "sync"

func main() {
	var wg sync.WaitGroup
	go func() {
		wg.Add(1) //由于并发不保证执行顺序，可能Add尚未执行，就已经退出
		println("HI")
	}()
	wg.Wait()       //没执行,wait判断WaitGroup的数量不为0则阻塞。
	println("exit") //exit
}
