package leetcode

import (
	"fmt"
	"testing"
)

/*
给定一个二叉树，判断其是否是一个有效的二叉搜索树。
假设一个二叉搜索树具有如下特征：
节点的左子树只包含小于当前节点的数。
节点的右子树只包含大于当前节点的数。
所有左子树和右子树自身必须也是二叉搜索树。
示例 1:

输入:
    2
   / \
  1   3
输出: true
示例 2:

输入:
    5
   / \
  1   4
     / \
    3   6
输出: false
解释: 输入为: [5,1,4,null,null,3,6]。
     根节点的值为 5 ，但是其右子节点值为 4 。
*/
var last *TreeNode

//线索二叉树 中序遍历是递增的 一边中序遍历一边判断递增
func isValidBST(root *TreeNode) bool {
	last = &TreeNode{Val: -1 << 63}
	return isBST(root)
}

func isBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if !isBST(root.Left) || root.Val <= last.Val {
		return false
	}
	last = root

	return isBST(root.Right)
}
func TestTweenThree(t *testing.T) {

	fmt.Println(last)
}
