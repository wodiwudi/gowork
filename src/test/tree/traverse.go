package tree

//中序遍历
func (node *Node) TraverseMiddle() {
	node.TraverseFunc(func(node *Node) {
		node.PrintValue()
	})
}

func (node *Node) TraverseFunc(f func(*Node)) { //函数作为参数
	if node == nil {
		return
	}
	node.Left.TraverseFunc(f) //递归调用
	f(node)                   //函数里面可以自己定义动作，不只是print操作，更加灵活
	node.Right.TraverseFunc(f)
}

//前序遍历
func (node *Node) TraverseFont() {
	if node == nil {
		return
	}
	node.PrintValue()
	node.Left.TraverseFont()
	node.Right.TraverseFont()
}

func (node *Node) TraverseWithChannel() chan *Node {
	out := make(chan *Node)
	go func() {
		node.TraverseFunc(func(node *Node) {
			out <- node
		})
		close(out)
	}()
	return out
}
