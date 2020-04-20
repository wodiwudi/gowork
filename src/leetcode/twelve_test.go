package leetcode

import (
	"fmt"
	"testing"
)

/*
给定一个有序整数数组，元素各不相同且按升序排列，编写一个算法，创建一棵高度最小的二叉搜索树。

示例:
给定有序数组: [-10,-3,0,5,9],
一个可能的答案是：[0,-3,9,-10,null,5]，它可以表示下面这个高度平衡二叉搜索树：

          0
         / \
       -3   9
       /   /
     -10  5

*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	if len(nums) == 1 {
		return &TreeNode{
			Val:   nums[0],
			Left:  nil,
			Right: nil,
		}
	}
	middle := len(nums) / 2
	head := &TreeNode{
		Val:   nums[middle],
		Left:  nil,
		Right: nil,
	}
	//这种递归的前提是数组是从大到小有序的
	head.Left = sortedArrayToBST(nums[:middle])
	head.Right = sortedArrayToBST(nums[middle+1:])
	return head
}
func TestTwelve(t *testing.T) {
	nums := []int{-10, -3, 0, 5, 9}
	fmt.Println(sortedArrayToBST(nums))
}
