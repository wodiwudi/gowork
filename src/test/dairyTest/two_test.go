package dairyTest

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestTwo(t *testing.T) {
	var lock sync.RWMutex //读写锁实现同步 避免读写操作同时进行
	m := make(map[string]int)
	go func() {
		for {
			lock.Lock()
			m["a"] += 1
			lock.Unlock()
			time.Sleep(time.Microsecond) //写操作
		}
	}()
	go func() {
		for {
			lock.RLock()
			tmp := m["a"] //读操作
			lock.RUnlock()
			fmt.Println(tmp)
			time.Sleep(time.Microsecond)
		}
	}()
	select {} //阻止进程退出
}

//提前分配空间有助于提升性能，减少扩张时内存的分配和重新哈希操作
func BenchmarkTest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		test()
	}
}
func BenchmarkTestCap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		testCap()
	}
}

func test() map[int]int {
	m := make(map[int]int)
	for i := 0; i < 1000; i++ {
		m[i] = i
	}
	return m
}
func testCap() map[int]int {
	m := make(map[int]int, 1000)
	for i := 0; i < 1000; i++ {
		m[i] = i
	}
	return m
}
