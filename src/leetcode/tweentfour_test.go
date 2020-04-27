package leetcode

import (
	"testing"
)

/*
请完成一个函数，输入一个二叉树，该函数输出它的镜像。
例如输入：

     4
   /   \
  2     7
 / \   / \
1   3 6   9
镜像输出：

     4
   /   \
  7     2
 / \   / \
9   6 3   1
*/
//思路：从根开始 如果为空节点返回nil 否则交换左右根地址，在递归左右子树。自顶向下的思想。
func TestTweentFour(t *testing.T) {

}

func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left, root.Right = root.Right, root.Left
	mirrorTree(root.Left)
	mirrorTree(root.Right)
	return root
}
