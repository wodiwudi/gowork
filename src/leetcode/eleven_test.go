package leetcode

/*
请编写一个函数，使其可以删除某个链表中给定的（非末尾）节点，你将只被给定要求被删除的节点。
现有一个链表 -- head = [4,5,1,9]，它可以表示为:

示例 1:

输入: head = [4,5,1,9], node = 5
输出: [4,1,9]
解释: 给定你链表中值为 5 的第二个节点，那么在调用了你的函数之后，该链表应变为 4 -> 1 -> 9.
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(node *ListNode) {
	node.Val = node.Next.Val   //把下一个节点的值覆盖到当前节点上
	node.Next = node.Next.Next //把node的下一个节点直接删掉就起到了删除，
	// 前提是要删除的节点不是最后一个节点，因为最后一个节点的next为nil,next的next就找不到了
}