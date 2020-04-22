package leetcode

import (
	"fmt"
	"testing"
)

func exchange(nums []int) []int {
	//快慢指针解法
	i, j := 0, 0
	length := len(nums)
	for j < length {
		if nums[j]&1 != 0 {
			//奇数
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
		j++
	}
	return nums
}

func TestFourtyOne(t *testing.T) {
	nums := []int{1, 2, 3, 4}
	fmt.Println(exchange(nums))
}
