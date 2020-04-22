package leetcode

/*
输入两个链表，找出它们的第一个公共节点。

输入：intersectVal = 2, listA = [0,9,1,2,4], listB = [3,2,4], skipA = 3, skipB = 1
输出：Reference of the node with value = 2
输入解释：相交节点的值为 2 （注意，如果两个列表相交则不能为 0）。从各自的表头开始算起，
链表 A 为 [0,9,1,2,4]，链表 B 为 [3,2,4]。在 A 中，相交节点前有 3 个节点；在 B 中，相交节点前有 1 个节点。
*/
//方法1 先对其 然后A B一起走看一不一样
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	lenA, lenB := getLength(headA), getLength(headB)
	if lenA > lenB {
		gap := lenA - lenB
		for gap > 0 {
			headA = headA.Next
			gap--
		}
	} else {
		gap := lenB - lenA
		for gap > 0 {
			headB = headB.Next
			gap--
		}
	}
	//现在两个链表一样长了
	for headA != nil && headB != nil {
		if headA == headB {
			return headA
		}
		headA = headA.Next
		headB = headB.Next
	}
	return nil
}
func getLength(head *ListNode) int {
	length := 0
	for head != nil {
		length++
		head = head.Next
	}
	return length
}

//方法2 一起走 然后如果A走到头再赋值成B的头结点 B也一样 当两个结点相等则退出
func getIntersectionNode2(headA, headB *ListNode) *ListNode {
	nodeA, nodeB := headA, headB
	for nodeA != nodeB {
		if nodeA != nil {
			nodeA = nodeA.Next
		} else {
			nodeA = headB
		}
		if nodeB != nil {
			nodeB = nodeB.Next
		} else {
			nodeB = headA
		}
	}
	return nodeA
}
