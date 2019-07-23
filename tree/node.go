package tree

import "fmt"

/**
|-------------------------------------------------------------------------------------------------
|	go 结构体 struct
|
|
|	结构体的创建
|		var root Node
|		root = Node{value: 3}
|		root.left = &Node{}
|		root.right = &Node{5, nil, nil}
|		root.left.right = new(Node)
|		不论地址还是结构本身，一律使用 . 来访问成员
|		调用结构体的方法时也使用 . 调用
|
|	在调用结构体中传值的方法和传引用两种不同的方法时，结构体实例不需要区分是否是指针，
|	当调用传旨的方法时，系统会自动在内存中拷贝一份，然后执行后面程序
|	当嗲用传引用的方法时，方法接收到的结构体实例在内存中的地址
|
|	封装
|		名字一般使用CamelCase
|		首字母大写：public
|		首字母小写：private
|
|	包
|		每个目录一个包
|		main 包包含可执行入口
|
|-------------------------------------------------------------------------------------------------
*/
// 定义一个结构体
type Node struct {
	Value int
	Left, Right *Node
}

// 给结构体定义方法
// (node Node) 为方法的接收者，相当于其他语言的 this
// funcName 为方法名
func (node Node) Print() {
	fmt.Print(node.Value, " ")
	fmt.Println()
}

// 设置 Node 这个 struct 中的 value 属性的值
func (node *Node) SetValue(value int) {
	if node == nil {
		fmt.Println("Setting value to nil node. Ignored")
		return
	}
	node.Value = value
}

// 给结构体创建节点
func CreateNote(value int) *Node {
	return &Node{Value: value}
}
