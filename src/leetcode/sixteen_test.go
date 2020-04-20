package leetcode

import (
	"fmt"
	"testing"
)

/*
给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。

示例:

输入: [0,1,0,3,12]
输出: [1,3,12,0,0]
*/

func moveZeroes(nums []int) []int {
	//1.把零干掉，拼接数组 2、末尾补零
	count := 0
	i := 0
	for i < len(nums) {
		fmt.Println(nums)
		if nums[count] == 0 {
			nums = append(nums[:count], nums[count+1:]...)
			nums = append(nums, 0)
			i++
			continue
		}
		count++
		i++
	}
	return nums
}

func moveZeroes2(nums []int) {
	//1.计零位 2.计零位交换
	zero := 0
	for k, _ := range nums {
		if nums[k] != 0 {
			nums[k], nums[zero] = nums[zero], nums[k]
			zero++
		}
	}
}

func TestSixTeen(t *testing.T) {
	//nums := []int{4, 2, 4, 0, 0, 3, 0, 5, 1, 0}
}
