package main

import "fmt"

type TreeNode struct {
	Val   int
	left  *TreeNode
	right *TreeNode
}

func createTree(val int) *TreeNode { //工厂函数，Go会考虑建立在堆上还是栈上
	return &TreeNode{Val: val}
}
func (node TreeNode) print() { //接收者相当于this，用于调用方法
	fmt.Println(node.Val)
}
func print1(node TreeNode) { //传参
	fmt.Println(node.Val)
}
func print2(node *TreeNode) {
	fmt.Println(node.Val) //Go没有->解引用，共用.
}
func (node *TreeNode) setValue(val int) {
	if node == nil {
		fmt.Println("nil")
		return
	}
	node.Val = val
}
func (node *TreeNode) traverse() {
	if node == nil {
		return
	}
	node.left.traverse()
	node.right.traverse()
}
func main() {
	var root TreeNode
	root = TreeNode{
		Val: 6,
	}
	root.left = &TreeNode{}
	root.right = &TreeNode{}
	root.left.left = new(TreeNode)
	root.left.right = createTree(2)
	root.print()
	print1(root)
	print2(&root)
	root.left.right.setValue(9)
	fmt.Println(root.left.right.Val)
}
