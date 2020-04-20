package main

import (
	"fmt"
	"test/tree"
) //tree包中有一个 Node结构体与PrintValue结构体引用方法
/*
type Node struct {
	Value       int
	Left, Right *Node
}
func (node *Node) PrintValue() {
	fmt.Println(node.Value)
}
*/
//扩充一个后根遍历，在不改动源包的情况下扩充
type myTreeNode struct { //私有struct
	node *tree.Node
}

func (myNode *myTreeNode) traverseBack() {
	if myNode == nil || myNode.node == nil {
		return
	}
	left := myTreeNode{myNode.node.Left}
	right := myTreeNode{myNode.node.Right}
	root := myNode.node //不能为myTreeNode{myNode.node} 因为对于这个结构体来说是在本.go文件下的，它不能看见tree包中的内容
	root2 := myTreeNode{myNode.node}
	fmt.Println(root)
	fmt.Println(root2)
	//myNode.node.Left.TraverseMiddle() //指针不能为接收者，需要定义变量来接送地址
	//myNode.node.Right.TraverseMiddle()
	//myNode.node.PrintValue()
	left.traverseBack()
	right.traverseBack()
	root.PrintValue()
}

func main() {
	root := tree.Node{Value: 3}
	root.Left = &tree.Node{5, nil, nil}
	root.Right = &tree.Node{}
	root.Left.Right = new(tree.Node) //new操作返回引用  等价于 &treeNode{}
	newRoot := myTreeNode{&root}     ///指针不能为接收者，需要定义变量来接送地址
	newRoot.traverseBack()           //0503
	/**********************************/
	fmt.Println()
	root.TraverseMiddle()
	nodeCount := 0
	root.TraverseFunc(func(*tree.Node) {
		nodeCount++
	})
	fmt.Println("Node Count", nodeCount)
	/**********************************/
	fmt.Println()
	MaxValue := 0
	c := root.TraverseWithChannel()
	for node := range c {
		if node.Value > MaxValue {
			MaxValue = node.Value
		}
	}
	fmt.Println("MaxValue :", MaxValue)
}
