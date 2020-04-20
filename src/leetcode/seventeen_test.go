package leetcode

import (
	"fmt"
	"testing"
)

func maxArea(height []int) int {
	//1.暴力破解法 n^2
	maxArea := 0
	for i := 0; i < len(height); i++ {
		for j := i + 1; j < len(height); j++ {
			area := (j - i) * getmin(height[i], height[j])
			if area > maxArea {
				maxArea = area
			}
		}
	}
	return maxArea
}
func getmin(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func maxArea2(height []int) int {
	//双指针法，i，j代表左右两边指针，初始为最大宽度，直到i,j相遇停止循环。
	//由于向内靠拢会宽度减少，对于面积来说，area = height * width 所以两边h[i]与h[j]中短板那一个要移动。
	//看是否有更长的板子，这样才有可能有更大的面积。
	maxArea, i, j := 0, 0, len(height)-1
	for i != j {
		hi, hj := height[i], height[j]
		area := (j - i) * getmin(hi, hj)
		if area > maxArea {
			maxArea = area
		}
		if hi > hj {
			j--
		} else {
			i++
		}
	}
	return maxArea
}

func TestSevenTeen(t *testing.T) {
	nums := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(nums))
}
