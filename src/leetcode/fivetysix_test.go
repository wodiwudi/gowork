package leetcode

import (
	"fmt"
	"testing"
)

func search(nums []int, target int) int {
	left, right, middle := 0, len(nums)-1, (len(nums)-1)/2
	for left <= right {
		if nums[middle] == target {
			return middle
		} else if target > nums[middle] {
			left = middle + 1
		} else {
			right = middle - 1
		}
		middle = (left + right) / 2
	}
	return -1
}
func TestFivetySix(t *testing.T) {
	fmt.Println(search([]int{-1, 0, 3, 5, 9, 12}, 9))
}
