package leetcode

import "math"

/*
输入一棵二叉树的根节点，判断该树是不是平衡二叉树。如果某二叉树中任意节点的左右子树的深度相差不超过1，
那么它就是一棵平衡二叉树。

示例 1:
给定二叉树 [3,9,20,null,null,15,7]

    3
   / \
  9  20
    /  \
   15   7
返回 true 。

示例 2:

给定二叉树 [1,2,2,3,3,null,null,4,4]

       1
      / \
     2   2
    / \
   3   3
  / \
 4   4
返回 false 。
*/
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	_, result := getDepth(root)
	return result
}

func getDepth(root *TreeNode) (int, bool) {
	if root == nil {
		return 0, true
	}
	left, lb := getDepth(root.Left)
	if lb == false {
		return 0, false
	}
	right, rb := getDepth(root.Right)
	if rb == false {
		return 0, false
	}
	result := left - right
	if math.Abs(float64(result)) > 1 {
		return 0, false
	} else {
		return getMax(left, right) + 1, true
	}
}
