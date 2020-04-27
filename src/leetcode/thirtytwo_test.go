package leetcode

import (
	"fmt"
	"testing"
)

/*
请实现一个函数，输入一个整数，输出该数二进制表示中 1 的个数。例如，
把 9 表示成二进制是 1001，有 2 位是 1。因此，如果输入 9，则该函数输出 2。
示例 1：

输入：00000000000000000000000000001011
输出：3
解释：输入的二进制串 00000000000000000000000000001011 中，共有三位为 '1'。
*/
func hammingWeight(num uint32) int {
	result := 0
	for num > 0 {
		//1/任何数都为0
		if num%2 == 1 {
			result++
		}
		num /= 2
	}
	return result
}
func TestThirtyTwo(t *testing.T) {
	var num uint32 = 11
	fmt.Println(hammingWeight(num))
}
