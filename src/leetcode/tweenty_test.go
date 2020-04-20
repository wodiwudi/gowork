package leetcode

import (
	"testing"
)

/*
给定一个 N 叉树，返回其节点值的后序遍历、前序遍历、层次遍历
层次遍历返回结果如下。
[
     [1],
     [3,2,4],
     [5,6]
]
*/
type Node struct {
	Val      int
	Children []*Node
}

func TestTweenTy(t *testing.T) {

}

//后根遍历
func postorder(root *Node) []int {
	if root == nil {
		return nil
	}
	res := []int{}
	for _, v := range root.Children {
		node := postorder(v)
		res = append(res, node...)
	}
	res = append(res, root.Val)
	return res
}

//先根遍历
func preorder(root *Node) []int {
	if root == nil {
		return nil
	}
	res := []int{}
	res = append(res, root.Val)
	for _, v := range root.Children {
		node := preorder(v)
		res = append(res, node...)
	}
	return res
}

//层次遍历
func levelOrder(root *Node) [][]int {
	res := [][]int{}
	length := maxDepthNode(root)
	for i := 0; i < length; i++ {

	}
	return res
}

//N叉树的深度
func maxDepthNode(root *Node) int {
	if root == nil {
		return 0
	}
	max := 0
	for _, v := range root.Children {
		depth := maxDepthNode(v)
		max = getMax(max, depth)
	}
	return max + 1
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
