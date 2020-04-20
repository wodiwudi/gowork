package leetcode

import "testing"

/*
给定两个二叉树，想象当你将它们中的一个覆盖到另一个上时，两个二叉树的一些节点便会重叠。
你需要将他们合并为一个新的二叉树。合并的规则是如果两个节点重叠，那么将他们的值相加作为节点合并后的新值，
否则不为 NULL 的节点将直接作为新二叉树的节点。
示例 1:

输入:
	Tree 1                     Tree 2
          1                         2
         / \                       / \
        3   2                     1   3
       /                           \   \
      5                             4   7
输出:
合并后的树:
	     3
	    / \
	   4   5
	  / \   \
	 5   4   7

*/
func TestThirtyOne(t *testing.T) {

}
func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil && t2 == nil {
		return nil
	}

	t3 := &TreeNode{}
	if t1 != nil && t2 != nil {
		t3.Val = t1.Val + t2.Val
	} else if t1 != nil && t2 == nil {
		return t1
	} else if t2 != nil && t1 == nil {
		return t2
	}
	t3.Left = mergeTrees(t1.Left, t2.Left)
	t3.Right = mergeTrees(t1.Right, t2.Right)
	return t3
}
