package linklist

import "fmt"

//普通链表实现
type Linklist struct {
	Data int
	Next *Linklist
}

type List struct {
	headNode *Linklist
}

//链表是否为空
func (this *List) IsEmptyLinkList() bool {
	if this.headNode == nil {
		return true
	} else {
		return false
	}
}

//链表长度
func (this *List) LinkLength() int {
	var length int
	cur := this.headNode
	for cur != nil {
		length++
		cur = cur.Next
	}
	return length
}

//头部添加节点
func (this *List) AddHeadLinkList(data int) {
	node := &Linklist{Data: data}
	node.Next = this.headNode
	this.headNode = node
}

//删除某个节点(type = 1 某个位置  type = 2 某个值)
func (this *List) DeleteLinkList(Type, value int) {
	cur := this.headNode
	switch Type {
	case 1:

	case 2:
		if cur.Data == value {
			this.headNode = cur.Next
		} else {
			for cur.Next != nil {
				if cur.Data == value {
					cur.Next = cur.Next.Next
				} else {
					cur = cur.Next
				}
			}
		}

	}
}

//清空链表
func (this *List) EmptyLinkList() {
	this.headNode = nil
}

//遍历链表
func (this *List) SearchLinkList() {
	cur := this.headNode
	for cur != nil {
		fmt.Println(cur.Data)
		cur = cur.Next
	}
}
