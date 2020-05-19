package ch8

import (
	"fmt"
	"sync"
	"testing"
)

//sync.pool是对象缓存 但是不适合用来做连接池 因为GC会将其生命周期回收掉，get时，
// 会先找私有对象，没有去找共享池，共享池是协程安全（但是有锁），
// 再没有就取其他processer的对象，都没有就调用New函数
func TestSyncPool(t *testing.T) {
	pool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("create new obj")
			return 100
		},
	}
	pool.Put(1)
	pool.Put(2)
	pool.Put(3)
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			v := pool.Get()
			fmt.Println(v)
			wg.Done()
		}()
	}
	wg.Wait()
}
