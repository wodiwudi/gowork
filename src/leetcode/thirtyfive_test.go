package leetcode

import (
	"fmt"
	"testing"
)

/*
从上到下按层打印二叉树，同一层的节点按从左到右的顺序打印，每一层打印到一行。

例如:
给定二叉树: [3,9,20,null,null,15,7],
    3
   / \
  9  20
    /  \
   15   7
返回其层次遍历结果：
[
  [3],
  [9,20],
  [15,7]
]
*/
func levelOrder2(root *TreeNode) [][]int {
	depth := maxDepth2(root)
	fmt.Println(depth)
	result := make([][]int, depth)
	return result
}
func TestThirtyFive(t *testing.T) {
	root := &TreeNode{}
	fmt.Print(levelOrder2(root))
}

func maxDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := maxDepth2(root.Left)
	right := maxDepth2(root.Right)
	return whosMax(left, right) + 1
}
func whosMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
