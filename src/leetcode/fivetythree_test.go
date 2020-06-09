package leetcode

/*
将一个按照升序排列的有序数组，转换为一棵高度平衡二叉搜索树。
本题中，一个高度平衡二叉树是指一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1。
示例:
给定有序数组: [-10,-3,0,5,9],
一个可能的答案是：[0,-3,9,-10,null,5]，它可以表示下面这个高度平衡二叉搜索树：
      0
     / \
   -3   9
   /   /
 -10  5
*/
import (
	"testing"
)

//二分法
func sortedArrayToBST2(nums []int) *TreeNode {
	return arrGetMax(nums, 0, len(nums)-1)
}

func arrGetMax(arr []int, start, end int) *TreeNode {
	if start > end {
		return nil
	}
	rootIndex := (start + end + 1) / 2
	root := &TreeNode{arr[rootIndex], nil, nil}
	root.Left, root.Right = arrGetMax(arr, start, rootIndex-1), arrGetMax(arr, rootIndex+1, end)
	return root
}

func TestFivetyThree(t *testing.T) {
	sortedArrayToBST2([]int{-10, -3, 0, 5, 9})
}
