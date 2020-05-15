package leetcode

import (
	"testing"
)

/*
编写一个程序，找出第 n 个丑数。
丑数就是质因数只包含 2, 3, 5 的正整数。

示例:
输入: n = 10
输出: 12
解释: 1, 2, 3, 4, 5, 6, 8, 9, 10, 12 是前 10 个丑数。
*/

//动态规划 n+1一定是0~n的丑数 * （2/3/5）再取其中最小值得来
func nthUglyNumber(n int) int {
	ix2, ix3, ix5 := 0, 0, 0
	results := make([]int, n)
	results[0] = 1
	for i := 1; i < n; i++ {
		a, b, c := results[ix2]*2, results[ix3]*3, results[ix5]*5
		results[i] = getMin(getMin(a, b), c)
		//这里不能用switch case 因为有相等的情况 就都要加一
		if results[i] == a {
			ix2++
		}
		if results[i] == b {
			ix3++
		}
		if results[i] == c {
			ix5++
		}
	}
	return results[len(results)-1]
}

func TestFourtyNine(t *testing.T) {
	nthUglyNumber(10)
}
