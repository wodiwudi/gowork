package main

import (
	"fmt"
	"testing"
	"time"
)

func init() {
	done := make(chan struct{})
	go func() {
		defer close(done)
		fmt.Println("init :", time.Now())
		time.Sleep(time.Second * 2)
	}()
	<-done
}

func TestInitChannel(t *testing.T) {
	fmt.Println("main :")
}
