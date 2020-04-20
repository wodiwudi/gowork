package one

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func returnMutilValue() (int, int) {
	return rand.Intn(10), rand.Intn(15)
}

func timeSpent(inner func(op int) int) func(op int) int {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("time spent : ", time.Since(start).Seconds())
		return ret
	}
}

func slowFunc(number int) int {
	time.Sleep(time.Second * 1)
	return number
}

func TestMutilValue(t *testing.T) {
	a, b := returnMutilValue()
	t.Log(a, b)
	newFunc := timeSpent(slowFunc)
	fmt.Printf("%T", newFunc(10))
	t.Log(newFunc(100))
}
