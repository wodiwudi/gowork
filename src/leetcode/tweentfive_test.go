package leetcode

/*
将两个升序链表合并为一个新的升序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。

示例：
输入：1->2->4, 1->3->4
输出：1->1->2->3->4->4
*/
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}
	l3 := &ListNode{}
	result := l3
	for l1 != nil && l2 != nil {
		if l1.Val > l2.Val {
			l3.Val = l2.Val
			l2 = l2.Next
		} else {
			l3.Val = l1.Val
			l1 = l1.Next
		}
		l3.Next = new(ListNode)
		l3 = l3.Next
	}
	if l1 != nil {
		l3.Val = l1.Val
		l3.Next = l1.Next
	}
	if l2 != nil {
		l3.Val = l2.Val
		l3.Next = l2.Next
	}
	return result
}
