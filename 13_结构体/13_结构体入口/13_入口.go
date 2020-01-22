package main

import (
	"fmt"
	"learn_go_ccmouse/13_结构体"
)

// 结构体包含结构体实现数据类型的扩展
type myTreeNode struct {
	node *tree.Node
}

// 内嵌方式扩展类型
type myTreeNode2 struct {
	*tree.Node	// Embedding 内嵌
	// 这种内嵌的方式省略了字段名，
	// 假如 var n myTreeNode2
	// 这是要调用内部的指针类型变量，取指针变量的结构体名称即可，即：n.Node 就可
	// 当使用 n.Node.xx 就可以访问 Node 下面的属性和方法，但是这时直接用 n.xx 也可以直接访问 Node 下面的属性和方法
	// 假如此时 n 下面有一个方法和 Node 下面的方法重名，这时 n.FuncName() 默认访问的是 n 的方法，n.Node.FuncName() 才是 Node 下面的方法
	// 总体来说，这种内嵌方式让代码看起来更简洁
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
