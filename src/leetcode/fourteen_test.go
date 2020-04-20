package leetcode

import "testing"

/*
求一个树的深度
   3
   / \
  9  20
    /  \
   15   7

返回3
*/
//递归思想，从叶子节点开始考虑

type TreesNode struct {
	Val   int
	Left  *TreesNode
	Right *TreesNode
}

func maxDepth(root *TreesNode) int {
	if root.Left == nil && root.Right == nil {
		return 1
	}
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)
	return max(left, right) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func TestFourteen(t *testing.T) {

	//	t.Log(maxDepth())
}
