package leetcode

import (
	"fmt"
	"testing"
)

/*
MinStack minStack = new MinStack();
minStack.push(-2);
minStack.push(0);
minStack.push(-3);
minStack.min();   --> 返回 -3.
minStack.pop();
minStack.top();      --> 返回 0.
minStack.min();   --> 返回 -2.
*/
type MinStack struct {
	data   []int //数据
	min    []int //辅助栈  以实现min函数的O(1)复杂度
	length int   //长度
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		data:   []int{},
		length: 0,
	}
}

func (this *MinStack) Push(x int) {
	//如果是第一次，两个都append，当x > min的最后一个时，就把min最后一个元素append到min上 <=时，把x加到min上
	if this.length == 0 || x <= this.min[this.length-1] {
		this.min = append(this.min, x)
	} else {
		this.min = append(this.min, this.min[this.length-1])
	}
	this.data = append(this.data, x)
	this.length++
}

func (this *MinStack) Pop() {
	//每一次pop都一起pop出去来保持一致
	if this.length > 0 {
		this.data = this.data[:this.length-1]
		this.min = this.min[:this.length-1]
		this.length--
	}
}

func (this *MinStack) Top() int {
	return this.data[this.length-1]
}

func (this *MinStack) Min() int {
	if this.length > 0 {
		return this.min[this.length-1]
	} else {
		return 0
	}
}

func TestMinStack(t *testing.T) {
	stack := Constructor()
	stack.Push(109)
	stack.Push(2)
	stack.Push(9)
	fmt.Println(stack.Top())
	fmt.Println(stack.data)
	fmt.Println(stack.Min())
}
