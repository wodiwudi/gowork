package leetcode

import (
	"fmt"
	"math"
	"testing"
)

/*
给定两个数组，编写一个函数来计算它们的交集。
示例 1:
输入: nums1 = [1,2,2,1], nums2 = [2,2]
输出: [2]
示例 2:

输入: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
输出: [9,4]
*/
func intersection(nums1 []int, nums2 []int) []int {
	len1, len2 := len(nums1), len(nums2)
	min := math.Min(float64(len1), float64(len2))
	m := make(map[int]int, 0)
	res := []int{}
	if min == float64(len1) {
		for i := 0; i < len1; i++ {
			m[nums1[i]] = i
		}
		for _, v := range nums2 {
			if _, ok := m[v]; ok {
				res = append(res, v)
				delete(m, v)
			}
		}
		return res
	} else {
		for i := 0; i < len2; i++ {
			m[nums2[i]] = i
		}
		for _, v := range nums1 {
			if _, ok := m[v]; ok {
				res = append(res, v)
				delete(m, v)
			}
		}
		return res
	}
}

func TestFivetyTwo(t *testing.T) {
	nums1, nums2 := []int{4, 9, 5}, []int{9, 4, 9, 8, 4}
	fmt.Println(intersection(nums1, nums2))
}
