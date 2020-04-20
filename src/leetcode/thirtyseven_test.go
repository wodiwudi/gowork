package leetcode

/*
从上到下打印出二叉树的每个节点，同一层的节点按照从左到右的顺序打印。

例如:
给定二叉树: [3,9,20,null,null,15,7],
    3
   / \
  9  20
    /  \
   15   7
返回：

[3,9,20,15,7]
*/
//广度优先遍历  BFS  建立队列，每次放进一个node，再
func levelOrderBFS(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	queue := []*TreeNode{root} //树队列
	res := []int{}
	push := func(root *TreeNode) {
		queue = append(queue, root)
	}
	pop := func() *TreeNode {
		node := queue[0]
		queue = queue[1:]
		return node
	}
	for len(queue) != 0 {
		node := pop()
		res = append(res, node.Val)
		if node.Left != nil {
			push(node.Left)
		}
		if node.Right != nil {
			push(node.Right)
		}
	}
	return res
}
