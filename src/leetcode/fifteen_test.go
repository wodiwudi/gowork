package leetcode

import (
	"fmt"
	"testing"
)

func removeDuplicates(nums []int) int {
	for i := 0; i < len(nums); i++ {
		j := i + 1
		if j >= len(nums) {
			return len(nums)
		}
		if nums[i] == nums[j] {
			nums = append(nums[:i+1], nums[j+1:]...)
			i--
		}
	}
	return len(nums)
}

func TestFifteen(t *testing.T) {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(removeDuplicates(nums))
}
func removeDuplicates2(nums []int) int {
	for i := len(nums) - 1; i > 0; i-- {
		if nums[i] == nums[i-1] {
			nums = append(nums[:i], nums[i+1:]...)
		}
		fmt.Println(nums)
	}
	return len(nums)
}
