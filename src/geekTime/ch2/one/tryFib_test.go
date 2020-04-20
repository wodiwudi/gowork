package one

import (
	"fmt"
	"testing"
)

func TestFib(t *testing.T) {
	var (
		a int = 1
		b     = 1
	)
	fmt.Print(a)
	for i := 0; i < 10; i++ {
		fmt.Print(" ", b)
		tmp := a
		a = b
		b = a + tmp
	}
	fmt.Println()
}
