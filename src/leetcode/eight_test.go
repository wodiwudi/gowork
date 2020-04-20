package leetcode

import (
	"testing"
)

/*
给你一个以行程长度编码压缩的整数列表 nums 。
考虑每对相邻的两个元素 freq, val] = [nums[2*i], nums[2*i+1]] （其中 i >= 0 ），每一对都表示解压后子列表中有 freq 个值为 val 的元素，你需要从左到右连接所有子列表以生成解压后的列表。
请你返回解压后的列表。

示例：
输入：nums = [1,2,3,4]
输出：[2,4,4,4]
解释：第一对 [1,2] 代表着 2 的出现频次为 1，所以生成数组 [2]。
第二对 [3,4] 代表着 4 的出现频次为 3，所以生成数组 [4,4,4]。
最后将它们串联到一起 [2] + [4,4,4] = [2,4,4,4]。
示例 2：

输入：nums = [1,1,2,3]
输出：[1,3,3]
*/

func decompressRLElist(nums []int) []int {
	var result []int
	if len(nums)%2 != 0 {
		return result
	}
	for i, j := 0, 1; i < len(nums) && j < len(nums); i, j = i+2, j+2 {
		tmp := []int{}
		for x := 0; x < nums[i]; x++ {
			tmp = append(tmp, nums[j])
		}
		result = append(result, tmp...)
	}
	return result
}

func TestEight(t *testing.T) {
	test := []int{1, 1, 2, 3}
	t.Log(decompressRLElist(test))
}
