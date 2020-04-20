package leetcode

import "testing"

/*
给定一个二叉树，返回它的中序 遍历。
示例:

输入: [1,null,2,3]
   1
    \
     2
    /
   3

输出: [1,3,2]
*/

type TreeNode2 struct {
	Val   int
	Left  *TreeNode2
	Right *TreeNode2
}

func inorderTraversal(root *TreeNode2) []int {
	result := make([]int, 0)
	if root == nil {
		result = append(result, root.Val)
	}
	inorderTraversal(root.Left)
	result = append(result, root.Val)
	inorderTraversal(root.Right)
	return result
}

func TestEighteen(t *testing.T) {

}
