package main

import (
	"fmt"
	"sync"
	"testing"
)

func TestSelectChannel(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(2)
	a, b := make(chan int), make(chan int)
	go func() {
		defer wg.Done()
		for {

			select {
			case x, ok := <-a:
				if !ok {
					a = nil
					break
				}
				name := "a"
				fmt.Println(x, name)
			case x, ok := <-b:
				if !ok {
					a = nil
					break
				}
				name := "b"
				fmt.Println(x, name)
			}
			if a == nil && b == nil {
				return
			}
		}
	}()

	go func() {
		defer wg.Done()
		defer close(a)
		for i := 0; i < 4; i++ {
			a <- i
		}
	}()

	go func() {
		defer wg.Done()
		defer close(b)
		for i := 0; i < 6; i++ {
			b <- i * 10
		}
	}()
	wg.Wait()
}

//select即使都是同一通道也会随机选择case执行
func TestSelectChannel2(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(2)
	c := make(chan int)

	go func() {
		defer wg.Done()
		var (
			ok bool
			x  int
		)
		for {
			select {
			case x, ok = <-c:
				fmt.Println("a1", x)
			case x, ok = <-c:
				fmt.Println("a2", x)
			}
			//最后会多收到一个0 ，因为接收关闭的通道会返回对应类型零值
			if !ok {
				return
			}
		}
	}()

	go func() {
		defer wg.Done()
		defer close(c)
		for i := 1; i < 11; i++ {
			select {
			case c <- i:
			case c <- i * 10:

			}
		}
	}()
	wg.Wait()
}
