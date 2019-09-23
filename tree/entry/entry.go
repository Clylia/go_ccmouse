package main

import (
	"fmt"
	"learn_go_ccmouse/tree"
)

// 定义一个内容中包含另一个结构体的结构体
type myTreeNode struct {
	node *tree.Node
}

// myTreeNode 这个结构体的方法
func (myNode *myTreeNode) postOrder()  {
	if myNode == nil || myNode.node == nil {
		return
	}

	left := myTreeNode{myNode.node.Left}
	right := myTreeNode{myNode.node.Right}
	left.postOrder()
	right.postOrder()
	myNode.node.Print()
}

// todo， 执行出错
// build command-line-arguments: cannot load learn_go_ccmouse/tree: cannot find module providing package learn_go_ccmouse/tree
func main() {
	// 创建结构体
	var root tree.Node
	root = tree.Node{Value: 3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNote(2)

	// 调用结构体的方法，以 . 调用
	root.Print()
	root.Right.Left.SetValue(4)
	root.Right.Left.Print()

	root.TraverseStruct()	// 遍历结构体

	myRoot := myTreeNode{&root}
	myRoot.postOrder()
	fmt.Println()
}
