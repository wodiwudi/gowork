package leetcode

import (
	"fmt"
	"testing"
)

/*
假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
注意：给定 n 是一个正整数。
示例 1：

输入： 2
输出： 2
解释： 有两种方法可以爬到楼顶。
1.  1 阶 + 1 阶
2.  2 阶
示例 2：

输入： 3
输出： 3
解释： 有三种方法可以爬到楼顶。
1.  1 阶 + 1 阶 + 1 阶
2.  1 阶 + 2 阶
3.  2 阶 + 1 阶
*/
//用数学归纳法分析，化成最小单位开始分析
//n=1 1
//n=2 2
//n=3 f(1)+f(2)
//n=4 f(2)+f(3)
//f(n) = f(n-2)+f(n-1)
//本质为斐波那契额数列，但是重复递归复杂度为2^n 改为动态规划
func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	arr := []int{}
	for i := 0; i < n; i++ {
		arr = append(arr, i+1)
	}
	f1, f2, f3 := 1, 2, 3
	for k, _ := range arr {
		if k < 2 {
			continue
		}
		f3 = f1 + f2
		f1 = f2
		f2 = f3
	}
	return f3
}
func TestTweenTone(t *testing.T) {
	fmt.Println(climbStairs(4))
}
