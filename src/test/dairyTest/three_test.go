package dairyTest

import (
	"fmt"
	"testing"
)

type file struct {
	name string
	attr struct {
		owner int
		perm  int
	}
}

//长度为零的对象通常都指向runtime.zerobase变量
func TestThree(t *testing.T) {
	a := [10]struct{}{}
	b := a[:]
	c := [0]int{}
	fmt.Printf("%p,%p,%p\n", &a[0], &b[0], &c) //0x653d88,0x653d88,0x653d88
}
