package leetcode

import (
	"fmt"
	"testing"
)

/*
给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有满足条件且不重复的三元组。
注意：答案中不可以包含重复的三元组。
示例：

给定数组 nums = [-1, 0, 1, 2, -1, -4]，

满足要求的三元组集合为：
[
  [-1, 0, 1],
  [-1, -1, 2]
]
*/
//排序加双指针解法，固定i，创造left和right指针 然后遍历
func threeSum(nums []int) [][]int {
	//1.排序 2.for循环 定两个指针left,right for left!=right 3.判断三数之和就加上4.消除重复
	length := len(nums)
	if length < 3 {
		return nil
	}
	result := [][]int{}
	for i := 0; i < length; i++ {
		for j := 0; j < length-1-i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	for i := 0; i < length-2; i++ {
		left, right := i+1, length-1
		for left != right && left < right {
			tmp := []int{}
			sum := nums[i] + nums[left] + nums[right]
			if sum > 0 {
				right--
			} else if sum < 0 {
				left++
			} else {
				tmp = append(tmp, nums[i], nums[left], nums[right])
				result = append(result, tmp)
				left++
				right--
			}
		}
	}
	if len(result) == 1 {
		return result
	}
	//去重
	for n := 0; n < len(result); n++ {
		for m := n + 1; m < len(result); m++ {
			if isEqual(result[n], result[m]) {
				result = append(result[:m], result[m+1:]...)
				m--
			}
		}
	}
	return result
}
func isEqual(leftArr, rightArr []int) bool {
	count := 0
	for i := 0; i < 3; i++ {
		if leftArr[i] == rightArr[i] {
			count++
		}
	}
	if count == 3 {
		return true
	}
	return false
}

func TestTweenSeven(t *testing.T) {
	nums := []int{5, -9, -11, 9, 9, -4, 14, 10, -11, 1, -13, 11, 10, 14, -3, -3, -4, 6, -15, 6, 6, -13, 7, -11, -15, 10, -8, 13, -14, -12, 12, 6, -6, 8, 0, 10, -11, -8, -2, -6, 8, 0, 12, 3, -9, -6, 8, 3, -15, 0, -6, -1, 3, 9, -5, -5, 4, 2, -15, -3, 5, 13, -11, 7, 6, -4, 2, 11, -5, 7, 12, -11, -15, 1, -1, -9, 10, -8, 1, 2, 8, 11, -14, -4, -3, -12, -2, 8, 5, -1, -9, -4, -3, -13, -12, -12, -10, -3, 6, 1, 12, 3, -3, 12, 11, 11, 10}
	fmt.Println(threeSum(nums))
}
