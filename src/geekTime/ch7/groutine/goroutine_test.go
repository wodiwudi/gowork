package groutine

import (
	"sync"
	"testing"
	"time"
)

func TestGoroutineUnThreadSafe(t *testing.T) {
	counter := 0
	for i := 0; i < 5000; i++ {
		go func() {
			counter++
		}()
	}
	time.Sleep(2 * time.Second) //避免goroutine子进程还没都结束主进程就停止了
	t.Log(counter)              //4693  由于在test的主进程中与各个goroutine进程中对counter这个共享变量存在竞争，导致线程不安全
}

func TestGoroutineThreadSafe(t *testing.T) {
	var mu sync.Mutex //加锁保证线程安全
	counter := 0
	for i := 0; i < 5000; i++ {
		go func() {
			defer func() {
				mu.Unlock() //每个goroutine结束后要解锁否则会一直锁住
			}()
			mu.Lock()
			counter++
		}()
	}
	time.Sleep(2 * time.Second)
	t.Log(counter) //5000  但是时间为(2.01s)
}

func TestGoRoutineWithWaitGroup(t *testing.T) {
	var mu sync.Mutex //加锁保证线程安全
	var wg sync.WaitGroup
	counter := 0
	for i := 0; i < 5000; i++ {
		wg.Add(1) //每个goroutine加1
		go func() {
			defer func() {
				mu.Unlock() //每个goroutine结束后要解锁否则会一直锁住
				wg.Done()   //每个defer调用说明一个goroutine结束，则释放1
			}()
			mu.Lock()
			counter++
		}()
	}
	//time.Sleep(2 * time.Second) 由于不知道真实时间造成等待 改成waitGroup
	wg.Wait()
	t.Log(counter) //5000  不仅线程安全且不耗时（0.00s）
}
