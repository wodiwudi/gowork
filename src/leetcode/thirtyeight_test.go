package leetcode

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
func levelOrderBFS2(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	queue := []*TreeNode{root}
	res := [][]int{}
	push := func(root *TreeNode) {
		queue = append(queue, root)
	}
	pop := func() *TreeNode {
		node := queue[0]
		queue = queue[1:]
		return node
	}
	for len(queue) != 0 { //跟前一题比多加了个中间数组，每回len(queue)后要给一个新变量 因为pop后长度会变
		tmp := []int{}
		length := len(queue)
		for i := 0; i < length; i++ {
			node := pop()
			tmp = append(tmp, node.Val)
			if node.Left != nil {
				push(node.Left)
			}
			if node.Right != nil {
				push(node.Right)
			}
		}
		res = append(res, tmp)
	}
	return res
}
