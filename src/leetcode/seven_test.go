package leetcode

import (
	"testing"
)

/*
给定一副牌，每张牌上都写着一个整数。
此时，你需要选定一个数字 X，使我们可以将整副牌按下述规则分成 1 组或更多组：
每组都有 X 张牌。
组内所有的牌上都写着相同的整数。
仅当你可选的 X >= 2 时返回 true。

示例 1：

输入：[1,2,3,4,4,3,2,1]
输出：true
解释：可行的分组是 [1,1]，[2,2]，[3,3]，[4,4]
示例 2：

输入：[1,1,1,2,2,2,3,3]
输出：false
解释：没有满足要求的分组。
示例 3：

输入：[1]
输出：false
解释：没有满足要求的分组。
示例 4：

输入：[1,1]
输出：true
解释：可行的分组是 [1,1]
示例 5：

输入：[1,1,2,2,2,2]
输出：true
解释：可行的分组是 [1,1]，[2,2]，[2,2]

提示：
1 <= deck.length <= 10000
0 <= deck[i] < 10000
*/
func hasGroupsSizeX(deck []int) bool {
	//如果是一个不足以构成题意
	if len(deck) == 1 {
		return false
	}
	//辗转相除法计算最大公约数
	gcd := func(x, y int) int {
		for x%y != 0 {
			x, y = y, x%y
		}
		return y
	}
	//记录每个数字出现的个数，一般这种记录都用 map[int]int 表示
	count := make(map[int]int, 0)
	for k, _ := range deck {
		count[deck[k]]++
	}
	var x int
	for _, v := range count { //遍历map，求每个数字次数和下一个数字出现次数的最大公约数，如果每个都大于等于2，则true
		if x == 0 {
			x = v
		}
		x = gcd(x, v)
	}
	return x > 1
}

func TestSeven(t *testing.T) {
	test := []int{1, 1, 2, 2, 2, 2}
	t.Log(hasGroupsSizeX(test))
}
