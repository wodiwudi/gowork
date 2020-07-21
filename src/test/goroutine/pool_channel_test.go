package main

import (
	"fmt"
	"testing"
)

//通道本身是并发安全的队列，因此可以做Pool和ID生成器等
type pool chan []byte

func newPool(cap int) pool {
	return make(chan []byte, cap)
}

func (p pool) get() []byte {
	c := []byte{}
	select {
	case c = <-p: //返回
		fmt.Println(string(c))
	default: //返回失败 新建
		c = make([]byte, 0)
		fmt.Println("返回失败，新建空数据")
	}
	return c
}

func (p pool) put(b []byte) {
	select {
	case p <- b:
	default:
		fmt.Println("已经满了，放不下了")
	}
}

func TestPoolChannel(t *testing.T) {
	pool := newPool(2)
	pool.put([]byte{97, 98, 99})
	pool.put([]byte{100, 101, 102})
	pool.put([]byte{103, 104})
	pool.get()
	pool.get()
	pool.get()
}
