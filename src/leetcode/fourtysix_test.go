package leetcode

/*
给定一个二叉树，找出其最小深度。
最小深度是从根节点到最近叶子节点的最短路径上的节点数量。
说明: 叶子节点是指没有子节点的节点。

示例:
给定二叉树 [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//DFS深度优先搜索递归写法
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	resMin := 1 << 62
	//注意 if root.Left != nil 必须要有，否则永远nil最小了
	if root.Left != nil {
		left := minDepth(root.Left)
		resMin = getMin(resMin, left)
	}
	if root.Right != nil {
		right := minDepth(root.Right)
		resMin = getMin(resMin, right)
	}
	return resMin + 1

}

func getMin(a, b int) int {
	if a > b {
		return b
	}
	return a
}

//BFS广度优先做 辅助队列实现
func minDepth2(root *TreeNode) int {
	//如果left,right都为空则停止
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	arr := []*TreeNode{root}
	pop := func() *TreeNode {
		res := arr[0]
		arr = arr[1:]
		return res
	}
	push := func(root *TreeNode) {
		arr = append(arr, root)
	}
	length := 1
	for len(arr) != 0 {
		queueLength := len(arr)
		for i := 0; i < queueLength; i++ {
			tmp := pop()
			if tmp.Left == nil && tmp.Right == nil {
				return length
			}
			if tmp.Left != nil {
				push(tmp.Left)
			}
			if tmp.Right != nil {
				push(tmp.Right)
			}
		}
		length++
	}
	return 0
}
