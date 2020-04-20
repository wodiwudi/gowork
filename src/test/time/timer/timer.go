package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	timer := time.AfterFunc(2*time.Second, func() {
		fmt.Println("after func callback, elaspe:", time.Now().Sub(start))
	})

	// time.Sleep(1 * time.Second)
	// time.Sleep(3*time.Second)
	timer.Stop()
	// Reset 在 Timer 还未触发时返回 true；触发了或Stop了，返回false
	if timer.Reset(3 * time.Second) {
		fmt.Println("timer has not trigger!")
	} else {
		fmt.Println("timer had expired or stop!")
	}

	time.Sleep(10 * time.Second)
}
