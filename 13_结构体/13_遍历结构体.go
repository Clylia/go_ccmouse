package tree

// 遍历结构体
func (node *Node) TraverseStruct()  {
	if node == nil {
		return
	}

	node.Left.TraverseStruct()
	node.TraverseStruct()
	node.Right.TraverseStruct()
}
