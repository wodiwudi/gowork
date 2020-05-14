package Singleton

import (
	"fmt"
	"sync"
	"testing"
)

type single struct {
}

var instance *single
var one sync.Once

func instanceSingle() *single {
	one.Do(func() {
		instance = new(single)
	})
	return instance
}

//sync.Once 可以确保在多协程或多线程中只运行一次
func TestSingleton(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			tmp := instanceSingle()
			fmt.Printf("%v,%p\n", tmp, tmp)
			wg.Done()
		}()
	}
	wg.Wait()
}
