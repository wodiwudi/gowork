package main

import (
	"testing"
)

//将发往通道的数据打包，减少传输次数，可以提高性能，因为通道队列是通过锁同步机制来实现的，
// 单次获取更多数据，可改善因频繁加锁造成的性能问题
const (
	max     = 50000000
	block   = 500
	bufsize = 100
)

func test() { //普通模式
	done := make(chan struct{})
	c := make(chan int, bufsize)
	go func() {
		defer close(done)
		count := 0
		for x := range c {
			count += x
		}
	}()

	for i := 0; i < max; i++ {
		c <- i
	}

	close(c)
	<-done
}

func testBlock() { //块模式
	done := make(chan struct{})
	c := make(chan [block]int, bufsize)
	go func() {
		defer close(done)
		count := 0
		for a := range c {
			for _, x := range a {
				count += x
			}
		}
	}()

	for i := 0; i < max; i += block {
		tmp := [block]int{}
		for j := 0; j < block; j++ {
			tmp[j] = i + j
		}
		c <- tmp
	}
	close(c)
	<-done
}

func BenchmarkTest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		test()
	}
}

func BenchmarkTestBlock(b *testing.B) {
	for i := 0; i < b.N; i++ {
		testBlock()
	}
}
