package channel

import (
	"fmt"
	"sync"
	"testing"
)

func TestProRec(t *testing.T) {
	ch := make(chan int, 0)
	wg := &sync.WaitGroup{}
	wg.Add(1)
	dataProducter(ch, wg)
	wg.Add(1)
	dataReceiver(ch, wg)
	wg.Wait()
}

func dataProducter(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
		wg.Done()
	}()
}
func dataReceiver(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for {
			if data, ok := <-ch; ok {
				fmt.Println(data)
			} else {
				break
			}
		}
		wg.Done()
	}()
}
