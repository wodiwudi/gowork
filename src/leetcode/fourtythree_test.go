package leetcode

import (
	"fmt"
	"testing"
)

/*
输入一个整型数组，数组里有正数也有负数。数组中的一个或连续多个整数组成一个子数组。求所有子数组的和的最大值。
要求时间复杂度为O(n)。

示例1:

输入: nums = [-2,1,-3,4,-1,2,1,-5,4]
输出: 6
解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。
*/
func maxSubArray(nums []int) int {
	//  动态规划法 列出状态方程
	// 以nums[i]为结尾的最大数设为dp[i] ,则 dp[i]要么等于dp[i-1] + nums[i] 要么等于nums[i]
	//如果dp[i-1]是负数 那么一定是nums[i]否则为dp[i-1] + nums[i]
	result, pre, max := nums[0], nums[0], 0
	for i := 1; i < len(nums); i++ {
		if pre < 0 {
			max = nums[i]
		} else {
			max = pre + nums[i]
		}
		pre = max
		result = getMax(max, result)
	}
	return result
}

func TestFourtyThree(t *testing.T) {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}
