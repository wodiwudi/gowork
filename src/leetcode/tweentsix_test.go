package leetcode

/*
给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。
你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。

示例:
给定 1->2->3->4, 你应该返回 2->1->4->3.
*/

//只要是有不断重复的内容且有终止条件，即可形成递归
//递归由三点组成 1.返回值 2.终止条件 3.调用单元内部的逻辑
//这里显然当链表为空或者只有一个链表的时候返回
//内部逻辑 观察最小单元开始，每次都是head重新指向一个新节点，然后head->next->next=head
//返回为head->next，因为最终head->next->next = head了
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	after := head.Next
	head.Next = swapPairs(after.Next)
	after.Next = head
	return after
}
