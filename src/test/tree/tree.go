package tree

import "fmt"

type Node struct {
	Value       int
	Left, Right *Node
}



func (node *Node) PrintValue() {
	fmt.Println(node.Value)
}
