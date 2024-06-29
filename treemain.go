package main

import (
	"fmt"
	"test1/tree"
)

type myTreeNode struct {
	node *tree.Node
}

type myTreeNode1 struct { //embedding 内嵌
	*tree.Node
}

func (myTreeNode) Traverse() {
	fmt.Println("shadowed traverse")
}
func (myNode myTreeNode1) postOrder1() {
	if myNode.Node == nil {
		return
	}
	myTreeNode{myNode.Left}.postOrder()
	myTreeNode{myNode.Right}.postOrder()
	myNode.Print()
}
func (myNode myTreeNode) postOrder() {
	if myNode.node == nil {
		return
	}
	myTreeNode{myNode.node.Left}.postOrder()
	myTreeNode{myNode.node.Right}.postOrder()
	myNode.node.Print()
}
func createTree(val int) *tree.Node { //工厂函数，Go会考虑建立在堆上还是栈上
	return &tree.Node{Val: val}
}

func print1(node tree.Node) { //传参
	fmt.Println(node.Val)
}
func print2(node *tree.Node) {
	fmt.Println(node.Val) //Go没有->解引用，共用.
}

func main() {
	var root tree.Node
	root = tree.Node{
		Val: 6,
	}
	root.Left = &tree.Node{Val: 3}
	root.Right = &tree.Node{Val: 9}
	root.Left.Left = new(tree.Node)
	root.Left.Left.Val = 8
	root.Left.Right = createTree(2)
	root.Print()
	print1(root)
	print2(&root)
	root.Left.Right.SetValue(12)
	fmt.Println(root.Left.Right.Val)
	root.Traverse()
	fmt.Println()
	(&root).Traverse()
	fmt.Println()
	myTreeNode{&root}.postOrder()
	fmt.Println()
	myTreeNode1{&root}.postOrder1()
	var root1 myTreeNode
	root1.Traverse()
	root1.node.Traverse()
}
