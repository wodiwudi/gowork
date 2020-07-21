package mutex

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

//通道并非用来取代锁的，通道倾向于解决逻辑层次的并发处理问题，锁用来保护局部范围内的数据安全
//mutex作为匿名字段时，相关方法必须实现为pointer-receiver否则会复制锁导致锁失效
type data struct {
	sync.Mutex //互斥锁
}

func (d *data) test(s string) {
	d.Lock()
	defer d.Unlock()
	for i := 0; i < 5; i++ {
		fmt.Println(s, i)
		time.Sleep(time.Second * 1)
	}
}
func TestMutex(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(2)
	var d data
	go func() {
		defer wg.Done()
		d.test("read")
	}()

	go func() {
		defer wg.Done()
		d.test("write")
	}()
	wg.Wait()
}
