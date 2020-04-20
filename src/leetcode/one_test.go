package leetcode

import (
	"fmt"
	"testing"
)

/*
给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。

示例:

给定 nums = [2, 7, 11, 15], target = 9

因为 nums[0] + nums[1] = 2 + 7 = 9
所以返回 [0, 1]
*/
func TestOne(t *testing.T) {
	nums := []int{1, 2, 3}
	fmt.Println(twoSum(nums, 4))
}

//map作法，根据map的if k,ok := xx ; ok 特性
func twoSum2(nums []int, target int) []int {
	result := map[int]int{}
	for key, value := range nums {
		if k, ok := result[target-value]; ok {
			return []int{k, key}
		}
		result[value] = key
	}
	return nil
}

//暴力做法，双循环
func twoSum(nums []int, target int) []int {
	length := len(nums)
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if nums[i]+nums[j] == target {
				return []int{nums[i], nums[j]}
			}
		}
	}
	return nil
}
