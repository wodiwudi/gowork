package leetcode

import (
	"testing"
)

/*
给你一个整数 n，请你帮忙计算并返回该整数「各位数字之积」与「各位数字之和」的差。

示例 1：

输入：n = 234
输出：15
解释：
各位数之积 = 2 * 3 * 4 = 24
各位数之和 = 2 + 3 + 4 = 9
结果 = 24 - 9 = 15
*/
func subtractProductAndSum(n int) int {
	var sub = 0
	var mul = 1
	for {
		num := n % 10
		sub += num
		mul *= num
		if n < 10 {
			break
		}
		n = n / 10
	}
	return mul - sub
}

func TestFive(t *testing.T) {
	t.Log(subtractProductAndSum(234))
}
