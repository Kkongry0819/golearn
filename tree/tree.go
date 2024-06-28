package tree

import "fmt"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func (node *Node) Traverse() {
	if node == nil {
		return
	}
	node.Left.Traverse()
	fmt.Printf("%d ", node.Val)
	node.Right.Traverse()
}
func (node Node) Print() { //接收者相当于this，用于调用方法
	fmt.Println(node.Val)
}
func (node *Node) SetValue(val int) {
	if node == nil {
		fmt.Println("nil")
		return
	}
	node.Val = val
}
