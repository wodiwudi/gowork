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
type TreeNode3 struct {
	Val   int
	Left  *TreeNode3
	Right *TreeNode3
}

//普通迭代方法
func preorderTraversal(root *TreeNode3) []int {
	result := []int{}
	if root == nil {
		return nil
	}
	result = append(result, root.Val)
	left := preorderTraversal(root.Left)
	right := preorderTraversal(root.Right)
	result = append(append(result, left...), right...)
	return result
}

func preorderTraversal2(root *TreeNode3) []int {
	//手写栈
	result := []int{}
	return result
}

func TestNineTeen(t *testing.T) {

}
