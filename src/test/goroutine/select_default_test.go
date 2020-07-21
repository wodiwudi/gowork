package main

import (
	"fmt"
	"testing"
)

func TestSelectDefault(t *testing.T) {
	done := make(chan struct{})
	data := []chan int{
		make(chan int, 3),
	}
	//创造数据
	go func() {
		defer close(done)
		for i := 0; i < 10; i++ {
			select {
			case data[len(data)-1] <- i:
				//default 可以用来做其他逻辑 当所有select不可用时会执行
			default:
				data = append(data, make(chan int, 3))
			}
		}
	}()
	<-done
	fmt.Println(len(data))
	for i := 0; i < len(data); i++ {
		c := data[i]
		close(c)
		for x := range c {
			fmt.Println(x)
		}
	}

}
