package stack

import (
	"errors"
	"fmt"
)

func main() {
	stack := make([]int, 0, 5)

	push := func(x int) error {
		length := len(stack)
		if cap(stack) <= length {
			return errors.New("stack is full")
		}
		stack = stack[:length+1]
		stack[length] = x
		return nil
	}

	pop := func() (x int, err error) {
		length := len(stack)
		if length == 0 {
			return 0, errors.New("stack is empty")
		}
		x = stack[length-1]
		stack = stack[:length-1]
		return x, nil
	}
	//入栈测试
	for i := 0; i < 7; i++ {
		fmt.Printf("push %d,%v,%v\n", i, push(i), stack)
	}

	for i := 0; i < 7; i++ {
		x, err := pop()
		fmt.Printf("pop %d,%v,%v\n", i, x, err)
	}
}
