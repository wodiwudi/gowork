package leetcode

/**
请实现一个函数，用来判断一棵二叉树是不是对称的。如果一棵二叉树和它的镜像一样，那么它是对称的。
例如，二叉树 [1,2,2,3,4,4,3] 是对称的。

    1
   / \
  2   2
 / \ / \
3  4 4  3
但是下面这个 [1,2,2,null,3,null,3] 则不是镜像对称的:

    1
   / \
  2   2
   \   \
   3    3
*/
//递归做法
func isSymmetric(root *TreeNode) bool {
	//自顶向下递归
	if root == nil {
		return true
	}
	return isEqualsTree(root.Left, root.Right)
}
func isEqualsTree(left, right *TreeNode) bool {
	//判断双空
	if left == nil && right == nil {
		return true
	}
	//如果一个为空另一个不为空，或两个节点Val不同
	if left == nil || right == nil || left.Val != right.Val {
		return false
	}
	return isEqualsTree(left.Left, right.Right) && isEqualsTree(left.Right, right.Left)
}

//迭代做法 广度优先 一次pop两个，一次push两个
func isSymmetric2(root *TreeNode) bool {
	if root == nil {
		return true
	}
	arr := []*TreeNode{root, root}
	pop := func() *TreeNode {
		tmp := arr[0]
		arr = arr[1:]
		return tmp
	}
	push := func(root1, root2 *TreeNode) {
		arr = append(arr, root1, root2)
	}
	for len(arr) != 0 {
		t1 := pop()
		t2 := pop()

		if t1 == nil && t2 == nil {
			continue
		}
		if t1 == nil || t2 == nil {
			return false
		}
		if t1.Val != t2.Val {
			return false
		}
		push(t1.Left, t2.Right)
		push(t1.Right, t2.Left)
	}
	return true
}
