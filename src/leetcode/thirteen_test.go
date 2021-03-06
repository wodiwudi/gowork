package leetcode

import (
	"fmt"
	"testing"
)

/*
0,1,,n-1这n个数字排成一个圆圈，从数字0开始，每次从这个圆圈里删除第m个数字。求出这个圆圈里剩下的最后一个数字。
例如，0、1、2、3、4这5个数字组成一个圆圈，从数字0开始每次删除第3个数字，则删除的前4个数字依次是2、0、4、1，因此最后剩下的数字是3。
示例 1：

输入: n = 5, m = 3
输出: 3
示例 2：

输入: n = 10, m = 17
输出: 2
*/

//约瑟夫环问题，公式f(N,M)=(f(N−1,M)+M)%n  是一个反推导的过程 详情：https://blog.csdn.net/u011500062/article/details/72855826
func lastRemaining(n int, m int) int {
	var last int
	for i := 2; i <= n; i++ {
		last = (last + m) % i
	}
	return last
}

func TestThirteen(t *testing.T) {
	n, m := 5, 3
	t.Log(lastRemaining(n, m))
	fmt.Println()
}
