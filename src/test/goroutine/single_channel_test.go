package main

import (
	"fmt"
	"sync"
	"testing"
)

func TestSingleChannel(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(2)
	c := make(chan int, 0)
	var send chan<- int = c
	var recv <-chan int = c

	go func() {
		defer wg.Done()
		for x := range recv {
			fmt.Println(x)
		}
	}()

	go func() {
		defer wg.Done()
		defer close(c)

		for i := 0; i < 3; i++ {
			send <- i
		}
	}()
	wg.Wait()
}
