package main

import (
	"fmt"
	"test1/tree"
)

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
	root.Left = &tree.Node{}
	root.Right = &tree.Node{}
	root.Left.Left = new(tree.Node)
	root.Left.Right = createTree(2)
	root.Print()
	print1(root)
	print2(&root)
	root.Left.Right.SetValue(9)
	fmt.Println(root.Left.Right.Val)
	root.Traverse()
	(&root).Traverse()
}
